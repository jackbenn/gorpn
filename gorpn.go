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
type UnaryOperator func(float64) float64

var binops = make(map[string]BinaryOperator)
var unops = make(map[string]UnaryOperator)

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
			"/",
			"^",
			"%":
			b := rpn.pop()
			a := rpn.pop()
			rpn.push(binops[token](a, b))
		case "sin",
			"cos",
			"tan",
			"sqrt",
			"ln",
			"exp":
			a := rpn.pop()
			rpn.push(unops[token](a))
		case "pi":
			rpn.push(math.Pi)
		case "e":
			rpn.push(math.E)
		case "$":
			b := rpn.pop()
			a := rpn.pop()
			rpn.push(b)
			rpn.push(a)
		case "?":
			fmt.Println("binary operators: + - * / ^ %")
			fmt.Println("Unary Operators: sin cos tan sqrt ln exp")
			fmt.Println("Constants: pi e")
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
	}
	fmt.Println(rpn.stack)
	return true
}

func main() {
	binops["+"] = func(x float64, y float64) float64 { return x + y }
	binops["-"] = func(x float64, y float64) float64 { return x - y }
	binops["*"] = func(x float64, y float64) float64 { return x * y }
	binops["/"] = func(x float64, y float64) float64 { return x / y }
	binops["^"] = math.Pow
	binops["%"] = math.Mod
	unops["sin"] = math.Sin
	unops["cos"] = math.Cos
	unops["tan"] = math.Tan
	unops["sqrt"] = math.Sqrt
	unops["ln"] = math.Log
	unops["exp"] = math.Exp

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("RPN calculator")
	fmt.Println("Enter numbers or operators, or q to quit or ? for help")
	// when I understand slices better I can do this better
	var rpn RPN
	repeat := true
	for repeat {
		fmt.Print("] ")
		text, _ := reader.ReadString('\n')
		repeat = rpn.parseTokens(text)
	}
}
