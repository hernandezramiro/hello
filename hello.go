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
	//usingPointerArray()
	//usingPointerToPointer()
	usingPointersInFunctions()
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

func usingPointerToPointer(){
	var a int
	var ptr *int
	var pptr **int
	
	a = 3000
	
	/* take the address of var */
	ptr = &a
	
	/* take the address of ptr using address of operator & */
	pptr = &ptr
	
	/* take the value using pptr */
	fmt.Printf("Value of a = %d\n", a )
	fmt.Printf("Value available at *ptr = %d\n", *ptr )
	fmt.Printf("Value available at **pptr = %d\n", *pptr)
	
	fmt.Printf("----------------------------\n")
	
	fmt.Printf("Memory address of a = %x\n", &a)
	fmt.Printf("Memory address of ptr = %x\n", ptr)
	fmt.Printf("Memory address of pptr = %x\n", pptr)
}

func usingPointersInFunctions(){
	/* local variable definition */
	var a int = 100
	var b int= 200
	
	fmt.Printf("Before swap, value of a : %d\n", a )
	fmt.Printf("Before swap, value of b : %d\n", b )
	
	/* calling a function to swap the values.
	* &a indicates pointer to a ie. address of variable a and 
	* &b indicates pointer to b ie. address of variable b.
	*/
	funcWhichAllowsPointers(&a, &b);
	
	fmt.Printf("----------------------------\n")
	
	fmt.Printf("After swap, value of a : %d\n", a )
	fmt.Printf("After swap, value of b : %d\n", b )
}

func funcWhichAllowsPointers(x *int, y *int) {
   var temp int
   temp = *x    /* save the value at address x */
   *x = *y      /* put y into x */
   *y = temp    /* put temp into y */
}
