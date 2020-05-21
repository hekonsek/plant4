package plantuml_test

import (
	"github.com/hekonsek/plant4/model"
	"github.com/hekonsek/plant4/plantuml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlotSystem(t *testing.T) {
	landscape, err := (&model.ModelParser{}).Load("../examples/landscape.yml")
	assert.NoError(t, err)
	diagrams, err := (&plantuml.Plotter{}).Plot(landscape)
	assert.NoError(t, err)
	assert.Contains(t, diagrams[0], "node CRM")
	assert.Contains(t, diagrams[0], "node Accounting")
	println(diagrams[0])
}

func TestPlotMultipleSystems(t *testing.T) {
	// Given
	m, err := model.NewModelParser().Load("../examples/landscape.yml")
	assert.NoError(t, err)

	// When
	diagrams, err := (&plantuml.Plotter{}).Plot(m)
	assert.NoError(t, err)

	// Then
	assert.Len(t, diagrams, 3)
}


func TestPlotContainer(t *testing.T) {
	container, err := model.NewModelParser().Load("../examples/container.yml")
	assert.NoError(t, err)
	diagrams, err := (&plantuml.Plotter{}).Plot(container)
	assert.NoError(t, err)
	//assert.Contains(t, diagrams[0], "node CRM")
	//assert.Contains(t, diagrams[0], "node Accounting")
	println(diagrams[1])
}

func TestPlotShared(t *testing.T) {
	container, err := (&model.ModelParser{}).Load("../examples/shared.yml")
	assert.NoError(t, err)
	diagrams, err := (&plantuml.Plotter{}).Plot(container)
	assert.NoError(t, err)
	//assert.Contains(t, diagrams[0], "node CRM")
	//assert.Contains(t, diagrams[0], "node Accounting")
	println(diagrams[1])
}