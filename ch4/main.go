package main

import "fmt"

func main() {
	divisibleBy3()
	fizzBuzz()

	i := 1
	// way to write for loop in go, I think its similar to while loop
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}
	// another way, the common way
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Println("even")
		} else if i % 3 == 0 {
			fmt.Println("divisible by 3")
		} else {
			fmt.Println("odd and not divisible by 3")
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		switch i {
		case 0: fmt.Println("zero")
		case 1: fmt.Println("one")
		case 2: fmt.Println("two")
		case 3: fmt.Println("three")
		case 4: fmt.Println("four")
		}
	}

}

func divisibleBy3() {
	for i := 0; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzBuzz(){
	for i := 0; i <= 100; i++ {
		if (i % 3 == 0) && (i % 5 == 0) {
			fmt.Println("Fizz Buzz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		}
	}
}