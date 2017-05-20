package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/frickiericker/acol/printers"
	"github.com/frickiericker/acol/utils"
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	RowMajor bool `cli:"r" usage:"use row-major ordering"`
	Spacing  int  `cli:"s" usage:"space between columns" dft:"2"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		in := os.Stdin
		out := os.Stdout
		printer := makePrinter(argv)
		return run(in, out, printer)
	})
}

func makePrinter(argv *argT) printers.Interface {
	width, err := getTerminalWidth()
	if err != nil {
		fmt.Fprintln(os.Stderr, "warning:", err)
		width = 0
	}
	spacing := argv.Spacing

	if argv.RowMajor {
		return printers.NewDenseRowMajor(width, spacing)
	}
	return printers.NewDenseColumnMajor(width, spacing)
}

func run(in io.Reader, out io.Writer, printer printers.Interface) error {
	items, err := utils.ReadLines(in)
	if err != nil {
		return fmt.Errorf("reading from input: %s", err)
	}
	cells := printers.MakeCells(items)
	return printer.Print(out, cells)
}

/*
Validate validates command-line arguments.
*/
func (argv *argT) Validate(ctx *cli.Context) error {
	if argv.Spacing < 0 {
		return errors.New("space between columns cannot be negative")
	}
	return nil
}
