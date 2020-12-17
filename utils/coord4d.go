package utils

import "math"

type Coord4d struct {
	X int
	Y int
	Z int
	W int
}

var Origin4d = Coord4d{X: 0, Y: 0, Z: 0, W: 0}

func (c Coord4d) Add(cc Coord4d) Coord4d {
	return Coord4d{
		c.X + cc.X,
		c.Y + cc.Y,
		c.Z + cc.Z,
		c.W + cc.W,
	}
}
func (c Coord4d) Sub(cc Coord4d) Coord4d {
	return Coord4d{
		c.X - cc.X,
		c.Y - cc.Y,
		c.Z - cc.Z,
		c.W - cc.W,
	}
}

func (c Coord4d) Scale(amount int) Coord4d {
	return Coord4d{
		c.X * amount,
		c.Y * amount,
		c.Z * amount,
		c.W * amount,
	}
}

func (c Coord4d) Adjacents() []Coord4d {
	out := []Coord4d{}
	Foreach4D(c.Sub(Coord4d{1, 1, 1, 1}), c.Add(Coord4d{1, 1, 1, 1}), func(pos Coord4d) {
		if pos == c {
			return
		}
		out = append(out, pos)
	})
	return out
}

func ExtentsOf4DBoolMap(in map[Coord4d]bool) (min, max Coord4d) {
	min = Coord4d{math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}
	max = Coord4d{math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64}
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
		if pt.W < min.W {
			min.W = pt.W
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
		if pt.W > max.W {
			max.W = pt.W
		}
	}
	return min, max
}

func Foreach4D(min, max Coord4d, cb func(Coord4d)) {
	for w := min.W; w <= max.W; w++ {
		for z := min.Z; z <= max.Z; z++ {
			for y := min.Y; y <= max.Y; y++ {
				for x := min.X; x <= max.X; x++ {
					cb(Coord4d{
						X: x,
						Y: y,
						Z: z,
						W: w,
					})
				}
			}
		}
	}

}
