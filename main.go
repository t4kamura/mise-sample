package main

import (
	"fmt"
	"os"
)

const (
	x = 1
	y = 2
)

func main() {
	env := os.Getenv("ENV_NAME")
	fmt.Println(fmt.Sprintf("The environment is %s", env))

	fmt.Println(fmt.Sprintf("%d + %d", x, y))
	r := calculate(x, y)
	fmt.Println(r)
}

func calculate(x, y int) int {
	return x + y
}
