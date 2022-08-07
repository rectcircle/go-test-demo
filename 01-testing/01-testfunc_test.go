package testingdemo

import "testing"

// go test -run ^TestIntAbs$ ./01-testing

func TestIntAbs(t *testing.T) {
	got := IntAbs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
	got = IntAbs(1)
	if got != 1 {
		t.Errorf("Abs(1) = %d; want 1", got)
	}
	t.Run("a", func(t *testing.T) {
		t.Log(t.Name())
	})
}
