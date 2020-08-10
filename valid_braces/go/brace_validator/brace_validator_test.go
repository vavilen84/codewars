package brace_validator

import (
	"testing"
)

func TestValidate(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Validate OK (){}[]",
			args{
				i: "(){}[]",
			},
			true,
		},
		{
			"Validate OK ([{}])",
			args{
				i: "([{}])",
			},
			true,
		},
		{
			"Validate FAIL (}",
			args{
				i: "(}",
			},
			false,
		},
		{
			"Validate FAIL 'not allowed'",
			args{
				i: "not allowed",
			},
			false,
		},
		{
			"Validate FAIL empty string",
			args{
				i: "",
			},
			false,
		},
		{
			"Validate FAIL closed first",
			args{
				i: "}",
			},
			false,
		},
		{
			"Validate FAIL opened last",
			args{
				i: "((",
			},
			false,
		},
		{
			"Validate FAIL wrong chaining [(])",
			args{
				i: "[(])",
			},
			false,
		},
		{
			"Validate FAIL wrong chaining [({})](]",
			args{
				i: "[({})](]",
			},
			false,
		},
		{
			"Validate FAIL wrong chaining (()",
			args{
				i: "(()",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.i); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateChain(t *testing.T) {
	opened, closed, _, types := getDataMaps()
	type args struct {
		s      []string
		opened map[string]bool
		closed map[string]bool
		types  map[string]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Validate OK chain",
			args{
				s:      []string{"(", "(", ")", ")"},
				opened: opened,
				closed: closed,
				types:  types,
			},
			true,
		},
		{
			"Validate OK chain with the same type, last closed",
			args{
				s:      []string{"(", ")"},
				opened: opened,
				closed: closed,
				types:  types,
			},
			true,
		},
		{
			"Validate OK chain",
			args{
				s:      []string{"(", "[", "]",")"},
				opened: opened,
				closed: closed,
				types:  types,
			},
			true,
		},
		{
			"Validate FAIL chain",
			args{
				s:      []string{"(", "}"},
				opened: opened,
				closed: closed,
				types:  types,
			},
			false,
		},
		{
			"Validate FAIL chain",
			args{
				s:      []string{"(", "}", ")", "{"},
				opened: opened,
				closed: closed,
				types:  types,
			},
			false,
		},
		{
			"Validate FAIL chain",
			args{
				s:      []string{"(", "{", ")", "}"},
				opened: opened,
				closed: closed,
				types:  types,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateChain(tt.args.s, tt.args.opened, tt.args.closed, tt.args.types); got != tt.want {
				t.Errorf("validateChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateFirstSybmol(t *testing.T) {
	opened, _, _, _ := getDataMaps()
	type args struct {
		i      []string
		opened map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Validate FAIL first closed symbol",
			args{
				i:      []string{")"},
				opened: opened,
			},
			false,
		},
		{
			"Validate OK",
			args{
				i:      []string{"("},
				opened: opened,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateFirstSybmol(tt.args.i, tt.args.opened); got != tt.want {
				t.Errorf("validateFirstSybmol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateInput(t *testing.T) {
	_, _, allowed, _ := getDataMaps()

	var allowedSlice []string
	for k := range allowed {
		allowedSlice = append(allowedSlice, k)
	}

	type args struct {
		i       []string
		allowed map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Validate FAIL not allowed symbols",
			args{
				i:       []string{"not allowed", "%", "+", "", " "},
				allowed: allowed,
			},
			false,
		},
		{
			"Validate OK",
			args{
				i:       allowedSlice,
				allowed: allowed,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateInput(tt.args.i, tt.args.allowed); got != tt.want {
				t.Errorf("validateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateLastSybmol(t *testing.T) {
	_, closed, _, _ := getDataMaps()
	type args struct {
		i      []string
		closed map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Validate FAIL last opened symbol",
			args{
				i:      []string{"("},
				closed: closed,
			},
			false,
		},
		{
			"Validate OK",
			args{
				i:      []string{")"},
				closed: closed,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateLastSybmol(tt.args.i, tt.args.closed); got != tt.want {
				t.Errorf("validateLastSybmol() = %v, want %v", got, tt.want)
			}
		})
	}
}
