// A simple RPN calculator
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RPN struct {
	stack []float64
}

func (rpn *RPN) push(x float64) {
	fmt.Println(append(rpn.stack, x))
	rpn.stack = append(rpn.stack, x)
}

func (rpn *RPN) pop() float64 {
	x := rpn.stack[len(rpn.stack)-1]
	rpn.stack = rpn.stack[:len(rpn.stack)-1]
	return x
}

const maxSize = 100

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("RPN calculator")
	fmt.Println("Enter a bunch of symbols")
	// when I understand slices better I can do this better
	var stack RPN
	repeat := true
	for repeat {
		fmt.Print("] ")
		text, _ := reader.ReadString('\n')
		tokens := strings.Fields(text)

		isNumber := regexp.MustCompile(`^\d*\.?\d*$`)
		for i := 0; i < len(tokens); i += 1 {
			token := tokens[i]

			if token == "" {
				continue
			} else if isNumber.MatchString(token) {
				// should I do this through error handling?
				fmt.Println("found a number")
				number, _ := strconv.ParseFloat(token, 64)
				stack.push(number)
			} else if token == "+" {
				b := stack.pop()
				a := stack.pop()
				stack.push(a + b)
			} else if token == "-" {
				b := stack.pop()
				a := stack.pop()
				stack.push(a - b)
			} else if strings.HasPrefix(token, "q") {
				fmt.Println("exiting")
				repeat = false
				break
			}

			fmt.Println(stack)

		}
	}

}
