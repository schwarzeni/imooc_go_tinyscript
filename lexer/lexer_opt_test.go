package lexer

import (
	"go-tinyscript/common"
	"reflect"
	"strings"
	"testing"
)

func TestMakeVarOrKeyword(t *testing.T) {
	// 读取一个单词测试
	type args struct {
		it *common.PeekIterator
	}
	tests := []struct {
		name    string
		args    args
		want    *Token
		wantErr bool
	}{
		{name: "test1", args: args{it: common.NewPeekIterator(strings.NewReader("var false"))}, wantErr: false, want: NewToken(KEYWORD, "var")},
		{name: "test2", args: args{it: common.NewPeekIterator(strings.NewReader("2var false"))}, wantErr: false, want: NewToken(VARIABLE, "2var")},
		{name: "test3", args: args{it: common.NewPeekIterator(strings.NewReader("false false"))}, wantErr: false, want: NewToken(BOOLEAN, "false")},
		{name: "test4", args: args{it: common.NewPeekIterator(strings.NewReader("true 123"))}, wantErr: false, want: NewToken(BOOLEAN, "true")},
		{name: "test5", args: args{it: common.NewPeekIterator(strings.NewReader("abc false"))}, wantErr: false, want: NewToken(VARIABLE, "abc")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeVarOrKeyword(tt.args.it)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeVarOrKeyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeVarOrKeyword() = %v, want %v", got, tt.want)
			}
		})
	}

	// 连续读取单词测试
	it := common.NewPeekIterator(strings.NewReader("var false abc23b true func 23 "))
	ts := []*Token{
		NewToken(KEYWORD, "var"),
		NewToken(BOOLEAN, "false"),
		NewToken(VARIABLE, "abc23b"),
		NewToken(BOOLEAN, "true"),
		NewToken(KEYWORD, "func"),
		NewToken(VARIABLE, "23"),
	}
	for _, token := range ts {
		got, _ := MakeVarOrKeyword(it)
		if !reflect.DeepEqual(got, token) {
			t.Errorf("MakeVarOrKeyword() = %v, want %v", got, token)
		}
		it.Next()
	}
}

func TestMakeString(t *testing.T) {
	type args struct {
		it *common.PeekIterator
	}
	tests := []struct {
		name    string
		args    args
		want    *Token
		wantErr bool
	}{
		{name: "test1", args: args{it: common.NewPeekIterator(strings.NewReader("\"abc\" 23"))}, want: NewToken(STRING, "\"abc\""), wantErr: false},
		{name: "test2", args: args{it: common.NewPeekIterator(strings.NewReader("\"abc 23"))}, want: nil, wantErr: true},
		{name: "test3", args: args{it: common.NewPeekIterator(strings.NewReader("'abc' 23"))}, want: NewToken(STRING, "'abc'"), wantErr: false},
		{name: "test4", args: args{it: common.NewPeekIterator(strings.NewReader("abc' 23"))}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeString(tt.args.it)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeOp(t *testing.T) {
	type args struct {
		it *common.PeekIterator
	}
	tests := []struct {
		name    string
		args    args
		want    *Token
		wantErr bool
	}{
		{name: "test1", args: args{it: common.NewPeekIterator(strings.NewReader("+ 2"))}, want: NewToken(OPERATOR, "+"), wantErr: false},
		{name: "test2", args: args{it: common.NewPeekIterator(strings.NewReader("++aa"))}, want: NewToken(OPERATOR, "++"), wantErr: false},
		{name: "test3", args: args{it: common.NewPeekIterator(strings.NewReader("+=56"))}, want: NewToken(OPERATOR, "+="), wantErr: false},
		{name: "test4", args: args{it: common.NewPeekIterator(strings.NewReader("-+"))}, want: NewToken(OPERATOR, "-"), wantErr: false},
		{name: "test5", args: args{it: common.NewPeekIterator(strings.NewReader("--d"))}, want: NewToken(OPERATOR, "--"), wantErr: false},
		{name: "test6", args: args{it: common.NewPeekIterator(strings.NewReader("-=>"))}, want: NewToken(OPERATOR, "-="), wantErr: false},
		{name: "test7", args: args{it: common.NewPeekIterator(strings.NewReader("*2"))}, want: NewToken(OPERATOR, "*"), wantErr: false},
		{name: "test8", args: args{it: common.NewPeekIterator(strings.NewReader("*=="))}, want: NewToken(OPERATOR, "*="), wantErr: false},
		{name: "test9", args: args{it: common.NewPeekIterator(strings.NewReader(">!"))}, want: NewToken(OPERATOR, ">"), wantErr: false},
		{name: "test10", args: args{it: common.NewPeekIterator(strings.NewReader(">=="))}, want: NewToken(OPERATOR, ">="), wantErr: false},
		{name: "test10", args: args{it: common.NewPeekIterator(strings.NewReader(">>343s"))}, want: NewToken(OPERATOR, ">>"), wantErr: false},
		{name: "test11", args: args{it: common.NewPeekIterator(strings.NewReader("<56"))}, want: NewToken(OPERATOR, "<"), wantErr: false},
		{name: "test12", args: args{it: common.NewPeekIterator(strings.NewReader("<=="))}, want: NewToken(OPERATOR, "<="), wantErr: false},
		{name: "test13", args: args{it: common.NewPeekIterator(strings.NewReader("<<343s"))}, want: NewToken(OPERATOR, "<<"), wantErr: false},
		{name: "test14", args: args{it: common.NewPeekIterator(strings.NewReader("==="))}, want: NewToken(OPERATOR, "=="), wantErr: false},
		{name: "test15", args: args{it: common.NewPeekIterator(strings.NewReader("=3!"))}, want: NewToken(OPERATOR, "="), wantErr: false},
		{name: "test16", args: args{it: common.NewPeekIterator(strings.NewReader("&56"))}, want: NewToken(OPERATOR, "&"), wantErr: false},
		{name: "test17", args: args{it: common.NewPeekIterator(strings.NewReader("&&="))}, want: NewToken(OPERATOR, "&&"), wantErr: false},
		{name: "test18", args: args{it: common.NewPeekIterator(strings.NewReader("&=343s"))}, want: NewToken(OPERATOR, "&="), wantErr: false},
		{name: "test19", args: args{it: common.NewPeekIterator(strings.NewReader("|56"))}, want: NewToken(OPERATOR, "|"), wantErr: false},
		{name: "test20", args: args{it: common.NewPeekIterator(strings.NewReader("||="))}, want: NewToken(OPERATOR, "||"), wantErr: false},
		{name: "test21", args: args{it: common.NewPeekIterator(strings.NewReader("|=343s"))}, want: NewToken(OPERATOR, "|="), wantErr: false},
		{name: "test22", args: args{it: common.NewPeekIterator(strings.NewReader("^56"))}, want: NewToken(OPERATOR, "^"), wantErr: false},
		{name: "test23", args: args{it: common.NewPeekIterator(strings.NewReader("^^="))}, want: NewToken(OPERATOR, "^^"), wantErr: false},
		{name: "test24", args: args{it: common.NewPeekIterator(strings.NewReader("^=343s"))}, want: NewToken(OPERATOR, "^="), wantErr: false},
		{name: "test25", args: args{it: common.NewPeekIterator(strings.NewReader("%2="))}, want: NewToken(OPERATOR, "%"), wantErr: false},
		{name: "test26", args: args{it: common.NewPeekIterator(strings.NewReader("%=343s"))}, want: NewToken(OPERATOR, "%="), wantErr: false},
		{name: "test27", args: args{it: common.NewPeekIterator(strings.NewReader("334"))}, want: nil, wantErr: true},
		{name: "test28", args: args{it: common.NewPeekIterator(strings.NewReader("a"))}, want: nil, wantErr: true},
		{name: "test29", args: args{it: common.NewPeekIterator(strings.NewReader("/=343s"))}, want: NewToken(OPERATOR, "/="), wantErr: false},
		{name: "test30", args: args{it: common.NewPeekIterator(strings.NewReader("/3"))}, want: NewToken(OPERATOR, "/"), wantErr: false},
		{name: "test31", args: args{it: common.NewPeekIterator(strings.NewReader("!=343s"))}, want: NewToken(OPERATOR, "!="), wantErr: false},
		{name: "test32", args: args{it: common.NewPeekIterator(strings.NewReader("!3"))}, want: NewToken(OPERATOR, "!"), wantErr: false},
		{name: "test33", args: args{it: common.NewPeekIterator(strings.NewReader(",23"))}, want: NewToken(OPERATOR, ","), wantErr: false},
		{name: "test34", args: args{it: common.NewPeekIterator(strings.NewReader(";;a"))}, want: NewToken(OPERATOR, ";"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeOp(tt.args.it)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeOp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeOp() = %v, want %v", got, tt.want)
			}
		})
	}
}
