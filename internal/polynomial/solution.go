package polynomial

func sqrt(n float32) float32 {
    if n == 0 {
        return 0
    }
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

func (e Equation) Solve() []Root {

	degree := e.Degree()
	if degree == 1 {
		return []Root{NewRoot(-e.scalar[0]/e.scalar[1], 0)}
	} else if degree != 2 {
		return []Root{}
	}

	a := e.scalar[2]
	b := e.scalar[1]

	d := e.Discriminant()

	if d >= 0 {
		sqrtD := sqrt(d)
		root1 := NewRoot((-b-sqrtD)/(2*a), 0)
		root2 := NewRoot((-b+sqrtD)/(2*a), 0)

		if d == 0 {
			return []Root{root1}
		}
		return []Root{root1, root2}
	}

	sqrtD := sqrt(-d)
	re := -b / (2 * a)
	im := sqrtD / (2 * a)
	return []Root{NewRoot(re, im), NewRoot(re, -im)}
}