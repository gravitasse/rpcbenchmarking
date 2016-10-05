package main

import (
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
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

	protodata := new(ping.EchoRequest)
	for i := 0; i < *numOfIter; i++ {
		protodata.ReqMsg = &msg
		data, err := proto.Marshal(protodata)
		if err != nil {
			fmt.Println("Error Marshalling data")
			return
		}
		dst := "127.0.0.1:" + *svrPort
		conn, err := net.Dial("tcp", dst)
		checkError(err)
		_, err = conn.Write(data)
		checkError(err)

		recvData := make([]byte, 8192)
		n, err := conn.Read(recvData)
		checkError(err)
		protoRecvData := new(ping.EchoReply)
		err = proto.Unmarshal(recvData[0:n], protoRecvData)
		checkError(err)
		if *verbose {
			fmt.Println("Recv Data:", protoRecvData)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
