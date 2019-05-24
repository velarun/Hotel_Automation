package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func (h *Handler) Movements() {

	exit := make(chan string)

	for {
		errors := make(chan error, 0)
		input := make(chan string, 0)
		fmt.Println("\nChecking Sensor Inputs...")
		go h.ParsingInput(input, errors)
		select {
		case err := <-errors:
			fmt.Println(err)
		case <-input:
			h.BuildHandler.PrintHotelStructure(h.InputHandler.Floor, h.InputHandler.MainCorridor, h.InputHandler.SubCorridor)
		case <-exit:
			os.Exit(0)
		}
	}
}

func (h *Handler) ParsingInput(input chan string, erro chan error) {

	char, err := reader.ReadString('\n')
	if err != nil {
		erro <- err
		return
	}

	if len(char) == 0 {
		erro <- errors.New("Sensor Input Error")
		return
	}

	var matches []string
	var flag bool
	r, _ := regexp.Compile("No.*Floor\\s([\\d]+),.*Sub\\scorridor\\s([\\d]+).*minute.*")

	if r.MatchString(string(char)) {
		matches = r.FindStringSubmatch(string(char))
	} else {
		r, _ = regexp.Compile(".*Floor\\s([\\d]+),.*Sub\\scorridor\\s([\\d]+).*")
		if r.MatchString(string(char)) {
			matches = r.FindStringSubmatch(string(char))
			flag = true
		} else {
			erro <- errors.New("Sensor Input Error")
			return
		}
	}

	if matches[1] != "" && matches[2] != "" {
		floor, _ := strconv.Atoi(matches[1])
		subCorridor, _ := strconv.Atoi(matches[2])

		if !h.MovementHandler(floor-1, subCorridor, flag) {
			erro <- errors.New("Sensor Input Error")
			return
		}
	} else {
		erro <- errors.New("Sensor Input Error")
		return
	}

	input <- "Done"
}

func (h *Handler) MovementHandler(floor, subCorridor int, flag bool) bool {

	var x = h.BuildHandler.SubCorridor[floor].CorridorCount[subCorridor]

	if flag {
		if !x.Light {
			x.Light = true
		} else {
			return false
		}
	} else {
		if x.Light {
			x.Light = false
		} else {
			return false
		}
	}

	h.BuildHandler.SubCorridor[floor].CorridorCount[subCorridor] = x

	return h.PowerRestriction(floor, subCorridor)
}
