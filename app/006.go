package main

import "fmt"

type _Stringer interface {
	String() string
}

type _Message string

func (m _Message) String() string {
	return fmt.Sprintf("%v", string(m))
}

type _Hex int

func (h _Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

type _Activated bool

func (a _Activated) String() string {
	return fmt.Sprintf("%v", bool(a))
}
