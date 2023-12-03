package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/dlclark/regexp2"
	"github.com/gijsbers/go-pcre"
)

func main() {
	r := pcre.MustCompile("(one|two|three|four|five|six|seven|eight|nine)|\\d", 0)
	lines := readFile(os.Args[1])
	total := 0
	//for each row
	for _, l := range lines {
		calval := ""
		matcher := r.MatcherString(l, 0)
		matches := matcher.ExtractString()
		fmt.Println(matches)
		calval = calval + normalize(matches[0])
		calval = calval + normalize(matches[len(matches)-1])
		if num, err := strconv.Atoi(string(calval)); err == nil {
			total += num
		}
	}
	fmt.Println(total)

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

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
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
