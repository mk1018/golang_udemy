package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 0
	if Debug {
		t.Skip("スキップ")
	}
	v := IsOne(i)

	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}
