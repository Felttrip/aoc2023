package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var values = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"1", "2", "3", "4", "5", "6", "7", "8", "9",
}

type metaNumber struct {
	value     string
	possition int
}

func main() {
	lines := readFile(os.Args[1])
	total := 0
	//for each row
	for _, l := range lines {
		first := metaNumber{
			value:     "",
			possition: 9999999,
		}
		last := metaNumber{
			value:     "",
			possition: -1,
		}

		for _, value := range values {
			newFirstPos, newLastPos := findFirstAndLast(l, value)
			if newFirstPos < first.possition && newFirstPos != -1 {
				fmt.Printf("first found %v at possition %v \n", value, newFirstPos)
				first.possition = newFirstPos
				first.value = value
			}
			if newLastPos > last.possition {
				fmt.Printf("last found %v at possition %v \n", value, newLastPos)
				last.possition = newLastPos
				last.value = value
			}
		}
		fmt.Printf("finals: first %+v last %+v\n", first, last)
		calval := normalize(first.value) + normalize(last.value)

		if num, err := strconv.Atoi(string(calval)); err == nil {
			total += num
		}
	}
	fmt.Println(total)

}

func findFirstAndLast(s, subString string) (int, int) {
	return strings.Index(s, subString), strings.LastIndex(s, subString)
}

func normalize(num string) string {
	switch num {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		//if it didnt match then it was just a regular number
		return num
	}
}

func readFile(fileName string) []string {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)
	l := []string{}
	for s.Scan() {
		l = append(l, s.Text())
	}
	return l
}
