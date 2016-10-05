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
	//fmt.Println("Starting ProtoBuf Server")
	verbosity := flag.Bool("Verbose", false, "Verbosity")
	svrPort := flag.String("SvrPort", "9500", "Server Port Number")
	flag.Parse()

	addr := "127.0.0.1:" + *svrPort
	listener, err := net.Listen("tcp", addr)
	checkError(err)
	for {
		if conn, err := listener.Accept(); err == nil {
			go handleProtoClient(conn, *verbosity)
		} else {
			continue
		}
	}
}

func handleProtoClient(conn net.Conn, verbosity bool) {
	//fmt.Println("Connection Established")
	defer conn.Close()

	data := make([]byte, 8192)
	n, err := conn.Read(data)
	checkError(err)
	//fmt.Println("Decoding protobuf message")

	protodata := new(ping.EchoRequest)
	err = proto.Unmarshal(data[0:n], protodata)
	checkError(err)
	if verbosity {
		fmt.Println("Recv Data:", protodata)
	}
	//c <- protodata

	echoReply := new(ping.EchoReply)
	msg := protodata.ReqMsg
	echoReply.RepMsg = msg
	data, err = proto.Marshal(echoReply)
	checkError(err)
	_, err = conn.Write(data)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
