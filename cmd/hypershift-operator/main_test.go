package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewControlPlaneOperatorCommand(t *testing.T) {
	controllerNames := []string{}
	for name := range controllerFuncs {
		controllerNames = append(controllerNames, name)
	}
	cmd := newControlPlaneOperator()
	assert.ElementsMatch(t, controllerNames, cmd.Controllers)
}
