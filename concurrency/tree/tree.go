// pseudo tree command
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func iter(path string, depth int) {
	_, file := filepath.Split(path)
	fmt.Printf("%s|-- %s\n", strings.Repeat("|  ", depth), file)

	dirs, _ := ioutil.ReadDir(path)
	for _, d := range dirs {
		if d.IsDir() {
			iter(filepath.Join(path, d.Name()), depth+1)
			continue
		}
		fmt.Printf("%s|-- %s\n", strings.Repeat("|  ", depth+1), d.Name())
	}
}

func main() {
	var (
	// depth = flag.Int("depth", 3, "depth to tree")
	)
	flag.Parse()
	path := flag.Arg(0)
	iter(path, 0)
}
