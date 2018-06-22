package gest

import "testing"

type It struct {
	*testing.T
}

func (it *It) Should(should string) *Expecting {
	return &Expecting{it, should}
}

type Expecting struct {
	*It
	should string
}

func (e *Expecting) Expect(expected interface{}) *Expected {
	return &Expected{e,expected}
}

type Expected struct {
	*Expecting
	expected interface{}
}

func (e *Expected) ToBe(actual interface{}) {
	if e.expected != actual {
		e.Errorf("Should %s. Expected: %v, Actual: %v", e.should, e.expected, actual)
	}
}

func I(t *testing.T) *It {
	return &It{t}
}