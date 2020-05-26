package common

import "testing"

func TestIsLetter(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{c: "a"}, want: true},
		{name: "test2", args: args{c: "A"}, want: true},
		{name: "test3", args: args{c: "2"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLetter(tt.args.c); got != tt.want {
				t.Errorf("IsLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{c: "2"}, want: true},
		{name: "test2", args: args{c: "5"}, want: true},
		{name: "test3", args: args{c: "a"}, want: false},
		{name: "test3", args: args{c: "*"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.c); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLiteral(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{c: "2"}, want: true},
		{name: "test2", args: args{c: "a"}, want: true},
		{name: "test3", args: args{c: "B"}, want: true},
		{name: "test4", args: args{c: "*"}, want: false},
		{name: "test5", args: args{c: "<"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLiteral(tt.args.c); got != tt.want {
				t.Errorf("IsLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOperator(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{c: "*"}, want: true},
		{name: "test2", args: args{c: "+"}, want: true},
		{name: "test3", args: args{c: "/"}, want: true},
		{name: "test4", args: args{c: "2"}, want: false},
		{name: "test5", args: args{c: "-"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOperator(tt.args.c); got != tt.want {
				t.Errorf("IsOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
