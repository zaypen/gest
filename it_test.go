package gest

import "testing"

func TestI(t *testing.T) {
	I(t).Should("work").Expect(true).ToBe(true)
}