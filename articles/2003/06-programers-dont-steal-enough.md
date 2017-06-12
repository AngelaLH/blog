Id: 1654
Title: Programmers don't steal enough
Date: 2003-06-30T14:39:38-07:00
Format: Markdown
--------------
Now that I have your attention: what I really mean is that software
designer do not take advantage of the Giants as much as they could.

It's sad that the act of building on the works of others, a necessity
for progress, is labeled as stealing by those who are blinded by dreams
of easy money so much that they fail to recognize that everyone,
including themselves, is better off in a world where we can freely build
on ideas of others. Unfortunately those people have enough money, bad
taste and political savvy to force upon us copyright terms of insane
lengths and damaging patent policies.

Wise people recognized long ago that we all owe a big debt to those who
came before us for their ideas and whatever our contribution to the pool
of knowledge might be, it'll be miniscule compared to collective
knowledge gathered so far that we were able to build upon.

> *If I have seen further it is by standing on the shoulders of Giants.\
>  - [Isaac Newton](http://www.warble.com/jherbert/giants.html)*

It's scary to think what will happen if patent system will degenerate to
the point when it'll allow monopolizing fairly obvious yet crucial
ideas. What would happen if someone secured a patent for text editing
system, making it impossible to create a competing text editor. What
would happen if someone secured a patent for client-server communication
making it impossible to create web (or one of hundreds useful
applications)? Everybody would suffer because the whole computer
industry would freeze. One lucky soul could charge monopolistic prices
for the only implementation of text editing idea but it would lack the
incentive to improve it. Both high-prices and lack of features would
lead to meager adoption. The vendor might think that he's in heaven but
everyone else would be much worse off than in an alternative universe
where thriving competition drives the prices down and quickly improves
the software making it useful and affordable for more people, indirectly
stimulating hardware purchases which in turn stimulates creating more
kinds of software. Ironically, the vendor would also end up with smaller
revenues than in the alternative universe because free competition would
create a market much bigger and one lucky winner that would dominate
this market would earn profits unheard of in any other kind of
enterprise.

So let's take advantage of the fact that we're not yet in a world like
that. Thousand flowers can still bloom and your mission is to take
existing flowers and make them slightly better. Let's rip off each
other's ideas, let's build on each other's work, let's compete - that's
the way to make users happy. I claim that we don't do it nearly enough.
You would think that because of fierce competition good ideas in
software would immediately get replicated but my observations rather
lead me to nod with understanding when I'm reminded of those words:

> *Don't worry about people stealing your ideas. If your ideas are any
> good, you'll have to ram them down people's throats\
>  - [Howard Aiken](http://www.creativityinmotion.org/quotes.html)*

Below are a few examples of neat software ideas that I've seen in at
least once and would like everyone else to copy in their products.

1\. Emacs has a great implementation of the "intelligent auto-completion"
feature. While it's a standard feature in many (although not all) IDEs
those days it usually only works for a specific kinds of text e.g. it
might auto-complete C/Java function/method names and parameters and it
requires pre-population of the auto-completion data. Emacs will build
the list of suggestions by parsing the current text so it learns from
what you've already typed. Works wonders - I became depended on this
feature a few hours after discovering it. Now every editor that doesn't
have it is painful to use. Please, o unknown programmer, steal this
idea.

2\. There are 2 popular ways to implement spell-checking:

-   user has to invoke a menu item, gets a dialog box and is forced to
    sequentially go through every misspelled word in the document and
    decide whether it should be replaced, ignored, added to dictionary
    etc. Gets annoying if you spell-check the same  part multiple times
    and have to ignore the same words over and over again.
-   pointing out misspelled words directly on the screen by e.g. putting
    read squiggly lines underneath and letting user to correct spelling
    using mouse-activated context-menus

The second method has been with us for quite some time and is obviously
better that the first one yet there are still programs that do it the
hard-way (e.g. Macromedia Dreamweaver).

3\. There are more than 2 ways to implement searching in a text editor.
One popular way is to have a permanent Find dialog box with a field to
enter text we're searching for and a button for Find Next. The dialog
box will sometimes obscure the text we're working on requiring us to
move it around. Some pervert designers also force us to use mouse in
order to switch the search directions (i.e. either forward or backward)
by making this option a drop-box or a radio button. Why this design is
still in use while there is a superior solution: pop-up the dialog only
for entering the text to search, pressing enter dismisses the dialog and
navigates to first match. We also have key strokes for Find Previous and
Find Next actions (I like to map them to F3/F4 keys). That way there's
no need to use the mouse, no need to take your hand away from the
keyboard - that's important for all the touch-typists out there.

Additionally Emacs has a small improvement of this idiom: it highlights
all the matches in the currently visible portion of the document.

4\. Bookmarks in a text editor are quite useful when writing code. For
example imagine that you just wrote the .h file with function
declarations and you're writing the implementation in a corresponding .c
file. I usually find myself unable to remember the exact details of what
is the name of the parameter x or what exactly is the name of a function
etc. so I have to frequently switch back-and-forth between the files to
refresh my memory. If that's only two files, you can usually get by with
other methods (just switching between buffers in Emacs or arranging the
display to see both files at the same time) but I found that if the
feature is well implemented then I prefer to use it even in two files
case and the more places you need to bookmark the more valuable it
becomes. But here's the tricky part: "well implemented". It's crucial
that navigating to bookmarked places is quick. I've seen only one editor
that implements it well: Source Insight. With a single keystroke
(Ctrl-M) you can both set a bookmark and navigate to a bookmarked place
(don't just believe me, download free evaluation version of Source
Insight and play with it). Sadly, Emacs implements bookmarks poorly -
key sequences are too long that managing bookmarks is more trouble than
it's worth it. So if you implement a text editor, here's an idea for
you: steal Source Insight's bookmarks management. It's excellent.

5\. People who use third party file managers (instead of built-in
Explorer) are usually power users. They might need to drop into command
line to do things that are not easy to do from GUI like grep through
files, launch compilation etc. I first saw this feature in Total
Commander: there's additional text edit field at the bottom of the
window. If you type something there and press enter, it'll be executed
as a command in the directory that is shown in the active pane. It's a
great feature that saved me a lot of time (as compared to explicitly
launching cmd window and mindlessly navigating to the directory that I
already have selected in my file manager). Yet not every file manager
has it. The result is that I won't be using those file managers: once
I've tasted the usefulness of it, I refuse to use a tool that doesn't
give it to me.

Maybe we should have a collection of such "best design for doing foo"
somewhere on the web. Then if, for example, someone created a new text
editor and it lacked some of the useful features or implemented them in
a way inferior to existing implementation, we could send him a link to
this page and say: dude, we're not gonna use your software until you
implement all those features. Please, do implement all the great ideas
you have that make your text editor better than others, but this is your
baseline. Anything worse than that is not gonna cut it.
