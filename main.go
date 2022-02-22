package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Println("You have too many or too few arguments. Try passing the name of a file with ants ")
		return
	}

	err := functions.CheckInput

}
