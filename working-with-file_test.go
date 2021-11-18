package main

import "testing"

func TestIgnoredFile(test *testing.T) {
	var path string = "/Users/nome/.Trash"
	var ignored_ones = []string{".Trash"}
	isIgnored := isIgnoredFile(&path, ignored_ones)
	if !isIgnored {
		test.Fatalf("Arquivo deveria ter sido ignorado. Função retornou: %v", isIgnored)
	}
}

func TestNotIgnoredFile(test *testing.T) {
	var path string = "/Users/nome/nome.txt"
	var ignored_ones = []string{".Trash"}
	isIgnored := isIgnoredFile(&path, ignored_ones)
	if isIgnored {
		test.Fatalf("Arquivo não deveria ter sido ignorado. Função retornou: %v", isIgnored)
	}
}

func TestIgnoredFileFunctionWithoutPath(test *testing.T) {
	var path string = ""
	var ignored_ones = []string{".Trash"}
	isIgnored := isIgnoredFile(&path, ignored_ones)
	if !isIgnored {
		test.Fatalf("Arquivo deveria ter sido ignorado, já que está em branco. Função retornou: %v", isIgnored)
	}
}

func TestBigFile(test *testing.T) {
	var fileSize int64 = 100
	var bigSize int64 = 100
	isItBig := isItaBigFile(&fileSize, bigSize)
	if !isItBig {
		test.Fatalf("Deveria retornar true. Retornou: %v", isItBig)
	}
}

func TestSmallFile(test *testing.T) {
	var fileSize int64 = 5
	var bigSize int64 = 100
	isItBig := isItaBigFile(&fileSize, bigSize)
	if isItBig {
		test.Fatalf("Deveria retornar false. Retornou: %v", isItBig)
	}
}

func TestReturnFileFolderInformation(test *testing.T) {
	anyPath := "/Users/nome/.Trash"
	var size int64 = 100
	message := returnFileFolderInformation(&anyPath, &size)
	if message == "" {
		test.Fatalf("Retornou veio em branco!")
	}
}
