package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
калькулятор на горутинах без примитимов синхронизации

=> 2 + 2, 3 * 7, 9 / 3
<= 4, 21, 3
*/

func main() {
	fmt.Print("Input an expression separated with ', ': ")
	expressions, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	trimNewline(&expressions[len(expressions)-1])

	resultsChannel := make(chan int)
	for _, exp := range expressions {
		go calculateExpression(exp, resultsChannel)
	}
	time.Sleep(5 * time.Second)

	for i := range resultsChannel {
		fmt.Println(i)
	}
}

func readInput() ([]string, error) {
	in := bufio.NewReader(os.Stdin)
	userInput, err := in.ReadString('\n')
	if err != nil {
		return []string{}, err
	}

	var expressions = strings.Split(userInput, ", ")
	return expressions, nil
}

func trimNewline(str *string) {
	*str = strings.TrimSuffix(*str, "\n")
}

func calculateExpression(exp string, resultChannel chan int) {
	result := 0
	resultChannel <- result
}
