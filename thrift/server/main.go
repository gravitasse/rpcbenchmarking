package main

import (
	"flag"
	"fmt"
	"rpcbenchmarking/thrift/server/rpc"
)

func main() {
	port := flag.String("svrPort", "9000", "Server listener Port")
	flag.Parse()
	fmt.Println("Starting listener...")
	serverIface := rpc.NewTESTHandler()
	rpc.StartServer(serverIface, *port)
}
