# git-fwdd

A 'store-and-forward' daemon for Git repositories.

This will be a proxy between local Git repositories and remote counterparts.

Basically, if a Git push is pushed to a Git remote, it will go
through - this will possibly need to be manually configured - to go
through git-fwdd, which stores it locally, and pushes it when the
system is next online.

It might be possible to automagically configure the upstream Git
remote based on the repo name and system username, concat'd with a URL
like `git@github.com`.

Probably a _bit_ useless, but I wanted to play around with this idea..

## License

Licensed under the Apache License 2.0.
