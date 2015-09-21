package main

import (
	"flag"
	"github.com/tallstreet/todographqlgo/server"
)

var listenAddr = flag.String("l", ":8080", "listen addr")

func main() {
	flag.Parse()
	server.Application = server.NewApp(*listenAddr)
	server.Application.RunServer()
}
