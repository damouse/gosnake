package gosnake

// WARN: PYTHONPATH must be set to demo before tests can run!

import (
	"testing"

	"github.com/sbinet/go-python"
	"github.com/stretchr/testify/assert"
)

//
// Python -> Go
func TestStringToGo(t *testing.T) {
	o := "there once was a man from peru"
	c, e := togo(python.PyString_FromString(o))

	assert.Nil(t, e)
	assert.Equal(t, o, c.(string))
}

func TestIntToGo(t *testing.T) {
	o := 123
	c, e := togo(python.PyInt_FromLong(o))

	assert.Nil(t, e)
	assert.Equal(t, o, c.(int))
}

// func TestBoolToGo(t *testing.T) {
// 	o := 123
// 	c, e := togo(python.PyInt_FromLong(o))

// 	assert.Nil(t, e)
// 	assert.Equal(t, o, c.(int))
// }

// func TestFloatToGo(t *testing.T) {
// 	o := float64(1234)
// 	c, e := togo(python.PyLong_FromDouble(o))

// 	assert.Nil(t, e)
// 	assert.Equal(t, o, c.(float64))
// }

// Finish this off, but the types on incoming (py -> go) may be different types

// func TestDoubleToGo(t *testing.T) {
// 	o := 123
// 	c, e := togo(python.PyInt_FromLong(o))

// 	assert.Nil(t, e)
// 	assert.Equal(t, o, c.(int))
// }

// func TestSliceToGo(t *testing.T) {
// 	o := 123
// 	c, e := togo(python.PyInt_FromLong(o))

// 	assert.Nil(t, e)
// 	assert.Equal(t, o, c.(int))
// }

// func TestDictToGo(t *testing.T) {
// 	o := 123
// 	c, e := togo(python.PyInt_FromLong(o))

// 	assert.Nil(t, e)
// 	assert.Equal(t, o, c.(int))
// }

// func TestNoneToGo(t *testing.T) {
// 	o := 123
// 	c, e := togo(python.PyInt_FromLong(o))

// 	assert.Nil(t, e)
// 	assert.Equal(t, o, c.(int))
// }

//
// Go -> Python
//

func TestIntToPy(t *testing.T) {
	o := 123
	c, e := topy(o)

	assert.Nil(t, e)
	assert.True(t, python.PyInt_Check(c))
}

func TestBoolToPy(t *testing.T) {
	o := true
	c, e := topy(o)

	assert.Nil(t, e)
	assert.True(t, python.PyBool_Check(c))
}

func TestStringToPy(t *testing.T) {
	c, e := topy("horses")
	assert.Nil(t, e)
	assert.True(t, python.PyString_Check(c))
}

func TestFloatToPy(t *testing.T) {
	c, e := topy(float32(1234))
	assert.Nil(t, e)
	assert.True(t, python.PyFloat_Check(c))
}

func TestLongToPy(t *testing.T) {
	c, e := topy(float64(1234))
	assert.Nil(t, e)
	assert.True(t, python.PyLong_Check(c))
}
