package main

import "testing"

func Test_MyFunc(t *testing.T) {
	res := myFunc()
	t.Run("ok", func(t *testing.T) {
		if res != "myFunc" {
			t.Failed()
		}
	})
}
