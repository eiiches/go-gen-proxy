package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/eiiches/go-gen-proxy/internal/genproxy"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "go-gen-proxy",
		Usage: "Generate proxy implementations for an interface",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "interface",
				Usage:    "Interface a generated struct should implement. Use PACKAGE_PATH.INTERFACE_NAME format. If you need to implement multiple interfaces, repeat this option. e.g. io.Reader, github.com/foo/bar.Foo",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "package",
				Usage:    "Package path to where generated struct should reside. This affects import directives of the generated file. e.g. github.com/foo/bar",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "package-name",
				Usage:    "Usually package name is deduced from taking suffix from --package argument. If you need to override the behavior use this option. e.g. foo",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "name",
				Usage:    "Name of a struct to generate. Recommended value is: <Interface>Proxy where <Interface> is the name of the interface to be implemented. e.g. FooProxy",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Usage:    "Write output to a file. Defaults to write to stdout.",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			interfaces := c.StringSlice("interface")

			name := c.String("name")

			pkgPath := c.String("package")

			pkgName := c.String("package-name")
			if pkgName == "" {
				pos := strings.LastIndex(pkgPath, "/")
				if pos < 0 {
					panic("cannot deduce package name. fix --package or use --package-name to override the package name.")
				}
				pkgName = pkgPath[pos+1:]
			}

			code, err := genproxy.Generate(pkgPath, pkgName, name, interfaces)
			if err != nil {
				panic(err)
			}

			output := c.String("output")
			if output == "" {
				fmt.Println(code)
			} else {
				os.WriteFile(output, []byte(code), 0644)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
