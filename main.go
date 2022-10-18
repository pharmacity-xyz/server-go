package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}

func main() {
	timeServer := new(TimeServer)
	err := rpc.Register(timeServer)
	if err != nil {
		log.Fatal("error")
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	errorS := http.Serve(l, nil)
	if errorS != nil {
		log.Fatal("error")
	}
}
