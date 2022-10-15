package greeting_module

import (
	"time"

	greeting1 "github.com/tenntenn/greeting"
	greeting2 "github.com/tenntenn/greeting/v2"
)

func Do1() string {
	return greeting1.Do()
}

func Do2() string {
	return greeting2.Do(time.Now())
}
