package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

//go:embed rights
var unRights embed.FS

func main() {
	if len(os.Args) == 1 {
		printFiles()
		os.Exit(0)
	}

	fileName := "UN-" + os.Args[1] + "-rights.txt"

	data, err := os.ReadFile("rights/" + fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}

func printFiles() {
	fmt.Println("contents:")
	err := fs.WalkDir(unRights, "rights",
		func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				_, fileName, _ := strings.Cut(path, "/")
				fmt.Println(fileName)
			}
			return nil
		})

	if err != nil {
		log.Fatal(err)
	}
}
