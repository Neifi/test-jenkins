package main

import "testing"

func Test(t *testing.T) {

	t.Run("Should fail", func(t *testing.T) {
		t.Fail()
	})
}
