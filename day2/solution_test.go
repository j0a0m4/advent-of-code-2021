package day2

import (
	"bufio"
	"fmt"
	assert "github.com/matryer/is"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSolution1(t *testing.T) {
	is := assert.New(t)
	inputs := Commands{
		Command{Forward, 5},
		Command{Down, 5},
		Command{Forward, 8},
		Command{Up, 3},
		Command{Down, 8},
		Command{Forward, 2},
	}

	horizontal, depth := 15, 10
	got := Solution1(inputs)
	is.Equal(horizontal, got.horizontal)
	is.Equal(depth, got.depth)
}

func TestRunSolution1(t *testing.T) {
	inputs := readInput()
	got := Solution1(inputs)
	result := got.depth * got.horizontal
	fmt.Println(result)
}

func TestSolution2(t *testing.T) {
	is := assert.New(t)
	inputs := Commands{
		Command{Forward, 5},
		Command{Down, 5},
		Command{Forward, 8},
		Command{Up, 3},
		Command{Down, 8},
		Command{Forward, 2},
	}

	horizontal, depth := 15, 60
	got := Solution2(inputs)
	is.Equal(horizontal, got.horizontal)
	is.Equal(depth, got.depth)
}

func TestRunSolution2(t *testing.T) {
	inputs := readInput()
	got := Solution2(inputs)
	result := got.depth * got.horizontal
	fmt.Println(result)
}

func TestPosition_Move(t *testing.T) {
	tests := []struct {
		name        string
		command     Command
		given, want *Position
	}{
		{
			name: "moves forward",
			command: Command{
				direction: Forward,
				units:     12,
			},
			given: &Position{
				horizontal: 3,
				depth:      5,
			},
			want: &Position{
				horizontal: 15,
				depth:      5,
			},
		},
		{
			name: "Moves down",
			command: Command{
				direction: Down,
				units:     5,
			},
			given: &Position{
				horizontal: 20,
				depth:      5,
			},
			want: &Position{
				horizontal: 20,
				depth:      10,
			},
		},
		{
			name: "Moves up",
			command: Command{
				direction: Up,
				units:     3,
			},
			given: &Position{
				horizontal: 20,
				depth:      5,
			},
			want: &Position{
				horizontal: 20,
				depth:      2,
				aim:        0,
			},
		},
	}
	for _, test := range tests {
		position := test.given
		command := test.command
		want := test.want
		t.Run(test.name, func(t *testing.T) {
			is := assert.New(t)
			position.Move(command)
			is.Equal(position, want)
		})
	}
}

func TestPosition_AimAndMove(t *testing.T) {
	tests := []struct {
		name        string
		command     Command
		given, want *Position
	}{
		{
			name: "Down increases aim",
			command: Command{
				direction: Down,
				units:     5,
			},
			given: &Position{
				aim:        0,
				horizontal: 0,
				depth:      0,
			},
			want: &Position{
				aim:        5,
				horizontal: 0,
				depth:      0,
			},
		},
		{
			name: "Up decreases aim",
			command: Command{
				direction: Up,
				units:     3,
			},
			given: &Position{
				aim:        5,
				horizontal: 0,
				depth:      0,
			},
			want: &Position{
				aim:        2,
				horizontal: 0,
				depth:      0,
			},
		},
		{
			name: "Forward increases horizontal and depth position",
			command: Command{
				direction: Forward,
				units:     3,
			},
			given: &Position{
				aim:        10,
				horizontal: 0,
				depth:      0,
			},
			want: &Position{
				aim:        10,
				horizontal: 3,
				depth:      30,
			},
		},
	}
	for _, test := range tests {
		position := test.given
		command := test.command
		want := test.want
		t.Run(test.name, func(t *testing.T) {
			is := assert.New(t)
			position.AimAndMove(command)
			is.Equal(position, want)
		})
	}
}

func readInput() Commands {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer closeFile(file)

	return scanInput(file)
}

func scanInput(reader io.Reader) (inputs Commands) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(command[1])
		inputs = append(inputs, Command{
			direction: command[0],
			units:     num,
		})
	}
	return
}

func closeFile(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		fmt.Printf("Error closing closer: %v\n", err)
	}
}
