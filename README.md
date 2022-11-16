# Example of testing go code with individual mysql containers using dockertest

Golang's test command builds the code under test into separate binaries for each package.
Therefore, it is necessary to start a resource such as mysql for each binary to be built.
The test code in this project shows how to launch a mysql container for each of the packages.

## How to run tests

Run tests outside of containers

```
$ go test -v -count=1 ./...
```

Run tests inside of containers (Compose, GitHub Actions, etc..)

```
# apt update; apt install -y iproute2  # to get default gw as host nif
# export DB_HOST=$(ip route | awk '$0 ~ /default via/ { print $3 }')
# go test -v -count=1 ./...
```
