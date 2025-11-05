package calculator

import (
	"fmt"
)

func Evaluate(exp Exp) (float64, error) {

	switch exp := exp.(type) {
	case *Literal:
		return exp.Value, nil
	case *BinaryOp:
		left, err_left := Evaluate(exp.Left)
		right, err_right := Evaluate(exp.Right)
		if err_left != nil {
			return 0, err_left
		}
		if err_right != nil {
			return 0, err_right
		}
		switch exp.Operator {
		case "+":
			return left + right, nil
		case "-":
			return left - right, nil
		case "*":
			return left * right, nil
		case "/":
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return left / right, nil
		}
	case *UnaryOp:
		operand, err := Evaluate(exp.Operand)
		if err != nil {
			return 0, err
		}
		switch exp.Operator {
		case "+":
			return +operand, nil
		case "-":
			return -operand, nil
		}
	}
	return 0, fmt.Errorf("unknown expression type")
}
