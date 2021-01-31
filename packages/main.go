package main
//C:\Users\SHAUN\Desktop\go_practice\practice\packages\strutil
import (
	"fmt"
	"math"
	"strutil"

)

func main ()  {
	
	//Math
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(25))

	// Reverse
	fmt.Println(strutil.Reverse("dlroW olleH"))
}