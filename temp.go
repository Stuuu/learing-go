package main 

import "fmt"

func main() {

	var wholeNumber int = 1
	var fractionalNumber float64 = 2.75
	var wholeNumber2 int = int(fractionalNumber) // ERROR
	var fractionalNumber2 float64 = float64(wholeNumber) // ERROR
	fmt.Println("wholeNumber2:", wholeNumber2)
	fmt.Println("fractionalNubmer2:", fractionalNumber2)


}