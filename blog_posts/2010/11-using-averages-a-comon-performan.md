Id: 340001
Title: Using averages - a common performance measurement mistake
Tags: programming
Date: 2010-11-24T22:06:05-08:00
Format: Markdown
--------------
**Summary:** when people want to get a more accurate result when
benchmarking a performance of given piece of code they often run the
same test multiple times and use an average as a final benchmark result.
It’s a mistake: using the time of the fastest run is more accurate.

By definition, a performance test must be deterministic: given the same
inputs it’ll execute exactly the same number of machine instructions,
read and write the same amount of data to/from disk etc. If it isn’t
deterministic, benchmarking it is pointless.

We all know, however, that execution time is not deterministic. Why is
that?

Multi-tasking nature of the operating system is to blame. Your code is
only one of the many processes that compete for fixed resources like cpu
time and i/o bandwidth. Operating system will interrupt your program and
start executing some other code, at random times and for unpredictable
amount of time.

If you measure just the time of execution, like most benchmarking
methods do, it’ll include not only the execution time of your code but
also of other programs executed by operating system during that
particular test run.

The same applies to other shared resources like a hard-drive: the
benchmarked program asks the OS to read a piece of data from disk but so
do other programs. The OS decides who gets to do I/O first in an
unpredictable and unaccountable way.

You have little control over that behavior. You can think of execution
time of your test as consisting of 2 components:

-   the deterministic execution time of only your code, which would
    happen if your process had exclusive access to all resources like
    cpu and hard-drive
-   random, unpredictable execution time of all other programs that OS
    decided to run during that particular test run

In other words, the benchmarked time is: the time your’re interested
in + a random overhead attributable to other processes.

This model explains why you should use the time of the fastest test run:
it’s the best approximation of the running time that is attributable
only to your code (and with the smallest part attributable to other
random processes).

Other coping mechanisms when doing performance tests involve trying to
minimize the random component by shutting down as many processes as
possible, so that less things will get in the way.
