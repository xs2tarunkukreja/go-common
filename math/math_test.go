package math

import "testing"

func TestAdd(t *testing.T) {
    want := 3
    if got := Add(1, 2); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}