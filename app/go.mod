module example.com/app

go 1.19

replace example.com/greetings => ../greetings

replace example.com/greeting_module => ../greeting_module

require example.com/greeting_module v1.0.0

require (
	github.com/tenntenn/greeting v1.0.0 // indirect
	github.com/tenntenn/greeting/v2 v2.2.1 // indirect
)
