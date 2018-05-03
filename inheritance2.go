package main

import "fmt"

type I interface {
	Talk(func())
	SameMethod()
}

type Person struct {
	I
  Name string
}

func (p *Person) Talk(callback func()) {
  callback()
  fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) SameMethod() {
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

func main() {
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
	
	pi := getAndroid("adaf")
	pi.SameMethod()
	pi.Talk(func(){fmt.Printf("lala")})
}