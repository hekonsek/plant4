package main

import (
	"github.com/hekonsek/plant4/model"
	"github.com/hekonsek/plant4/plantuml"
)

func main() {
	m, err := model.NewModelParser().Load("gi.yml")
	if err != nil {
		panic(err)
	}
	res, err := (&plantuml.Plotter{}).Plot(m)
	if err != nil {
		panic(err)
	}

	println(res[0])
	println()

	println(res[1])
	println()

	println(res[2])
	println()
}
