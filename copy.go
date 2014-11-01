package main

import (
	"fmt"
	"io"
	"os"
)

// http://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
func copyFileContents(src, dst string) (err error) {
	fmt.Printf("Copying: %v\n", src)
	fmt.Printf("     to: %v\n", dst)

	in, err := os.Open(src)

	if err != nil {
		return
	}

	defer in.Close()

	out, err := os.Create(dst)

	if err != nil {
		return
	}

	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	if _, err = io.Copy(out, in); err != nil {
		return
	}

	err = out.Sync()

	return
}
