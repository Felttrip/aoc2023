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
func (s schematic) clone() schematic {
	duplicate := make(schematic, len(s))
	for i := range s {
		duplicate[i] = make([]item, len(s[i]))
		copy(duplicate[i], s[i])
	}
	return duplicate
}

func main() {
	//read into 2d array
	s := parseSchamtic("input_2.txt")
	sg := s.clone()
	fmt.Printf("Part 1: %v\n", processSchematic(s))
	fmt.Printf("Part 2: %v\n", sumGears(sg))

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
func (s schematic) lookAround(x, y int) []int {
	//account for index out of bounds
	rowUp := y - 1
	rowDown := y + 1
	columnLeft := x - 1
	columnRight := x + 1
	found := []int{}
	//diag up left
	if s.shouldProcess(columnLeft, rowUp) && unicode.IsNumber(s[rowUp][columnLeft].char) {
		found = append(found, s.getNumber(columnLeft, rowUp))
	}
	//up
	if s.shouldProcess(x, rowUp) && unicode.IsNumber(s[rowUp][x].char) {
		found = append(found, s.getNumber(x, rowUp))
	}
	//diag up right
	if s.shouldProcess(columnRight, rowUp) && unicode.IsNumber(s[rowUp][columnRight].char) {
		found = append(found, s.getNumber(columnRight, rowUp))
	}

	//left
	if s.shouldProcess(columnLeft, y) && unicode.IsNumber(s[y][columnLeft].char) {
		found = append(found, s.getNumber(columnLeft, y))
	}
	//right
	if s.shouldProcess(columnRight, y) && unicode.IsNumber(s[y][columnRight].char) {
		found = append(found, s.getNumber(columnRight, y))
	}
	//diag down left
	if s.shouldProcess(columnLeft, rowDown) && unicode.IsNumber(s[rowDown][columnLeft].char) {
		found = append(found, s.getNumber(columnLeft, rowDown))
	}
	//down
	if s.shouldProcess(x, rowDown) && unicode.IsNumber(s[rowDown][x].char) {
		found = append(found, s.getNumber(x, rowDown))
	}
	//diag down right
	if s.shouldProcess(columnRight, rowDown) && unicode.IsNumber(s[rowDown][columnRight].char) {
		found = append(found, s.getNumber(columnRight, rowDown))
	}
	// fmt.Println(numPartNumbers, needed)
	return found
}

func processSchematic(s schematic) int {
	total := 0
	for y := range s {
		for x, i := range s[y] {
			if !unicode.IsNumber(i.char) && i.char != '.' {
				nums := s.lookAround(x, y)
				// fmt.Printf("symbol %v possition %v, %v, num %v\n", string(i.char), x, y, num)
				for _, num := range nums {
					total += num
				}
			}
		}
	}
	return total
}
func sumGears(s schematic) int {
	total := 0
	for y := range s {
		for x, i := range s[y] {
			if i.char == '*' {
				nums := s.lookAround(x, y)
				if len(nums) == 2 {
					total += nums[0] * nums[1]

				}
				// fmt.Printf("symbol %v possition %v, %v, num %v\n", string(i.char), x, y, num)
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
