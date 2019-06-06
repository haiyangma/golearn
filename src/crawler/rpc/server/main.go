package main

import (
	"crawler/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdaemon.DemoService{})

	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error : %v", err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
