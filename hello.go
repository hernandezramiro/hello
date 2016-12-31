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
	//usingPointersInFunctions()
	//usingStructs()
	//sendingStructsToFunctions()
	//usingSlices()
	//usingSlicesSecondPart()
	usingRanges()
}

type Books struct {
   title string
   author string
   subject string
   book_id int
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

func usingStructs(){
	
	/*
	type Books struct {
		title string
		author string
		subject string
		book_id int
	}*/
	
	/* Declare Book1 of type Book */
	var Book1 Books
	
	/* Declare Book2 of type Book */
	var Book2 Books

	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407
	
	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700
	
	/* print Book1 info */
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	fmt.Printf( "Book 1 author : %s\n", Book1.author)
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)
	
	fmt.Printf("----------------------------\n")
	
	/* print Book2 info */
	fmt.Printf( "Book 2 title : %s\n", Book2.title)
	fmt.Printf( "Book 2 author : %s\n", Book2.author)
	fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
	fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}

func sendingStructsToFunctions(){
	
	/* Declare Book1 of type Book */
	var Book1 Books
	
	/* Declare Book2 of type Book */
	var Book2 Books
	
	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407
	
	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700
	
	/* print Book1 info */
	funcWhichAllowsStructs(Book1)
	
	fmt.Printf("----------------------------\n")
	
	/* print Book2 info */
	funcWhichAllowsStructsMixingPointers(&Book2)
}

func funcWhichAllowsStructs(book Books) {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}

func funcWhichAllowsStructsMixingPointers(pointerToBook *Books) {
	fmt.Printf( "Pointer Book title : %s\n", pointerToBook.title);
	fmt.Printf( "Pointer Book author : %s\n", pointerToBook.author);
	fmt.Printf( "Pointer Book subject : %s\n", pointerToBook.subject);
	fmt.Printf( "Pointer Book book_id : %d\n", pointerToBook.book_id);
}

func usingSlices(){
	/* create a slice */
	numbers := []int{0,1,2,3,4,5,6,7,8}   
	printSlice(numbers)
	
	/* print the original slice */
	fmt.Println("numbers ==", numbers)
	
	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	
	/* missing lower bound implies 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])
	
	/* missing upper bound implies len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])
	
	numbers1 := make([]int,0,5)
	printSlice(numbers1)
	
	/* print the sub slice starting from index 0(included) to index 2(excluded) */
	number2 := numbers[:2]
	printSlice(number2)
	
	/* print the sub slice starting from index 2(included) to index 5(excluded) */
	number3 := numbers[2:5]
	printSlice(number3)
}

func usingSlicesSecondPart(){
	var numbers []int
	printSlice(numbers)
	
	/* append allows nil slice */
	numbers = append(numbers, 0)
	printSlice(numbers)
	
	/* add one element to slice*/
	numbers = append(numbers, 1)
	printSlice(numbers)
	
	/* add more than one element at a time*/
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)
	
	/* create a slice numbers1 with double the capacity of earlier slice*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	
	/* copy content of numbers to numbers1 */
	copy(numbers1,numbers)
	printSlice(numbers1)   
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func usingRanges(){
	/* create a slice */
	numbers := []int{0,1,2,3,4,5,6,7,8} 
	
	/* print the numbers */
	for i:= range numbers {
		fmt.Println("Slice item",i,"is",numbers[i])
	}
	
	fmt.Printf("----------------------------\n")
	
	/* create a map*/
	countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo"}
	
	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
	
	fmt.Printf("----------------------------\n")
	
	/* print map using key-value*/
	for country,capital := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",capital)
	}
}

