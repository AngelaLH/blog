Id: 1859
Title: A collaborative text editor for Windows
Date: 2004-08-30T13:36:57-07:00
Format: Markdown
--------------
[SubEthaEdit](http://www.codingmonkeys.de/subethaedit/), a collaborative
text editor for Mac OS, received a lot of buzz and for a good reason. It
does basic text editing well, is a well designed application and adds a
killer feature not available in other editors: collaborative editing.

Unfortunately for me it only works under Mac OS while I mostly use
Windows so I was looking for similar Windows app.
[multi-editoro](http://me.sphere.pl/) is such an app. I only tried it
for a while for non-collaborative editing so can't yet tell how it works
for its indendent purpose but I do have a few comments on the app.

The app is much less polished that SubEthaEdit. Even though it's a
Windows app, for the UI it seems to be using some sort of curses-based
interface. This is not acceptable for a Windows app. It's ok for toy
apps but not ok for end-user apps intendent for average users. I have no
doubts that the reason for that is that it's much harder to write UI in
native Windows and that Unix version is also possible but users don't
care.

The main page on a website is a big flash animation. Seriously, when
will people learn? You would think that since books like [Philip and
Alex's Guide to Web Publishing](http://philip.greenspun.com/panda/) or
resources like [Alertbox](http://www.useit.com/alertbox/) are available
for free on the web, people doing websites would read them and stopped
using Flash gratuitiously since it makes website less usable in many
ways than a simple html page, not to mention the maintenance problems
for webmasters when they have to change anything on the website.

A better name wouldn't hurt. There's already at least one multi-editor
out there and one letter difference in a name isn't quite enough.

An installation file should also be provided - currently it's just a zip
file.

By default, the text I write should not be highlighted.

So the final verdict is that multi-editoro is an interesting proof of
concept that is worth keeping an eye on but it's far from being a
usable, end-user application.

I cannot be 100% sure but it seems that the author wrote his own text
editor engine which to me is a big waste of time. Writing a
high-quality, bug free text editor and display engine with a lot of
features takes years. Collaborative editing part is comparatively small
part of the overall code. I don't understand why people choose to waste
a lot of time re-implementing the code that they can get [for
free](http://www.scintilla.org/). The best way to get a high-quality
collaborative editor would be to just extend
[SciTe](http://www.scintilla.org/SciTE.html) (or some other mature,
open-source text editor) and add collaborative features to it.
