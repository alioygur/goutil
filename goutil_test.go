package goutil

import (
	"reflect"
	"testing"
)

func TestStrToSlice(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"comma sep", args{"one,two,three", ","}, []string{"one", "two", "three"}},
		{"should ignore extra sep", args{",,one,two,three,,", ","}, []string{"one", "two", "three"}},
		{"should ignore spaces", args{",,one , two, three , ,", ","}, []string{"one", "two", "three"}},
		{"should return nil", args{",,,", ","}, nil},
		{"space sep", args{"  one  two ", " "}, []string{"one", "two"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToSlice(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
