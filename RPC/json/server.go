package json

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}
type Arith int

func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func JsonStart() {
	arith := new(Arith)
	_ = rpc.Register(arith)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":9090")
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("error")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
