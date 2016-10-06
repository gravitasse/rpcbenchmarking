package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	"rpcbenchmarking/protobuf/ping"
)

func main() {
	svrPort := flag.String("svrPort", "9500", "Server Port Number")
	numOfIter := flag.Int("NumOfIter", 1, "Number of Iteration")
	numOfBytes := flag.Int("NumOfBytes", 100, "Number of Bytes")
	verbose := flag.Bool("Verbose", false, "Verbosity")

	flag.Parse()

	var msg string
	for i := 0; i < *numOfBytes; i++ {
		msg += "A"
	}

	dst := "127.0.0.1:" + *svrPort
	conn, err := grpc.Dial(dst, grpc.WithInsecure())
	checkError(err)
	defer conn.Close()

	c := ping.NewPingClient(conn)

	for i := 0; i < *numOfIter; i++ {
		r, err := c.SendEcho(context.Background(), &ping.EchoRequest{ReqMsg: &msg})
		checkError(err)
		if *verbose {
			fmt.Println("Recv Message:", *r.RepMsg)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
