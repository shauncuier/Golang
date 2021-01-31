// Variable 

package main

import "fmt"

func main() {

// Using Variable
var title = "Mr"

// Using Variable Single Line
var firstName, lastName = "Jone", "Doy"

//var lastName = "Doy"

//Bool Variable
var isCool = true

// ShortHand
age := 24

// ShortHand Multiple Var
size , email := 3.2,"test@abc.com"
// Print/ Show Variable
fmt.Println(title, firstName, lastName)
fmt.Println("He is", age, "years old.")
fmt.Println("Size:", size, "email:", email)

//Data Type Check
fmt.Printf("%T\n", title)
fmt.Printf("%T\n", age)
fmt.Printf("%T\n", isCool)
fmt.Printf("%T\n", email)
fmt.Printf("%T\n", size)
}

/*
# Output

Mr Jone Doy
He is 24 years old.
Size: 3.2 email: test@abc.com
string
int
bool
string
float64
*/