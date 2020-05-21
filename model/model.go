package model

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Model struct {
	Name    string
	Systems []*System
	Shared  []*Shared
}

type System struct {
	Name string
	Containers []*Container
	Relations []*Relation
}

type Relation struct {
	Target string
	Name string
}

type Container struct {
	Name string
	Relations []*Relation
}

type Shared struct {
	Name string
}

type ModelParser struct {
}

func NewModelParser() *ModelParser {
	return &ModelParser{}
}

func (*ModelParser) Load(path string) (*Model, error) {
	definition, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var landscape Model
	err = yaml.Unmarshal(definition, &landscape)
	if err != nil {
		return nil, err
	}
	return &landscape, nil
}