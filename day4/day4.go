package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	numbers, bingoCards := readInput()
	fmt.Println(numbers)
	fmt.Println(bingoCards)
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
