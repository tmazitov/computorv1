package polynomial

func sqrt(n float32) float32 {
    if n == 0 {
        return 0
    }
    x := n
    if x < 1 {
        x = 1
    }
    for {
        next := (x + n/x) / 2
        if next == x {
            return x
        }
        x = next
    }
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