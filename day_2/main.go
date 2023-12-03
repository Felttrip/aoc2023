package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
*/

type games []game

type game struct {
	number int
	rounds []round
}
type round struct {
	red   int
	green int
	blue  int
}

func main() {
	//parse games into structs
	g := parseGames("input_2.txt")
	// fmt.Printf("%+v\n\n", g)

	vg := getValidGames(g)
	// fmt.Printf("%+v\n\n", vg)

	fmt.Printf("Valid Game sums %d\n\n", sumGames(vg))
	fmt.Printf("Power %d\n\n", getTotalPower(g))

	//get valid games
	//sum up numeners
}

func sumGames(gs games) int {
	total := 0
	for _, g := range gs {
		total += g.number
	}
	return total
}

func getTotalPower(gs games) int {
	total := 0
	for _, g := range gs {
		total += power(getMaxColors(g))
	}
	return total
}

func power(red, green, blue int) int {
	return red * green * blue
}

func getMaxColors(g game) (int, int, int) {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	for _, r := range g.rounds {
		if r.red > maxRed {
			maxRed = r.red
		}
		if r.green > maxGreen {
			maxGreen = r.green
		}
		if r.blue > maxBlue {
			maxBlue = r.blue
		}
	}
	return maxRed, maxGreen, maxBlue
}

// 12 red cubes, 13 green cubes, and 14 blue cubes?
func getValidGames(gs games) games {
	validGames := games{}
	for _, g := range gs {
		isValid := true
		for _, r := range g.rounds {
			if r.red > 12 || r.green > 13 || r.blue > 14 {
				// fmt.Printf("invalid game %v, round %+v\n\n", g.number, r)
				isValid = false
				break
			}
		}
		if isValid {
			validGames = append(validGames, g)
		}
	}
	return validGames
}

func parseGames(fileName string) games {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)
	g := games{}
	for s.Scan() {
		g = append(g, parseLine(s.Text()))
	}
	return g
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func parseLine(line string) game {
	initial := strings.Split(line, ":")
	gameNumSplit := strings.Split(initial[0], " ")
	number, _ := strconv.Atoi(gameNumSplit[1])
	g := game{
		number: number,
	}
	//3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	roundStr := strings.Split(initial[1], ";")
	for _, rounds := range roundStr {
		r := round{
			red:   0,
			green: 0,
			blue:  0,
		}
		//3 blue, 4 red
		colorCountStr := strings.Split(rounds, ",")

		for _, colorAndCount := range colorCountStr {
			//3 blue
			final := strings.Split(colorAndCount, " ")
			switch final[2] {
			case "red":
				if num, err := strconv.Atoi(final[1]); err == nil {
					r.red += num
				}
				break
			case "blue":
				if num, err := strconv.Atoi(final[1]); err == nil {
					r.blue += num
				}
				break
			case "green":
				if num, err := strconv.Atoi(final[1]); err == nil {
					r.green += num
				}
				break
			}
		}
		g.rounds = append(g.rounds, r)
	}

	return g
}
