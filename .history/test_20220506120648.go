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
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	for {
		n, err := fmt.Fscanln(r, &a, &b, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
	}
}