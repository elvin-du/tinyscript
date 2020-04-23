package gen

import (
	"strconv"
	"strings"
	"tinyscript/translator"
)

type OpCodeProgram struct {
	Entry        *int
	Instructions []*Instruction
	Comments     map[int]string //注释；行号：注释内容
}

func NewOpCodeProgram() *OpCodeProgram {
	return &OpCodeProgram{Entry: nil, Instructions: make([]*Instruction, 0), Comments: make(map[int]string)}
}

func (o *OpCodeProgram) Add(instr *Instruction) {
	o.Instructions = append(o.Instructions, instr)
}
func (o *OpCodeProgram) String() string {
	prts := make([]string, 0, len(o.Instructions))
	for i, instr := range o.Instructions {
		if c, ok := o.Comments[i]; ok {
			prts = append(prts, "#"+c)
		}
		str := instr.String()
		if o.Entry != nil && *o.Entry == i {
			str = "MAIN:" + str
		}
		prts = append(prts, str)
	}

	return strings.Join(prts, "\n")
}

func (o *OpCodeProgram) SetEntry(entry *int) {
	o.Entry = entry
}

//当前指令的位置添加一行注释
func (o *OpCodeProgram) AddComment(comment string) {
	o.Comments[len(o.Instructions)] = comment
}

func (o *OpCodeProgram) ToByteCode() []int {
	codes := []int{}
	for _, instr := range o.Instructions {
		codes = append(codes, instr.ToByteCode())
	}

	return codes
}

//从三地址代码中获取静态符号表中的值，存起来在虚拟机实例化时写入内存静态区
func (o *OpCodeProgram) GetStaticArea(taProgram *translator.TAProgram) []int {
	l := []int{}
	for _, symbol := range taProgram.StaticTable.Symbols {
		i, err := strconv.Atoi(symbol.Lexeme.Value)
		if nil != err {
			panic(err)
		}
		l = append(l, i)
	}
	return l
}
