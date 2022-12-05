package main

import "testing"

func TestMain(t *testing.T) {
	if 1+1 != 2 {
		t.Error("math isn't math")
	}
}
