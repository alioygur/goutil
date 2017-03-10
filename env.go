package goutil

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// EnvMustGet env must be exists and its value != "", else panic
func EnvMustGet(k string) string {
	env := os.Getenv(k)
	if env == "" {
		panic(fmt.Sprintf("%s env not exists", k))
	}
	return env
}

// EnvMustInt converts env value to int then returns, else panic
func EnvMustInt(k string) int {
	i, err := strconv.Atoi(os.Getenv(k))
	if err != nil {
		panic(fmt.Errorf("can't convert `%s` env to string: %v", os.Getenv(k), err))
	}
	return i
}

// EnvMustSliceStr returns a slice of string seperated by sep
func EnvMustSliceStr(k, sep string) []string {
	var s []string
	envVal := strings.Trim(EnvMustGet(k), " "+sep)
	for _, p := range strings.Split(envVal, sep) {
		p = strings.Trim(p, " "+sep)
		if p == "" {
			continue
		}
		s = append(s, p)
	}
	return s
}

// EnsureEnvExists ensures given env exists.
func EnsureEnvExists(envs map[string]string) error {
	for k, v := range envs {
		if os.Getenv(k) == "" { // not exists
			if v != "" { // is there a fallback value? then set
				os.Setenv(k, v)
			} else { // else return err
				return fmt.Errorf("missing env: %s", k)
			}
		}
	}
	return nil
}
