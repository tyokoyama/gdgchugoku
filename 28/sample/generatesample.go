package main

import "fmt"

type Sample int

//go:generate stringer -type=Sample
const (
	hoge Sample = iota
	fuga
	piyo
)

func main() {
	fmt.Printf("%v\n", hoge)
	fmt.Printf("%v\n", fuga)
	fmt.Printf("%v\n", piyo)
}
