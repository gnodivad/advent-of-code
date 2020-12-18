package utils

type Vector interface {
	Neighbors() []Vector
}

type Vector3 struct {
	X, Y, Z int
}

type Vector4 struct {
	X, Y, Z, W int
}

func (v Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func (v Vector3) Neighbors() []Vector {
	neighbors, i := make([]Vector, 26), 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if !(x == 0 && y == 0 && z == 0) {
					neighbors[i] = v.Add(Vector3{X: x, Y: y, Z: z})
					i++
				}
			}
		}
	}

	return neighbors
}

func (v Vector4) Add(v2 Vector4) Vector4 {
	return Vector4{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
		W: v.W + v2.W,
	}
}

func (v Vector4) Neighbors() []Vector {
	neighbors, i := make([]Vector, 80), 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if !(x == 0 && y == 0 && z == 0 && w == 0) {
						neighbors[i] = v.Add(Vector4{X: x, Y: y, Z: z, W: w})
						i++
					}
				}
			}
		}
	}

	return neighbors
}
