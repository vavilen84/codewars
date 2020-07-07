package _go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidIp(t *testing.T) {
	invalidInputs := []string{
		``,
		`invalid input`,
		`.1`,
		`1.2.3`,
		`1.2.3.4.`,
		`1.2.3.4.5`,
		`123.456.78.90`,
		`123.045.067.089`,
	}

	var valid bool

	for _, v := range invalidInputs {
		valid = IsValidIp(v)
		assert.False(t, valid)
	}

	valid = IsValidIp(`123.123.123.123`)
	assert.True(t, valid)
}
