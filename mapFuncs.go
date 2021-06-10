package main

import (
    "fmt"
)

func print1() error{
	fmt.Println("fun 1")
	return nil
}

func print2() error{
	fmt.Println("fun 2")
	return nil
}

func main() {
    mFuncs := map[string]func() error{}
	mFuncs["func1"] = print1
	mFuncs["func2"] = print2
	fmt.Println(mFuncs)

	mFuncs["func2"]()
}