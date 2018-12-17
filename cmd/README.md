dp command
==========

Helper command to work with Digital Publishing services

### Getting started

1. Add `DP_SSH_USER` and `DP_CONFIG` to your bash/zsh profile
2. `go get github.com/ONSdigital/dp/cmd/dp`
3. `dp`

### Known issues

* `remote` commands (`allow` and `deny`) get current IP from ipify, but this means a `deny`
    will fail if the IP has changed since running `allow` (for SGs, ACLs use rule numbers)