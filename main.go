package main

import (
	"myapp/packageone"
)

var myVar = "This is a package level variable"

func main() {
	// variables
	var blockVar = "This is a block level variable"

	packageone.PrintMe(myVar, blockVar)
}
