package brace_validator

import (
	"strings"
)

const (
	_ = iota
	ROUND
	CURLY
	SQUARE
)

func Validate(i string) bool {

	s := strings.Split(i, "")
	opened, closed, allowed, types := getDataMaps()

	if !validateInput(s, allowed) {
		return false
	}
	if !validateFirstSybmol(s, opened) {
		return false
	}
	if !validateLastSybmol(s, closed) {
		return false
	}
	if !validateChain(s, opened, closed, types) {
		return false
	}
	return true
}

func getDataMaps() (opened map[string]bool, closed map[string]bool, allowed map[string]bool, types map[string]int) {
	types = map[string]int{
		"(": ROUND,
		")": ROUND,
		"{": CURLY,
		"}": CURLY,
		"[": SQUARE,
		"]": SQUARE,
	}
	opened = map[string]bool{
		"(": true,
		"[": true,
		"{": true,
	}
	closed = map[string]bool{
		")": true,
		"]": true,
		"}": true,
	}
	allowed = map[string]bool{
		"(": true,
		")": true,
		"{": true,
		"}": true,
		"[": true,
		"]": true,
	}
	return
}

func validateChain(s []string, opened map[string]bool, closed map[string]bool, types map[string]int) bool {
	type Counter struct {
		Opened map[int]int
		Closed map[int]int
	}
	c := Counter{make(map[int]int), make(map[int]int)}
	for k := range s {
		if ok := opened[s[k]]; ok {
			c.Opened[types[s[k]]]++
		} else if ok := closed[s[k]]; ok {
			c.Closed[types[s[k]]]++
		}
		if k == 0 {
			continue
		}
		if ok := opened[s[k-1]]; ok {
			if ok = closed[s[k]]; ok {
				if types[s[k-1]] != types[s[k]] {
					return false
				}
			}
		}
	}
	allTypes := []int{ROUND, CURLY, SQUARE}
	for k := range allTypes {
		if c.Opened[allTypes[k]] != c.Closed[allTypes[k]] {
			return false
		}
	}
	return true
}

func validateLastSybmol(i []string, closed map[string]bool) bool {
	if ok := closed[i[len(i)-1]]; !ok {
		return false
	}
	return true
}

func validateFirstSybmol(i []string, opened map[string]bool) bool {
	if ok := opened[i[0]]; !ok {
		return false
	}
	return true
}

func validateInput(i []string, allowed map[string]bool) bool {
	if len(i) < 1 {
		return false
	}
	for _, v := range i {
		if ok := allowed[v]; !ok {
			return false
		}
	}
	return true
}
