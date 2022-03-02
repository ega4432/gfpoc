package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "image.png", "Name of target filename")
	flag.Parse()
}

func main() {
	if _, err := os.Stat(filename); err != nil {
		handleError(err)
		return
	}

	if err := os.Remove(filename); err != nil {
		handleError(err)
		return
	}
	fmt.Println("Deleted!!")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
