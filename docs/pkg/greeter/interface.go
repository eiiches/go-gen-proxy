package greeter

type Greeter interface {
	SayHello(name string) (string, error)
}
