package main

import (
     "flag"
     "fmt"
)

var (
     arg1 string
     arg2 int
     arg3 bool
)

func init() {
      flag.IntVar(&arg2, "arg2", 12, "arg2")
      flag.BoolVar(&arg3, "arg3", false, "arg3")
      flag.Parse()
}

func main() {
  fmt.Print()
	flag.String("I", "", "search `directory` for include files")
  fmt.Println("arg2 = ", arg2)
  fmt.Println("arg3 = ", arg3)
}