Id: 1650
Title: Why consistency is important in software design
Date: 2003-06-24T18:24:36-07:00
Format: Markdown
--------------
I have a problem: none of the available file managers really meet my
needs. There's always something that bothers me either because a feature
is missing or (I think) could work better. That's why I'm constantly
looking for something better. I've tried many file managers: Salamander,
Total Commander, [Frigate](http://www.frigate3.com/) (and a few others
that I cannot recall so easily). Out of those I tried Frigate is the
best but I keep looking and that's why recently I tried another one:
Directory Opus. I remember Directory Opus from my Amiga days: it was THE
file manager on that platform, packed with features and nice to use.
Amiga went down and took the whole market for Amiga software with it but
it looks like Directory Opus folks averted the disaster by porting their
software to Windows. This post is about how they lost a potential
costumer (me) because of lack of consistency with Windows.

File managers have long history, dating back to DOS days and Norton
Commander who ruled that market (and was cloned multiple times). On
Windows we have two general styles of file management:

-   one that works just like Explorer (built into Windows)
-   second that works more or less like Norton Commander

By now we all expect some things to work a certain way: F5 is copy, F6
is move, F2 is rename, Ins selects a single file, you have two panes,
you can drag-and-drop files between them using a mouse etc. There's no
particular reason to use F2 instead of Ctrl-C except that everyone
excepts Ctrl-C to mean "Copy from a clipboard".

Every file managers seems to implement the golden standard set by Norton
Commander. That is: every except Directory Opus. You might think that
this is a small thing but it's not. I was quickly infuriated by things
not working the way I was expected them to behave. Among other things:

-   there's not right-click-drag-and-drop (which usually allows me to
    choose move/copy/create shortcut when I perform drop)
-   F5, F6 (and probably many others) aren't mapped to the same
    operations as in NC
-   Ins doesn't select/deselect a file (in fact, there doesn't seem to
    be possible to do it not using the mouse which renders keyboard
    navigation largely useless)

Customization is possible so I was able to fix at least some of the
keyboard mappings but I quickly gave up on Directory Opus even though it
seems to be a very solid product, has all the basic capabilities and
maybe even has features not available anywhere else.

The moral of the story: don't be gratuitously different than anyone
else. Be as consistent with the platform as possible. There are cases
where you simply find a better way of doing things and that justifies
being different, but keyboard mappings aren't the place to innovate.
