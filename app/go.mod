module example.com/app

go 1.19

replace example.com/greetings => ../greetings

replace example.com/greeting_module => ../greeting_module

replace example.com/imageconverter => ../imageconverter

require example.com/imageconverter v0.0.0-00010101000000-000000000000
