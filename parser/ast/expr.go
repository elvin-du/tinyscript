package ast

import (
	"tinyscript/lexer"
)

var DefaultExpr = MakeExpr()

type Expr struct {
	*node
}

func MakeExpr() *Expr {
	e := &Expr{MakeNode()}
	return e
}

func NewExpr2(typ NodeType, token *lexer.Token) *Expr {
	expr := MakeExpr()
	expr.SetType(typ)
	expr.SetLexeme(token)
	return expr
}

type ExprHOF func() ASTNode

//left: E(k) -> E(k) op(k) E(k+1) | E(k+1)
//right:
//		E(k) -> E(k+1) E_(k)
//		E_(k) -> op(k) E(k+1) E_(k) | ⍷
// 最高优先级：
// 		E(t) -> F E_(k) | U E_(k)
//		E_(t) -> op(t) E(t) E_(t) | ⍷

func (e *Expr) E(stream *PeekTokenStream, k int) ASTNode {
	if k < PriorityTable.Size()-1 {
		return e.combine(
			stream,
			func() ASTNode {
				return e.E(stream, k+1)
			},
			func() ASTNode {
				return e.E_(stream, k)
			},
		)
	}

	return e.race(
		stream,
		func() ASTNode {
			return e.combine(
				stream,
				func() ASTNode {
					return e.F(stream)
				},
				func() ASTNode {
					return e.E_(stream, k)
				},
			)
		},
		func() ASTNode {
			return e.combine(
				stream,
				func() ASTNode {
					return e.U(stream)
				},
				func() ASTNode {
					return e.E_(stream, k)
				},
			)
		},
	)
}

func (e *Expr) U(stream *PeekTokenStream) ASTNode {
	token := stream.Peek()
	value := token.Value
	if value == "(" {
		stream.NextMatch("(")
		expr := e.E(stream, 0)
		stream.NextMatch(")")
		return expr
	} else if value == "++" || value == "--" || value == "!" {
		t := stream.Peek()
		stream.NextMatch(value)
		unaryExpr := NewExpr2(ASTNODE_TYPE_UNARY_EXPR, t)
		unaryExpr.AddChild(e.E(stream, 0))
		return unaryExpr
	}

	return nil
}

func (e *Expr) F(stream *PeekTokenStream) ASTNode {
	factor := FactorParse(stream)
	if nil == factor {
		return nil
	}

	if stream.HasNext() && stream.Peek().Value == "(" {
		return CallExprParse(factor, stream)
	}

	return factor
}

func (e *Expr) E_(stream *PeekTokenStream, k int) ASTNode {
	token := stream.Peek()
	value := token.Value
	if PriorityTable.IsContain(k, value) {
		expr := NewExpr2(ASTNODE_TYPE_BINARY_EXPR, stream.NextMatch(value))
		expr.AddChild(
			e.combine(
				stream,
				func() ASTNode {
					return e.E(stream, k+1)
				},
				func() ASTNode {
					return e.E_(stream, k)
				},
			),
		)
		return expr
	}

	return nil
}

func (e *Expr) race(stream *PeekTokenStream, af ExprHOF, bf ExprHOF) ASTNode {
	if !stream.HasNext() {
		return nil
	}

	a := af()
	if nil != a {
		return a
	}

	return bf()
}

func (e *Expr) combine(stream *PeekTokenStream, af ExprHOF, bf ExprHOF) ASTNode {
	a := af()
	if nil == a {
		if stream.HasNext() {
			return bf()
		}
		return nil
	}

	var b ASTNode = nil
	if stream.HasNext() {
		b = bf()
		if nil == b {
			return a
		}
	} else {
		return a
	}

	expr := NewExpr2(ASTNODE_TYPE_BINARY_EXPR, b.Lexeme())
	expr.AddChild(a)
	expr.AddChild(b.GetChild(0))

	return expr
}
func ExprParse(stream *PeekTokenStream) ASTNode {
	return DefaultExpr.E(stream, 0)
}
