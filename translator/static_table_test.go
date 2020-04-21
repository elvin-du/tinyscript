package translator

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"tinyscript/parser"
)

func TestStaticTable(t *testing.T) {
	source := `if(a){a=1}else{b=a+1*5}`
	node := parser.Parse(source)
	program := NewTranslator().Translate(node)
	assert.Equal(t, program.StaticTable.Size(), 2)
}
