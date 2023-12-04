package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

/*
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
*/

type schematic [][]item

type item struct {
	char    rune
	checked bool
}

func (s schematic) String() string {
	out := ""
	for y, _ := range s {
		for _, i := range s[y] {
			out += string(i.char)
		}
		out += "\n"
	}
	return out
}

func main() {
	//read into 2d array
	s := parseSchamtic("input_2.txt")
	// fmt.Println(s)
	fmt.Println(processSchematic(s))
}

func (s schematic) shouldProcess(x, y int) bool {
	/*
		in schematic
		 0 <= x < len(s[y])
		 0 <= y < len(s)
	*/
	inSchematic := x >= 0 && x < len(s[y]) && y >= 0 && y < len(s)

	return inSchematic && !s[y][x].checked
}

// givin a starting possition find the number
func (s schematic) getNumber(x, y int) int {

	curr := x
	//move all the way left until theres no number
	for s.shouldProcess(curr-1, y) && unicode.IsNumber(s[y][curr-1].char) {
		curr--
	}

	//from the start build up the number untill theres no numbers
	numStr := ""
	for s.shouldProcess(curr, y) && unicode.IsNumber(s[y][curr].char) {
		numStr = numStr + string(s[y][curr].char)
		s[y][curr].checked = true
		curr++
	}

	num, _ := strconv.Atoi(numStr)
	//return numbers
	return num
}
func (s schematic) lookAround(x, y int) int {
	//account for index out of bounds
	rowUp := y - 1
	rowDown := y + 1
	columnLeft := x - 1
	columnRight := x + 1
	total := 0
	//diag up left
	if s.shouldProcess(columnLeft, rowUp) && unicode.IsNumber(s[rowUp][columnLeft].char) {
		total += s.getNumber(columnLeft, rowUp)
	}
	//up
	if s.shouldProcess(x, rowUp) && unicode.IsNumber(s[rowUp][x].char) {
		total += s.getNumber(x, rowUp)
	}
	//diag up right
	if s.shouldProcess(columnRight, rowUp) && unicode.IsNumber(s[rowUp][columnRight].char) {
		total += s.getNumber(columnRight, rowUp)
	}

	//left
	if s.shouldProcess(columnLeft, y) && unicode.IsNumber(s[y][columnLeft].char) {
		total += s.getNumber(columnLeft, y)
	}
	//right
	if s.shouldProcess(columnRight, y) && unicode.IsNumber(s[y][columnRight].char) {
		total += s.getNumber(columnRight, y)
	}
	//diag down left
	if s.shouldProcess(columnLeft, rowDown) && unicode.IsNumber(s[rowDown][columnLeft].char) {
		total += s.getNumber(columnLeft, rowDown)
	}
	//down
	if s.shouldProcess(x, rowDown) && unicode.IsNumber(s[rowDown][x].char) {
		total += s.getNumber(x, rowDown)
	}
	//diag down right
	if s.shouldProcess(columnRight, rowDown) && unicode.IsNumber(s[rowDown][columnRight].char) {
		total += s.getNumber(columnRight, rowDown)
	}
	return total
}

func processSchematic(s schematic) int {
	total := 0
	for y := range s {
		for x, i := range s[y] {
			if !unicode.IsNumber(i.char) && i.char != '.' {
				num := s.lookAround(x, y)
				// fmt.Printf("symbol %v possition %v, %v, num %v\n", string(i.char), x, y, num)
				total += num
			}
		}
	}
	return total
}

func parseSchamtic(fileName string) schematic {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)
	schem := schematic{}
	for s.Scan() {
		line := []item{}

		for _, r := range s.Text() {
			i := item{
				char:    r,
				checked: false,
			}
			line = append(line, i)
		}
		schem = append(schem, line)
	}
	return schem
}
