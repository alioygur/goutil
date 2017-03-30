package goutil

import "strings"

// StrToSlice converts string to slice of string by given sep
// Example:
// StrToSlice("one, two, three") = []string{"one", "two", "three"}
func StrToSlice(s string, sep string) []string {
	var slice []string
	s = strings.Trim(s, " "+sep)
	for _, p := range strings.Split(s, sep) {
		p = strings.Trim(p, " "+sep)
		if p == "" {
			continue
		}
		slice = append(slice, p)
	}
	return slice
}
