package gosnake

// WARN: PYTHONPATH must be set to demo before tests can run!

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	b := NewBinding()
	e := b.Import("adder")
	assert.Nil(t, e)
}
