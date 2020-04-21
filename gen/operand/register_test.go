package operand

import (
	"testing"
)

func TestString(t *testing.T) {
	op := NewRegister("cc", 22)
	Foo(t, op)
}

func Foo(t *testing.T, o Operand) {
	f := o.(*Register)
	t.Log(f.Name)
}
