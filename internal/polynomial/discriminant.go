package polynomial

func (e Equation) Discriminant() float32 {
	a := e.scalar[2]
	b := e.scalar[1]
	c := e.scalar[0]

	return b*b - 4*a*c
}
