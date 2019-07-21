package main 

import "fmt"

var a = "a"

func main() {
	var b = "b"
	{
		var c = "c"
		{ 
		var d = "d"
			fmt.Println(a, b, c, d)
		}
		fmt.Println(a, b, c, d)
	}
	fmt.Println(a, b, c, d)
}