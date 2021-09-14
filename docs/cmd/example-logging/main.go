package main

import (
	"fmt"

	"github.com/eiiches/go-gen-proxy/examples"
	"github.com/eiiches/go-gen-proxy/pkg/interceptor"
	"github.com/foo/bar/pkg/greeter"
)

type GreeterImpl struct{}

func (this *GreeterImpl) SayHello(name string) (string, error) {
	return fmt.Sprintf("Hello %s!", name), nil
}

func main() {
	p := &greeter.GreeterProxy{
		Handler: &interceptor.InterceptingInvocationHandler{
			Delegate:    &GreeterImpl{},
			Interceptor: &examples.LoggingInterceptor{},
		},
	}
	msg, err := p.SayHello("James")
	fmt.Println(msg, err)
}
