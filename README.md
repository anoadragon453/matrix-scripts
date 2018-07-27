# Matrix Scripts

A collection of handy scripts for developing on matrix.

## DummyAS

A python script acting as a pseudo-application service that homeservers can be
tested against.

## Create And Login

Quickly register an account on a homeserver. The script prints the user details
after it finishes, as well as saves the details as exported environment
variable. Make sure to `source` the script if you plan on using this
functionality.

## Wipe Dendrite DB

A script that drops and recreates all postgres dendrite databases.

## Benchmarks

A collection of scripts for benchmarking various bits and pieces of
functionality. Homeserver-agnostic. Run with `go run ./scriptname.go
portnumber`, where portnumber is the HTTP listening port for the homeserver. It
is currently assumed the homeserver is running on localhost.
