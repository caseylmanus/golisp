package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseExpression(t *testing.T) {
	exp := "(* 9 9)"
	nine := Symbol("9")
	expect := Expression{
		Operator: "*",
		Args: []Atom{{
			Symbol: &nine,
		}, {
			Symbol: &nine,
		}},
	}
	res, err := ParseExpression(exp)
	require.NoError(t, err)
	require.Equal(t, res, expect)
}
