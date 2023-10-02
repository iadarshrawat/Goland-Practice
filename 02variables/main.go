package main

import "fmt"

// withoutVar := 45    throw error we have to use var globallly

const Cnst string = "hello world" // first letter of Cnst i.e, "C" denotes the public constant

func main() {
	var username string = "adarsh"
	fmt.Println(username)
	fmt.Printf("variable is of type:  %T \n", username)

	var isBoolean bool = true
	fmt.Println(isBoolean)
	fmt.Printf("variable is of type:  %T \n", isBoolean)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type:  %T \n", smallVal)

	var smallFloat float32 = 255.45654745678568 // in flaot 32, only 5 values are seen after the ".";
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type:  %T \n", smallFloat)

	var bigFloat float64 = 255.456547456785388456487495860 // in flaot 64, only 13 values are seen after the "." with round off;
	fmt.Println(bigFloat)
	fmt.Printf("variable is of type:  %T \n", bigFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable) // can not take any garbage value take 0
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "learnwithme.in" // if we can not give the datatype and lexer automatically detect
	fmt.Println(website)
	// website = 34     // but after that we cannot change the datatype (based on values)

	// without var keyword
	withOutVar := "learnwithme.in" // we also give use without "var" keyword but not globally
	fmt.Println(withOutVar)

	// use constant
	fmt.Println(Cnst)
	fmt.Printf("variable is of type: %T \n", Cnst)
}
