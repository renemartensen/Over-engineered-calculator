package calculator

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"simple addition", "3 + 5", []string{"3", "+", "5"}},
		{"expression involving all operators", "3 + 5 * (2 - 8) / 5", []string{"3", "+", "5", "*", "(", "2", "-", "8", ")", "/", "5"}},
		{"Invalid chararcters", "3 + 5 & - 2", nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tokens, _ := Tokenize(test.input)
			if !reflect.DeepEqual(tokens, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, tokens)
			}
		})
	}
}

func TestParser(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected Exp
	}{
		{
			"simple addition",
			[]string{"3", "+", "5"},
			&BinaryOp{
				Left:     &Literal{Value: 3},
				Operator: "+",
				Right:    &Literal{Value: 5},
			},
		},
		{
			"precendence +- vs */",
			[]string{"3", "+", "5", "*", "2"},
			&BinaryOp{
				Left:     &Literal{Value: 3},
				Operator: "+",
				Right: &BinaryOp{
					Left:     &Literal{Value: 5},
					Operator: "*",
					Right:    &Literal{Value: 2},
				},
			},
		},
		{
			"parentheses precendence",
			[]string{"3", "*", "(", "3", "+", "5", ")"},
			&BinaryOp{
				Left:     &Literal{Value: 3},
				Operator: "*",
				Right: &BinaryOp{
					Left:     &Literal{Value: 3},
					Operator: "+",
					Right:    &Literal{Value: 5},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parser := &Parser{tokens: test.input}
			ast := parser.Parse()
			if !reflect.DeepEqual(ast, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, ast)
			}
		})
	}
}

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name     string
		input    Exp
		expected float64
	}{
		{
			"simple addition",
			&BinaryOp{
				Left:     &Literal{Value: 3},
				Operator: "+",
				Right:    &Literal{Value: 5},
			},
			8,
		},
		{
			"precendence +- vs */",
			&BinaryOp{
				Left:     &Literal{Value: 3},
				Operator: "+",
				Right: &BinaryOp{
					Left:     &Literal{Value: 5},
					Operator: "*",
					Right:    &Literal{Value: 2},
				},
			},
			13,
		},
		{
			"parentheses precendence",
			&BinaryOp{
				Left:     &Literal{Value: 3},
				Operator: "*",
				Right: &BinaryOp{
					Left:     &Literal{Value: 3},
					Operator: "+",
					Right:    &Literal{Value: 5},
				},
			},
			24,
		},
		{
			"unary minus",
			&BinaryOp{
				Left:     &Literal{Value: 3},
				Operator: "*",
				Right: &UnaryOp{
					Operator: "-",
					Operand:  &Literal{Value: 5},
				},
			},
			-15,
		},
		{
			"division by zero",
			&BinaryOp{
				Left:     &Literal{Value: 10},
				Operator: "/",
				Right:    &Literal{Value: 0},
			},
			0, // expect error
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Evaluate(test.input)
			if test.name == "division by zero" && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if test.name == "division by zero" && err != nil {
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestFullCalculator(t *testing.T) {
	expression := "3 + 5 * (2-8) / 5"
	result, err := EvaluateExpression(expression)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := -3.0
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
