package utils

import "math"

type Coord struct {
	X int
	Y int
}

var Origin = Coord{X: 0, Y: 0}
var Up = Coord{X: 0, Y: -1}
var UpLeft = Coord{X: -1, Y: -1}
var UpRight = Coord{X: 1, Y: -1}
var Right = Coord{X: 1, Y: 0}
var Down = Coord{X: 0, Y: 1}
var DownLeft = Coord{X: -1, Y: 1}
var DownRight = Coord{X: 1, Y: 1}
var Left = Coord{X: -1, Y: 0}

var Cardinals = []Coord{Up, Right, Down, Left}
var Ordinals = []Coord{UpLeft, UpRight, DownRight, DownLeft}
var CardinalsAndOrdinals = []Coord{UpLeft, Up, UpRight, Right, DownRight, Down, DownLeft, Left}

func (c Coord) Manhattan() int {
	return int(math.Abs(float64(c.X)) + math.Abs(float64(c.Y)))
}

func (c Coord) Add(cc Coord) Coord {
	return Coord{
		c.X + cc.X,
		c.Y + cc.Y,
	}
}
func (c Coord) Sub(cc Coord) Coord {
	return Coord{
		c.X - cc.X,
		c.Y - cc.Y,
	}
}
func (c Coord) ThreeD(z int) Coord3d {
	return Coord3d{
		c.X,
		c.Y,
		z,
	}
}
func (c Coord) Simplify() Coord {

	absX := AbsInt(c.X)
	absY := AbsInt(c.Y)

	if absX == 0 && absY == 0 {
		return c
	}
	if absX == 0 {
		return Coord{0, c.Y / absY}
	}
	if absY == 0 {
		return Coord{c.X / absX, 0}
	}

	lcf := LargestCommonFactor(absX, absY)
	return Coord{c.X / lcf, c.Y / lcf}
}

func (c Coord) Adjacents() []Coord {
	out := make([]Coord, 8)
	out[0] = c.Add(UpLeft)
	out[1] = c.Add(Up)
	out[2] = c.Add(UpRight)
	out[3] = c.Add(Left)
	out[4] = c.Add(Right)
	out[5] = c.Add(DownLeft)
	out[6] = c.Add(Down)
	out[7] = c.Add(DownRight)
	return out
}

func (c Coord) IsInside(min, max Coord) bool {
	return c.X >= min.X && c.X <= max.X && c.Y >= min.Y && c.Y <= max.Y
}

func ExtentsOfIntMap(in map[Coord]int) (min, max Coord) {
	min = Coord{math.MaxInt64, math.MaxInt64}
	max = Coord{math.MinInt64, math.MinInt64}
	for pt := range in {
		if pt.X < min.X {
			min.X = pt.X
		}
		if pt.Y < min.Y {
			min.Y = pt.Y
		}
		if pt.X > max.X {
			max.X = pt.X
		}
		if pt.Y > max.Y {
			max.Y = pt.Y
		}
	}
	return min, max
}
func ExtentsOfBoolMap(in map[Coord]bool) (min, max Coord) {
	min = Coord{math.MaxInt64, math.MaxInt64}
	max = Coord{math.MinInt64, math.MinInt64}
	for pt := range in {
		if pt.X < min.X {
			min.X = pt.X
		}
		if pt.Y < min.Y {
			min.Y = pt.Y
		}
		if pt.X > max.X {
			max.X = pt.X
		}
		if pt.Y > max.Y {
			max.Y = pt.Y
		}
	}
	return min, max
}

func ExtentsOfStringMap(in map[Coord]string) (min, max Coord) {
	min = Coord{math.MaxInt64, math.MaxInt64}
	max = Coord{math.MinInt64, math.MinInt64}
	for pt := range in {
		if pt.X < min.X {
			min.X = pt.X
		}
		if pt.Y < min.Y {
			min.Y = pt.Y
		}
		if pt.X > max.X {
			max.X = pt.X
		}
		if pt.Y > max.Y {
			max.Y = pt.Y
		}
	}
	return min, max
}
