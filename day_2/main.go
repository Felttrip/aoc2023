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
	number   int
	rounds   []round
	possible bool
}
type round struct {
	red   int
	green int
	blue  int
}

func main() {
	//parse games into structs
	g := parseGames("input_1.txt")
	fmt.Printf("%+v", g)
	//get valid games
	//sum up numeners
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
	return g
}
