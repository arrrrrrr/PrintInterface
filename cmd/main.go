package main

import (
	"PrintInterface/cmd/Printer"
	"fmt"
)

func printStuff(printer Printer.Printer) {
	printer.Print()
}

func main() {
	text := []string{
		"This is the header",
		"This is the body",
		"This is the footer",
	}

	pp := Printer.NewPlainPrinter(text)
	sp := Printer.NewSquarePrinter(text)

	printStuff(&pp)
	fmt.Println()
	printStuff(&sp)

}
