package main

import "fmt"

func main() {
	var ret int
	
	/* calling a function to get max value */
	ret = max(10,15)
	fmt.Printf( "Max value is : %d\n", ret)
}

/* function returning the max between two numbers */
func max(num1, num2 int)int{
	/* local variable declaration */
	var result int;
	
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result 
}

