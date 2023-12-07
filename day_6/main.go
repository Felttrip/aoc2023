package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time     int
	distance int
}

func main() {
	races := parseRaces("input_2.txt")
	fmt.Printf("%+v\n", parseRaces("input_1.txt"))
	numWaysToWin := -1
	for _, r := range races {
		w := getWinners(r)
		if numWaysToWin == -1 {
			numWaysToWin = len(w)
		} else {
			numWaysToWin = numWaysToWin * len(w)
		}
	}
	fmt.Println("part 1 num ways to win ", numWaysToWin)

	r := parseSingleRace("input_2.txt")
	w := getWinners(r)
	fmt.Println("part 2 num ways to win ", len(w))

}

func getWinners(r race) []int {
	winners := []int{}
	for timeHeld := 0; timeHeld < r.time; timeHeld++ {
		distanceCovered := timeHeld * (r.time - timeHeld)
		if distanceCovered > r.distance {
			winners = append(winners, distanceCovered)
		}
		distanceCovered = 0
	}
	return winners
}

func parseSingleRace(fileName string) race {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)
	//times
	s.Scan()
	time := parseSingleRaceLine(s.Text())
	s.Scan()
	distance := parseSingleRaceLine(s.Text())

	return race{
		time:     time,
		distance: distance,
	}

}

func parseSingleRaceLine(s string) int {
	s = strings.Replace(s, "Time:", "", -1)
	s = strings.Replace(s, "Distance:", "", -1)
	s = strings.ReplaceAll(s, " ", "")
	num, _ := strconv.Atoi(s)
	return num
}

func parseRaces(fileName string) []race {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)
	races := []race{}
	//times
	s.Scan()
	times := parseLine(s.Text())
	s.Scan()
	distances := parseLine(s.Text())

	for i, val := range times {
		r := race{
			time:     val,
			distance: distances[i],
		}
		races = append(races, r)
	}
	return races
}

/*
Time:      7  15   30
Distance:  9  40  200
*/
func parseLine(s string) []int {
	s = strings.Replace(s, "Time:", "", -1)
	s = strings.Replace(s, "Distance:", "", -1)
	sArr := strings.Split(s, " ")
	ret := []int{}
	for _, val := range sArr {
		if num, err := strconv.Atoi(val); err == nil {
			ret = append(ret, num)
		}
	}
	return ret
}
