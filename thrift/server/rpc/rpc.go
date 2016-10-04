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

func (h *TESTHandler) Ping(str string) (string, error) {
	//fmt.Println("Server: Recvd Msg: ", str)
	return "Echo Reply", nil
}

type TESTHandler struct {
}

func NewTESTHandler() *TESTHandler {
	h := new(TESTHandler)
	return h
}
