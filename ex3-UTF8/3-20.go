package main

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
	var d Direction = 0
	for d = 0; d < 4; d++ {
		println(d.String())
	}
}
