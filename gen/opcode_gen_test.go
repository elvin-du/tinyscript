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
	expected := `#p0 = 5 + 1
LW S0 STATIC 2
LW S1 STATIC 3
ADD S2 S0 S1
SW S2 SP -1
#p1 = 2 * p0
LW S0 STATIC 1
LW S1 SP -1
MULT S0 S1
MFLO S2
SW S2 SP -2
#p2 = 3 * p1
LW S0 STATIC 0
LW S1 SP -2
MULT S0 S1
MFLO S2
SW S2 SP -3
#a = p2
LW S0 SP -3
SW S0 SP 0`

	assert.Equal(t, prog.String(), expected)
}

func TestFuncEvaluate(t *testing.T) {
	node := parser.ParseFromFile("../tests/add.ts")
	taprog := translator.NewTranslator().Translate(node)
	g := NewOpCodeGen()
	prog := g.Gen(taprog)
	expected := `#FUNC_BEGIN
SW RA SP 0
#p1 = a + b
LW S0 SP -1
LW S1 SP -2
ADD S2 S0 S1
SW S2 SP -3
#RETURN p1
LW S0 SP -3
SW S0 SP 1
RETURN 
#FUNC_BEGIN
MAIN:SW RA SP 0
#PARAM 10 3
LW S0 STATIC 0
SW S0 SP -3
#PARAM 20 4
LW S0 STATIC 1
SW S0 SP -4
#SP -2
SUBI SP 2
#CALL L0
JR L0
#SP 2
ADDI SP 2
#RETURN
SW S0 SP 1
RETURN 
#SP -1
SUBI SP 1
#CALL L1
JR L1
#SP 1
ADDI SP 1`

	assert.Equal(t,prog.String(),expected)
}
