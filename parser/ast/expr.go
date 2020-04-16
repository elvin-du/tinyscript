package ast

var _ ASTNode = &Expr{}

type Expr struct {
	*node
}

func NewExpr() *Expr {
	return &Expr{NewNode()}
}
