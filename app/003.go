package main

func sliceSum() {
	var numbers = []int{19, 86, 1, 12}
	var sum int

	// for _, number := range numbers {
	// 	sum += number
	// }

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	println(sum)
}

type Score struct {
	UserId string
	GameId string
	Value  int
}

func isOdd(number int) bool {
	return number%2 == 0
}

func oddOrEven2() {
	for i := 1; i <= 100; i++ {
		print(i)

		if isOdd(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}

func swap(x, y int) (int, int) {
	return y, x
}

func swap2(xp, yp *int) {
	*xp, *yp = *yp, *xp

	/*
	 * 以下ではスワップ出来ない
	 */
	// *xp = *yp
	// *yp = *xp
}

type MyInt int

func (n *MyInt) Inc() { *n++ }
