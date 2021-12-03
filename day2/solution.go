package day2

const (
	Forward = "forward"
	Down    = "down"
	Up      = "up"
)

type Commands []Command

type Command struct {
	direction string
	units     int
}

type Position struct {
	aim, horizontal, depth int
}

func (p *Position) forward(units int) {
	p.horizontal += units
}

func (p *Position) up(units int) {
	p.depth -= units
}

func (p *Position) down(units int) {
	p.depth += units
}

func (p *Position) Move(command Command) {
	units := command.units
	switch command.direction {
	case Forward:
		p.forward(units)
	case Up:
		p.up(units)
	case Down:
		p.down(units)
	}
}

func (p *Position) increaseAim(units int) {
	p.aim += units
}

func (p *Position) decreaseAim(units int) {
	p.aim -= units
}

func (p *Position) AimAndMove(command Command) {
	units := command.units
	switch command.direction {
	case Forward:
		depth := p.aim * units
		p.forward(units)
		p.down(depth)
	case Down:
		p.increaseAim(units)
	case Up:
		p.decreaseAim(units)
	}
}

func Solution1(commands Commands) *Position {
	position := &Position{}
	for _, command := range commands {
		position.Move(command)
	}
	return position
}

func Solution2(commands Commands) *Position {
	position := &Position{}
	for _, command := range commands {
		position.AimAndMove(command)
	}
	return position
}
