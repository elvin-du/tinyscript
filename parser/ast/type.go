package ast

type NodeType int

const (
	ASTNODE_TYPE_BLOCK NodeType = iota

	ASTNODE_TYPE_BINARY_EXPR // 1+1
	ASTNODE_TYPE_UNARY_EXPR  //++1
	ASTNODE_TYPE_CALL_EXPR

	ASTNODE_TYPE_VARIABLE
	ASTNODE_TYPE_SCALAR // 1.0 true

	ASTNODE_TYPE_IF_STMT
	ASTNODE_TYPE_WHILE_STMT
	ASTNODE_TYPE_FOR_STMT
	ASTNODE_TYPE_RETURN_STMT
	ASTNODE_TYPE_ASSIGN_STMT
	ASTNODE_TYPE_FUNCTION_DECLARE_STMT
	ASTNODE_TYPE_DECLARE_STMT
)

var NodeTypeStringMap = map[NodeType]string{
	ASTNODE_TYPE_BLOCK:                 "block",
	ASTNODE_TYPE_ASSIGN_STMT:           "assign_stmt",
	ASTNODE_TYPE_BINARY_EXPR:           "binary_expr",
	ASTNODE_TYPE_UNARY_EXPR:            "unary_expr",
	ASTNODE_TYPE_CALL_EXPR:             "call_expr",
	ASTNODE_TYPE_DECLARE_STMT:          "declare_stmt",
	ASTNODE_TYPE_FOR_STMT:              "for_stmt",
	ASTNODE_TYPE_FUNCTION_DECLARE_STMT: "function_declare_stmt",
	ASTNODE_TYPE_IF_STMT:               "if_stmt",
	ASTNODE_TYPE_RETURN_STMT:           "return_stmt",
	ASTNODE_TYPE_SCALAR:                "scalar",
	ASTNODE_TYPE_VARIABLE:              "variable",
	ASTNODE_TYPE_WHILE_STMT:            "while_stmt",
}

func (nt NodeType) String() string {
	return NodeTypeStringMap[nt]
}
