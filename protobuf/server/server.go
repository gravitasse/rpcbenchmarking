package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"os"
	"rpcbenchmarking/protobuf/ping"
)

type server struct{}

var ver bool

func (s *server) SendEcho(ctx context.Context, in *ping.EchoRequest) (*ping.EchoReply, error) {
	//fmt.Println("Server Recvd Msg: ", *in.ReqMsg)
	//msg := *in.ReqMsg + "Reply"
	return &ping.EchoReply{RepMsg: in.ReqMsg}, nil
	//return &ping.EchoReply{RepMsg: &msg}, nil
}

func main() {
	//fmt.Println("Starting ProtoBuf Server")
	verbosity := flag.Bool("Verbose", false, "Verbosity")
	svrPort := flag.String("svrPort", "9500", "Server Port Number")
	flag.Parse()

	ver = *verbosity

	addr := "127.0.0.1:" + *svrPort
	listener, err := net.Listen("tcp", addr)
	checkError(err)
	s := grpc.NewServer()
	ping.RegisterPingServer(s, &server{})
	s.Serve(listener)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
