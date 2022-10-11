package main

import (
	"fmt"
	"math/rand"
	"time"
)

func oddOrEven() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d-偶数\n", i)
		} else {
			fmt.Printf("%d-奇数\n", i)
		}
	}

	for i := 0; i <= 100; i++ {
		switch {
		case i%2 == 0:
			fmt.Printf("%d-偶数\n", i)
		default:
			fmt.Printf("%d-奇数\n", i)
		}
	}
}

func omikuji() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	n := rand.Intn(6) // 0-5

	switch n + 1 {
	case 6:
		print("大吉")
	case 5, 4:
		print("中吉")
	case 3, 2:
		print("吉")
	default:
		print("凶")
	}
}
