package main

import "fmt"
//import "strings"

func main() {
	
	
	/* calling a function to get max value */
	/*
	var ret int
	ret = max(10,15)
	fmt.Printf( "Max value is : %d\n", ret)
	*/
	
	/* calling a function to get more than one value*/
	/*
	a, b := swap("Hola", "Mundo")
	fmt.Println(a, b)
	*/
	
	/*
	greetings :=  []string{"Hello","world!", "16", "20"}
	fmt.Println(strings.Join(greetings, " "))
	*/
	
	//Sample of using pointers
	/*
	var a int = 10
	fmt.Printf("Address of a variable: %x\n", &a  )
	*/
	
	//usingPointers()
	usingPointerArray()
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

/* function returning more than one value*/
func swap(x, y string) (string, string) {
   return y, x
}

/*function that uses pointers*/
func usingPointers(){
	/* actual variable declaration */
	var a int = 20
	
	/* pointer variable declaration */
	var ip *int
	
	/* store address of a in pointer variable*/
	ip = &a
	fmt.Printf("Address of a variable: %x\n", &a  )
	
	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip )
	
	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip )
}

func usingPointerArray(){
	const MAX int = 3
	
	a := []int{10,100,200}
	var i int
	var ptr [MAX]*int;
	for  i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer. */
	}
	
	for  i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i,*ptr[i] )
	}
}



