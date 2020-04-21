package translator

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"tinyscript/parser/ast"
)

func TestIsInstanceOf(t *testing.T) {
	var i interface{}
	i = &ast.Expr{}
	assert.Equal(t, IsInstanceOfExpr(i), true)
}
