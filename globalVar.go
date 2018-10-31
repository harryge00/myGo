package main

import (
	"github.com/harryge00/myGo/pkg1"
	"github.com/harryge00/myGo/pkg2"
)

func main() {
	pkg2.PrintName()
	pkg1.SetName("ddddd")
	pkg2.PrintName()
}