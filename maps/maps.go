package main

import "fmt"

func main() {

	// Define map
	/*emails:= make(map[string]string)

	//Assing Kv

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	*/


	// Decalre map ang key Value

	emails:= map[string]string{
		"Bob": "bob@gmail.com", 
		"Sharon": "sharon@gmail.com",
		"Mike": "mike@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete From map

	delete(emails, "Bob")

	fmt.Println(emails)
}