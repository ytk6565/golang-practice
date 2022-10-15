package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	var hex Hex = 100
	f := Hex.String
	fmt.Printf("%T\n%s\n", f, f(hex))
}
