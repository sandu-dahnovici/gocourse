package main

import "fmt"

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintf("Here is an error : %v", c.info)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	var c customErr
	foo(c)
}
