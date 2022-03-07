package functions

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// saved Input file
var Savedinput = []string{}

// Ants - storts the nuber of Ants
var Ants int

// names of rooms
var nameRooms []string

// names of room plus coordinates
var cordRooms []string

var indexStart int = 0
var indexEnd int = 0

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

func antCheck() error {
	if len(Savedinput) == 0 {
		return errors.New("Error: Empty file")
	}

	nAants, err := strconv.Atoi(Savedinput[0])
	Ants = nAants
	if err != nil {
		return errors.New("Error: Ivalid number of Ants")
	}
	if nAants <= 0 {
		return errors.New("Error: Invalid number of Ants too few")
	}
	if nAants > 100000 {
		return errors.New("Error: inValid Number of Ants too many")
	}
	return nil
}

// Checking for the presence of a start and end room and the number of such rooms
func roomCheck() error {
	isstart := false
	isend := false
	cstart := 0 // Count of Start identifer (##start)
	cend := 0   // Count of End  indtifier (##end)

	// Checking for the presence of a start and end room and the number of such rooms
	for _, v := range Savedinput {
		if v == "##start" {
			cstart++
			isstart = true
		}
		if v == "##end" {
			cend++
			isend = true
		}
	}

	if !isstart || !isend {
		return errors.New("ERROR: invalid data format, no start room or no end room found")
	}
	if cstart > 1 || cend > 1 {
		return errors.New("ERROR: invalid data format, more than 1 start or end rooms found")
	}
	// Checking for duplicate rooms
	// Checking for the presence of start and end rooms, as well as links between rooms
	for i, v := range Savedinput {
		if v == "##start" {
			indexStart = i
			if i+1 != len(Savedinput) {
				roomstart := Savedinput[indexStart+1]
				if strings.HasPrefix(roomstart, " ") || strings.HasPrefix(roomstart, "L") || strings.HasPrefix(roomstart, "#") {
					return errors.New("Error: invalid format, invalid start room")
				}
			}
		}
		if v == "##end" {
			indexEnd = i
			if i+1 != len(Savedinput) {
				roomend := Savedinput[indexEnd+1]
				if strings.HasPrefix(roomend, " ") || strings.HasPrefix(roomend, "L") || strings.HasPrefix(roomend, "#") {
					return errors.New("Error: invalid format, invalid end room")
				}
			}
		}
	}
}
