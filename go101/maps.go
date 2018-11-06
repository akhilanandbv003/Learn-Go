package main

import "fmt"

func main() {
	map1 := map[string] int{
		"java" :5,
		"python" :10,
	}

	fmt.Println(map1)
	fmt.Println(map1["java"])
	fmt.Println(map1["scala"])//This returns 0 ->default value for int even if its not in the map

	fmt.Println("With error handling")
	v,ok :=map1["scala"]
	fmt.Println(v)
	fmt.Println(ok) // This returns false if the key is not in the list

	fmt.Println("idiomatic way of error handling") //Called comma-ok idiom
	if v,ok := map1["java"]; ok{
		fmt.Println(v)
	}

	fmt.Println("Add new values to Map")
	map1["scala"] = 7
	map1["javascript"] = 2
	fmt.Println(map1)

	fmt.Println("Looping over Map")
	for k, v := range map1{
		fmt.Println(k)
		fmt.Println(v)
	}

	fmt.Println("Delete from Map")
	delete(map1,"javascript") // Deleting a value that doesn't exist totally works but just returns nothing
	fmt.Println(map1)

}
