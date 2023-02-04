package main

import "fmt"

const (
	NoError      = iota
	GeneralError = iota
	InternaError = iota
)

func main() {
	fmt.Print(NoError, GeneralError, InternaError)
}
