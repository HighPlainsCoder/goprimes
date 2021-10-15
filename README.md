# go primes


## This is a program that generates primes

We dont need to generate primes, but its a tool to learn go,
And, maybe, a tool for BDD, or Testing, or other things.
                                           
### Evolution

v1: the simplest. Start at 3, divide by every number possible

v2: add signals so I can catch an interrupt, and goroutine, so I can loop until I do.

v2.1: cleaned up the signalling; refactored how the select is coded; added proper wait for shutdown.
I also added some sleeps around, so the program runs slowly enough to watch with my old human eyes.
                             
v2.9: I pulled the code out into functions, and started the first tiny improvement to getNextCandidate,
which only needs to check the odd numbers after 2.  I also moved it into Intj, where it works fine.

                                                                                                
## The go features used
* goroutines
* channels
* signal
* time
* fmt
* global variables
