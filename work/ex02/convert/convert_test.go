package convert_test

import (
	"fmt"
	"gopher-dojo/work/ex02/convert"
	"testing"
)

func TestSearchfile(t *testing.T) {
	var tests = []struct {
		name  string
		path  string
		pre   string
		after string
		err   error
	}{
		{name: "no file", path: "./", pre: "jpg", after: "png", err: fmt.Errorf("there is no target file : jpg")},
	}

	i := new(convert.ImageFormat)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			i.Path = tt.path
			i.PreFormat = tt.pre
			i.AfterFormat = tt.after
			got := i.Convert()
			if got.Error() != tt.err.Error() {
				t.Errorf("correct : ImageFormat() = %v, got : ImageFormat() = %v", got, tt.err)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	var tests = []struct {
		name  string
		path  string
		pre   string
		after string
	}{
		{name: "jpg to png", path: "../", pre: "jpg", after: "png"},
		{name: "jpg to gif", path: "../", pre: "jpg", after: "gif"},
		{name: "png to jpg", path: "../", pre: "png", after: "jpg"},
		{name: "png to jpeg", path: "../", pre: "png", after: "jpeg"},
	}

	i := new(convert.ImageFormat)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			i.Path = tt.path
			i.PreFormat = tt.pre
			i.AfterFormat = tt.after
			//t.Error(i.Path)
			got := i.Convert()
			if got != nil {
				t.Errorf("%T %v", got, got)
				t.Errorf("correct : ImageFormat() = nil, got : ImageFormat() = %v", got)
			}
		})
	}
}
