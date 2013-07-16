package main

import (
	"fmt"
	"go-stfl"	
)

func init() {
	stfl.Init()
}

func main() {
	f := stfl.Create("<example.stfl>")

	f.Set("value_a", "This is a little")
	f.Set("value_b", "test for STFL!")

	event := ""
	for event != "ESC" {
		event = f.Run(0)
	}

	stfl.Reset()

	fmt.Printf("A: %s\n", f.Get("value_a"))
	fmt.Printf("B: %s\n", f.Get("value_b"))
	fmt.Printf("C: %s\n", f.Get("value_c"))
}
