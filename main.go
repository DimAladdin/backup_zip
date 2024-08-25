package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func name_file(path string) []string {
	word := strings.Split(path, "/")

	word_1 := word[len(word)-1]
	word_last := strings.Split(word_1, ".")
	return word_last
}

func checking_name(path string) string {
	var s string
	words := strings.Split(path, "/")
	for _, v := range words[len(words)-1] {
		if v == '.' {
			word := strings.Split(path, ".")
			s = word[0]
		} else {
			s = words[len(words)-1]
		}
	}
	return s
}

func path_archiv(path string) {

	file, err := os.Create(checking_name(path) + ".zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// f, err := w.Create(name_file(path)[0] + "." + name_file(path)[1])
		f, err := w.Create(checking_name(path))
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}
	err = filepath.Walk(path, walker)
	if err != nil {
		panic(err)
	}
}

func main() {
	path_archiv("/Users/dmitriiiks/project/go.mod")
	header, _ := tar.FileInfoHeader(fileInfo, "")
}
