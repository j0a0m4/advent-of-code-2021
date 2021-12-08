package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var inputNumbers, bingoCards = readInput()

func Solution1() int {
	return FirstToWin(inputNumbers, bingoCards)
}

func Solution2() int {
	return LastToWin(inputNumbers, bingoCards)
}

func FirstToWin(numbers []int, cards [][][]int) int {
	var chosenNumbers []int
	for i, number := range numbers {
		chosenNumbers = append(chosenNumbers, number)
		if i < 5 {
			continue
		}
		for _, card := range cards {
			if CheckBingo(chosenNumbers, card) {
				return CalculateScore(chosenNumbers, card)
			}
		}
	}
	return -1
}

func LastToWin(numbers []int, cards [][][]int) int {
	var chosenNumbers []int
	for i, number := range numbers {
		chosenNumbers = append(chosenNumbers, number)
		if i < 5 {
			continue
		}
		for j, card := range cards {
			if len(cards) == 1 && CheckBingo(chosenNumbers, card) {
				return CalculateScore(chosenNumbers, card)
			} else if len(cards) == 1 {
				continue
			} else if CheckBingo(chosenNumbers, card) {
				cards = append(cards[:j], cards[j+1:]...)
			}
		}
	}
	return -1
}

func CalculateScore(numbers []int, card [][]int) int {
	var sum int
	set := make(map[int]bool)

	for _, row := range card {
		for _, value := range row {
			set[value] = true
			sum += value
		}
	}

	for _, value := range numbers {
		if _, found := set[value]; found {
			sum -= value
		}
	}

	return sum * numbers[len(numbers)-1]
}

func CheckBingo(chosenNumbers []int, card [][]int) bool {
	for _, combination := range append(card, Rotate(card)...) {
		if Bingo(chosenNumbers, combination) {
			return true
		}
	}
	return false
}

func Bingo(chosenNumbers, bingoNumbers []int) bool {
	set := make(map[int]bool)
	for _, value := range chosenNumbers {
		set[value] = true
	}

	for _, value := range bingoNumbers {
		if _, found := set[value]; !found {
			return false
		}
	}

	return true
}

func Rotate(inputs [][]int) (rotated [][]int) {
	rowSize := len(inputs[0])
	for i := 0; i < rowSize; i++ {
		var buffer []int
		for _, row := range inputs {
			buffer = append(buffer, row[i])
		}
		rotated = append(rotated, buffer)
	}
	return
}

func SplitRow(row string) (res []string) {
	for _, num := range strings.Split(row, " ") {
		if num == "" {
			continue
		}
		res = append(res, num)
	}
	return
}

func ReadRow(row []string) ([]int, error) {
	var numbers []int
	for _, s := range row {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func ReadCard(card []string) ([][]int, error) {
	var bingo [][]int
	for _, c := range card {
		row, err := ReadRow(SplitRow(c))
		if err != nil {
			return nil, err
		}
		bingo = append(bingo, row)
	}
	return bingo, nil
}

func readInput() ([]int, [][][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer closeFile(file)

	return scanInput(file)
}

func scanInput(reader io.Reader) (numbers []int, inputs [][][]int) {
	var memory [][]int
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if numbers == nil {
			numbers = scanNumbers(scanner)
			continue
		} else if len(scanner.Text()) == 0 && memory == nil {
			continue
		}
		if len(scanner.Text()) == 0 && len(memory) != 0 {
			//memory = append(memory, Rotate(memory)...)
			inputs = append(inputs, memory)
			memory = [][]int{}
		} else {
			r, _ := ReadRow(SplitRow(scanner.Text()))
			memory = append(memory, r)
		}
	}
	return
}

func scanNumbers(scanner *bufio.Scanner) (numbers []int) {
	line := strings.Split(scanner.Text(), ",")
	for _, s := range line {
		num, _ := strconv.Atoi(s)
		numbers = append(numbers, num)
	}
	return
}

func closeFile(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		fmt.Printf("Error closing closer: %v\n", err)
	}
}
