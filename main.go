package main

import (
	"log"

	"github.com/CodeGophercises/secrets/cmd"
)

func main() {
	Must(cmd.Execute())
}

func Must(err error) {
	if err != nil {
		log.Fatalf("Something went wrong: %s\n", err.Error())
	}
}
