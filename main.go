package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	shortPtr := flag.Bool("s", false, "Shorthand")

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	if *shortPtr {
		s := strings.Split(currDir, "/")
		fmt.Println(s[len(s)-1])
	} else {
		fmt.Println(currDir)
	}
}
