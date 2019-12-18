// +build aix darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package main

import (
	"os"
	"path/filepath"
)

func findExecutables(name string) []string {
	var ps []string
	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		if dir == "" {
			dir = "."
		}
		p := filepath.Join(dir, name)
		if executableExists(p) {
			ps = append(ps, p)
		}
	}
	return ps
}

func executableExists(p string) bool {
	stat, err := os.Stat(p)
	if err != nil {
		return false
	}
	mode := stat.Mode()
	return !mode.IsDir() && mode&0o111 > 0
}
