package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	docopt "github.com/docopt/docopt-go"
)

func parentDir(s string) string {
	lastIndex := strings.LastIndex(s, "/")

	if lastIndex == -1 {
		return s
	}

	return s[0:lastIndex]
}

type Candidate struct {
	path string
	name string
	size int64
}

func (c *Candidate) copyTo(dir string) error {
	dst := path.Join(dir, c.name)
	return copyFileContents(c.path, dst)
}

func cult(filename string) error {
	fmt.Printf("Searching for '%v' to clone ...\n", filename)

	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	root := parentDir(cwd)

	candidates := map[int64]Candidate{}
	counts := map[int64]int{}

	updateCounts := func(c Candidate) {
		count, ok := counts[c.size]

		if !ok {
			count = 0
		}

		count = count + 1
		counts[c.size] = count
	}

	absorb := func(path string, info os.FileInfo, err error) error {
		if info.Name() == filename {
			size := info.Size()

			c := Candidate{
				path: path,
				name: info.Name(),
				size: size,
			}

			candidates[c.size] = c
			updateCounts(c)
		}

		return nil
	}

	err = filepath.Walk(root, absorb)

	if err != nil {
		return err
	}

	highest := -1
	best := int64(-1)

	for size, count := range counts {
		if count > highest {
			highest = count
			best = size
		}
	}

	if highest == -1 {
		fmt.Printf("Could not find any qualifying copies of %v.  Sorry!\n", filename)
		return nil
	}

	winner, ok := candidates[best]

	if !ok {
		panic("WTF?!?!?")
	}

	return winner.copyTo(cwd)
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

	fmt.Println("BEFORE OPT")

	args, _ := docopt.Parse(usage, nil, true, "cargo 0.1", false)

	fmt.Println("BEFORE ARGS")
	if _, ok := args["cult"]; ok {
		filenames := args["<filename>"].([]string)

		for _, filename := range filenames {
			err := cult(filename)
			if err != nil {
				log.Printf("Failure on filename '%v': %v\n", filename, err)
			}
		}
	} else {
		fmt.Println(args)
	}
}
