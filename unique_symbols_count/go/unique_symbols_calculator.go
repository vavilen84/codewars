package main

import "strings"

func calculate(i string) int {
	a := strings.Split(i, "")
	m := make(map[string]bool, len(a))
	for _, v := range a {
		m[v] = true
	}
	return len(m)
}
