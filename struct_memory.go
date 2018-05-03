package main

import (
	"fmt"
	"reflect"
	"unsafe"
)


type I interface {
	Talk(func())
	SameMethod()
}

type Person struct {
	I
  Name string
  Sss string
}

type empty struct {

}

type checkbool struct {
	a bool
	b bool
	c string
}

type checkbool2 struct {
	a bool
	b string
	c bool
}

func (e *empty) Talk(str string) string {
	return "hello world " + str
}

func (p Person) Talk(callback func()) {
  callback()
  fmt.Println("Hi, my name is", p.Name)
}

func (p Person) SameMethod() {
  fmt.Println("just the same")
}

type Android struct {
  Person
  Model string
  Name string
}

func getPerson(name string) I {
	return &Person {
		Name: name,
	}
}

func getAndroid(name string) I {
	return &Android{
		Person: Person {
			Name: name,
		},
	}
}

func (a *Android) Internal() {
	fmt.Printf("Android Internal")
}

const (
	Invalid int = iota + 2 // type is invalid

	// predeclared types
	Bool
	Int
	Int8
	)

func main() {
	fmt.Println(Invalid, Bool, Int, Int8)

	p := Person{
		Name: "sbcd",	
	}
	p.Talk(func(){return})
	a := Android{
		Person: p,
	}
	
	a.Talk(func(){return})
	fmt.Println(a.Name)
	a.Name = "new name"
	fmt.Println(a.Name)
	a.Talk(func(){return})
	
	typAndroid := reflect.TypeOf(a)
	fmt.Printf("Struct of Android is %d bytes long\n", typAndroid.Size())
	fmt.Printf("Sizeof Android is %d bytes long\n", unsafe.Sizeof(a))
	// We can run through the fields in the structure in order
	n := typAndroid.NumField()
	for i := 0; i < n; i++ {
	        field := typAndroid.Field(i)
	        fmt.Printf("%s at offset %v, size=%d, align=%d\n",
	            field.Name, field.Offset, field.Type.Size(), 
	            field.Type.Align())
	}

	typPersion := reflect.TypeOf(p)
	fmt.Printf("Struct of Person is %d bytes long\n", typPersion.Size())
	fmt.Printf("Sizeof Person is %d bytes long\n", unsafe.Sizeof(p))
	// We can run through the fields in the structure in order
	n = typPersion.NumField()
	for i := 0; i < n; i++ {
	        field := typPersion.Field(i)
	        fmt.Printf("%s at offset %v, size=%d, align=%d\n",
	            field.Name, field.Offset, field.Type.Size(), 
	            field.Type.Align())
	}

	typeEmpty := reflect.TypeOf(empty{})
	fmt.Printf("typeEmpty is %d bytes long\n", typeEmpty.Size())
	fmt.Printf("Sizeof is %d bytes long\n", unsafe.Sizeof(empty{}))

	var i I
	typeI := reflect.TypeOf(i)
	if typeI != nil {
		fmt.Printf("Non nil typeI is %d bytes long\n", typeI.Size())
		fmt.Printf("Sizeof I is %d bytes long\n", unsafe.Sizeof(i))		
	}
	i = p
	i.SameMethod()
	typeI = reflect.TypeOf(i)
	fmt.Printf("typeI is %d bytes long\n", typeI.Size())
	fmt.Printf("Sizeof I is %d bytes long\n", unsafe.Sizeof(i))

	pi := getAndroid("adaf")
	pi.SameMethod()
	pi.Talk(func(){fmt.Printf("lala")})

	pp := &Person{}
	pp.SameMethod()
}