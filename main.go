package main

import 	"fmt"

func sum(num1 int, num2 int ) int {
	return  num1 + num2
}

func main() {
	var num1 = 5
	var num2 = 5
	fmt.Printf( "Resultado da soma de %d + %d é: %d", num1, num2, sum(num1,num2) )	
}
 