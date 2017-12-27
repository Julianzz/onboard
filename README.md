## Onboard Coding 

The task is to implement a basic RESTful HTTP service in Go for a simplified Tantan backend:
Adding users and swiping other people in order to find a match.

## Setup Env

* install brew on mac, please following the link https://brew.sh/
* brew upgrade && brew update && brew install golang 
* brew install postgres
* set GOPATH variable to go local path, liking "export GOPATH=~/gopath"
 

## Build & Run

* clone code into golang path path: $GOPATH/src/github.com/p1cn/onboard/liuzhenzhong
* start postgres
* change  postgres database config: conf/config.yml
* create table: go run cmd/database/initdb.go
* start server: go run cmd/server/onboard.go -port=":8090" 

