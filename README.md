# rpcbenchmarking

This repository contains a message ping-pong Client Server Implementation using 

1. Apache Thrift version 0.9.3 (https://thrift.apache.org/)
2. gRPC (http://www.grpc.io/) and protobuf-3

Programming language: Golang

This repository also contains some performance tests. 

- **Test 1** : Keep the message size fixed and vary the number of times message is sent across. 

    - **Number of Server**: 1
    - **Number of Clients**: 1
    - **Message Size** = 10000 bytes
    - **Result**:
    
        |          Number of Iterations         |      10     |     100     |     1000    |    10000    |    100000    |
        |:-------------------------------------:|:-----------:|:-----------:|:-----------:|:-----------:|:------------:|
        |       Time taken (Thrift) in sec      | 0.052241493 | 0.064553807 | 0.228196454 | 1.229499966 | 11.433889190 |
        | Time taken (gRPC & Protobuf) in secs  | 0.085299737 | 0.122767140 | 0.371312381 | 3.350658766 | 32.749954770 |
           

- **Test 2** : Keep the number of times message is sent across fixed and vary the message size.

    - **Number of Server**: 1
    - **Number of Clients**: 1
    - **Number of Iteration** = 1000
    - **Result**:
    
        |            Number of Bytes            |      10     |     100     |     1000    |    10000    |    100000   |
        |:-------------------------------------:|:-----------:|:-----------:|:-----------:|:-----------:|:-----------:|
        |       Time taken (Thrift) in sec      | 0.082524018 | 0.103245114 | 0.070778594 | 0.191037682 | 2.232637674 |
        | Time taken (gRPC & Protobuf) in secs  | 0.257207749 | 0.307096547 | 0.262112270 | 0.455775537 | 3.531310872 |
