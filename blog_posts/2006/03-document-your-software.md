Id: 1944
Title: Document your software
Date: 2006-03-11T18:43:03-08:00
Format: Markdown
--------------
The most important thing you can do to make your software more useful is to
write docs for it. This is especially important for Unix'y command-line tools.

Ok, maybe not **the** most important thing, but I want to make a point.

I was wasting my time reading weblogs and one link led me to [cook][1], a
better make than make. I hate make and a like trying new software (even if
ultimately I don't use it) so I wanted to check it out. I didn't because of
lack of any useful documentation. I assume the author does care about other
people using his software since:

  * he put a lot of effort into writing the code
  * he maintains a mailing list for users
  * he maintains a website for the program

At first I was hopeful, because there is a link to "Cook Reference Manual" on
the webpage but it's very unhelpful. First, it's a PDF, even though the
content is just a plain text rendering of a couple of man pages. Second, man
pages are the worst documentation idea that refuses to die.

Earth to people writing documentation: what I want to read about is "a gentle
and concise introduction to using cook" i.e. how to compile the simplest
project using cook. When I do that and it works, I might graduate to slightly
bigger examples. When I'm comfortable with that, I will need a detailed
reference guide, describing each possible option and command line switch.

A man page just slaps me in the face with its detailed reference of all
possible options, arranged in alphabetic order, with no examples of how to
achieve the most frequent tasks.

It's a cultural thing. When Unix programmer needs to solve documentation
problems, he writes a man page. Now he has two problems.

It's quite amazing:

  * the inability of intelligent people to see how awful man pages are and
break free from harmful tradition that originated in times when restrictions
of computers was forcing design decisions
  * compared to writing code, writing docs is much easier and takes less time,
but people are willing to spend infinite amounts of time on code but none on
reasonably useful docs

   [1]: http://www.canb.auug.org.au/~millerp/cook/cook.html


