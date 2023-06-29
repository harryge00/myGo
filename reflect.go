package main

import (
	"fmt"
	"reflect"
	"net/http"
	"time"
	"unsafe"
)

type Foo struct {
	s  string
	i  int
	hc *http.Client
}

func test123() {
	x := Foo{"hello", 2, http.DefaultClient}
	v := reflect.ValueOf(x)

	s := v.FieldByName("s")
	fmt.Printf("%T %v\n", s.String(), s.String())

	i := v.FieldByName("i")
	fmt.Printf("%T %v\n", i.Int(), i.Int())

	hc := v.FieldByName("hc")
	cli, ok := hc.Interface().(*http.Client)
	fmt.Printf("%v %v\n", cli, ok)
}

func test2() {
	var s = Foo{"hello", 2, &http.Client{
		Timeout: 123 * time.Millisecond,
	}}
	var i *http.Client

	rs := reflect.ValueOf(s)

	rs2 := reflect.New(rs.Type()).Elem()
	rs2.Set(rs)
	rf := rs2.Field(2)
	// Fails with "reflect.Value.UnsafeAddr of unaddressable value" because rs isn't addressable:
	// rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()

	// But we can work around that:
	rfNew := reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()

	ri := reflect.ValueOf(&i).Elem() // i, but writeable
	ri.Set(rfNew)

	hc, ok := ri.Interface().(*http.Client)
	fmt.Println(hc, ok)
}

func main() {
    test2()
}