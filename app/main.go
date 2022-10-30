package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

type Hex2 = struct{ Hex }

func main() {
	var s Stringer
	var h = Hex(100)
	s = h
	fmt.Println(s.String())

	var h2 = Hex2{h}
	s = h2
	fmt.Println(s.String())
}
