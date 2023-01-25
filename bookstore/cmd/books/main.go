package main

import "fmt"

func main() {
	basicTypeRules()
	defaultValues()
	customValues()
}

// Basic type rules
func basicTypeRules() {
	var title, author string
	var copies int
	var inStock bool
	var royaltyPercentage float64
	title = "The Happiness Project"
	copies = 99
	author = "Dave Normal"
	inStock = true
	royaltyPercentage = 12.5
	fmt.Println(title)
	fmt.Println(author)
	fmt.Println(copies)
	fmt.Println(inStock)
	fmt.Println(royaltyPercentage)
}

// Variables of given types have default values in GO, so we should use var declaration
// if we want a variable to be initialised using a default value
func defaultValues() {
	var noValInt int
	// It is 0 for ints and floats
	fmt.Println("int default is ", noValInt)
	var noValFloat float64
	fmt.Println("float default is ", noValFloat)
	// It is empty string for strings
	var noValString string
	fmt.Println("string default is \"", noValString, "\"")
	fmt.Println("Length of string default is ", len(noValString))
	// It is false for booleans
	var noValBool bool
	fmt.Println("bool default is ", noValBool)
}

// And if we want to initialise a variable to a different value other than
// the default, we should use the syntax ":=" instead to combine var declaration
// and assignment in one go. It means less typing as the GO compiler infers the
// type from the assigned value given.
//
// Note: Whole number assignment means a variable gets treated as an int
func customValues() {
	customInt := 1
	fmt.Println("custom int is ", customInt)
	customFloat := 1.2
	fmt.Println("custom float is ", customFloat)
	// It is empty string for strings
	customString := "Custom String innit"
	fmt.Println("custom string is ", customString)
	// It is false for booleans
	customBool := true
	fmt.Println("custom bool is ", customBool)
}
