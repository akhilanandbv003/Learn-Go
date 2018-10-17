package main

import "fmt"

func main() {

var z string = `my name is "within 
double quotes"`  // When we have ` everything within it is parsed as raw text including several lines

var x = 9 // Auto inference of data type . We need not specify var x int specifically
fmt.Println(z)
	fmt.Printf("%T\n",z) // %T tells the type
	fmt.Println(x)
	fmt.Printf("%T\n",x) // %T tells the type

	s:= fmt.Sprintf("%X\t%b\t%#x", x,x,x) // Will assign this string formatting to a variable s
	fmt.Println(s)

	typesExample()
	
}

func typesExample()  {
	type mytype string
	var name mytype = " Akhil"
	fmt.Println(name)
}

/*
Go is a statically typed language
It will infer some of the data types if you are assigning value to it

Similar printf we sprintf and fprintf .
*/