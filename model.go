package gosnake

// Models to expose interface

// Representation of a python module
type Module struct {
	name string
}

// Python equivalent: import name
func Import(name string) (*Module, error) {
	// Ensure we're initialized
	initializeBinding()

	m := &Module{
		name: name,
	}

	return m, nil
}

// Python equivalent: module.function(args)
// If the call raises an exception in python its passed along as a go error
func (b *Module) Call(function string, args ...interface{}) (interface{}, error) {
	op := &Operation{
		module:     b,
		target:     function,
		args:       "",
		returnChan: make(chan string),
	}

	opChan <- op

	ret := <-op.returnChan
	return ret, nil
}
