package main

import "fmt"

func greeteing(name string)  string {
	return "hello "+ name
}
func gitSum(num1, num2 int) int {
	return num1+num2
}

func main() {
    fmt.Println(greeteing("Brad"))
    fmt.Println(gitSum(4,7))
}
