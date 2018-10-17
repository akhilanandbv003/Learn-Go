package main // Always should be main for it to be runnable

import "fmt"

func main()  {
	//Declare and assign simultaneously
	var y = 43
	z := 43
	v , err := fmt.Println("Hello World")
	println(v)
	println(err)
	println(y)
	println(z)

}


/* Notes
:= Called as short Declaration operator
Go throws error if there are any unused variables

Difference between var and := is more like global and local variable.
We can use var as a global variable. var's scope is beyond the func defnition.
But := is like a local variable. Can't access it outside func.
*/