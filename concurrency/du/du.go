// pseudo du
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walk(dir string, fileSizes chan<- int64) error {
	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, d := range dirs {
		if d.IsDir() {
			walk(filepath.Join(dir, d.Name()), fileSizes)
			continue
		}
		fileSizes <- d.Size()
	}
	return nil
}

func main() {
	flag.Parse()
	dirs := flag.Args()
	if len(dirs) == 0 {
		dirs = []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, dir := range dirs {
			if err := walk(dir, fileSizes); err != nil {
				fmt.Fprintf(os.Stderr, "du: walkdir failed: %s\n", err)
				continue
			}
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
