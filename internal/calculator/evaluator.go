package calculator

import (
	"fmt"
)

func Evaluate(exp Exp) float64 {

	switch exp := exp.(type) {
	case *Literal:
		return exp.Value
	case *BinaryOp:
		left := Evaluate(exp.Left)
		right := Evaluate(exp.Right)
		switch exp.Operator {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		case "/":
			return left / right
		}
	case *UnaryOp:
		operand := Evaluate(exp.Operand)
		switch exp.Operator {
		case "+":
			return +operand
		case "-":
			return -operand
		}
	default:
		fmt.Println("Unknown expression type", exp)
	}
	return 0
}
