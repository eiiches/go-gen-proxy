package test

import (
	"testing"
)

func test(t *testing.T) {
	sut := &FooProxy{
		Handler: nil,
	}
	sut.VariadicArgNoReturn("foo", "a", "b")
}
