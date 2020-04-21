package translator

import (
	"fmt"
	"tinyscript/lexer"
	"tinyscript/parser/ast"
	"tinyscript/translator/symbol"
)

type Translator struct {
}

func NewTranslator() *Translator {
	return &Translator{}
}

func (t *Translator) Translate(node ast.ASTNode) *TAProgram {
	program := NewTAProgram()

	table := symbol.NewTable()
	for _, child := range node.Children() {
		t.TranslateStmt(program, child, table)
	}

	return program
}

func (t *Translator) TranslateStmt(program *TAProgram, node ast.ASTNode, table *symbol.Table) {
	switch node.Type() {
	case ast.ASTNODE_TYPE_ASSIGN_STMT:
		t.TranslateAssignStmt(program, node, table)
		return
	case ast.ASTNODE_TYPE_DECLARE_STMT:
		t.TranslateDeclareStmt(program, node, table)
		return
	case ast.ASTNODE_TYPE_BLOCK:
		t.TranslateBlock(program, node, table)
		return
	case ast.ASTNODE_TYPE_IF_STMT:
		t.TranslateIfStmt(program, node.(*ast.IfStmt), table)
		return
	}

	panic("unknown node type" + node.Type().String())
}

func (t *Translator) TranslateDeclareStmt(program *TAProgram, node ast.ASTNode, table *symbol.Table) {
	lexeme := node.GetChild(0).Lexeme()
	if table.Exists(lexeme) {
		panic("Syntax Error, Identifier " + lexeme.Value + " is already defined")
	}
	assigned := table.CreateSymbolByLexeme(lexeme)
	expr := node.GetChild(1)
	addr := t.TranslateExpr(program, expr, table)
	program.Add(NewTAInstruction(TAINSTR_TYPE_ASSIGN, assigned, "=", addr, nil))
}

func (t *Translator) TranslateAssignStmt(program *TAProgram, node ast.ASTNode, table *symbol.Table) {
	assigned := table.CreateSymbolByLexeme(node.GetChild(0).Lexeme())
	expr := node.GetChild(1)
	addr := t.TranslateExpr(program, expr, table)
	program.Add(NewTAInstruction(TAINSTR_TYPE_ASSIGN, assigned, "=", addr, nil))
}

/*
SDD:
	E -> E1 op E2
	E -> F
*/
func (t *Translator) TranslateExpr(program *TAProgram, node ast.ASTNode, table *symbol.Table) *symbol.Symbol {
	if node.IsValueType() {
		addr := table.CreateSymbolByLexeme(node.Lexeme())
		node.SetProp("addr", addr)
		return addr
	} else if node.Type() == ast.ASTNODE_TYPE_CALL_EXPR {
		panic("not now")
	}

	for _, child := range node.Children() {
		t.TranslateExpr(program, child, table)
	}

	if node.Prop("addr") == nil {
		node.SetProp("addr", table.CreateVariable())
	}

	instr := NewTAInstruction(
		TAINSTR_TYPE_ASSIGN,
		node.Prop("addr").(*symbol.Symbol),
		node.Lexeme().Value,
		node.GetChild(0).Prop("addr").(*symbol.Symbol),
		node.GetChild(1).Prop("addr").(*symbol.Symbol),
	)

	program.Add(instr)
	return instr.Result
}

func (t *Translator) TranslateBlock(program *TAProgram, node ast.ASTNode, parent *symbol.Table) {
	table := symbol.NewTable()
	parent.AddChild(table)
	parentOffset := table.CreateVariable()
	parentOffset.Lexeme = lexer.NewToken(lexer.INTEGER, fmt.Sprintf("%d", parent.LocalSize()))

	pushRecord := NewTAInstruction(TAINSTR_TYPE_SP, nil, "", nil, nil)
	program.Add(pushRecord)
	for _, stmt := range node.Children() {
		t.TranslateStmt(program, stmt, table)
	}

	popRecord := NewTAInstruction(TAINSTR_TYPE_SP, nil, "", nil, nil)
	program.Add(popRecord)

	pushRecord.Arg1 = -parent.LocalSize()
	popRecord.Arg1 = parent.LocalSize()
}

func (t *Translator) TranslateIfStmt(program *TAProgram, node *ast.IfStmt, table *symbol.Table) {
	expr := node.GetExpr()
	exprAddr := t.TranslateExpr(program, expr, table)
	ifOpCode := NewTAInstruction(TAINSTR_TYPE_IF, nil, "", exprAddr, nil)
	program.Add(ifOpCode)

	t.TranslateBlock(program, node.GetBlock(), table)

	var gotoInstr *TAInstruction = nil
	if node.GetChild(2) != nil {
		gotoInstr = NewTAInstruction(TAINSTR_TYPE_GOTO, nil, "", nil, nil)
		program.Add(gotoInstr)
		labelEndIf := program.AddLabel()
		ifOpCode.Arg2 = labelEndIf.Arg1
	}

	if node.GetElseBlock() != nil {
		t.TranslateBlock(program, node.GetElseBlock(), table)
	} else if node.GetElseIfStmt() != nil {
		t.TranslateIfStmt(program, node.GetElseIfStmt().(*ast.IfStmt), table)
	}

	labelEnd := program.AddLabel()
	if node.GetChild(2) == nil {
		ifOpCode.Arg2 = labelEnd.Arg1
	} else {
		gotoInstr.Arg1 = labelEnd.Arg1
	}
}
