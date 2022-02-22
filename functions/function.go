package functions

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

// saved Input file
var Savedinput = []string{}

//CheckInput Checks that the input file is Valid
func CheckInput(file string) error {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return err
	}

	var input []string

	scanner := bufio.NewScanner(f)

	//Stores all strings in an array without \n
	for scanner.Scan() {
		input = append(input, scanner.Text())
		Savedinput = append(Savedinput, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func invalidFormat() error {
	if len(Savedinput) == 0 {
		return errors.New("Error: Empty file")
	}

	nAants, err := strconv.Atoi(Savedinput[0])
	if err != nil {
		return errors.New("Error: Ivalid number of Ants")
	}
	el
}
