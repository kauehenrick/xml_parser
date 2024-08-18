package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	readFiles(".")
}

func readFiles(dir string)  {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		endsWith := strings.HasSuffix(file.Name(), ".xml")

		if endsWith {
			fmt.Println(file.Name())
		}
	}
}
