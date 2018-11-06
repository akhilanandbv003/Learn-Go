package main

import "fmt"

func main() {
	a:= 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n",a) //int
	fmt.Printf("%T\n" ,&a)//*int pointer to an int

	fmt.Println(*&a)

	b := &a // b has the address of a's value
	*b = 43
	fmt.Println("Now a is :" ,a)


whypointer()
}

func whypointer()  {
	// Pointers come in handy when we have to pass around heavy objects. instead we pass pointers
	x := 88
	fmt.Println(x)
	fmt.Println(&x)
	foo(&x) // We are passing the address of x
	fmt.Println(x) // This will be updated by the foo function
	fmt.Println(&x)
}

func foo(inp *int)  {
	fmt.Println(inp)
	fmt.Println(*inp)
	*inp = 9882 // Updating the value at the address
	fmt.Println(inp)
	fmt.Println(*inp)

}