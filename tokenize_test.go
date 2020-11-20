package main

import (
	"reflect"
	"testing"
)

func Test_tokenize(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "(* 8 8 )",
			args: "(* 8 8 )",
			want: []string{"(* 8 8 )"},
		},
		{
			name: "(* 8 8 ) 8",
			args: "(* 8 8 ) 8",
			want: []string{"(* 8 8 )", "8"},
		},
		{
			name: "run (* 8 8 ) 8",
			args: "run (* 8 8 ) 8",
			want: []string{"run", "(* 8 8 )", "8"},
		},
		{
			name: "run (* 8 8 ) 8 stop",
			args: "run (* 8 8 ) 8 stop",
			want: []string{"run", "(* 8 8 )", "8", "stop"},
		},
		{
			name: "run (* (sq(4)) 8 ) 8 stop",
			args: "run (* (sq(4)) 8 ) 8 stop",
			want: []string{"run", "(* (sq(4)) 8 )", "8", "stop"},
		},
		{
			name: "run (* (sq(4)) 8 ) 8 'stop'",
			args: "run (* (sq(4)) 8 ) 8 'stop'",
			want: []string{"run", "(* (sq(4)) 8 )", "8", "'stop'"},
		},
		{
			name: `run (* (sq(4)) 8 ) 8 "shoot()" "stop"`,
			args: `run (* (sq(4)) 8 ) 8 "shoot()" "stop"`,
			want: []string{"run", "(* (sq(4)) 8 )", "8", `"shoot()"`, `"stop"`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tokenize(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
