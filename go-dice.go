package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func rollDice(input string) {
	tokens := strings.Split(input, "+")
	first := strings.Split(tokens[0], "d")

	numRolls, err := strconv.Atoi(first[0])
	if err != nil {
		fmt.Println("Invalid syntax: item before 'd' must be a number.")
		return
	}

	numSides, err := strconv.Atoi(first[1])
	if err != nil {
		fmt.Println("Invalid syntax: item after 'd' must be a number.")
		return
	}
	if numSides < 1 {
		fmt.Println("Invalid option: the number of sides must be at least 1.")
		return
	}

	res := 0
	for x := 0; x < numRolls; x++ {
		res += rand.Intn(numSides) + 1
	}
	if len(tokens) > 1 {
		bonus, err := strconv.Atoi(tokens[1])
		if err == nil {
			res += bonus
		} else {
			fmt.Println("Invalid syntax: item after '+' must be a number.")
			return
		}
	}
	fmt.Println(res)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) > 1 {
		rollDice(os.Args[1])
	} else {
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print(">>>> ")
			input, err := reader.ReadString('\n')
			if err == nil {
				input = strings.Replace(strings.TrimSpace(input), " ", "", -1)
			} else {
				fmt.Println("Error while reading input")
				continue
			}

			if len(input) > 0 {
				rollDice(input)
			}
		}
	}
}
