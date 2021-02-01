package main

import "fmt"

func main() {
	x:= 17
	y:= 15

	//  If else

	if (x < y) {
		fmt.Printf("%d is less than %d\n", x, y)
	}else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if

	color := "blue"

	if color == "red" {
		fmt.Println("Color is Red")
	} else if color =="blue" {
		fmt.Println("Color is Blue")
	} else {
		fmt.Println("Color is not Red or Blue")
	}

	// Swich
	switch color {
	case "red":
		fmt.Println("Color is Red")
	case "blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is not Red or Blue")
	}
	
}






/*
# Output


*/