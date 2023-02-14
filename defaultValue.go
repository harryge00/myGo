// Golang program to assign
// default values to a struct
// using tags

package main

import (
	"fmt"
	"reflect"
)

// declaring a person struct
type Person struct {

	// setting the default value
	// of name to "geek"
	name string `default:"geek"`
	age  int    `default:16`
}

func default_tag(p Person) string {

	// TypeOf returns type of
	// interface value passed to it
	typ := reflect.TypeOf(p)

	// checking if null string
	if p.name == "" {

		// returns the struct field
		// with the given parameter "name"
		f, _ := typ.FieldByName("name")

		// returns the value associated
		// with key in the tag string
		// and returns empty string if
		// no such key in tag
		p.name = f.Tag.Get("default")
	}

	return fmt.Sprintf("%s", p.name)
}

// main function
func main() {

	// prints out the default name
	fmt.Println("Default name is:", default_tag(Person{}))
	var p Person
	fmt.Println("2 Default name is:", p.name)
}
