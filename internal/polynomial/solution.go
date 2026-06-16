package polynomial

func sqrt(n float32) float32 {
    lo, hi := float32(0), n
    if n < 1 {
        hi = 1
    }
    for hi-lo > 0.00001 {
        mid := (lo + hi) / 2
        if mid*mid > n {
            hi = mid
        } else {
            lo = mid
        }
    }
    return (lo + hi) / 2
}

func (e Equation) Solve() []float32 {
	
	degree := e.Degree()
	if degree == 1 {
		return []float32{-e.scalar[0] / e.scalar[1]}
	} else if degree != 2 {
		return []float32{}
	}

	a := e.scalar[2]
	b := e.scalar[1]

	d := e.Discriminant()

	if d >= 0 {

		sqrtDiscriminant := sqrt(d)
		root1 := (-b - sqrtDiscriminant) / (2 * a)
		root2 := (-b + sqrtDiscriminant) / (2 * a)

		if d == 0 {
			return []float32{root1}
		}
		return []float32{root1, root2}
	}



	return nil
}