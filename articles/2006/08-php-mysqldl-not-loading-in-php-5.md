Id: 1957
Title: php_mysql.dll not loading in PHP 5.1.4 and Apache 2.2
Date: 2006-08-06T22:31:07-07:00
Format: Markdown
--------------
Installing software is always fun. A recent purchase of new computer
forced to redo the setup of my test web developement environment (i.e.
Apache + PHP + MySQL) under Windows XP. Something I did at least 3 times
in the past.\
\
Murphy, as always, won, and things that could go wrong, went.\
\
But this is an educational tale, so let's get to the point.\
\
The official Windows installation for latest Apache 2.0.x is so bad that
after vanilla installation with no changes at all, Apache refused to
run. I didn't feel like debugging problems caused by a broken installer,
so I tried the latest 2.2.x installer. That worked out-of-the box
(although the Apache monitor that they start by default and stick into
tray seems pointless to me).\
\
Second problem was that the latest version of Apache (2.2.x) isn't
binary-compatible with 2.0 and the latest PHP 5.1.4 only comes with 2.0
adapter so I had to hunt down binary 2.2. Thank god for [Apache
Lounge](http://www.apachelounge.com/).\
\
Third problem was that even after setting correct extension\_dir in
php.ini and uncommenting extension=php\_mysql.dll, I was getting
mysterious error:\
\

> PHP Warning:  PHP Startup: Unable to load dynamic library
> 'C:\\\\php-514\\\\ext\\\\php\_mysql.dll' - The specified module could
> not be found.\\r\\n in Unknown on line 0\

\
It's mysterious because the file clearly is there. Lucky for me that I
misspent part of my life doing Windows programming, which includes
intricacies of DLL loading, and I know that a DLL might not get loaded
even if it exists when it refers to another DLL that is not present in
%PATH%.\
\
After some googling I found that others had the same problem and the
solution proposed was to copy libmysql.dll to c:\\windows or
c:\\windows\\system32. Polluting system directories is a gross solution
but at the time I would be happy if things just worked. However, copying
libmysql.dll to c:\\windows didn't solve the problem.\
\
In a flash of inspiration I decided to run [Dependency
Walker](http://www.dependencywalker.com/) on php\_mysql.dll to see what
else it might depend on. Turns out it also needs php5ts.dll and probably
in some cases msjava.dll (that is nowhere to be found and fortunately
doesn't seem necessary). I copied php5ts.dll to c:\\windows as well and
voila, things work.\
\
I, however, am left with a bad taste in my mouth. More on that later,
now let's recap wisdom gained:\

-   Apache logs are essential for diagnostic of Apache-related problems
-   Dependency Walker is one of those tools that you need once a year,
    but when you do, it saves you. Learn to use and learn about DLLs in
    Windows. Apparently installing PHP requires that.\

\
Now about bad taste in my mouth.\
\
Apache and PHP are amongst the most popular open-source project and yet
the Windows experience of using them is awful.\
\
Managing dependencies is a nightmare (3 versions of Apache + a dozen
versions of PHP, good luck figuring out what's compatible with what).\
\
Open-source world, in general, is not good about maintaining binary
compatibility. Apache 2.0 was supposed to be the big rewrite, what was
so important to add to Apache 2.2 to justify a 3rd, incompatible binary
interface.\
\
Forcing people to copy DLL to system folders to make things work is bad
and easily fixable. PHP could easily figure out the path of their main
DLL and add it to path so that DLLs that it knows live there could be
loaded by system.\
\
Let's not forget misleading error messages. Could not be found? But it's
there! Windows is partly to blame, because that's probably the error
that it returns to PHP, but given that it seems to be a common problem,
it could be smart enough to stat the file to see whether it's really not
there or it's not loading because of some other reason.\
\
Given those problems it's only a minor issue that Apache Foundations
seems to change their minds about where Apache is being installed about
once a year.\

