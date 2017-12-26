## Onboard Coding 

The task is to implement a basic RESTful HTTP service in Go for a simplified Tantan backend:
Adding users and swiping other people in order to find a match.

## Setup Env

* golang  go1.9.2 darwin/amd64
* install brew on mac, please follow the link https://brew.sh/
* brew upgrade && brew update && brew install golang 
* brew install postgres
* set GOPATH variable to go local path, export GOPATH=~/gopath
 

## Build & Run

* copy the code into the path: $GOPATH/src/github.com/p1cn/onboard/liuzhenzhong
* start postgres
* change  postgres database config in conf/config.yml
* init database go run cmd/database/initdb.go
* start server go run cmd/server/onboard.go -port=":8090" 

