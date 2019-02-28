// RPC函数的实现方式
// func (opt) MethodName(args Arg, arg *T) error
package http

import (
	"errors"
	"net/rpc"
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

func Start() {
	atirh := new(Arith)
	rpc.Register(atirh)
	rpc.HandleHTTP()
}
