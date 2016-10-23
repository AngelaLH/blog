Id: 912
Title: Open-source and windows
Date: 2005-10-12T20:19:18-07:00
Format: Markdown
--------------
Do you run an open-source project? Want it to support Windows? Wonder why
windows programmers are not jumping in to help you?

It's because you're making it hard for them.

As an example, see this recent [Inkscape status report][1]:

Special mention needs to be given to the Windows-specific Inkscape
bugs. This is probably the weakest area of our project right now;
even the relatively new OSX port is advancing more rapidly. In part the
trouble is simply lack of one or two coders to put attention into the
port. If you're on Windows, can code, and love Inkscape, then check out
the bug tracker and wiki to see how you can lend a hand towards making
the Windows port of Inkscape robust!

So I looked at what it would take to do that. The first step would be
to compile and run the program. Turns out that in order to do that I
need to follow a [lengthy tutorial][2]
and install several other packages, including MingW, a hack that turns
gcc toolchain into something that can create windows executables.

It's an easy cop-out for Unix developers, because they don't have to
worry about writing their code in a way that supports compilers other
than gcc but it's way too much trouble for windows developers.

First, it requires them to take on a very steep learning curve of
completely different, alien and complicated tool-chain (autotools, gcc,
gdb etc.).

Assuming that they don't shoot themselves out of despair in the
process, even if they climb that hill it leaves them with primitive,
under-powered toolchain that doesn't play with other development
windows tools like decent profilers and debuggers, mostly because MingW
cannot generate *.pdb files which are required for debugging.

As an example, I once had a problem with Gaim on Windows - when
gadu-gadu protocol was Gaim would crash from time to time. No biggie, I
though, I'm a programmer so I can fix it myself, thanks to the
greatness of open-source. I spent a day downloading all the necessary
files, setting up build MingW environment and compiled a debug version.
It had a buffer over-run that was probably a reason for the crash later
on and visual c runtime (which MingW uses) was nice enough to point it
out to me when I was running it under Microsoft's free WinDBG debugger.

If I knew which memory was overwritten, it would have been an easy fix.
If the binary was built with native Windows tools like Visual Studio,
WinDBG would have told me what was overwritten. But since it was
compiled with MingW I had no such luck. And due to the nature of memory
corruption bugs, the crash whose callstack I could get happened in
completely unrelated code. So instead of an easy fix I was faced with
lots of work (trying to find a bug by eye-balling the code or
instrumenting random pieces). Thanks, but no thanks.

I've looked at many open-source projects and found that those that have
good Windows support usually have the windows development support on par with
Unix one i.e. all a Windows developer needs to do in order to compile
them natively on Windows is to:

  * checkout CVS
  * open Visual Studio project
  * hit go

There are a couple of very prominent open-source projects (Gaim, GIMP,
Inkscape) that have very weak windows support and they're also the ones
that require a lot of hoop jumping in order to contribute.

This is more of an attitude problem than a technical problem. It's
possible to write code that compiles both under Visual C under Windows
and gcc under Unix. It's possible to maintain a parallel build system
(Visual Studio project files and/or nmake makefiles and autotools/gmake
makefiles). But often maintainers of open-source projects originated on
Unix will FUD you to death with extreme exaggerations of how that will
lead to messy code and, in consequence, end of the world.

So: if you're running an open-source project and want a Windows port
that is alive and kicking, beg Windows developers to contribute a
native build system and Windows portability fixes instead of forcing
them to learn Unix and use sub-standard (on Windows) tools. It's just one
small example of optimizing [development for fun][3].

   [1]: http://www.inkscape.org/status/en/status_20051001.php (Inkscape status report)
   [2]: http://inkscape.org/win32/win32buildnotes.html (lengthy tutorial)
   [3]: http://www.oreillynet.com/pub/wlg/7996 (development for fun)

