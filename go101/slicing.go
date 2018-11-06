package main

import "fmt"

func main() {
	x := []int{9,22,99,32,76}
	fmt.Println(x)
	//fmt.Println(x[5])  panic: runtime error: index out of range

	fmt.Println("Using range!")
	for i , v := range x{
		fmt.Println(i,v)
	}

	fmt.Println("good old for loop ")
	for i:=0; i<len(x);i++{
		fmt.Println(i,x[i])
	}

	//Slicing a slice
	fmt.Println(x) //[9 22 99 32 76]
	fmt.Println(x[1 : 3]) //[22 99]From the specified(min) index until the max. Max won't be included.
	fmt.Println(x[0 : 2]) //[9 22]From the specified(min) index until the max. Max won't be included.
	fmt.Println(x[2 : 5]) //[99 32 76]From the specified(min) index until the max. Max won't be included.
//Appending

a := []int{98,71,622,791,19}
	fmt.Println(a)
	a =  append(a, 01) // Append is always added to the tail of list
	fmt.Println(a)
//...T -> It can take unlimited num of items of the type T
//T... -> Take something which has values in it
	x = append(x,a...) // a... is variadic parameter Also they should be of same type
	fmt.Println(x)


	fmt.Println("Deleting elements from a slice")
	d := []int{8,13,1,44,23}

	d = append(d[:2], d[4:]...)
	fmt.Println(d)

	fmt.Println("Make function to create slices") // This is memory optimised as we specify the size of the array and capacity when create it.

	mk := make([]int , 5 , 6)
	fmt.Println(mk)
	fmt.Println(len(mk))
	fmt.Println(cap(mk))
	fmt.Println(mk)
	mk[1] = 99
	mk[3] = 1
	fmt.Println(mk)
	//mk[5] = 3 //panic: runtime error: index out of range . We can't add new values as its out of range but we can append values
	mk = append(mk,9)
	fmt.Println(cap(mk))
	fmt.Println(mk)
	mk = append(mk,17) // When we add elements more than the capacity, it dynamically increases the size of the slice . Basically copies it a new memory allocation and stuff
	fmt.Println(cap(mk))
	fmt.Println(mk)


	fmt.Println("Multi Dimensional slices")

	proglang := []string{"c", "java" , "c#"}
	fmt.Println(proglang)
	funcprog := []string{"scala", "haskell" , "LISP"}
	fmt.Println(funcprog)
	allproglang := [][]string{proglang,funcprog}
	fmt.Println(allproglang)
	}
