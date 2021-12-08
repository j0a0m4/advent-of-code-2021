package day4

import (
	"fmt"
	assert "github.com/matryer/is"
	"testing"
)

var numbers = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

func TestRunSolution1(t *testing.T) {
	got := Solution1()
	fmt.Println(got)
}

func TestRunSolution2(t *testing.T) {
	got := Solution2()
	fmt.Println(got)
}

func TestReadCard(t *testing.T) {
	is := assert.New(t)
	given := []string{
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
	}
	want := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}
	got, err := ReadCard(given)
	is.NoErr(err)
	is.Equal(got, want)
}

func TestRotate(t *testing.T) {
	is := assert.New(t)
	given := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}
	want := [][]int{
		{22, 8, 21, 6, 1},
		{13, 2, 9, 10, 12},
		{17, 23, 14, 3, 20},
		{11, 4, 16, 18, 15},
		{0, 24, 7, 5, 19},
	}
	got := Rotate(given)
	is.Equal(got, want)
}

func TestFirstToWin(t *testing.T) {
	is := assert.New(t)
	cards := [][][]int{
		{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		{
			{3, 15, 0, 2, 22},
			{9, 18, 13, 17, 5},
			{19, 8, 7, 25, 23},
			{20, 11, 10, 24, 4},
			{14, 21, 16, 12, 6},
		},
		{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		},
	}
	want := 4512
	got := FirstToWin(numbers, cards)
	is.Equal(got, want)
}

func TestLastToWin(t *testing.T) {
	is := assert.New(t)
	cards := [][][]int{
		{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		{
			{3, 15, 0, 2, 22},
			{9, 18, 13, 17, 5},
			{19, 8, 7, 25, 23},
			{20, 11, 10, 24, 4},
			{14, 21, 16, 12, 6},
		},
		{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		},
	}
	want := 1924
	got := LastToWin(numbers, cards)
	is.Equal(got, want)
}

func TestBingo(t *testing.T) {
	tests := []struct {
		name  string
		given [][]int
		want  bool
	}{
		{
			name: "has bingoed",
			given: [][]int{
				{2, 6, 23, 22, 1, 21, 33, 22, 8},
				{22, 8, 21, 6, 1},
			},
			want: true,
		},
		{
			name: "has not bingoed",
			given: [][]int{
				{23, 6, 23, 29, 1, 21, 33, 22, 10},
				{22, 8, 21, 6, 1},
			},
			want: false,
		},
		{
			name: "chosen numbers are smaller than bingo numbers",
			given: [][]int{
				{22, 8, 21},
				{22, 8, 21, 6, 1},
			},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := assert.New(t)
			got := Bingo(test.given[0], test.given[1])
			is.Equal(got, test.want)
		})
	}
}

func TestSplitRow(t *testing.T) {
	is := assert.New(t)
	given := " 2  0 12  3  7"
	want := []string{"2", "0", "12", "3", "7"}
	got := SplitRow(given)
	is.Equal(got, want)
}

func TestReadRow(t *testing.T) {
	is := assert.New(t)
	given := []string{"2", "0", "12", "3", "7"}
	want := []int{2, 0, 12, 3, 7}
	got, err := ReadRow(given)
	is.NoErr(err)
	is.Equal(got, want)
}
