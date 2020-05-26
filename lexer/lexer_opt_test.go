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
