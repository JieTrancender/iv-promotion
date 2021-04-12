package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

var flagvar int
var newFlagVar int

func main() {
	pflag.IntVar(&flagvar, "flagVar", 0, "flag var")
	pflag.IntVar(&newFlagVar, "newFlagVar", 0, "flag var")
	pflag.Parse()

	fmt.Println("flagvar has value ", flagvar)
	newFlagVar, err := pflag.CommandLine.GetInt("newFlagVar")
	if err != nil {
		fmt.Println("lookup fail")
	}

	fmt.Println("flag lookup", newFlagVar)
	fmt.Println(pflag.Lookup("newFlagVar"))
}
