package main

import "fmt"

var x string = "hello"
var (
	a = "a"
	b = "b"
	c = "c"
)
func main() {
	x = x + " world"
	fmt.Println(x)

	fahrenheitToCelcius()
}

func fahrenheitToCelcius() {

	fmt.Println("enter fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	res := (input-32)*5/9
	fmt.Print("celcius: ", res)
}