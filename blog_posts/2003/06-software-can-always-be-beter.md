Id: 1653
Title: Software can always be better
Date: 2003-06-27T22:15:03-07:00
Format: Markdown
--------------
I told you that Frigate is the best file manager I know but it doesn't
mean that I think that it couldn't be better.

Recently I wrote about the need for being consistent. Ironically,
there's another, opposite side to this story: it seems that usually
people just stick with the tried ideas which makes for suboptimal
progress in software.

The best place to observe that is in a category with lots of competition
e.g. file manager, simple text editors or, to use a recent example,
Wikis or RSS aggregators. It looks like the first entrants create such a
strong precedent that the followers are usually content with copying
them and adding just a few, relatively small features. I rarely see a
radical improvements or a significant number of new features in any
given applications.

They slowly grow as a whole because everyone copies other people's
features (which is good) but when I look at the difference between
Frigate and Norton Commander I'm really disappointed by the progress in
the state-of-the-art of file management apps in more than 10 years.

There's probably more than one reason for that:

-   writing software, in general, is hard
-   there's probably not much money in the "just another file manager"
    business (given the insane competition from free stuff (including
    Explorer built into Windows) and other file managers) so it's
    probably fairly easy to whip up something that is comparable to
    existing stuff but much harder to justify investing into making it
    much better

I'll make a claim that part of the problem is software designer's mental
block caused by existing implementations that are good enough. The
solution to that problem? A conscious effort on finding and implementing
new, useful features. (Here I have to note that Apple, with their iApps,
is doing a good job by significantly improving upon tired (you would
think) concepts like music player or chat.)

Here are my ideas for a few features that I haven't seen in existing
file managers.

Build a database of files. It's not, by any means, a novel idea but I
haven't seen anyone doing it. The reason is that it makes it possible to
improve some features. For example, finding files by name on a
hard-drive is still painfully slow because scanning the file system of a
large hard-drive is slow. If we had names and other vital properties
indexed in a database it would be very fast (probably even possible to
provide interactive performance). The same holds for calculating the
size of a directory: it's slow if you re-scan all files in a directory
with lots of files. If we had this data pre-calculated, there's no
reason to have it always displayed. Of course, there are problems with
this approach: you have re-scan the hard-drive periodically and you
cannot guarantee being always in perfect sync with the hard-drive but
those disadvantages are minor compared to what they enable.

Virtual folders. In a way virtual folders are already in: most file
managers enable you too treat a zip file as a directory but I want more.
Once we have file names, their sizes and other important meta-data in a
database, why shouldn't I be able to create virtual folder that shows
the files matching an arbitrary SQL query. That's ultimate flexibility.
How would it be useful? For example I frequently view PDF files, create
text files or Excel spreadsheets, view Word files, PowerPoint files etc.
The files are never in the same directory so when I want to switch from
within Adobe Acrobat Reader between a currently viewed PDF files in my
d:downloads directory to a PDF file in "My Documents" directory I have
some folder navigation to do. Annoying if you do it often. Applications
usually try to make my life easier by providing a list of a few recently
used documents but those are application-specific fixes and I usually
find myself accessing more documents than the length of this list. If I
had my imaginary friend I could create a virtual folder with 10 (or 20
or whatever is the appropriate limit for me) recently accessed Excel
spreadsheets or Word documents or PDF documents. Or I could have one
folder for all those types of documents. Or I could have a folder with
manually maintained list of files that I access most frequently (program
could even help me choose which files are those by watching which files
I access most frequently). Or I could have a virtual folder than only
shows jpeg files from all my hard-drives. I could view the files in this
folder in a traditional, hierarchical view or as a flat list. I could
also sort them as in a regular folder i.e. by name or data. Or I could
search for files matching given name but only restricted to the folder
with jpeg files. And it all would be fast. Think how easy would it be to
browse through my collection of digital photos, how easy would it be to
find recently taken pictures.

Integrated ssh/sftp. It's the Internet era, god damn it. I've seen ftp
integration but not ssh/sftp (there's
[WinSCP](http://winscp.vse.cz/eng/), which is mostly about ssh but it's
a lousy file manager - I want first-class file manager with first-class
support for ssh).

Break out with 2-pane tradition. Here I'm not convinced that it's even a
good idea but it's certainly something that could be tried. Why do we
usually have 2+ panes: because that's makes it easy to move files
between directories (we see both the files we selected in the source
directory and the destination directory). Explorer has only one pane and
you can copy files by selecting them in source directory, doing Ctrl-C,
navigating to destination directory and doing Ctrl-V. Possible but
inferior to the 2-pane way because we might loose track of what is it
that we're copying and it's not possible to select files from multiple
directories at once. Can we improve upon that? Imagine that we only have
one pane but it gets split vertically when you select files. The bottom
shows currently selected files (you unselect files by removing them from
that bottom view). This allows us to select files from multiple
directories at once, we always know what files are selected and to
copy/move files to a different directory we just navigate there and
press F5/F6. We could improve upon that by adding tabs so we can easily
navigate between frequently accessed folders (the bottom view with
selected files wouldn't change).
