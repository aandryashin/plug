package main

import (
	"fmt"
	"github.com/aandryashin/plug/so"
	"os"
	"plugin"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n\t%s file1.so file2.so...\n", os.Args[0])
	}
	for _, f := range os.Args[1:] {
		p, err := plugin.Open(f)
		if err != nil {
			fmt.Printf("%s: load %s: %v\n", os.Args[0], f, err)
			os.Exit(1)
		}

		f, err := p.Lookup("New")
		if err != nil {
			fmt.Printf("%s: lookup interface: %v\n", os.Args[0], err)
			os.Exit(1)
		}

		fn, ok := f.(func() so.SO)
		if !ok {
			fmt.Printf("%s: verify signature: %v\n", os.Args[0], err)
			os.Exit(1)
		}

		fn().Run()
	}
}
