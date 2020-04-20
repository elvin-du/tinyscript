package translator

import (
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
