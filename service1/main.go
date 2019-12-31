package main

import (
	"fmt"
	"github.com/juxuny/go-build-trial/tools"
)

func main() {
	fmt.Println(tools.Add(10, 100))
	tools.LogS("Hello")
}
