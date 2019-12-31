package main

import (
	"fmt"

	"go-build-trial/tools"
)

func main() {
	fmt.Println(tools.Add(10, 100))
	tools.LogS("Hello")
}
