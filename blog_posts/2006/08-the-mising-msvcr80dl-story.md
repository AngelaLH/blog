Id: 1958
Title: The missing msvcr80.dll story
Tags: programming,msvc
Date: 2006-08-07T14:32:16-07:00
Format: Markdown
--------------
It all started [with a
complaint](http://blog.kowalczyk.info/forum_sumatra/topic.php?TopicId=29)
that [Sumatra PDF v0.2](http://blog.kowalczyk.info/software/sumatrapdf/)
(a yellow PDF viewer that I've just released) doesn't run.\
\
After some psychic debugging and a little help from my friends I figured
out that the reason is that it won't run on machines that lack
msvcr80.dll.\
\
Programs usually use C library calls that usually reside in a shared
library to maximize the amount of shared code between different apps
running at the same time.\
\
For ages (which means up to Visual Studio 6) you could pretty much rely
on that DLL being present on people's machines.\
\
Visual Studio 2002/2003/2005 each use their own version of
msvcr70.dll/msvcr71.dll/msvcr80.dll. And that is a problem. Two
problems, actually.\
\
The first problem is that you can't expect that msvcrt80.dll is
available on people's machine. This is nasty, because a developer, be
virtue of installing Visual Studio 2005, has msvcrt80.dll installed, so
he doesn't see the problem. Only when he gives the binary to someone to
run on a machine without msvcr80.dll installed, it mysteriously fails.
I've seen it a couple of times in the past and still make this mistake.
You can diagnose this (and other DLL-loading related problems) using
[Dependency Walker](http://www.dependencywalker.com/).\
\
What's the solution to problem one?\
\
Distribute msvcr80.dll (and others you depend on that are not guranteed
to be installed everywhere) with your app. MSDN has articles on that
e.g. [this one](http://msdn2.microsoft.com/en-us/library/ms235265.aspx)
and [this one](http://msdn2.microsoft.com/en-us/library/8kche8ah.aspx).\
\
The solution I've chosen is static linking. I changed C++\\Code
Generation\\Runtime Library setting from multi-threaded DLL (/MD) to
just multi-threaded (/MT). This caused Visual Studio to complain about
the conflict with libcmtd.lib. I don't understand why or is it serious
or not, but just in case I've added libcmtd.lib to the "Linker\\Ignore
Specific Library" list.\
\
Problem number two is worse. Due to how C library is implemented, you
can't use different versions of C library at the same time. The need for
that happens surprising often, especially in the days of abundance of
open-source code. Imagine your app wants to use libxml2 and you
downloaded libxml2 library and headers files that were compiled in
Visual Studio 6. libxml2 opens files so it had to link to VS 6 version
of C library. Imagine that you're compiling your app with Visual Studio
2005 (e.g. because you can no longer buy VS 6) and you also need to use
C library. It's not possible to link from Visual Studio 2005 to the
older C library so you have to link to msvcr80.dll. Now you're in
trouble.\
\
When you call both old C library and new one, things might work but they
also can break mysterioiusly. I once debugged a problem caused exactly
by that. A python app compiled with Visual Studio 2005 was calling into
an extension compiled with VS 6. As it happend, the code compiled with
VS 2005 was closing a file descriptor (i.e. calling close() inside
msvcr80.dll) that was allocated by an extension (using open() call
inside VS 6 version of C lib dll). And the app would die right there.
Turns out that file descriptor (which are just integers) to Windows file
handle mapping is part of C lib. The same file descriptor in one dll
maps to a completely different file handle in the other dll.\
\
There is no easy fix for this problem. If you can, your best choice is
to re-compile everything under the same compiler. For many open-source
project it's easier said than done. Because of strong Unix heritage they
often care little about the state of Windows build (if it compiles under
Cygwin, it's clearly good enough) so a simple act of compilation can
turn into multiple-day endeavour of creating a build system and fixing
small but annoying portability problems.\
\
\

