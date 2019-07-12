// A simple RPN calculator
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var isNumber = regexp.MustCompile(`^\d*\.?\d*$`)

const maxSize = 100

type RPN struct {
	stack []float64
}

type BinaryOperator func(float64, float64) float64

var ops = make(map[string]BinaryOperator)

func (rpn *RPN) push(x float64) {
	rpn.stack = append(rpn.stack, x)
}

func (rpn *RPN) pop() float64 {
	x := rpn.stack[len(rpn.stack)-1]
	rpn.stack = rpn.stack[:len(rpn.stack)-1]
	return x
}

func (rpn *RPN) parseToken(token string) {
	if token == "" {
		return
	} else if isNumber.MatchString(token) {
		// should I do this through error handling?
		number, _ := strconv.ParseFloat(token, 64)
		rpn.push(number)
	} else {
		switch token {
		case "+",
			"-",
			"*",
			"/":
			b := rpn.pop()
			a := rpn.pop()
			rpn.push(ops[token](a, b))
		case "^":
			b := rpn.pop()
			a := rpn.pop()
			rpn.push(math.Pow(a, b))
		}
	}
}

func (rpn *RPN) parseTokens(text string) bool {
	tokens := strings.Fields(text)

	for i := 0; i < len(tokens); i += 1 {
		token := tokens[i]

		if strings.HasPrefix(token, "q") {
			fmt.Println("exiting")
			return false
		} else {
			rpn.parseToken(token)
		}
		fmt.Println(rpn)
	}
	return true

}

func main() {
	ops["+"] = func(x float64, y float64) float64 { return x + y }
	ops["-"] = func(x float64, y float64) float64 { return x - y }
	ops["*"] = func(x float64, y float64) float64 { return x * y }
	ops["/"] = func(x float64, y float64) float64 { return x / y }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("RPN calculator")
	fmt.Println("Enter a bunch of symbols")
	// when I understand slices better I can do this better
	var rpn RPN
	repeat := true
	for repeat {
		fmt.Print("] ")
		text, _ := reader.ReadString('\n')
		repeat = rpn.parseTokens(text)
	}

}
