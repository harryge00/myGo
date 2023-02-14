package main

import "fmt"

type S struct {
	nums []int
	name string
}

func Do(arr *[]int) {
	*arr = make([]int, 4)
	(*arr)[0] = 123
}

type R struct {
	name string
	ss   S
}

func main() {

	r := R{}
	s := S{
		name: "aaaa",
		nums: []int{3, 5},
	}
	r.ss = s
	fmt.Println(r)
	s.name = "abc"
	s.nums[0] = 999
	fmt.Println("r", r, "s", s)
	test()
}

func test() {
	type Person struct {
		Name string
		Age  int
	}

	alice1 := &Person{"Alice", 30}
	alice2 := *alice1
	fmt.Println(alice1 == &alice2) // => false, they have different addresses

	alice3 := (*alice1)
	fmt.Println(&alice3 == alice1)

	alice2.Age += 10
	fmt.Println(alice1, alice2)

	pp := []Person{
		Person{
			Name: "abc",
			Age:  33,
		},
	}
	fmt.Println(pp)
	p0 := pp[0]
	p0.Name = "def"
	fmt.Println(p0, pp)
}
