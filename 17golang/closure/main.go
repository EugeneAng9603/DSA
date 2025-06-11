package main

import "fmt"

var globalVar = 0

func plus() func() {
	sum := 0
	return func() {
		println(sum)
		sum += 1
	}
}

func pluss() {
	globalVar += 1
	println(globalVar)
}
func main() {
	plusFunc := plus()
	plusFunc()
	plusFunc()
	plusFunc()
	plusFunc()
	fmt.Print("globalVar here")
	pluss()
	pluss()
	pluss()
}
