package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"rpcx-test/service"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {

	//client.Apple("google", "apple")

	flag.Parse()
	s := server.NewServer()
	//s.Register(new(service.Arith), "")
	s.RegisterName("Arith", new(service.Arith), "")
	s.Serve("tcp", *addr)
	//s := server.NewServer()
	//s.RegisterName("Apple", new(client.Apple), "")
	//s.Serve("tcp", addr)

}
