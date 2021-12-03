package day3

type Report []Bits

func (r Report) LifeSupport() (int, error) {
	oxygen, err := r.OxygenGeneratorRate()
	if err != nil {
		return 0, err
	}
	co2, err := r.CO2ScrubberRate()
	if err != nil {
		return 0, err
	}
	return oxygen * co2, nil
}

func (r Report) PowerConsumption() (int, error) {
	gamma, err := r.GammaRate()
	if err != nil {
		return 0, err
	}
	epsilon, err := r.EpsilonRate()
	if err != nil {
		return 0, err
	}
	return gamma * epsilon, nil
}

func (r Report) GammaRate() (int, error) {
	gamma := r.mapBits(MostCommonBit)
	return gamma.ToDecimal()
}

func (r *Report) EpsilonRate() (int, error) {
	epsilon := r.mapBits(LeastCommonBit)
	return epsilon.ToDecimal()
}

func (r Report) OxygenGeneratorRate() (int, error) {
	bits := r.reduceToBits(MostCommonBit)
	decimal, err := bits.ToDecimal()
	if err != nil {
		return 0, err
	}
	return decimal, nil
}

func (r Report) CO2ScrubberRate() (int, error) {
	bits := r.reduceToBits(LeastCommonBit)
	decimal, err := bits.ToDecimal()
	if err != nil {
		return 0, err
	}
	return decimal, nil
}

func (r Report) reduceToBits(mapper BitMapper) Bits {
	for i := 0; i < len(r[0]); i++ {
		if len(r) == 1 {
			break
		}
		r = r.removeRowsWithoutBit(mapper, i)
	}
	return r[0]
}

func (r Report) mapBits(mapper BitMapper) (bits Bits) {
	for i := 0; i < len(r[0]); i++ {
		bit := r.getBitFromColumn(mapper, i)
		bits = append(bits, bit)
	}
	return
}

func (r Report) getBitFromColumn(mapper BitMapper, i int) Bit {
	bits := r.extractColumn(i)
	return mapper(bits)
}

func (r Report) extractColumn(i int) (bits Bits) {
	for _, row := range r {
		bits = append(bits, row[i])
	}
	return
}

func (r Report) removeRowsWithoutBit(mapper BitMapper, col int) (report Report) {
	bit := r.getBitFromColumn(mapper, col)
	for _, row := range r {
		if row[col] == bit {
			report = append(report, row)
		}
	}
	return
}
