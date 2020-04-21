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

func TestBlock(t *testing.T) {

	sourc := `var a = 1
{
var b = 1 * 100
}
{
var b = a * 100
}`

	ast := parser.Parse(sourc)
	translator := NewTranslator()
	program := translator.Translate(ast)
	t.Log(program)
	//	expected := `a = 1
	//SP -1
	//p1 = 1 * 100
	//b = p1
	//SP 1
	//SP -1
	//p1 = a * 100
	//b = p1
	//SP 1
	//`
	//	assert.Equal(t, program, expected)
}

func TestTranslator_TranslateIfStmt(t *testing.T) {
	source := `if(a){
b=1
}`

	astNode := parser.Parse(source)
	translator := NewTranslator()
	program := translator.Translate(astNode)
	t.Log(program)
}

func TestTranslator_TranslateIfElseStmt(t *testing.T) {
	source := `if(a){
b=1
}else{
b=2
}`

	astNode := parser.Parse(source)
	translator := NewTranslator()
	program := translator.Translate(astNode)
	t.Log(program)
}

func TestTranslator_TranslateIfElseIfStmt(t *testing.T) {
	astNode := parser.ParseFromFile("../tests/complex-if.ts")
	translator := NewTranslator()
	program := translator.Translate(astNode)
	t.Log(program)
}
