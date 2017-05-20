package printers

/*
Common interface of printers. A printer displays given items (slice of strings)
in a specific format.
*/
type Interface interface {
	Print(items []string)
}
