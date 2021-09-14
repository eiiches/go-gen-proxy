package main

import (
	"fmt"

	"github.com/foo/bar/pkg/greeter"
	"github.com/stretchr/testify/mock"
)

type MockHandler struct {
	mock.Mock
}

func (this *MockHandler) Invoke(method string, args []interface{}) []interface{} {
	return this.Mock.MethodCalled(method, args...)
}

func main() {
	mock := &MockHandler{}
	p := &greeter.GreeterProxy{
		Handler: mock,
	}
	mock.On("SayHello", "James").Return("Hello James!", nil)

	msg, err := p.SayHello("James")
	fmt.Println(msg, err)
}
