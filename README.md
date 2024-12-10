# README

A simple C2 server and client code written in golang, that takes list of user agents from a file and communicates to the C2.

This is the result of an idea to host the C2 server in Github Codespaces, since Github is generous enough to have free version.

# Pre-requisites

1. `make` should be installed in the system
2. `go` obviously ! 

# Building the project

1. Clone this repository
2. run the following commands 

```
make clean build
```

The previous command will place the binaries inside the `/bin` directory, from there we can execute the server and client

## Starting the server

```
./bin/server -cert=/path/to/server.crt -key=/path/to/server.key -addr="HOST:PORT"
```

## Starting the client

```
./bin/client -server="HOST:PORT" -useragents=/path/to/useragents.txt
```