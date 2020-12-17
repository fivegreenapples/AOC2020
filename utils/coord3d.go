package utils

import (
	"math"
)

type Coord3d struct {
	X int
	Y int
	Z int
}

var Origin3d = Coord3d{X: 0, Y: 0, Z: 0}
var In = Coord3d{X: 0, Y: 0, Z: 1}
var Out = Coord3d{X: 0, Y: 0, Z: -1}

func (c Coord3d) Manhattan() int {
	return int(math.Abs(float64(c.X)) + math.Abs(float64(c.Y)) + math.Abs(float64(c.Z)))
}

func (c Coord3d) Add(cc Coord3d) Coord3d {
	return Coord3d{
		c.X + cc.X,
		c.Y + cc.Y,
		c.Z + cc.Z,
	}
}
func (c Coord3d) Sub(cc Coord3d) Coord3d {
	return Coord3d{
		c.X - cc.X,
		c.Y - cc.Y,
		c.Z - cc.Z,
	}
}

func (c Coord3d) TwoD() Coord {
	return Coord{
		c.X,
		c.Y,
	}
}

func (c Coord3d) Scale(amount int) Coord3d {
	return Coord3d{
		c.X * amount,
		c.Y * amount,
		c.Z * amount,
	}
}

func (c Coord3d) Adjacents() []Coord3d {
	out := []Coord3d{}
	Foreach3D(c.Sub(Coord3d{1, 1, 1}), c.Add(Coord3d{1, 1, 1}), func(pos Coord3d) {
		if pos == c {
			return
		}
		out = append(out, pos)
	})
	return out
}

func ExtentsOf3DIntMap(in map[Coord3d]int) (min, max Coord3d) {
	min = Coord3d{math.MaxInt64, math.MaxInt64, math.MaxInt64}
	max = Coord3d{math.MinInt64, math.MinInt64, math.MinInt64}
	for pt := range in {
		if pt.X < min.X {
			min.X = pt.X
		}
		if pt.Y < min.Y {
			min.Y = pt.Y
		}
		if pt.Z < min.Z {
			min.Z = pt.Z
		}
		if pt.X > max.X {
			max.X = pt.X
		}
		if pt.Y > max.Y {
			max.Y = pt.Y
		}
		if pt.Z > max.Z {
			max.Z = pt.Z
		}
	}
	return min, max
}
func ExtentsOf3DBoolMap(in map[Coord3d]bool) (min, max Coord3d) {
	min = Coord3d{math.MaxInt64, math.MaxInt64, math.MaxInt64}
	max = Coord3d{math.MinInt64, math.MinInt64, math.MinInt64}
	for pt := range in {
		if pt.X < min.X {
			min.X = pt.X
		}
		if pt.Y < min.Y {
			min.Y = pt.Y
		}
		if pt.Z < min.Z {
			min.Z = pt.Z
		}
		if pt.X > max.X {
			max.X = pt.X
		}
		if pt.Y > max.Y {
			max.Y = pt.Y
		}
		if pt.Z > max.Z {
			max.Z = pt.Z
		}
	}
	return min, max
}

func Foreach3D(min, max Coord3d, cb func(Coord3d)) {
	for z := min.Z; z <= max.Z; z++ {
		for y := min.Y; y <= max.Y; y++ {
			for x := min.X; x <= max.X; x++ {
				cb(Coord3d{
					X: x,
					Y: y,
					Z: z,
				})
			}
		}
	}

}
