package main

import (
	"fmt"
	"strings"
)

/*
An s-expression takes the form of (operator ..args) and is the primary sytax of a lisp-dialect
Examples:
Multiply: (* 2 3)
*/

// Symbol represents a symbol in the code
type Symbol string

// Expression represents and s-expression
type Expression struct {
	Operator Symbol
	Args     []Atom
}

// Atom
type Atom struct {
	Symbol     *Symbol
	Expression *Expression
}

// IsExpresion tests to see if a fragment might be an expression
func IsExpression(fragment string) bool {
	s := strings.TrimSpace(fragment)
	return strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")")
}

// ParseExpression will parse an expression
func ParseExpression(expression string) (Expression, error) {
	expression = strings.TrimSpace(expression)
	if !IsExpression(expression) {
		return Expression{}, fmt.Errorf("%s is not a valid expression", expression)
	}
	expression = strings.TrimPrefix(expression, "(")
	expression = strings.TrimSuffix(expression, ")")
	tokens := tokenize(expression)
	result := Expression{}
	for i, token := range tokens {
		switch {
		case i == 0:
			result.Operator = Symbol(token)
		case IsExpression(token):
			exp, err := ParseExpression(token)
			if err != nil {
				return Expression{}, err
			}
			result.Args = append(result.Args, Atom{Expression: &exp})
		default:
			s := Symbol(token)
			result.Args = append(result.Args, Atom{Symbol: &s})
		}
	}
	return result, nil
}
