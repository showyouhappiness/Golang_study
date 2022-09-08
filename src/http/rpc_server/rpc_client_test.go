package rpc_server

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
)

const serverAddress = "localhost"

func TestRpcClient(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// Synchronous call
	args := Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
