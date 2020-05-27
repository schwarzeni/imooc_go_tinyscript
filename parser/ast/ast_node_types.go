package ast

type ASTNodeType string

const (
	NONE                  ASTNodeType = "none"
	BLOCK                 ASTNodeType = "block"
	BINARY_EXPR           ASTNodeType = "binaryExpr" // 2+3
	UNARY_EXPR            ASTNodeType = "unaryExpr"  // ++i
	CALL_EXPTR            ASTNodeType = "call"
	VARIABLE              ASTNodeType = "variable"
	SCALAR                ASTNodeType = "scalar" // integer, float, boolean, string
	IF_STMT               ASTNodeType = "if"
	WHILE_STMT            ASTNodeType = "while"
	FOR_STMT              ASTNodeType = "for"
	RETURN_STMT           ASTNodeType = "return"
	ASSIGN_STMT           ASTNodeType = "assign"
	FUNCTION_DECLARE_STMT ASTNodeType = "func"
	FUNCTION_ARGS         ASTNodeType = "function args"
	DECLARE_STMT          ASTNodeType = "declare"
)
