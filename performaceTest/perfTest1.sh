#!/bin/bash

# 1 Server
# 1 Client
# Msg Size = 10000 bytes
# Num of Iterations = 10, 1000, 10000, 100000, 1000000


echo "TEST1 CASE1 : Num of Server = 1, Num of Client = 1, Msg Size = 10000 bytes, Num of Iterations = 10"

#echo "THRIFT TEST"
#echo "Starting Thrift Server"
../thrift/server/server_thrift -Verbose=false -svrPort=9000 > /dev/null &
sleep 2
#echo "Starting Thrift Client"
startTime=`date +%s.%N` 
../thrift/client/client_thrift -Verbose=false -svrPort=9000 -NumOfBytes=10000 -NumOfIter=10 > /dev/null
endTime=`date +%s.%N`
runTime=$(echo "$endTime-$startTime" | bc)
echo "THRIFT TEST1 CASE1: Runtime: $runTime"
sleep 2
pkill -x server_thrift

#echo "PROTOBUF TEST"
#echo "Starting Protobuf Server"
../protobuf/server/server_protobuf -Verbose=false -svrPort=9001 > /dev/null &
sleep 2
#echo "Starting Protobuf Client"
startTime=`date +%s.%N` 
../protobuf/client/client_protobuf -Verbose=false -svrPort=9001 -NumOfBytes=10000 -NumOfIter=10 > /dev/null
endTime=`date +%s.%N`
runTime=$(echo "$endTime-$startTime" | bc)
echo "GRPC PROTOBUF TEST1 CASE1: Runtime: $runTime"
sleep 2
pkill -x server_protobuf


echo "TEST1 CASE2: Num of Server = 1, Num of Client = 1, Msg Size = 10000 bytes, Num of Iterations = 100"

#echo "THRIFT TEST"
#echo "Starting Thrift Server"
../thrift/server/server_thrift -Verbose=false -svrPort=9000 > /dev/null &
sleep 2
#echo "Starting Thrift Client"
startTime=`date +%s.%N` 
../thrift/client/client_thrift -Verbose=false -svrPort=9000 -NumOfBytes=10000 -NumOfIter=100 > /dev/null
endTime=`date +%s.%N`
runTime=$(echo "$endTime-$startTime" | bc)
echo "THRIFT TEST1 CASE2: Runtime: $runTime"
sleep 2
pkill -x server_thrift

#echo "PROTOBUF TEST"
#echo "Starting Protobuf Server"
../protobuf/server/server_protobuf -Verbose=false -svrPort=9001 > /dev/null &
sleep 2
#echo "Starting Protobuf Client"
startTime=`date +%s.%N` 
../protobuf/client/client_protobuf -Verbose=false -svrPort=9001 -NumOfBytes=10000 -NumOfIter=100 > /dev/null
endTime=`date +%s.%N`
runTime=$(echo "$endTime-$startTime" | bc)
echo "GRPC PROTOBUF TEST1 CASE2: Runtime: $runTime"
#echo "TEST2: Runtime: $runTime"
sleep 2
pkill -x server_protobuf

echo "TEST1 CASE3: Num of Server = 1, Num of Client = 1, Msg Size = 10000 bytes, Num of Iterations = 1000"

#echo "THRIFT TEST"
#echo "Starting Thrift Server"
../thrift/server/server_thrift -Verbose=false -svrPort=9000 > /dev/null &
sleep 2
#echo "Starting Thrift Client"
startTime=`date +%s.%N` 
../thrift/client/client_thrift -Verbose=false -svrPort=9000 -NumOfBytes=10000 -NumOfIter=1000 > /dev/null
endTime=`date +%s.%N`
runTime=$(echo "$endTime-$startTime" | bc)
echo "THRIFT TEST1 CASE3: Runtime: $runTime"
sleep 2
pkill -x server_thrift

#echo "PROTOBUF TEST"
#echo "Starting Protobuf Server"
../protobuf/server/server_protobuf -Verbose=false -svrPort=9001 > /dev/null &
sleep 2
#echo "Starting Protobuf Client"
startTime=`date +%s.%N` 
../protobuf/client/client_protobuf -Verbose=false -svrPort=9001 -NumOfBytes=10000 -NumOfIter=1000 > /dev/null
endTime=`date +%s.%N`
runTime=$(echo "$endTime-$startTime" | bc)
echo "GRPC PROTOBUF TEST1 CASE3: Runtime: $runTime"
sleep 2
pkill -x server_protobuf


echo "TEST1 CASE4: Num of Server = 1, Num of Client = 1, Msg Size = 10000 bytes, Num of Iterations = 10000"

#echo "THRIFT TEST"
#echo "Starting Thrift Server"
../thrift/server/server_thrift -Verbose=false -svrPort=9000 > /dev/null &
sleep 2
#echo "Starting Thrift Client"
startTime=`date +%s.%N` 
../thrift/client/client_thrift -Verbose=false -svrPort=9000 -NumOfBytes=10000 -NumOfIter=10000 > /dev/null
endTime=`date +%s.%N`
runTime=$(echo "$endTime-$startTime" | bc)
echo "THRIFT TEST1 CASE4: Runtime: $runTime"
sleep 2
pkill -x server_thrift

#echo "PROTOBUF TEST"
#echo "Starting Protobuf Server"
../protobuf/server/server_protobuf -Verbose=false -svrPort=9001 > /dev/null &
sleep 2
#echo "Starting Protobuf Client"
startTime=`date +%s.%N` 
../protobuf/client/client_protobuf -Verbose=false -svrPort=9001 -NumOfBytes=10000 -NumOfIter=10000 > /dev/null
endTime=`date +%s.%N`
runTime=$(echo "$endTime-$startTime" | bc)
echo "GRPC PROTOBUF TEST1 CASE4: Runtime: $runTime"
#echo "TEST4: Runtime: $runTime"
sleep 2
pkill -x server_protobuf


echo "TEST1 CASE5: Num of Server = 1, Num of Client = 1, Msg Size = 10000 bytes, Num of Iterations = 100000"

#echo "THRIFT TEST"
#echo "Starting Thrift Server"
../thrift/server/server_thrift -Verbose=false -svrPort=9000 > /dev/null &
sleep 2
#echo "Starting Thrift Client"
startTime=`date +%s.%N` 
../thrift/client/client_thrift -Verbose=false -svrPort=9000 -NumOfBytes=10000 -NumOfIter=100000 > /dev/null
endTime=`date +%s.%N`
runTime=$(echo "$endTime-$startTime" | bc)
echo "THRIFT TEST1 CASE5: Runtime: $runTime"
sleep 2
pkill -x server_thrift

#echo "PROTOBUF TEST"
#echo "Starting Protobuf Server"
../protobuf/server/server_protobuf -Verbose=false -svrPort=9001 > /dev/null &
sleep 2
#echo "Starting Protobuf Client"
startTime=`date +%s.%N` 
../protobuf/client/client_protobuf -Verbose=false -svrPort=9001 -NumOfBytes=10000 -NumOfIter=100000 > /dev/null
endTime=`date +%s.%N`
runTime=$(echo "$endTime-$startTime" | bc)
echo "GRPC PROTOBUF TEST1 CASE5: Runtime: $runTime"
#echo "TEST4: Runtime: $runTime"
sleep 2
pkill -x server_protobuf
