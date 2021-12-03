package day3

import "strconv"

type BitMapper func(bits Bits) Bit

type Bits []Bit

func (b Bits) String() (s string) {
	for _, bit := range b {
		s += strconv.Itoa(int(bit))
	}
	return
}

func MostCommonBit(b Bits) Bit {
	return b.reduceBits(func(ones, zeros int) bool {
		return ones >= zeros
	})
}

func LeastCommonBit(b Bits) Bit {
	return b.reduceBits(func(ones, zeros int) bool {
		return ones < zeros
	})
}

func (b Bits) ToDecimal() (int, error) {
	num, err := strconv.ParseInt(b.String(), 2, 64)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

func (b Bits) reduceBits(predicate func(ones, zeros int) bool) Bit {
	ones, zeros := b.countBits()
	if predicate(ones, zeros) {
		return Bit(1)
	}
	return Bit(0)
}

func (b Bits) countBits() (ones int, zeros int) {
	for _, bit := range b {
		if bit.IsZero() {
			zeros++
		} else {
			ones++
		}
	}
	return
}
