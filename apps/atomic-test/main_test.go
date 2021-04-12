package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	m1 := loadConfig()
	m2 := loadConfig()

	assert.NotEqual(t, m1["id"], m2["id"], "they should not be equal")
}
