#!/bin/bash

GOPATH="/home/wb/myproj/web-back-end"
#& running background
run_go_project() {
    go build -o server main.go
    ./server &
}

kill_go_project() {
    SERVER_PID=$(pgrep server)
    if [ -n "$SERVER_PID" ]; then
        kill "$SERVER_PID"
    fi
}

kill_and_exit() {
    SERVER_PID=$(pgrep server)
    if [ -n "$SERVER_PID" ]; then
        kill "$SERVER_PID"
    fi
    exit
}

#SIGINT means ctrl+c signal
trap kill_and_exit SIGINT

go build -o server main.go
./server &

while true; do
    inotifywait -r -e modify $GOPATH
    kill_go_project
    run_go_project
done
