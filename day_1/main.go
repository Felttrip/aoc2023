package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	lines := readFile(os.Args[1])
	total := 0
	//for each row
	for _, l := range lines {
		calval := ""
		//forwards
		for _, r := range l {
			if unicode.IsDigit(r) {
				//convert rune to string
				calval = calval + string(r)
				break
			}
		}
		//backwards
		for i := len(l) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(l[i])); err == nil {
				calval = calval + string(l[i])
				break
			}
		}
		if num, err := strconv.Atoi(string(calval)); err == nil {
			total += num
		}
	}
	fmt.Println(total)

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
