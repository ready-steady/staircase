// Package staircase provides a step function with smooth transitions between
// steps based on cubic Hermite splines.
//
// https://en.wikipedia.org/wiki/Step_function
//
// https://en.wikipedia.org/wiki/Cubic_Hermite_spline
package staircase

import (
	"errors"
)

// Staircase represents a step function with smooth transitions.
type Staircase struct {
	lengths    []float64
	heights    []float64
	distances  []float64
	transition float64
}

// New returns a step function with smooth transitions.
func New(lengths, heights []float64, transition float64) (*Staircase, error) {
	if transition <= 0 || transition > 0.5 {
		return nil, errors.New("the transition length should be in (0, 0.5]")
	}

	n := len(lengths)

	distances := make([]float64, n)
	distances[0] = lengths[0]
	for i := 1; i < n; i++ {
		distances[i] = distances[i-1] + lengths[i]
	}

	staircase := &Staircase{
		lengths:    lengths,
		heights:    heights,
		distances:  distances,
		transition: transition,
	}

	return staircase, nil
}

// Evaluate computes the function at a point.
func (self *Staircase) Evaluate(point float64) float64 {
	l, d, h := self.lengths, self.distances, self.heights
	n := len(l)

	if point < 0 {
		return h[0]
	}
	if point > d[n-1] {
		return h[n-1]
	}

	k := 0
	for point > d[k] {
		k++
	}

	if k > 0 {
		point -= d[k-1]
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
