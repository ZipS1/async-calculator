package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
калькулятор на горутинах без примитимов синхронизации

=> 2 + 2, 3 * 7, 9 / 3
<= 4, 21, 3
*/

type expression struct {
	exp    string
	err    error
	result int
}

func (e *expression) GetExpression() string {
	return e.exp
}

func (e *expression) GetError() error {
	return e.err
}

func (e *expression) GetResult() int {
	return e.result
}

func (e *expression) SetError(err error) {
	e.err = err
}

func (e *expression) HasError() bool {
	return e.err != nil
}

func (e *expression) SetResult(result int) {
	e.result = result
}

type Expression interface {
	GetExpression() string
	GetError() error
	GetResult() int
	SetError(err error)
	HasError() bool
	SetResult(result int)
}

func NewExpression(exp ...string) Expression {
	structExp := expression{
		exp:    "",
		err:    nil,
		result: 0,
	}

	if len(exp) != 0 {
		structExp.exp = exp[0]
	}

	return &structExp
}

func main() {
	fmt.Print("Input an expressions separated with ', ': ")
	expressions, err := ReadInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	TrimSuffices(&expressions[len(expressions)-1])

	resultsChannel := make(chan Expression)
	for _, exp := range expressions {
		go CalculateExpression(exp, resultsChannel)
	}

	for i := 0; i < len(expressions); i++ {
		exp := <-resultsChannel
		if exp.HasError() == false {
			fmt.Printf("%s = %d\n", exp.GetExpression(), exp.GetResult())
		} else {
			fmt.Printf("%s evaluated with error: %s\n", exp.GetExpression(), exp.GetError())
		}
	}
}

func ReadInput() ([]string, error) {
	in := bufio.NewReader(os.Stdin)
	userInput, err := in.ReadString('\n')
	if err != nil {
		return []string{}, err
	}

	var expressions = strings.Split(userInput, ", ")
	return expressions, nil
}

func TrimSuffices(str *string) {
	*str = strings.TrimSuffix(*str, "\n")
	*str = strings.TrimSuffix(*str, "\r")
}

func CalculateExpression(exp string, resultChannel chan Expression) {
	resultingExpression := NewExpression(exp)
	tokenAmountError := fmt.Errorf("expression must contain 3 tokens")
	invalidTokenError := fmt.Errorf("invalid token value or order")

	tokens := strings.Fields(exp)
	if len(tokens) != 3 {
		resultingExpression.SetError(tokenAmountError)
	}

	result, err := EvaluateTokens(tokens)
	if err != nil {
		resultingExpression.SetError(invalidTokenError)
	}

	resultingExpression.SetResult(result)
	resultChannel <- resultingExpression
}

func CastNumbers(tokens []string) (int, int, error) {
	firstNum, firstErr := strconv.Atoi(tokens[0])
	secondNum, secondErr := strconv.Atoi(tokens[2])
	var err error = nil
	if firstErr != nil {
		err = firstErr
	} else if secondErr != nil {
		err = secondErr
	}

	return firstNum, secondNum, err
}

func EvaluateTokens(tokens []string) (int, error) {
	firstNum, secondNum, err := CastNumbers(tokens)
	if err != nil {
		return 0, err
	}

	switch tokens[1] {
	case "+":
		return firstNum + secondNum, nil
	case "-":
		return firstNum - secondNum, nil
	case "*":
		return firstNum * secondNum, nil
	case "/":
		return firstNum / secondNum, nil
	default:
		return 0, fmt.Errorf("")
	}
}
