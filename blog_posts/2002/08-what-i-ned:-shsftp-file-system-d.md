Id: 1379
Title: What I need: ssh/sftp file system driver for Windows
Date: 2002-08-17T18:36:27-07:00
Format: Markdown
Deleted: yes
--------------
What would be great to have is a ssh/sftp file system driver for
Windows. Currently to exchange files between my servers running ssh and
my primary Windows machine I use [WinSCP](http://winscp.vse.cz/eng/),
and for remote access I use
[Putty](http://www.chiark.greenend.org.uk/~sgtatham/putty/) (both
programs are free).

What would be even more useful is an ability to mount the filesystem of
a remote machine into a drive on Windows. Technically it's possible.
Some operations (changing small parts of the file) would probably suck
because I don't think the protocol supports the full i/o sementics so
the filesystem driver would have to cache the file locally and every
change would require writing the whole file out. It would still work
great for file copying (and I could use any app I want for that and not
be limited by what WinSCP can do) and editing remote files.
