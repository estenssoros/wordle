package main

import (
	"log"
	"time"

	"github.com/estenssoros/wordle/cmd"
)

func run() error {
	return cmd.Execute()
}

func main() {
	start := time.Now()
	defer func() {
		log.Printf("took %v", time.Since(start))
	}()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
