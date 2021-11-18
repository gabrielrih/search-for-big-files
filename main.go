package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

const BIG_SIZE = 314572800 // 300 MB
const SOURCE_FOLDER = "/Users/gr14171/"

var ignored_ones = []string{"./", ".Trash", ".git", ".DS_Store", "Library"}

func main() {

	fmt.Println("Searching in", SOURCE_FOLDER, "...")

	// Navega pelos arquivos / diretórios a partir de um diretório base
	// Reference: https://pkg.go.dev/path/filepath#Walk
	countFile := 0
	err := filepath.Walk(SOURCE_FOLDER, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir // ignora
		}
		if !isIgnoredFile(&path, ignored_ones) {
			size := info.Size()
			if isItaBigFile(&size, BIG_SIZE) {
				isItFolder := info.IsDir()
				if !isItFolder {
					fmt.Println(returnFileFolderInformation(&path, &size))
					countFile += 1
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if countFile > 0 {
		fmt.Println("Big files quantity: ", countFile)
	} else {
		fmt.Println("No big files were founded!")
	}
}
