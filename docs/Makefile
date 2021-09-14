cmds := $(wildcard cmd/*)
bins := $(patsubst cmd/%,%,$(cmds))

.PHONY: all
all: $(bins) pkg/greeter/greeter_proxy.go

%: cmd/% pkg/greeter/greeter_proxy.go
	go build -o $@ ./$<

.PHONY: clean
clean:
	$(RM) $(bins)

pkg/greeter/greeter_proxy.go: pkg/greeter/interface.go
	go run github.com/eiiches/go-gen-proxy/cmd/go-gen-proxy --interface github.com/foo/bar/pkg/greeter.Greeter --package github.com/foo/bar/pkg/greeter --name GreeterProxy --output pkg/greeter/greeter_proxy.go
