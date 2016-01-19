package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	for _, entry1 := range hash {
		for size, list1 := range entry1 {
			if len(list1) <= 1 || size <= 0 {
				continue
			}
			for i, path := range list1 {
				if i == 0 {
					fmt.Printf("\nrem \"%s\"\n", path)
				} else {
					fmt.Printf("del \"%s\"\n", path)
				}
			}
		}
	}
}
