package json

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func Client() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(client)
	args := Args{17, 8}
	var res int
	err = client.Call("Arith.Multiply", args, &res)
	if err != nil {
		fmt.Println("Call error")
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, res)
	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Printf("Divied %d/%d=%d..%d", args.A, args.B, quo.Quo, quo.Rem)
}
