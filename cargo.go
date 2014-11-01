package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"

	docopt "github.com/docopt/docopt-go"
)

type Finder func(map[string]interface{}) (string, error)
type FinderMap map[string]Finder

type Candidate struct {
	path string
	name string
	size int64
}

func (c *Candidate) copyTo(dir string) error {
	dst := path.Join(dir, c.name)
	return copyFileContents(c.path, dst)
}

func cult(root string, filenames []string) error {
	fmt.Printf("Searching under: %v\n", root)
	fmt.Printf("            for: %v\n", filenames)

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
		for _, filename := range filenames {
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
		}

		return nil
	}

	err := filepath.Walk(root, absorb)

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
		fmt.Printf("Somehow we found no highest count.  Maybe there were no matches at all?")
		return nil
	}

	winner := candidates[best]

	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	return winner.copyTo(cwd)
}

func atFinder(args map[string]interface{}) (string, error) {
	dir := args["<dir>"].(string)
	return dir, nil
}

func fromFinder(args map[string]interface{}) (string, error) {
	dir := args["<dir>"].(string)

	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	d := cwd
	for true {
		parentDir, fn := path.Split(d)
		// TODO: parentDir[0:-1] ?

		if parentDir == d {
			break
		}

		if fn == dir {
			return parentDir, nil
		}

		d = parentDir
	}

	return "", errors.New(fmt.Sprintf("Could not find %v in ancestor paths.", dir))
}

func upFinder(args map[string]interface{}) (string, error) {
	levelStr := args["<levels>"].(string)

	levels, err := strconv.Atoi(levelStr)

	if err != nil {
		return "", err
	}

	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	for i := 0; i < levels; i++ {
		dir, _ := path.Split(dir)
		fmt.Printf("1 => %v\n", dir)
		dir = dir[0 : len(dir)-2]
		fmt.Printf("2 => %v\n", dir)
	}

	return "", nil
}

func finderFromArgs(args map[string]interface{}) Finder {
	finders := FinderMap{
		"at":   atFinder,
		"from": fromFinder,
		"up":   upFinder,
	}

	for cmd, finder := range finders {
		_, ok := args[cmd]
		if ok {
			return finder
		}
	}

	return nil
}

func main() {
	usage := `cargo.

Usage:
  cargo at <dir> <filename>...
  cargo from <dir> <filename>...
  cargo up <levels> <filename>...
  cargo -h | --help
  cargo --version

Options:
  -h --help           Show this screen.
  -e --early-out=<n>  Short-circuit at <n> identical copies.
  --version           Show version.`

	args, _ := docopt.Parse(usage, nil, true, "cargo 0.1", false)

	finder := finderFromArgs(args)

	if finder == nil {
		fmt.Println("How did you end up with no finder?  You sly dog, you...")
		return
	}

	dir, err := finder(args)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	filenames := args["<filename>"].([]string)

	cult(dir, filenames)
}
