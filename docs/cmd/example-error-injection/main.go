package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiiches/go-gen-proxy/examples"
	"github.com/eiiches/go-gen-proxy/pkg/interceptor"
	"github.com/foo/bar/pkg/greeter"
)

type GreeterImpl struct{}

func (this *GreeterImpl) SayHello(name string) (string, error) {
	return fmt.Sprintf("Hello %s!", name), nil
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := &greeter.GreeterProxy{
		Handler: &interceptor.InterceptingInvocationHandler{
			Delegate: &GreeterImpl{},
			Interceptor: &examples.ErrorInjectingInterceptor{
				Random:             r,
				FailureProbability: 0.5,
			},
		},
	}
	msg, err := p.SayHello("James")
	fmt.Println(msg, err)
}
