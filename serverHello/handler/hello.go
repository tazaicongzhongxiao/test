package handler

import (
	"context"
	"fmt"
	"test"
)

type Imp struct {
}

func (imp *Imp) SayHello(ctx context.Context, req test.ReqHello) (res test.ResHello, err error) {
	res.Greeting = "hello" + req.Name
	fmt.Println(res)
	return
}
