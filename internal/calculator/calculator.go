package calculator

func EvaluateExpression(input string) (float64, error) {
	tokens, err := Tokenize(input)
	if err != nil {
		return 0, err
	}
	parser := &Parser{tokens: tokens}
	ast := parser.Parse()
	return Evaluate(ast), nil
}
