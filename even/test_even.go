package even

import "testing"

func TestEven(t *testing.T) {
	got := Even(6)
	want := true
	if got!= want {
		t.Errorf("error in Even Func, got %t; want %t", got, want)
	}
}
