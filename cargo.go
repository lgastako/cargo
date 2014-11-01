package main

import (
	"fmt"

	docopt "github.com/docopt/docopt-go"
)

func cult() {
}

func main() {
	usage := `cargo.

Usage:
  cargo cult <filename>...
  cargo -h | --help
  cargo --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "cargo 0.1", false)

	if _, ok := arguments["cult"]; ok {
	} else {
		fmt.Println(arguments)
	}
}
