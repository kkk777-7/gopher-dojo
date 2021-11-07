package main

import (
	"flag"
	"testing"
)

var tests = []struct {
	name  string
	path  string
	pre   string
	after string
	want  bool
}{
	{name: "jpg->png", path: ".", pre: "jpg", after: "png", want: true},
	{name: "aaa->png", path: ".", pre: "aaa", after: "png", want: false},
	{name: "jpg->aaa", path: ".", pre: "jpg", after: "aaa", want: false},
}

func TestFlag(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			flag.Set("p", tt.pre)
			flag.Set("a", tt.after)
			flag.Set("path", tt.path)
			if got := CheckFormat(tt.path, tt.pre, tt.after); got != tt.want {
				t.Errorf("CheckFormat(%s,%s,%s) = %v", tt.path, tt.pre, tt.after, got)
			}
		})
	}
}
