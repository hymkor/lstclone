package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattn/go-isatty"
	"github.com/zetamatta/nyagos/dos"
)

var hash = map[string]map[int64][]string{}

func eachfile(path string, info os.FileInfo, err error) error {
	if info == nil || info.IsDir() {
		return nil
	}
	basename := filepath.Base(path)
	entry1, entry1ok := hash[basename]
	if !entry1ok {
		entry1 = make(map[int64][]string)
		hash[basename] = entry1
	}
	list1, list1ok := entry1[info.Size()]
	if list1ok {
		list1 = append(list1, path)
	} else {
		list1 = []string{path}
	}
	entry1[info.Size()] = list1
	return nil
}

func main() {
	for _, arg1 := range os.Args[1:] {
		filepath.Walk(arg1, eachfile)
	}

	isatty1 := isatty.IsTerminal(os.Stdout.Fd())
	var newline []byte
	if isatty1 {
		newline = []byte{'\n'}
	} else {
		newline = []byte{'\r', '\n'}
	}

	for _, entry1 := range hash {
		for size, list1 := range entry1 {
			if len(list1) <= 1 || size <= 0 {
				continue
			}
			for i, path := range list1 {
				var s string
				if i == 0 {
					os.Stdout.Write(newline)
					s = fmt.Sprintf("rem \"%s\"", path)
				} else {
					s = fmt.Sprintf("del \"%s\"", path)
				}
				if isatty1 {
					os.Stdout.WriteString(s)
				} else {
					if data, err := dos.UtoA(s); err != nil {
						fmt.Fprintf(os.Stderr, "%s: %s\n", path, err)
						break
					} else {
						// trim '\0'
						length := len(data)
						if length >= 2 {
							os.Stdout.Write(data[:length-1])
						}
					}
				}
				os.Stdout.Write(newline)
			}
		}
	}
}
