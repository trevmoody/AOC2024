package util

type Point struct {
	Horizontal int
	Vertical   int
}

func (p Point) MoveUp(distance int) Point {
	return Point{p.Horizontal, p.Vertical - distance}
}
func (p Point) MoveDown(distance int) Point {
	return Point{p.Horizontal, p.Vertical + distance}
}
func (p Point) MoveLeft(distance int) Point {
	return Point{p.Horizontal - distance, p.Vertical}
}
func (p Point) MoveRight(distance int) Point {
	return Point{p.Horizontal + distance, p.Vertical}
}
func (p Point) Move(instruction string) Point {
	switch instruction {
	case "^":
		return p.MoveUp(1)
	case ">":
		return p.MoveRight(1)
	case "V":
		return p.MoveDown(1)
	case "<":
		return p.MoveLeft(1)
	}
	return p
}
