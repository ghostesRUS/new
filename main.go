package main

import "fmt"
import "os"

var (
	m0 = 33
)

func main() {
	fmt.Println("hello,\nworld")

	fmt.Println("1 + 1 =", 1+1)

	fmt.Println(m0)

	return

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
