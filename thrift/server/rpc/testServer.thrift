namespace go testServer

struct EchoRequest {
	1: string Request
}

struct EchoReply {
	1: string Reply
}

service TESTServices {
        EchoReply Ping(1: EchoRequest msg)
}
