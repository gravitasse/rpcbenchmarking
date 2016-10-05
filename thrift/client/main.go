package main

import (
	"../../../../../generated/src/testServer"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	//"time"
	"utils/ipcutils"
)

type TestClientBase struct {
	Address            string
	Transport          thrift.TTransport
	PtrProtocolFactory *thrift.TBinaryProtocolFactory
}

type TestClient struct {
	TestClientBase
	ClientHdl *testServer.TESTServicesClient
}

var testClient TestClient
var verbose bool

func connect_to_testServer(port string) {
	testClient.Address = "localhost:" + port
	testClient.Transport, testClient.PtrProtocolFactory, _ = ipcutils.CreateIPCHandles(testClient.Address)
	testClient.ClientHdl = testServer.NewTESTServicesClientFactory(testClient.Transport, testClient.PtrProtocolFactory)
}

func sendPing(msg string) {
	//fmt.Println("Client: Sending Msg:", msg)
	requestMsg := testServer.EchoRequest{
		Request: msg,
	}
	replyMsg, err := testClient.ClientHdl.Ping(&requestMsg)
	if err != nil {
		fmt.Println("Error sending msg")
	} else {
		if verbose {
			fmt.Println("Client: recv msg:", replyMsg)
		}
	}
}

func main() {
	port := flag.String("svrPort", "9000", "Server port number")
	numOfItr := flag.Int("NumOfIter", 1, "Number of Iteration")
	numOfBytes := flag.Int("NumOfBytes", 100, "Number of Bytes")
	ver := flag.Bool("Verbose", false, "Enable logging")
	flag.Parse()
	verbose = *ver
	var msg string
	for i := 0; i < *numOfBytes; i++ {
		msg = msg + "A"
	}
	connect_to_testServer(*port)
	//t := time.Now()
	for i := 0; i < *numOfItr; i++ {
		sendPing(msg)
	}
	//fmt.Println(time.Since(t))
}
