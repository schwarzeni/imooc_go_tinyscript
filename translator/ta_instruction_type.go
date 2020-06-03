package translator

// TAInstructionType 三地址指令类型
type TAInstructionType string

const (
	ASSIGN     TAInstructionType = "assign"
	GOTO       TAInstructionType = "goto"
	IF         TAInstructionType = "if"
	LABEL      TAInstructionType = "label"
	CALL       TAInstructionType = "call"
	RETURN     TAInstructionType = "return"
	SP         TAInstructionType = "sp"
	PARAM      TAInstructionType = "param"
	FUNC_BEGIN TAInstructionType = "function begin"
)
