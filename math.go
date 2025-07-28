package main 

import (
	"fmt"
	"flag"
)

func main() {
	number1 := flag.Int("n1",0,"First number")
	number2 := flag.Int("n2",0,"Second number")
	operator := flag.String("op","+","Operator : +,-,*,/")
	flag.Parse()
	
	var result int
	switch(*operator) {
	case "+":
		result = *number1 + *number2
	case "-":
		result = *number1 - *number2
	case "*":
		result = *number1 * *number2
	case "/":
		if *number2 == 0 {
			fmt.Println("Division by zero not possible")
			return
		} else {
					result = *number1 / *number2
		
		}
	case "%"	
	default:
		fmt.Println("Invalid Operator")
		return
	}
	fmt.Println(result)
}
