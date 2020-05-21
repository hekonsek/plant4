package model_test

import (
	"github.com/hekonsek/plant4/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLandscape(t *testing.T) {
	landscape, err := (&model.ModelParser{}).Load("../examples/landscape.yml")
	assert.NoError(t, err)
	assert.Len(t, landscape.Systems, 2)
}

func TestParseLandscapeName(t *testing.T) {
	landscape, err := (&model.ModelParser{}).Load("../examples/landscape.yml")
	assert.NoError(t, err)
	assert.Equal(t, "Happy Customer Cloud", landscape.Name)
}

func TestParseSystemRelation(t *testing.T) {
	landscape, err := (&model.ModelParser{}).Load("../examples/landscape.yml")
	assert.NoError(t, err)
	assert.NotNil(t, landscape.Systems[0].Relations)
}
