package Printer

import (
	"fmt"
	"strings"
)

type Printer interface {
	Print()
}

type SquarePrinter struct {
	text []string
}

type PlainPrinter struct {
	text []string
}

func NewSquarePrinter(text []string) SquarePrinter {
	return SquarePrinter{
		text,
	}
}

func NewPlainPrinter(text []string) PlainPrinter {
	return PlainPrinter{
		text,
	}
}

func (sp *SquarePrinter) Print() {
	maxLen := calculateMaxLength(sp.text)
	borderSlice := make([]string, maxLen+2)
	for i := range borderSlice {
		borderSlice[i] = "-"
	}
	borderTemplate := fmt.Sprintf("|%s|", strings.Join(borderSlice, ""))
	rowTemplate := fmt.Sprintf("| %%-%ds |\n", maxLen)

	fmt.Println(borderTemplate)
	for _, v := range sp.text {
		fmt.Printf(rowTemplate, v)
	}
	fmt.Println(borderTemplate)
}

func (sp *PlainPrinter) Print() {
	for _, v := range sp.text {
		fmt.Println(v)
	}
}

func calculateMaxLength(text []string) int {
	var max_len int
	for _, v := range text {
		if max_len < len(v) {
			max_len = len(v)
		}
	}
	return max_len
}
