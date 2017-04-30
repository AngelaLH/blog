package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
)

type secureCookieValue struct {
	TwitterUser string
	TwitterTemp string
}

// IsAdmin returns true if the user is an admin
func IsAdmin(r *http.Request) bool {
	return getSecureCookie(r).TwitterUser == "kjk"
}

func getLogInOutURL(r *http.Request) string {
	url := r.URL.Path
	if IsAdmin(r) {
		return "/logout?redirect=" + url
	}
	return "/login?redirect=" + url
}

func setSecureCookie(w http.ResponseWriter, cookieVal *secureCookieValue) {
	val := make(map[string]string)
	val["twuser"] = cookieVal.TwitterUser
	val["twittertemp"] = cookieVal.TwitterTemp
	if encoded, err := secureCookie.Encode(cookieName, val); err == nil {
		// TODO: set expiration (Expires    time.Time) long time in the future?
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	} else {
		logger.Errorf("setSecureCookie(): error encoding secure cookie %s\n", err)
	}
}

const weekInSeconds = 60 * 60 * 24 * 7

// to delete the cookie value (e.g. for logging out), we need to set an
// invalid value
func deleteSecureCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "deleted",
		MaxAge: weekInSeconds,
		Path:   "/",
	}
	http.SetCookie(w, cookie)
}

func getSecureCookie(r *http.Request) *secureCookieValue {
	ret := &secureCookieValue{}
	if cookie, err := r.Cookie(cookieName); err == nil {
		// detect a deleted cookie
		if "deleted" == cookie.Value {
			return &secureCookieValue{}
		}
		val := make(map[string]string)
		if err = secureCookie.Decode(cookieName, cookie.Value, &val); err != nil {
			// most likely expired cookie, so ignore. Ideally should delete the
			// cookie, but that requires access to http.ResponseWriter, so not
			// convenient for us
			//logger.Noticef("Error decoding cookie %q, error: %s", cookie.Value, err)
			return &secureCookieValue{}
		}
		var ok bool
		if ret.TwitterUser, ok = val["twuser"]; !ok {
			logger.Errorf("Error decoding cookie, no 'twuser' field")
			return &secureCookieValue{}
		}
		if ret.TwitterTemp, ok = val["twittertemp"]; !ok {
			logger.Errorf("Error decoding cookie, no 'twittertemp' field")
			return &secureCookieValue{}
		}
	}
	return ret
}

func decodeUserFromCookie(r *http.Request) string {
	return getSecureCookie(r).TwitterUser
}

func decodeTwitterTempFromCookie(r *http.Request) string {
	return getSecureCookie(r).TwitterTemp
}

// getTwitter gets a resource from the Twitter API and decodes the json response to data.
func getTwitter(cred *oauth.Credentials, urlStr string, params url.Values, data interface{}) error {
	if params == nil {
		params = make(url.Values)
	}
	oauthClient.SignParam(cred, "GET", urlStr, params)
	resp, err := http.Get(urlStr + "?" + params.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyData, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("Get %s returned status %d, %s", urlStr, resp.StatusCode, bodyData)
	}
	//fmt.Printf("getTwitter(): json: %s\n", string(bodyData))
	return json.Unmarshal(bodyData, data)
}

// GET /oauthtwittercb?redirect=$redirect
func handleOauthTwitterCallback(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("handleOauthTwitterCallback()\n")
	redirect := strings.TrimSpace(r.FormValue("redirect"))
	if redirect == "" {
		httpErrorf(w, "Missing redirect value for /login")
		return
	}
	tempCred := oauth.Credentials{
		Token: r.FormValue("oauth_token"),
	}
	tempCred.Secret = decodeTwitterTempFromCookie(r)
	if "" == tempCred.Secret {
		http.Error(w, "Error getting temp token secret from cookie, ", 500)
		return
	}
	//fmt.Printf("  tempCred.Secret: %s\n", tempCred.Secret)
	tokenCred, _, err := oauthClient.RequestToken(http.DefaultClient, &tempCred, r.FormValue("oauth_verifier"))
	if err != nil {
		http.Error(w, "Error getting request token, "+err.Error(), 500)
		return
	}

	//fmt.Printf("  tokenCred.Token: %s\n", tokenCred.Token)

	var info map[string]interface{}
	if err := getTwitter(
		tokenCred,
		"https://api.twitter.com/1.1/account/verify_credentials.json",
		nil,
		&info); err != nil {
		http.Error(w, "Error getting timeline, "+err.Error(), 500)
		return
	}
	if user, ok := info["screen_name"].(string); ok {
		cookie := getSecureCookie(r)
		cookie.TwitterUser = user
		setSecureCookie(w, cookie)
	}
	http.Redirect(w, r, redirect, 302)
}

// GET /login?redirect=$redirect
func handleLogin(w http.ResponseWriter, r *http.Request) {
	redirect := strings.TrimSpace(r.FormValue("redirect"))
	if redirect == "" {
		httpErrorf(w, "Missing redirect value for /login")
		return
	}
	q := url.Values{
		"redirect": {redirect},
	}.Encode()

	cb := "http://" + r.Host + "/oauthtwittercb" + "?" + q
	//fmt.Printf("handleLogin: cb=%s\n", cb)
	tempCred, err := oauthClient.RequestTemporaryCredentials(http.DefaultClient, cb, nil)
	if err != nil {
		http.Error(w, "Error getting temp cred, "+err.Error(), 500)
		return
	}
	cookie := &secureCookieValue{TwitterTemp: tempCred.Secret}
	setSecureCookie(w, cookie)
	http.Redirect(w, r, oauthClient.AuthorizationURL(tempCred, nil), 302)
}

// GET /logout?redirect=$redirect
func handleLogout(w http.ResponseWriter, r *http.Request) {
	redirect := strings.TrimSpace(r.FormValue("redirect"))
	if redirect == "" {
		httpErrorf(w, "Missing redirect value for /logout")
		return
	}
	deleteSecureCookie(w)
	http.Redirect(w, r, redirect, 302)
}
