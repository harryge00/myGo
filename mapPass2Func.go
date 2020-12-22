package main
import "fmt"

func test_map(m map[string]string){
	fmt.Printf("inner: %v, %p\n",m, m)
	m["a"]="11"
	fmt.Printf("inner: %v, %p\n",m, m)
}

func main() {
	m := map[string]string{
		"a":"1",
		"b":"2",
		"c":"3",
	}

	fmt.Printf("outer: %v, %p\n",m, m)
	test_map(m)
	fmt.Printf("outer: %v, %p\n",m, m)
}
