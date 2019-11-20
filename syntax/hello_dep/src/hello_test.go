package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Ciao mondo."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
