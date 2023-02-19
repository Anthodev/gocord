package gocord

import "strings"

func ContainsUserID(elems []string, v string) bool {
	for _, s := range elems {
		if strings.Contains(s, v) {
			return true
		}
	}

	return false
}
