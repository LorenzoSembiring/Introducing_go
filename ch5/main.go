package main

import "fmt"

func main() {
	// array
	x := [6]float64{90,91,92,93,94,95,}
	// x[0] = 91
	// x[1] = 92
	// x[2] = 93
	// x[3] = 94
	// x[4] = 95
	fmt.Println(x)
	var total float64
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	slice()
	sliceAppend()
	sliceCopy()

	maps()
	objectLike()

	smallestNumber()
}

func slice() {
	var y[]int		// empty slice
	z := make([]int, 5, 10)		// another slice, that associated with array that have length of 10
	fmt.Println(y, z)

	arr := [5]float64{1,2,3,4,5}
	a := arr[1:2]	// another way to create slice, take an array, and give params of start and end

	fmt.Println(arr, a)

}

func sliceAppend() {
	// append
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	
	fmt.Println(slice1, slice2)
}

func sliceCopy() {
	slice1 := []int{1,2,3}
	slice2 := make([]int,10)
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
}

func maps() {
	x := make(map[string]int)
	x["a"] = 10
	x["b"] = 20
	x["c"] = 30
	x["d"] = 40
	x["e"] = 50

	y := map[string]int {
		"a" : 1,
		"b" : 2,
		"c" : 3,
		"d" : 4,
		"e" : 5,
	}
	if name, ok := x["a"]; ok {
		fmt.Println(name, ok)
	}
	fmt.Println(y)
}

func objectLike() {
	courses := map[string]map[string]string{
		"CS01" : map[string]string{
			"lecturer" : "ARB",
			"student" : "10",
		},
		"CS02" : map[string]string{
			"lecturer" : "ABG",
			"student" : "20",
		},
		"CS03" : map[string]string{
			"lecturer" : "AHY",
			"student" : "25",
		},
	}

	fmt.Println(courses["CS01"]["student"])
}

func smallestNumber() {
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	var smallest int = x[0]
	for i := 0; i < len(x); i++ {
		if smallest > x[i] {
			smallest = x[i]
		}
	}
	fmt.Println(smallest)
}