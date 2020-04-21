package gen

import "testing"

func TestF(t *testing.T) {
	var b byte = 1
	code := 0
	code |= int(b) << 9
	t.Logf("%b", code)
}
