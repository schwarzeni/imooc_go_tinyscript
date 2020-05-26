package common

import "regexp"

var (
	ptnLetter   = regexp.MustCompile("^[a-zA-Z]$")
	ptnNumber   = regexp.MustCompile("^[0-9]$")
	ptnLiteral  = regexp.MustCompile("^[_a-zA-Z0-9]$")
	ptnOperator = regexp.MustCompile("^[*+\\-<>=!&|^%/,]$")
)

// IsLetter 判断是否为英文字符
func IsLetter(c string) bool {
	return ptnLetter.MatchString(c)
}

// IsNumber 判断是否为数字
func IsNumber(c string) bool {
	return ptnNumber.MatchString(c)
}

// IsLiteral 判断是否为字符和数字
func IsLiteral(c string) bool {
	return ptnLiteral.MatchString(c)
}

// IsOperator 判断是否为运算符
func IsOperator(c string) bool {
	return ptnOperator.MatchString(c)
}
