package main

import (
	"PrintInterface/cmd/Printer"
	"fmt"
)

func main() {
	text := []string{
		"This is the header",
		"This is the body",
		"This is the footer",
	}

	pp := Printer.NewPlainPrinter(text)
	sp := Printer.NewSquarePrinter(text)

	pp.Print()
	fmt.Println()
	sp.Print()
}
