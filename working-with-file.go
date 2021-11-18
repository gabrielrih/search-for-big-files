package main

import (
	"strconv"
	"strings"
)

func isIgnoredFile(path *string, ignored_ones []string) bool {
	if *path == "" {
		return true
	}
	for _, name := range ignored_ones {
		if strings.Contains(*path, name) {
			return true
		}
	}
	return false
}

func isItaBigFile(fileSize *int64, bigSize int64) bool {
	return *fileSize >= bigSize
}

func returnFileFolderInformation(path *string, size *int64) string {
	size_in_mb := float64(int(*size) / 1024)
	return "File " + *path + " - " + strconv.FormatFloat(size_in_mb, 'f', 0, 32) + " " + "MB"
}
