package service

import (
	"context"
	"fmt"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith)Apple(ctx context.Context, args *Args, replay *Reply) error {
	fmt.Println(args.A, args.B)
	replay.C = args.B + args.A
	return nil
}