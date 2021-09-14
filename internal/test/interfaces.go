package test

import (
	"io"
	"time"

	alias "github.com/eiiches/go-gen-proxy/pkg/handler"
	"github.com/eiiches/go-gen-proxy/pkg/interceptor"
)

//go:generate go run github.com/eiiches/go-gen-proxy/cmd/go-gen-proxy --interface io.Closer --interface github.com/eiiches/go-gen-proxy/internal/test.Foo --package github.com/eiiches/go-gen-proxy/internal/test --name FooProxy --output generated.go

type StructA struct{}
type StructB struct{}
type StructC struct{}
type StructD struct{}
type StructE struct{}
type StructF struct{}
type StructG struct{}
type StructH struct{}
type StructI struct{}
type StructJ struct{}
type StructK struct{}
type StructL struct{}
type StructM struct{}

type InterfaceA interface{}
type InterfaceB interface{}

type AliasA = io.Reader

type Foo interface {
	// test arguments
	SingleArgNoReturn(format string)
	MultiArgNoReturn(string, string)
	VariadicArgNoReturn(format string, args ...string)

	// test return values
	NoArgSingleReturn() string
	NoArgMultiReturn() (string, string)

	// test types
	InterfaceValue(InterfaceA) InterfaceA
	StructValue(StructA) StructA
	AliasValue(AliasA) AliasA
	PointerValue(*StructF) *StructF
	MapValue(map[StructB]StructC) map[StructB]StructC
	SliceValue([]StructD) []StructD
	ArrayValue([1]StructE) [1]StructE
	ChanValue(chan StructG) chan StructG
	AnonymousStructValue(struct {
		StructI
		A StructH
	}) struct {
		StructI
		A StructH
	}
	AnonymousInterfaceValue(interface {
		InterfaceB
		A(StructK) StructJ
	}) interface {
		InterfaceB
		A(StructK) StructJ
	}
	AnonymousFuncValue(func(StructL) StructM) func(StructL) StructM

	// test imports
	ThisPackage(StructA)
	LocalPackageAlias(alias.InvocationHandler)
	LocalPackageNoAlias(interceptor.Interceptor)
	StandardPackage(time.Time)
}
