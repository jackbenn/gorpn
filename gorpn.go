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

const max_size = 100

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("RPN calculator")
	fmt.Println("Enter a bunch of symbols")
	// when I understand slices better I can do this better
	var stack []float64
	repeat := true
	for repeat {
		fmt.Print("] ")
		text, _ := reader.ReadString('\n')
		tokens := strings.Fields(text)

		is_number := regexp.MustCompile(`^\d*\.?\d*$`)
		for i := 0; i < len(tokens); i += 1 {
			token := tokens[i]

			if token == "" {
				continue
			} else if is_number.MatchString(token) {
				// should I do this through error handling?
				number, _ := strconv.ParseFloat(token, 64)
				stack = append(stack, number)
			} else if token == "+" {
				a := stack[len(stack)-2]
				b := stack[len(stack)-1]
				stack = append(stack[:len(stack)-2], a+b)
			} else if token == "-" {
				a := stack[len(stack)-2]
				b := stack[len(stack)-1]
				stack = append(stack[:len(stack)-2], a-b)
			} else if strings.HasPrefix(token, "q") {
				fmt.Println("exiting")
				repeat = false
				break
			}
			fmt.Println(token[0])

			fmt.Println(stack)

		}
	}

}
