package lexer

// Tokens2Arrays []*Token --> []interface{}
func Tokens2Arrays(tokens []*Token) (arr []interface{}) {
	arr = make([]interface{}, len(tokens))
	for idx, _ := range arr {
		arr[idx] = tokens[idx]
	}
	return
}
