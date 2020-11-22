package main

import (
	"fmt"
)

func main() {
	// Declaration of array
	var numberList [5]int
	var stringList [2]string
	var float64List [4]float64

	//Defining value
	numberList = [5]int{98, 93, 77, 82, 83}
	stringList = [2]string{"maneesh", "xyz"}
	float64List = [4]float64{98.1, 93.2, 77.5, 82.4}
	fmt.Println(numberList)
	fmt.Println(stringList)
	fmt.Println(float64List)

	//Delaring and Defining value into array
	x := [5]int{98, 93, 77, 82, 83}

	//In this array determine size automatically
	y := [...]int{98, 93, 77, 82, 83}
	fmt.Println(x)
	fmt.Println(y)

	var buffer [5]string
	buffer[0] = "Robert"
	buffer[1] = "Michael"
	buffer[2] = "James"
	buffer[3] = "Aaron"
	buffer[4] = "Simon"

	fmt.Println(buffer)

	//Slice are just like array but it is Mutable and dynamic
	// It contain 1 2 3 element
	small_buffer := buffer[1:4]
	fmt.Println(small_buffer)

	//when Slice modify then underlying  array also modified
	small_buffer[1] = "Rodriguez"
	fmt.Println(buffer)
	fmt.Println(small_buffer)

	// Growing Slice
	// When modify g h then it not affect y
	g := []string{"John", "Paul"}
	h := []string{"George", "Ringo", "Pete"}
	y := append(g, h...)
	fmt.Println(y)

	// Map is similar to python dictionary
	var retval = map[string]int{}
	myHomeTown := "Brisbane By The Bay"
	var tokens = strings.Fields(myHomeTown)
	fmt.Println(tokens)
	for i := 0; i < len(tokens); i++ {
		retval[tokens[i]] = len(tokens[i])
	}
	fmt.Println(retval)

	//Map using Structure
	type User struct {
		name, client string
		age          int
	}

	var users = map[string]User{
		"settermjd": {name: "Matthew Setter", client: "Sitepoint", age: 38},
	}
	fmt.Println(users)

}
