# git-fwdd

A 'store-and-forward' daemon for Git repositories.

To be a proxy between remote Git repositories and local Git repositories.

Basically, if a Git commit is pushed to a remote proxy, it goes through
git-fwdd, which forwards the push when the system is next online.

Probably a _bit_ useless, but I wanted to play around with this idea..
