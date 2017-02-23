package main

import (
	"fmt"
	"github.com/aandryashin/plug/so"
)

type Plugin struct {
	name  string
	value string
}

func (p *Plugin) Run() {
	fmt.Printf("I am %s plugin [%s]\n", p.name, p.value)
}

func New() so.SO {
	return &Plugin{"String", "Hello, world!"}
}
