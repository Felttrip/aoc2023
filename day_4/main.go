package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
*/

type card struct {
	winners    []int
	numbers    []int
	score      float64
	numWinners int
}
type cards []card
type cardMap map[int]int

func main() {
	cs := parseCards("input_2.txt")
	scoredCards := scoreCards(cs)
	fmt.Printf("part 1 %+v\n", sumScores(scoredCards))
	fmt.Printf("part 2 %+v\n", fanOut(scoredCards))

}

func fanOut(scoredCards cards) int {
	cardMap := map[int]int{}
	for currentCardIndex, c := range scoredCards {
		//every card has atleast one copy
		cardMap[currentCardIndex]++
		for processingCardIndex := currentCardIndex + 1; processingCardIndex <= currentCardIndex+c.numWinners && processingCardIndex < len(scoredCards); processingCardIndex++ {
			cardMap[processingCardIndex] += cardMap[currentCardIndex]
		}

	}
	total := 0
	for _, val := range cardMap {
		total += val
	}

	return total
}

func sumScores(cs cards) int {
	total := 0
	for _, c := range cs {
		total += int(c.score)
	}
	return total
}
func scoreCards(cs cards) cards {
	scoredCards := cards{}
	for _, c := range cs {
		scoredCards = append(scoredCards, calcScore(c))
	}
	return scoredCards
}
func calcScore(c card) card {
	scoredCard := card{
		winners:    c.winners,
		numbers:    c.numbers,
		score:      0,
		numWinners: 0,
	}
	matches := -1
	for _, w := range c.winners {
		if contains(c.numbers, w) {
			matches++
		}
	}

	if matches != -1 {
		scoredCard.score = math.Pow(2, float64(matches))
	}
	scoredCard.numWinners = matches + 1
	return scoredCard

}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func parseCards(fileName string) cards {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)
	cs := cards{}
	for s.Scan() {
		c := parseLine(s.Text())
		cs = append(cs, c)
	}
	return cs
}

func parseLine(line string) card {
	//Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	nameAndNums := strings.Split(line, ":")
	winnersAndNums := strings.Split(nameAndNums[1], "|")
	c := card{
		winners: []int{},
		numbers: []int{},
		score:   0,
	}
	for _, num := range strings.Split(winnersAndNums[0], " ") {

		if n, err := strconv.Atoi(num); err == nil {
			c.winners = append(c.winners, n)
		}
	}
	for _, num := range strings.Split(winnersAndNums[1], " ") {
		if n, err := strconv.Atoi(num); err == nil {
			c.numbers = append(c.numbers, n)
		}
	}
	// fmt.Printf("card %+v ", c)
	return c
}
