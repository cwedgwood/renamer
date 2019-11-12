//

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	fFrom    = flag.String("from", "", "current suffix (to be changed)")
	fTo      = flag.String("to", "", "new suffix (to end up with)")
	fDryRun  = flag.Bool("dryrun", false, "Show what we will do")
	fVerbose = flag.Bool("verbose", false, "Show what we did")
)

func main() {
	flag.Parse()

	if *fFrom == "" || *fTo == "" {
		fmt.Printf("need from -from and -to\n")
		os.Exit(1)
	}

	for _, path := range flag.Args() {
		files, err := filepath.Glob(path)
		chk(err)

		for _, file := range files {
			newName := strings.Replace(file, *fFrom, *fTo, 1)
			if newName == file {
				continue
			}
			if *fDryRun {
				fmt.Printf("Dryrun: %s -> %s\n", file, newName)
				continue
			}
			err = os.Rename(file, newName)
			if err == nil && *fVerbose {
				fmt.Printf("Renamed %s -> %s\n", file, newName)
			}
		}
	}
}
