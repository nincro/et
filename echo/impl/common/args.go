package common

import "errors"

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

func (a *Arith)Divide(args *Args, reply *int) error {
	if args.B == 0 {
		return errors.New("divided by zero!")
	}
	return nil;
}