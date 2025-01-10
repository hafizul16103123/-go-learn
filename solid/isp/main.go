package main

import "fmt"

/*
Interfaces should be client-specific and not force implementations of unused methods.
*/

// Printer interface
type Printer interface {
	Print()
}

// Scanner interface
type Scanner interface {
	Scan()
}

// Specific Printer
type PDFPrinter struct{}

func (pp PDFPrinter) Print() {
	fmt.Println("Printing to PDF...")
}

// Specific Scanner
type DocumentScanner struct{}

func (ds DocumentScanner) Scan() {
	fmt.Println("Scanning document...")
}

func main() {
	// Example
	printer := PDFPrinter{}
	scanner := DocumentScanner{}

	printer.Print()
	scanner.Scan()
}
