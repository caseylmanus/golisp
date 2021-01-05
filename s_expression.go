package main

import (
	"fmt"
	"strconv"
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
func (exp Expression) Eval() (Symbol, error) {
	return apply(exp.Operator, exp.Args...)
}

type OperatorFunc func(rgs ...Symbol) (Symbol, error)

func apply(operator Symbol, atoms ...Atom) (Symbol, error) {
	funcs := map[Symbol]OperatorFunc{
		"+": func(args ...Symbol) (Symbol, error) {
			var total float64
			for _, i := range args {

				v, err := strconv.ParseFloat(string(i), 10)
				if err != nil {
					return "", err
				}
				total = total + v
			}
			return Symbol(fmt.Sprint(total)), nil
		},
	}
	f, ok := funcs[operator]
	if !ok {
		return "", fmt.Errorf("Unknown Operator %s", operator)
	}
	var args []Symbol
	for _, atom := range atoms {
		if atom.Symbol != nil {
			args = append(args, *atom.Symbol)
		} else if atom.Expression != nil {
			v, err := atom.Expression.Eval()
			if err != nil {
				return "", err
			}
			args = append(args, v)
		}

	}
	return f(args...)
}
