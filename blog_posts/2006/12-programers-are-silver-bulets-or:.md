Id: 1971
Title: Programmers are silver bullets (or: after all this years, C still kicks Python ass)
Date: 2006-12-07T17:12:51-08:00
Format: Markdown
--------------
Another day,
[another](http://www.knowing.net/PermaLink,guid,a3319866-bcd1-4f0a-bce6-c7138e16b240.aspx)
[debate](http://wesnerm.blogs.com/net_undocumented/2006/12/lego_programmin.html)
about the existence (or lack of it) of [silver
bullets](http://www.joelonsoftware.com/items/2006/12/05.html) in
programming.\
\
The topic is: are there improvements in programming technology (i.e.
silver bullets) that can account for an order of magnitude improvements
in programmers productivity.\
\
In other words: if a programmer learned a new technology, could he write
an average program an order of magnitude faster?\
\
One technology that is peddled as a silver bullet are higher-level
languages, where C/C++ is today considered a low level language, Java a
bit more high level and Python even higher level. It seems to be a 
no-brainer to claim that you can write code much faster in Python than
in Java and in Java faster than in C/C++. And often it is true and there
even are some studies showing that using small-scale programming
exercises.\
\
How well does this theory of Python supreme productivity over C/C++
plays out in the real world?\
\
Today BitTorrent, Inc. announced acquisition or [uTorrent
BitTorrent](http://forum.utorrent.com/viewtopic.php?id=17279) client.
BitTorrent protocol wasw invented by Bram Cohen (now of BitTorrent,
Inc.) and implemented in Python, the uber productive programming
language. On Windows, it used UI based on Gtk and later wxWindows. In
both cases the executable was large, using a lot of memory and UI was
ugly. For example Windows installer for latest 5.0.3 version is 5.9
megabytes.\
\
BitTorrent is also open-source, so at  least in theory it should benefit
from the open-source magic fairy, contributing improvements and bug
fixes.\
\
When BitTorrent, the protocol, became a runaway success, other clients
were developed. Those that are popular today are Azeurus (written in
Java), BitComet (written in C++) and uTorrent. It's unclear what
language was used to develop uTorrent, but given the (packed) exe is 170
kilobytes, it's a safe bet that it was written in C or C++, the assembly
language of 21st century.\
\
If you use BitTorrent today, you'll see that almost no-one uses the
original BitTorrent client. The people have spoken and switched to
Azeurus, BitComet or uTorrent and uTorrent acquisition by BitTorrent,
Inc. makes it very clear.\
\
This is the best study of programmer productivity you can come up with:
the same basic idea (BitTorrent client) implemented by multiple teams
using different technologies.\
\
Free market has decided that the winner of this race is uTorrent,
Azeurus and BitComet. The looser is BitTorrent.\
\
Is uTorrent an order of magnitude better than BitTorrent? If you compare
functionality, yes in some respects (resource usage) but in general, no
(both will do the job of downloading BitTorrent files just fine). But
according to only metric that matters i.e. popularity, uTorrent is
orders of magnitude better because software tends to have strong network
effects i.e. "winner takes all".\
\
Let's recap. BitTorrent:\

-   written in high-level language (Python)
-   a huge (years) head-start on other clients
-   benefiting from open-source magic fairy\
-   a resource hog
-   lost popularity battle

uTorrent:\

-   written in low-level language (C or C++, most likely)
-   tiny and uses little resources
-   wins popularity battle

Even if we take into account the fact that uTorrent had it much easier
(it had BitTorrent Python code to study), BitTorrent had a huge
head-start and should still improve at a much faster rate.\
\
But the truth is that a good, dedicated programmer using C/C++ sword
will write better apps than another programmer using high-level language
like Python.\
\
Another truth is that despite all those "performance doesn't matter"
claims, Python and Ruby and Java are not viable languages for shipping
high-quality Windows apps (C\# is getting there). Few counter-examples
like Azeurus or Eclipse do not a revolution make. Also, the fact that
those languages are successfully used in other domains like web apps
doesn't invalidate the point about client applications. Show me a
spreadsheet close to Excel in speed and features written in a high-level
language.\
\
The only thing that (occasionally) can deliver order of magnitude
productivity improvements are humans.\
\
Human being is the only silver bullet so far.\
\

