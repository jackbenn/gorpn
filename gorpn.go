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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("RPN calculator")
	fmt.Println("Enter a bunch of symbols")
	var stack []float64
	for {
		fmt.Print("] ")
		text, _ := reader.ReadString('\n')
		tokens := strings.Split(text, " ")

		is_number := regexp.MustCompile(`^\d*$`)
		for i := 0; i < len(tokens); i += 1 {
			token := tokens[i]
			fmt.Println(token)

			if token == "" {
				continue
			} else if is_number.MatchString(token) {
				number, _ := strconv.ParseFloat(token, 64)
				stack = append(stack, number)
			}
			fmt.Println(stack)

		}
	}

}
