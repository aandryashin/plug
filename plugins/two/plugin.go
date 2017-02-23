package main

import (
	"fmt"
	"github.com/aandryashin/plug/so"
)

type Plugin struct {
	name  string
	value int
}

func (p *Plugin) Run() {
	fmt.Printf("I am %s plugin [%d]\n", p.name, p.value)
}

func New() so.SO {
	return &Plugin{"Integer", 17}
}
