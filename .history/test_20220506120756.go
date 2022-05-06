package main

import (
	// "flag"
	"fmt"
	"io"
	"strings"
)

var (
	arg1 string
	arg2 int
	arg3 bool
)

// func init() {
//       flag.StringVar(&arg1, "arg1", "defaultArg1", "arg1")
//       flag.IntVar(&arg2, "arg2", 12, "arg2")
//       flag.BoolVar(&arg3, "arg3", false, "arg3")
//       flag.Parse()
// }

func main() {
	var name string
	var age int
	n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)
}
