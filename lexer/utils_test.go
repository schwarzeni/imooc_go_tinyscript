package lexer

import (
	"reflect"
	"testing"
)

func TestTokens2Arrays(t *testing.T) {
	type args struct {
		tokens []*Token
	}
	tests := []struct {
		name    string
		args    args
		wantArr []interface{}
	}{
		{name: "test1", args: args{[]*Token{
			NewToken(KEYWORD, "var"),
			NewToken(INTEGER, "1"),
			NewToken(BOOLEAN, "false"),
		}}, wantArr: []interface{}{
			NewToken(KEYWORD, "var"),
			NewToken(INTEGER, "1"),
			NewToken(BOOLEAN, "false"),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArr := Tokens2Arrays(tt.args.tokens); !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("Tokens2Arrays() = %v, want %v", gotArr, tt.wantArr)
			}
		})
	}
}
