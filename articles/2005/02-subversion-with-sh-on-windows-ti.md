Id: 895
Title: Subversion with SSH on Windows tip
Date: 2005-02-09T01:12:23-08:00
Format: Markdown
--------------
If you're running Subversion with SSH on Windows, you need to have an
SSH client installed (I use plink which is part of putty) and set
SVN\_SSH apropriately. Here's a tip: make sure to run at least one query
against the server that holds your source code e.g. run a simple command
like ls: `plink -pw $YourPassword $YourLogin@$YourServer ls`. To
increase security SSH calculates a checksum for the server you're trying
to connect to (so that e.g. if someone hi-jacks DNS server and redirects
your SSH connection to their own server with the intent of capturing
your login/password, SSH will warn you that server's signature doesn't
match previously seen signature). The problem is that it's an
interactive process i.e. SSH displays a warning and waits for you to
accept the new server signature. It also does that when the signature
doesn't exist (i.e. this is the first time you connect to a given server
via SSH). Subversion doesn't show that message but still waits for
user's confirmation which looks like it hung.

I've just spent an hour trying to figure out such problem.
