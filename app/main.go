package main

import (
	"fmt"

	"example.com/greeting_module"
)

func main() {
	message1 := greeting_module.Do1()
	fmt.Println(message1)
	message2 := greeting_module.Do2()
	fmt.Println(message2)
}
