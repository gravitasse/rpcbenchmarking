package rpc

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"testServer"
)

func StartServer(handler *TESTHandler, port string) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	//serverTransport, err := thrift.NewTServerSocket("localhost:" + strconv.Itoa(9000))
	serverTransport, err := thrift.NewTServerSocket("localhost:" + port)
	if err != nil {
		fmt.Println("StartServer: NewTServerSocket failed with error:", err)
		return
	}
	processor := testServer.NewTESTServicesProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	err = server.Serve()
	if err != nil {
		fmt.Println("Failed to start the listener, err:", err)
	}
	fmt.Println("Start the listener successfully")
	return
}

func (h *TESTHandler) Ping(msg *testServer.EchoRequest) (*testServer.EchoReply, error) {
	if h.Verbose {
		fmt.Println("Server: Recvd Msg: ", msg.Request)
	}
	var replyMsg testServer.EchoReply
	replyMsg.Reply = msg.Request
	return &replyMsg, nil
}

type TESTHandler struct {
	Verbose bool
}

func NewTESTHandler(ver bool) *TESTHandler {
	h := new(TESTHandler)
	h.Verbose = ver
	return h
}
