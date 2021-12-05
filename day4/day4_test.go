package day4

import (
	assert "github.com/matryer/is"
	"testing"
)

var numbers = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
var input = [][]string{
	{
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
	},
	{
		" 3 15  0  2 22",
		" 9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
	},
	{
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	},
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

func Test(t *testing.T) {
	Solution()
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
