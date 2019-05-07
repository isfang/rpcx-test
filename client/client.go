package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx-test/service"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {

	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xcient := client.NewXClient("Arith",
		client.Failtry,
		client.RandomSelect,
		d,
		client.DefaultOption)
	defer xcient.Close()

	args := service.Args{
		A: 100,
		B: 200,
	}

	for {
		reply := &service.Reply{}
		err := xcient.Call(context.Background(), "Apple", args, reply)
		if err != nil {
			fmt.Println("~~")
			log.Fatal(err)
		}

		log.Println(args.A, args.B, reply.C)

		//time.Sleep(1e2)
	}
}
