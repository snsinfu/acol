package main

import (
	"fmt"
	"os"

	"github.com/frickiericker/acol/printers"
	"github.com/frickiericker/acol/utils"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	items, err := utils.ReadLines(os.Stdin)
	if err != nil {
		return fmt.Errorf("reading from standard input:", err)
	}

	printer, err := makePrinter()
	if err != nil {
		return fmt.Errorf("creating printer:", err)
	}

	printer.Print(items)
	return nil
}

func makePrinter() (printers.Interface, error) {
	width, err := getTerminalWidth()
	if err != nil {
		fmt.Fprintln(os.Stderr, "warning:", err)
		width = 0
	}
	spacing := 2
	printer := printers.NewDenseColumnMajor(width, spacing)
	return printer, nil
}
