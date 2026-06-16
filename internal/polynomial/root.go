
package polynomial

import "fmt"

type Root struct {
	re float32
	im float32
}

func NewRoot(re, im float32) Root {
	return Root{re: re, im: im}
}

func (r Root) Re() float32 { return r.re }
func (r Root) Im() float32 { return r.im }

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func roundInt(f float32) int {
	if f >= 0 {
		return int(f + 0.5)
	}
	return int(f - 0.5)
}

func toFraction(f float32) (int, int) {
	const precision = 10000
	num := roundInt(f * precision)
	den := precision
	g := gcd(absInt(num), den)
	return num / g, den / g
}

func formatReal(num, den int) string {
	if den == 1 {
		return fmt.Sprintf("%d", num)
	}
	return fmt.Sprintf("%d/%d", num, den)
}

func formatImag(num, den int) string {
	abs := absInt(num)
	switch {
	case abs == 1 && den == 1:
		return "i"
	case abs == 1:
		return fmt.Sprintf("i/%d", den)
	case den == 1:
		return fmt.Sprintf("%di", abs)
	default:
		return fmt.Sprintf("%di/%d", abs, den)
	}
}

func (r Root) ToString() string {
	if r.im == 0 {
		return fmt.Sprintf("%.6f", r.re)
	}

	imNum, imDen := toFraction(r.im)
	imStr := formatImag(imNum, imDen)

	if r.re == 0 {
		if imNum < 0 {
			return "-" + imStr
		}
		return imStr
	}

	reNum, reDen := toFraction(r.re)
	reStr := formatReal(reNum, reDen)

	if imNum < 0 {
		return fmt.Sprintf("%s - %s", reStr, imStr)
	}
	return fmt.Sprintf("%s + %s", reStr, imStr)
}
