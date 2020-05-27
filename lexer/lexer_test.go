package lexer

import (
	"reflect"
	"strings"
	"testing"
)

func Test_all(t *testing.T) {
	str1 := `
  var a1=22;
  var a2=.3;
  var a3=-2.4;
  var a4;
  // this is a function
  // sum
  func sum(v1 int,v2 int) float {
    var v3 = v1+v2 -10.3*.2 + (v2<<2);
    return -v3;
  }
  /*测试注释*/
  /*
  多行注释
  */
  /*



  多行注释2*/
  a4=  1-sum(a1,a2)*a3;
  a3-1
  a2 = "abcdfrf"
  abcd='33x33'`

	l := Lexer{}
	tokens, err := l.Analyze(strings.NewReader(str1))
	if err != nil {
		t.Errorf("got error %v", err)
	}
	expected := []*Token{
		NewToken(KEYWORD, "var"),
		NewToken(VARIABLE, "a1"),
		NewToken(OPERATOR, "="),
		NewToken(INTEGER, "22"),
		NewToken(OPERATOR, ";"),
		NewToken(KEYWORD, "var"),
		NewToken(VARIABLE, "a2"),
		NewToken(OPERATOR, "="),
		NewToken(FLOAT, ".3"),
		NewToken(OPERATOR, ";"),
		NewToken(KEYWORD, "var"),
		NewToken(VARIABLE, "a3"),
		NewToken(OPERATOR, "="),
		NewToken(FLOAT, "-2.4"),
		NewToken(OPERATOR, ";"),
		NewToken(KEYWORD, "var"),
		NewToken(VARIABLE, "a4"),
		NewToken(OPERATOR, ";"),
		NewToken(KEYWORD, "func"),
		NewToken(VARIABLE, "sum"),
		NewToken(BRACKET, "("),
		NewToken(VARIABLE, "v1"),
		NewToken(KEYWORD, "int"),
		NewToken(OPERATOR, ","),
		NewToken(VARIABLE, "v2"),
		NewToken(KEYWORD, "int"),
		NewToken(BRACKET, ")"),
		NewToken(KEYWORD, "float"),
		NewToken(BRACKET, "{"),
		NewToken(KEYWORD, "var"),
		NewToken(VARIABLE, "v3"),
		NewToken(OPERATOR, "="),
		NewToken(VARIABLE, "v1"),
		NewToken(OPERATOR, "+"),
		NewToken(VARIABLE, "v2"),
		NewToken(OPERATOR, "-"),
		NewToken(FLOAT, "10.3"),
		NewToken(OPERATOR, "*"),
		NewToken(FLOAT, ".2"),
		NewToken(OPERATOR, "+"),
		NewToken(BRACKET, "("),
		NewToken(VARIABLE, "v2"),
		NewToken(OPERATOR, "<<"),
		NewToken(INTEGER, "2"),
		NewToken(BRACKET, ")"),
		NewToken(OPERATOR, ";"),
		NewToken(KEYWORD, "return"),
		NewToken(OPERATOR, "-"),
		NewToken(VARIABLE, "v3"),
		NewToken(OPERATOR, ";"),
		NewToken(BRACKET, "}"),
		NewToken(VARIABLE, "a4"),
		NewToken(OPERATOR, "="),
		NewToken(INTEGER, "1"),
		NewToken(OPERATOR, "-"),
		NewToken(VARIABLE, "sum"),
		NewToken(BRACKET, "("),
		NewToken(VARIABLE, "a1"),
		NewToken(OPERATOR, ","),
		NewToken(VARIABLE, "a2"),
		NewToken(BRACKET, ")"),
		NewToken(OPERATOR, "*"),
		NewToken(VARIABLE, "a3"),
		NewToken(OPERATOR, ";"),
		NewToken(VARIABLE, "a3"),
		NewToken(OPERATOR, "-"),
		NewToken(INTEGER, "1"),
		NewToken(VARIABLE, "a2"),
		NewToken(OPERATOR, "="),
		NewToken(STRING, "\"abcdfrf\""),
		NewToken(VARIABLE, "abcd"),
		NewToken(OPERATOR, "="),
		NewToken(STRING, "'33x33'"),
	}
	if len(expected) != len(tokens) {
		t.Errorf("the size is not equal: expected %d, but got %d", len(expected), len(tokens))
		return
	}
	for idx, token := range tokens {
		if !reflect.DeepEqual(token, expected[idx]) {
			t.Errorf("[%d] not equal: expected %v, but got %v", idx, expected[idx], token)
		}
	}
}

func Test_badcomment(t *testing.T) {
	strs := []string{"aavv /*    /  d", "aavv /*      /", "aavv /*/"}
	for _, str := range strs {
		l := Lexer{}
		if _, err := l.Analyze(strings.NewReader(str)); err == nil {
			t.Errorf("comment %s is not right", str)
		}
	}
}

func Test_expression(t *testing.T) {
	str1 := "(a+b)^100.12==+100-20"
	l := Lexer{}
	tokens, err := l.Analyze(strings.NewReader(str1))
	expected := []*Token{
		NewToken(BRACKET, "("),
		NewToken(VARIABLE, "a"),
		NewToken(OPERATOR, "+"),
		NewToken(VARIABLE, "b"),
		NewToken(BRACKET, ")"),
		NewToken(OPERATOR, "^"),
		NewToken(FLOAT, "100.12"),
		NewToken(OPERATOR, "=="),
		NewToken(INTEGER, "+100"),
		NewToken(OPERATOR, "-"),
		NewToken(INTEGER, "20"),
	}
	if err != nil {
		t.Errorf("got error %v\n", err)
		return
	}
	if len(expected) != len(tokens) {
		t.Errorf("the size is not equal: expected %d, but got %d", len(expected), len(tokens))
		return
	}
	for idx, token := range tokens {
		if !reflect.DeepEqual(token, expected[idx]) {
			t.Errorf("[%d] not equal: expected %v, but got %v", idx, expected[idx], token)
		}
	}
}

func Test_function(t *testing.T) {
	str1 := "func foo(a, b){\n" +
		"print(a+b)\n" +
		"}\n" +
		"foo(-100.0, 100)"

	l := Lexer{}
	tokens, err := l.Analyze(strings.NewReader(str1))
	if err != nil {
		t.Errorf("got error %v\n", err)
		return
	}
	expected := []*Token{
		NewToken(KEYWORD, "func"),
		NewToken(VARIABLE, "foo"),
		NewToken(BRACKET, "("),
		NewToken(VARIABLE, "a"),
		NewToken(OPERATOR, ","),
		NewToken(VARIABLE, "b"),
		NewToken(BRACKET, ")"),
		NewToken(BRACKET, "{"),
		NewToken(VARIABLE, "print"),
		NewToken(BRACKET, "("),
		NewToken(VARIABLE, "a"),
		NewToken(OPERATOR, "+"),
		NewToken(VARIABLE, "b"),
		NewToken(BRACKET, ")"),
		NewToken(BRACKET, "}"),
		NewToken(VARIABLE, "foo"),
		NewToken(BRACKET, "("),
		NewToken(FLOAT, "-100.0"),
		NewToken(OPERATOR, ","),
		NewToken(INTEGER, "100"),
		NewToken(BRACKET, ")"),
	}
	if len(expected) != len(tokens) {
		t.Errorf("the size is not equal: expected %d, but got %d", len(expected), len(tokens))
		return
	}
	for idx, token := range tokens {
		if !reflect.DeepEqual(token, expected[idx]) {
			t.Errorf("[%d] not equal: expected %v, but got %v", idx, expected[idx], token)
		}
	}
}

// fmt.Printf("NewToken(%s, \"%s\"),\n", strings.ToUpper(string(token.GetType())), token.GetValue())
