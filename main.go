package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		panic("no expression provided")
	}

	fmt.Println(sum(1, 1.3))
	parse("1 + 2")
}

type TokenType int

const (
	Value TokenType = iota
	Operator
)

type Token struct {
	val       float64
	TokenType TokenType
	Operation func(float64, float64) float64
}

type Expression struct {
	tokens []*Token
}

func isNumber(r rune) bool {
	return (r >= '0' && r <= '9') || r == '.'
}

func isOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}

func operatorToken(r rune) *Token {
	switch r {
	case '+':
		return &Token{
			val:       0,
			TokenType: Operator,
			Operation: sum,
		}
	case '-':
		return &Token{
			val:       0,
			TokenType: Operator,
			Operation: sub,
		}
	case '*':
		return &Token{
			val:       0,
			TokenType: Operator,
			Operation: multiply,
		}
	case '/':
		return &Token{
			val:       0,
			TokenType: Operator,
			Operation: division,
		}
	}
	return nil
}

func isWhitespace(r rune) bool {
	return r == ' '
}

func parse(expr string) *Expression {
	number := ""
	tokens := []*Token{}
	for _, r := range expr {
		if isWhitespace(r) {
			v, err := strconv.ParseFloat(number, 64)
			if err != nil {
				panic(err)
			}

			tokens = append(tokens, &Token{
				val:       v,
				TokenType: Value,
			})

			number = ""
			continue
		}

		if isNumber(r) {
			number += string(r)
		}

		if isOperator(r) {
			tokens = append(tokens, operatorToken(r))
		}
	}

	return &Expression{
		tokens: tokens,
	}
}

func sum(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func division(a, b float64) float64 {
	return a / b
}
