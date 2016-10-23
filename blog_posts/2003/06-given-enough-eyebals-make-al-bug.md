Id: 1640
Title: Given enough eyeballs make all bugs shallow
Date: 2003-06-04T21:11:20-07:00
Format: Markdown
--------------
One famous argument made by [Eric
Raymond](http://armedndangerous.blogspot.com/) in favor of open-source
is "[given enough eyeballs, all bugs are
shallow.](http://www.openresources.com/documents/cathedral-bazaar/node4.html)"
The argument goes like this: in open-source you get the source code to
software so people can look at it and discover bugs. This kind of
examination isn't possible in closed source development therefore open
source is better and leads to software of higher quality.

There's only one problem: **in real world you don't have enough
eyeballs**, therefore the argument remains purely theoretical.

Here's a little story: I found an interesting (judging by screenshots),
open-source software meld (which is a GUI for CVS). It only runs under
Linux so I fired up my RedHat 8 (a mainstream Linux distribution) and
tried to run it. I've downloaded the sources and tried to run it but
failed because it depends on pygtk that I didn't have installed. Looking
for rpms for obscure packages is usually a waste of time so I just
grabbed the latest sources for pygtk, compiled and installed it (at
which point I did something that a majority of users isn't capable of
(and they shouldn't be)). Then I run meld:

    [kjk@localhost meld]$ ./meld
    Segmentation fault (core dumped)

Which basically says that there was something wrong with the software.
I'm a software developer by profession and I'm capable of fixing the
problem, whatever it might be. Segmentation fault is usually caused by
trying to access memory you're not allowed to access and, if you have
things set up properly, you'll get dumped into a debugger in a place in
the program exactly where the problem happened. From that it's usually
simple to look up the callstack, see the sequence of calls and figure
out why you're playing with invalid pointer. The thing is: I won't
bother. This software is not crucial for me, I can live without it. If I
had a machine set up for debugging, had everything compiled with
debugging symbols and assuming that the debugger would work (which is
not given at all, for a long time gdb didn't support threads) it would
probably take me a few hours to figure out the problem in the code I see
for the first time. But I don't have this so it would take me a few days
to get to the point I can start useful debugging. So I give up without
thinking.

In real life:

-   your environment isn't comfortably set up for immediate debugging
-   getting started on any software project of decent size has a steep
    learning curve; in practice this means that no-one will do that just
    to fix a bug - too much work for too little gain

There's no argument that having the source is better than not having it
but I argue that, with a few exceptions, it doesn't change the quality
of open source software because no-one bothers to look at the source and
contribute small fixes. Just look at any open source project that isn't
Linux or Apache (those are 2 examples of very few exceptions) - they are
all desperately (and mostly in vain) looking for contributors.
