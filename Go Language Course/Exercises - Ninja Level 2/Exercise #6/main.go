package main

import "fmt"

const (
	now = 2021 + iota
	afterOneYear
	afterTwoYears
	afterThreeYears
	afterFourYears
)

func main() {
	fmt.Println(now)
	fmt.Println(afterOneYear)
	fmt.Println(afterTwoYears)
	fmt.Println(afterThreeYears)
	fmt.Println(afterFourYears)
}
