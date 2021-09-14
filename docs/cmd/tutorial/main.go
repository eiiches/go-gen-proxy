package main

import (
	"fmt"

	"github.com/foo/bar/pkg/greeter"
)

type handler struct{}

func (*handler) Invoke(method string, args []interface{}) []interface{} {
	switch method {
	case "SayHello":
		return []interface{}{fmt.Sprintf("Hello %s!", args[0]), nil}
	default:
		panic("not implemented")
	}
}

func main() {
	p := &greeter.GreeterProxy{
		Handler: &handler{},
	}
	msg, err := p.SayHello("James")
	fmt.Println(msg, err)
}
