## Onboard Coding 

The task is to implement a basic RESTful HTTP service in Go for a simplified Tantan backend:
Adding users and swiping other people in order to find a match.

## Setup Env

* install brew on mac, please following the link https://brew.sh/
        
        brew upgrade && brew update && brew install golang 


* set GOPATH variable to go local path, liking "export GOPATH=~/gopath"

        export GOPATH=~/gopath
        mkdir -p $GOPATH

* install and start postgres
    
        brew install postgres
        export DATAPATH=/tmp/postdata
        mkdir  -p $DATAPATH
        initdb -D $DATAPATH
        pg_ctl -D $DATAPATH start



## Build & Run

* clone code into golang path path: 

        export ONBOARD_PATH="$GOPATH/src/github.com/p1cn/onboard"
        mkdir -p $ONBOARD_PATH
        cd $ONBOARD_PATH
        git clone https://github.com/Julianzz/onboard liuzhenzhong
        

* change  postgres database config: conf/config.yml

        db_settings:
            host: "127.0.0.1:5432"
            database: "test"
            user: "liuzhenzhong"
            password: ""

* create table: 
    
        go run cmd/database/initdb.go

* start server:

        go run cmd/server/onboard.go -port=":8080" 


* test visit:

        pip install requests
        cd $ONBOARD_PATH/liuzhenzhong/scripts
        python test.py