Id: 898
Title: How to delete a file you get from urllib.urlretrieve()
Tags: python
Date: 2005-05-04T23:40:58-07:00
Format: Markdown
--------------
I just hit an interesting corner case in Python. Urllib module has a very
useful function urllib.urlretrieve(url, filePath) which will get a given url
and save it to a file. This function, however, might fail, in which case we
would like to delete it so that we don't get confused by partially downloaded,
corrupted file.

To detect failure of urlretrieve() we can simply wrap it
inside exception handling. As experience showed, at least on windows, calling
os.remove() on the file immediately after the exception happened will fail
with 'Access denied'. According to docs, this error is returned if file is
still being used. That is plausible (although not something I would expect).

My solution to the problem was to retry os.remove() few times, sleeping for a
second between tries, therefore giving python time to close the file used in
urlretrieve(). This seems to work. And here's the full code:

```python
def fFileExists(filePath):
  try:
    st = os.stat(filePath)
  except OSError:
    # TODO: should check that Errno is 2
    return False
  return True

failedDownload = False
try:
  urllib.urlretrieve(url, filePath)
  break
except:
  print "exception: n  %s, n  %s, n  %s n  when downloading %s" %
(sys.exc_info()[0], sys.exc_info()[1], sys.exc_info()[2], url)
  failedDownload = True
  # remove potentially only partially downloaded file

if failedDownload and fFileExists(filePath):
  removeRetryCount = 0
while removeRetryCount < 3:
  time.sleep(1) # try to sleep to make the time for the file not be used anymore
  try:
    os.remove(filePath)
    break
  except:
    print "exception: n  %s, n  %s, n  %s n  when trying to remove
file %s" % (sys.exc_info()[0], sys.exc_info()[1], sys.exc_info()[2], filePath)
    removeRetryCount += 1
```
