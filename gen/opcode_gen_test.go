package gen

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"tinyscript/parser"
	"tinyscript/translator"
)

func TestExprEvaluate(t *testing.T) {
	source := "var a = 3 * 2*(5+1)"
	node := parser.Parse(source)
	taprog := translator.NewTranslator().Translate(node)
	assert.Equal(t, taprog.StaticTable.String(), `0:3
1:2
2:5
3:1`)

	g := NewOpCodeGen()
	prog := g.Gen(taprog)
	t.Log(prog.String())
	//todo
}

func TestFuncEvaluate(t *testing.T) {
	node := parser.ParseFromFile("../tests/add.ts")
	taprog := translator.NewTranslator().Translate(node)
	g := NewOpCodeGen()
	prog := g.Gen(taprog)
	t.Log(prog)
}
