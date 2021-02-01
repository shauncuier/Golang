package main

import "fmt"

func main() {
	//Arrays
	// var friutArr [3] string

	// // Assing Value

	// friutArr[0] = "Apple"
	// friutArr[1] = "Orange"
	// friutArr[2] = "Banana"

	// Decalre and Assing

	friutArr:=[2] string {"apple", "orange"}

	fmt.Println(friutArr)
	fmt.Println(friutArr[1])

	// Slice 

	fruitSlice := [] string {"apple", "orange","Banana","Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}

/*
#OutPut

[apple orange]
orange
[apple orange Banana Grape]
4
[orange Banana]
*/
