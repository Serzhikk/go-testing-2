package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sam")
	want := "Hello, Sam!"
	if got != want {
		t.Errorf("Hello(\"Sam\")=%s; want %s", got, want)
	}
}
