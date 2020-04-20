package translator

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"tinyscript/parser"
	"tinyscript/translator/symbol"
)

func TestExprTranslator(t *testing.T) {
	source := `a+(b-c)+d*(b-c)*2`
	p := parser.Parse(source)
	exprNode := p.GetChild(0)
	translator := NewTranslator()
	table := symbol.NewTable()
	program := NewTAProgram()
	translator.TranslateExpr(program, exprNode, table)
	expected := `p0 = b - c
p1 = b - c
p2 = p1 * 2
p3 = d * p2
p4 = p0 + p3
p5 = a + p4`

	assert.Equal(t, program.String(), expected)
}

func TestAssignStmt(t *testing.T) {
	source := "a=1.0*2.0*3.0"
	node := parser.Parse(source)
	translator := NewTranslator()
	program := translator.Translate(node)

	expected := `p0 = 2.0 * 3.0
p1 = 1.0 * p0
a = p1`
	assert.Equal(t, program.String(), expected)
}

func TestAssignStmt2(t *testing.T) {
	source := "a=1"
	node := parser.Parse(source)
	translator := NewTranslator()
	program := translator.Translate(node)

	assert.Equal(t, program.String(), "a = 1")
}
