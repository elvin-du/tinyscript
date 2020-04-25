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

func TestTranslator_TranslateDeclareStmt(t *testing.T) {
	source := "var a=1.0*2.0*3.0"
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
	expected := `a = 1
p1 = 1 * 100
b = p1
p1 = a * 100
b = p1`
	assert.Equal(t, program.String(), expected)
}

func TestTranslator_TranslateIfStmt(t *testing.T) {
	source := `if(a){
b=1
}`

	astNode := parser.Parse(source)
	translator := NewTranslator()
	program := translator.Translate(astNode)
	expected := `IF a ELSE L0
b = 1
L0:`
	assert.Equal(t, program.String(), expected)
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
	expected := `IF a ELSE L0
b = 1
GOTO L1
L0:
b = 2
L1:`
	assert.Equal(t, program.String(), expected)
}

func TestTranslator_TranslateIfElseIfStmt(t *testing.T) {
	astNode := parser.ParseFromFile("../tests/complex-if.ts")
	translator := NewTranslator()
	program := translator.Translate(astNode)
	expected := `p0 = a == 1
IF p0 ELSE L0
b = 100
GOTO L5
L0:
p1 = a == 2
IF p1 ELSE L1
b = 500
GOTO L4
L1:
p2 = a == 3
IF p2 ELSE L2
p1 = a * 1000
b = p1
GOTO L3
L2:
b = -1
L3:
L4:
L5:`
	assert.Equal(t, program.String(), expected)
}

func TestSimpleFunction(t *testing.T) {
	node := parser.ParseFromFile("../tests/function.ts")
	translator := NewTranslator()
	program := translator.Translate(node)
	expected := `L0:
FUNC_BEGIN
p1 = a + b
RETURN p1`
	assert.Equal(t, program.String(), expected)
}

func TestRecursionFunc(t *testing.T) {
	node := parser.ParseFromFile("../tests/recursion.ts")
	translator := NewTranslator()
	program := translator.Translate(node)
	expected := `L0:
FUNC_BEGIN
p1 = n == 0
IF p1 ELSE L1
RETURN 1
L1:
p2 = n - 1
PARAM p2 6
SP -5
CALL L0
SP 5
p4 = p3 * n
RETURN p4`
	assert.Equal(t, program.String(), expected)
}
