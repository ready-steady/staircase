// Package staircase provides a step function with smooth transitions based on
// cubic Hermite splines.
//
// https://en.wikipedia.org/wiki/Step_function
//
// https://en.wikipedia.org/wiki/Cubic_Hermite_spline
package staircase

// Staircase represents a step function with smooth transitions.
type Staircase struct {
	lengths    []float64
	heights    []float64
	transition float64
}

// New returns a step function with smooth transitions. The transition parameter
// specifies the fraction of a step’s length dedicated for smoothing on each end
// of the step; the parameter should be in (0, 0.5].
func New(lengths, heights []float64, transition float64) *Staircase {
	return &Staircase{
		lengths:    lengths,
		heights:    heights,
		transition: transition,
	}
}

// Evaluate computes the function at a point.
func (self *Staircase) Evaluate(point float64) float64 {
	h := self.heights

	if point < 0 {
		return h[0]
	}

	l := self.lengths
	n := len(l)

	k := 0
	for k < n && point > l[k] {
		point -= l[k]
		k++
	}

	if k == n {
		return h[n-1]
	}

	if z, t := point/l[k], self.transition; k > 0 && z <= t {
		return hermite((t*l[k-1]+point)/(t*l[k-1]+t*l[k]), h[k-1], h[k])
	} else if k < n-1 && z >= 1-t {
		return hermite((point-(1-t)*l[k])/(t*l[k]+t*l[k+1]), h[k], h[k+1])
	}

	return h[k]
}

func hermite(t, a, b float64) float64 {
	t2 := t * t
	t3 := t * t2
	return (2*t3-3*t2+1)*a + (-2*t3+3*t2)*b
}
