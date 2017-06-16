Id: 999
Title: Announcing fofou - forum software for Google App Engine
Tags: appengine
Date: 2008-07-06T11:33:13-07:00
Format: Markdown
--------------
Elevator pitch: [Fofou](/software/fofou/) is
forum software optimized for software product support/discussions. Free
to use, easy to install and free to run, thanks to [Google's App
Engine](http://code.google.com/appengine).

I believe that every software product (be it open source or commercial)
should have associated forums so that users can ask questions about the
software, suggest improvements or simply tell you how much they love
your software.

Unfortunately most forum software is bad. I call it phpBB-disease
(although it's probably UBB that is the root of all evil in this case).
The sad fact of life is that popular software gets cloned. That, by
itself, isn't bad, except in cases where the popular software is bad.
phpBB is bad but people who write their own forum software mindlessly
copy its design, just because it's popular, with only minor variations.
The end result is that we get more and more of the same awfulness.

The first reason why most forum software is bad for software support is
that it requires registration. Software forums are mostly for casual
interaction: ask one question, report one bug, suggest one feature.
Asking someone to register just to do that is stupid and a great way to
turn away a big percentage of people. I've been in that situation myself
and it makes me angry when a software vendor makes me jump through hoops
just to report a bug in their software.

Another reason, harder to describe succinctly, is that most forum
software is way too complicated. They show way too much information, for
example a date when user registered. On every page. First: who cares?
Second: I always confuse that with the post date, because it's more
visible.

They encourage creating multiple forums, which doesn't make sense in
low-frequency forums. But the functionality is right there, the
interface is optimized for showing multiple forums so hapless developer
creates a separate forum for bug reports, another one for announcements,
another one for feature requests, another for miscellaneous discussion
about the product and yet another about miscellaneous discussion about
everything else and then over months experiences what it's like to host
dead forums with nary a post per month. One forum is enough for
everything until the volume justifies creating more forums.

That's just two example out of many - most developers endlessly add more
and more features and few are willing to [Get
Real](http://gettingreal.37signals.com/toc.php) 37signals-style.

But once in a while someone comes in and shows that there is a better
way. In case of forum software it was Joel Spolsky who wrote forum
software from scratch for the needs of his software company. Even more
importantly, he [described
eloquently](http://www.joelonsoftware.com/articles/BuildingCommunitieswithSo.html)
what kind of mistakes most forum implementations make and how his forum
fix those mistakes.

The code for his implementation is not available but Wayne Venables
wrote [FruitShow](http://sourceforge.net/projects/fruitshow), a
re-implementation of Joel's ideas in PHP.

Which brings me to another reason why most forum software is bad: it's
hard to install. Even when the code is free, it's still a hassle to run
your own server, install Apache, PHP and MySQL and make sure it all runs
fine.

This is where Fofou comes in: it re-implements Joel's forums on Google
App Engine infrastructure, so not only is it easy to deploy your own
copy but it's also free to run (it's very unlikely that you'll go over
Google's free quota).

Fofou is also [an open-source project](http://github.com/kjk/fofou), so
you're free to mess with it, change it any way you want and contribute
improvements.

I believe that thanks to ease of development, ease of deployment and
generous free quotas, App Engine has a potential to revolutionize the
landscape of low-end web hosting. App Engine is plenty powerful to host
your website, blog, wiki, forums and what not, for free, with better
quality guarantees than most low-end web hosts.

It doesn't make sense to pay even $5/month for low-end web hosting and
mess with configuring web servers and databases and PHP software.

It doesn't make sense to pay $5/month for TypePad blog, if you could
install TypePad-like software on App Engine and run it for free.

It's even attractive to use App Engine instead of free (read: ad
supported) services like blogger or wordpress.org, because you get rid
of ads, you can install any blog software you like and you get more
control over the website.

It's still early in App Engine lifetime so we're not yet at a point
where this is true. The biggest missing piece is availability of
high-quality, free re-implementations of the most common types of
software people use: blogs, wikis, forums etc. Fofou is here to give you
one more choice when it comes to forum software.
