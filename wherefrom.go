package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)
	dups := flag.Bool("dups", false, "Include duplicate entries")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}

	name := flag.Arg(0)
	if strings.ContainsRune(name, os.PathSeparator) {
		log.Fatal("name must not contain path separators")
	}

	seen := make(map[string]struct{})
	for _, p := range findExecutables(name) {
		if _, ok := seen[p]; !ok || *dups {
			fmt.Println(p)
		}
		seen[p] = struct{}{}
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `usage: wherefrom [flags] <name>

Wherefrom prints every executable in $PATH matching name in the order
they are listed. Flags:
`)
	flag.PrintDefaults()
	os.Exit(2)
}
