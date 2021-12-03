package day3

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
	inputs := Report{
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
	}

	gamma, epsilon := 22, 9
	power := gamma * epsilon
	got, err := Solution1(inputs)
	is.Equal(got, power)
	is.NoErr(err)
}

func TestRunSolution1(t *testing.T) {
	is := assert.New(t)
	inputs := readInput()
	got, err := Solution1(inputs)
	fmt.Printf("Power Consumption: %d \n", got)
	is.NoErr(err)
}

func TestSolution2(t *testing.T) {
	is := assert.New(t)
	inputs := Report{
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
	}

	o2, co2 := 23, 10
	lifeSupport := o2 * co2
	got, err := Solution2(inputs)
	is.Equal(got, lifeSupport)
	is.NoErr(err)
}

func TestRunSolution2(t *testing.T) {
	is := assert.New(t)
	inputs := readInput()
	got, err := Solution2(inputs)
	fmt.Printf("Life support rating: %d \n", got)
	is.NoErr(err)
}

func TestReport_OxygenGeneratorRate(t *testing.T) {
	is := assert.New(t)
	report := Report{
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
	}
	res, err := report.OxygenGeneratorRate()
	is.Equal(res, 23)
	is.NoErr(err)
}

func TestReport_CO2ScrubberRate(t *testing.T) {
	is := assert.New(t)
	report := Report{
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
	}
	res, err := report.CO2ScrubberRate()
	is.Equal(res, 10)
	is.NoErr(err)
}

func TestBits_MostCommonBit(t *testing.T) {
	tests := []struct {
		name  string
		given Bits
		want  Bit
	}{
		{
			name:  "ZERO is the most common bit",
			given: Bits{0, 0, 0, 1, 0},
			want:  0,
		},
		{
			name:  "ONE is the most common bit",
			given: Bits{1, 1, 1, 0, 1},
			want:  1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := assert.New(t)
			got := MostCommonBit(test.given)
			is.Equal(got, test.want)
		})
	}
}

func TestBits_LeastCommonBit(t *testing.T) {
	tests := []struct {
		name  string
		given Bits
		want  Bit
	}{
		{
			name:  "ZERO is the least common bit",
			given: Bits{1, 1, 1, 0, 1},
			want:  0,
		},
		{
			name:  "ONE is the least common bit",
			given: Bits{0, 0, 0, 1, 0},
			want:  1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := assert.New(t)
			got := LeastCommonBit(test.given)
			is.Equal(got, test.want)
		})
	}
}

func TestBit_IsZero(t *testing.T) {
	tests := []struct {
		name  string
		given Bit
		want  bool
	}{
		{
			name:  "bit is zero",
			given: Bit(0),
			want:  true,
		},
		{
			name:  "bit is not zero",
			given: Bit(1),
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := assert.New(t)
			got := test.given.IsZero()
			is.Equal(got, test.want)
		})
	}
}

func TestBits_String(t *testing.T) {
	is := assert.New(t)
	b := Bits{1, 0, 1, 1, 0}
	want := "10110"
	is.Equal(b.String(), want)
}

func TestBits_ToDecimal(t *testing.T) {
	is := assert.New(t)
	b := Bits{1, 0, 1, 1, 0}
	want := 22
	decimal, _ := b.ToDecimal()
	is.Equal(decimal, want)
}

func readInput() Report {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer closeFile(file)

	return scanInput(file)
}

func scanInput(reader io.Reader) (inputs Report) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		var bits Bits
		line := strings.Split(scanner.Text(), "")
		for _, char := range line {
			num, _ := strconv.Atoi(char)
			bits = append(bits, Bit(num))
		}
		inputs = append(inputs, bits)
	}
	return
}

func closeFile(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		fmt.Printf("Error closing closer: %v\n", err)
	}
}
