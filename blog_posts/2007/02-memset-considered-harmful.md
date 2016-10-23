Id: 1973
Title: memset() considered harmful
Tags: programming
Date: 2007-02-15T17:08:14-08:00
Format: Markdown
--------------
Conclusion: use memzero(addr, size) instead of memset(addr, 0, size) to
avoid making frequent memset() mistake of giving parameters in the wrong
order.\
\
memset() is another example where following established convention leads
people to doing the wrong thing. One of the most frequent uses for
memset() is to zero-out a piece of memory. The problem with memset() is
that it's easy to swap the "value to set" with "count of values to set"
parameters. They're both ints, so the compiler won't complain and humans
are not that good at remembering things of that nature. Combined with
laziness, another human affliction, it leads to many cases where
memset() is misused and ends-up being a no-op (because 0, the value, is
passed as size, so it ends up doing nothing). I've seen this in my own
code, I've seen this in other people's code and you can see for yourself
how many people got it wrong with [this
query](http://www.google.com/codesearch?hl=en&lr=&q=memset%5Cs*%5C%28.%2B%2C%5Cs*0%5Cs*%5C%29&btnG=Search).\
\
And a solution is so simple: write a trivial wrapper memzero(void \*s,
int size). Not only is it a better name but removes possibility of this
particular mistake to ever happen again. On Windows you don't even have
to write it since it's there as ZeroMemory(). Some unixes have bzero(),
but it's an ugly name and I don't think it's widely available so writing
memzero() utility function is a good idea anyway.\
\

