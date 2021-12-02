package day1

import (
	"bufio"
	"fmt"
	assert "github.com/matryer/is"
	"io"
	"os"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name    string
		inputs  Measurements
		handler Handler
		want    int
	}{
		{
			name:    "measurements that are larger than the previous measurement",
			inputs:  Measurements{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			handler: Solution1,
			want:    7,
		},
		{
			name:    "sums that are larger than the previous sum",
			inputs:  Measurements{607, 618, 618, 617, 647, 716, 769, 792},
			handler: Solution2,
			want:    5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := assert.New(t)
			got := RunSolution(test.inputs, test.handler)
			is.Equal(test.want, got.count)
		})
	}
}

func TestSum(t *testing.T) {
	is := assert.New(t)
	numbers := Measurements{10, 6, 4}
	want := Measurement(20)
	got := Sum(numbers)
	is.Equal(got, want)
}

func TestHasIncreased(t *testing.T) {
	tests := []struct {
		name string
		a, b Measurements
		want bool
	}{
		{
			name: "measure a is larger than b",
			a:    Measurements{Measurement(5)},
			b:    Measurements{Measurement(4)},
			want: false,
		},
		{
			name: "measure b is larger than a",
			a:    Measurements{Measurement(4)},
			b:    Measurements{Measurement(5)},
			want: true,
		},
		{
			name: "measurements a is larger than b",
			a:    Measurements{50, 20},
			b:    Measurements{20, 30},
			want: false,
		},
		{
			name: "measurements b is larger than a",
			a:    Measurements{30, 40},
			b:    Measurements{40, 50},
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := assert.New(t)
			got := test.b.HasIncreasedOver(test.a)
			is.Equal(got, test.want)
		})
	}
}

func TestRunSolution1(t *testing.T) {
	inputs := readInput()
	got := RunSolution(inputs, Solution1)
	fmt.Printf("Larger: %d\n", got.count)
}

func TestRunSolution2(t *testing.T) {
	inputs := readInput()
	got := RunSolution(inputs, Solution2)
	fmt.Printf("Larger: %d\n", got.count)
}

func readInput() Measurements {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer closeFile(file)

	return scanInput(file)
}

func scanInput(reader io.Reader) (inputs Measurements) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, Measurement(num))
	}
	return
}

func closeFile(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		fmt.Printf("Error closing closer: %v\n", err)
	}
}
