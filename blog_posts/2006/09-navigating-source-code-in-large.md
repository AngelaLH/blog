Id: 1969
Title: Navigating source code in large programs
Date: 2006-09-21T12:35:37-07:00
Format: Markdown
--------------
[Navigating Linux Source Code](http://www.ddj.com/184401358) is an
article that tells you, in many words, how to spend hours to get a
somewhat usable way of navigating large code base.\
\
A feel compelled to point out that Windows users have a superior
solution: [Source Insight](http://www.sourceinsight.com/) editor. I've
learned about it many, many years ago when I started my first
programming job at Microsoft. Everyone on my team (with one exception)
was using Source Insight as their programming editor. Many teams at
Microsoft were using it. At that time Visual Studio 6 (or any other
editor) was no competition. Visual Studio 2005 closes the gap but I
still feel blind when I'm not using Source Insight when writing or
browsing the code. Source Insight works best for C/C++/Java/C\# code,
because it has built-in semantic parsers for those languages (as opposed
to simplistic syntax coloring available in most editors). When it comes
to speed and ease of use when browsing unfamiliar code base, nothing
I've seen so far comes even close to Source Insight. It's really that
good.\
\
A text editor is not something that can be appreciated in a day or two.
You have to stick with it for some time before you learn the good way of
doing things, but Source Insight is well worth the investement.\
\
The crucial thing are projects: those are the units best understood by
Source Insight and only when you create a project you get access to the
time saving features. Project is just a set of files and you have to
create it explicitly. Obviously, files in a project usually are all
files for a given program. Once you create a project, Source Insight
parses the files and builds a database of all symbols. It's like tags
except Source Insight parser is almost as good as compiler's parser and
the database is updated while you change the files. Once you have the
symbol database, navigating the source code in many ways is just a
snap.\
\
Another crucial thing is that Source Insight is optimized for really
large projects. I've created projects with millions lines of code (linux
kernel, mozilla, webkit etc.) and Source Insight handled them blazingly
fast.\
\
As it usually is with editors, no-one really changes their habits. Once
they become vi or emacs or foo addicts, nothing will make them switch.
But do try Source Insight if you ever find yourself working on large
C/C++/Java/C\# project.\
\
The bottom line is that Source Insight saves time and cuts down
annoyance. When working with large code we have to switch between
various files, find out what is the exact definition of function foo or
find that function that has "append" in the name but we forgot what
exactly was the name and how this function worked. In Source Insight
doing those things is fast. I've watched an experience vi user navigate
between multiple source files and felt sorry for him seeing how much
effort he had to spend doing the same thing that I could do in Source
Insight with ctrl-o and a few keystrokes. I also feel sorry (or
infuriated, depending on the day) when I can't use Source Insight and
have to use dumb programming editors that can't tell a difference
between a function, a keyword or a variable.\
\

