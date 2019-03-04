package main

import (
	"flag"
	"fmt"
	"frequentlyUsedWords/internal/frequentlyusedwords"
	"log"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "internal/frequentlyusedwords/test_fixtures/mobydick.txt", "text file you wish parse")
	flag.Parse()
	if len(file) == 0 {
		log.Fatal("You must specify a file")
	}

	out, err := frequentlyusedwords.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
