# git-fwdd

A 'store-and-forward' daemon for Git repositories.

This will be a proxy between local Git repositories and remote counterparts.

Basically, if a Git push is pushed to a Git remote, it will go through - this
will have to be manually configured - to go through git-fwdd, which stores it
locally, and pushes it when the system is next online.

Probably a _bit_ useless, but I wanted to play around with this idea..
