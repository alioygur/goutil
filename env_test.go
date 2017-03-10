package goutil

import (
	"os"
	"reflect"
	"testing"
)

func TestEnvMustGet(t *testing.T) {
	var ek, ev = "ENV1", "env1 val"
	os.Setenv(ek, ev)
	defer os.Unsetenv(ek)

	if got := EnvMustGet(ek); got != ev {
		t.Errorf("EnvMustGet() = %v, want %v", got, ev)
	}

	t.Run("should return panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fail()
			}
		}()

		EnvMustGet("NOT_EXISTS")
	})
}

func TestEnvMustInt(t *testing.T) {
	var ek, ev, evi = "ENV1", "1", 1
	os.Setenv(ek, ev)
	defer os.Unsetenv(ek)

	if got := EnvMustInt(ek); got != evi {
		t.Errorf("EnvMustInt() = %v, want %v", got, evi)
	}

	t.Run("should return panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fail()
			}
		}()
		var ek, ev = "ENV2", "string"
		os.Setenv(ek, ev)
		defer os.Unsetenv(ek)

		EnvMustInt(ek)
	})
}

func TestEnvMustSliceStr(t *testing.T) {
	type args struct {
		k   string
		sep string
	}
	tests := []struct {
		name   string
		args   args
		want   []string
		envVal string
	}{
		{"comma sep", args{"env1", ","}, []string{"one", "two", "three"}, "one,two,three"},
		{"should ignore extra sep", args{"env2", ","}, []string{"one", "two", "three"}, ",,one,two,three,,"},
		{"should ignore spaces", args{"env3", ","}, []string{"one", "two", "three"}, ",,one , two, three , ,"},
		{"should return nil", args{"env4", ","}, nil, ",,,"},
		{"space sep", args{"env5", " "}, []string{"one", "two"}, "  one  two "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.args.k, tt.envVal)
			defer os.Unsetenv(tt.args.k)
			if got := EnvMustSliceStr(tt.args.k, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnvMustSliceStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnsureEnvExists(t *testing.T) {
	os.Setenv("exists1", "exists 1 val")
	defer os.Unsetenv("exists1")
	os.Setenv("exists2", "exists 2 val")
	defer os.Unsetenv("exists2")

	type args struct {
		envs map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"all env that given are exists", args{map[string]string{"exists1": "", "exists2": ""}}, false},
		{"all env that given aren't exists but has fallback value", args{map[string]string{"exists1": "", "exists2": "", "fallback1": "fallback 1 val"}}, false},
		{"with not exists so should return err", args{map[string]string{"exists1": "", "exists2": "", "fallback1": "fallback 1 val", "notexists": ""}}, true},
		{"empty list shouldn't return err", args{nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EnsureEnvExists(tt.args.envs); (err != nil) != tt.wantErr {
				t.Errorf("EnsureEnvExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
