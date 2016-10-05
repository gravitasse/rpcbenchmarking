package main

import (
	"flag"
	"fmt"
	"rpcbenchmarking/thrift/server/rpc"
)

func main() {
	port := flag.String("svrPort", "9000", "Server listener Port")
	ver := flag.Bool("Verbose", false, "Verbosity")
	flag.Parse()
	fmt.Println("Starting listener...")
	serverIface := rpc.NewTESTHandler(*ver)
	rpc.StartServer(serverIface, *port)
}
