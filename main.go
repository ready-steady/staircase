// Package staircase provides a step function with smooth transitions between
// steps based on the logistic function.
//
// https://en.wikipedia.org/wiki/Step_function
// https://en.wikipedia.org/wiki/Logistic_function
package staircase

import (
	"errors"
	"math"
)

// Staircase represents a step function with smooth transitions.
type Staircase struct {
	lengths    []float64
	distance   []float64
	heights    []float64
	transition float64
	steepness  float64
}

// New returns a step function with smooth transitions.
func New(lengths, heights []float64, transition, steepness float64) (*Staircase, error) {
	if transition < 0 || transition > 0.5 {
		return nil, errors.New("the transition length should be from 0 to 0.5")
	}
	if steepness <= 0 {
		return nil, errors.New("the steepness should be positive")
	}

	n := len(lengths)

	distance := make([]float64, n)
	distance[0] = lengths[0]
	for i := 1; i < n; i++ {
		distance[i] = distance[i-1] + lengths[i]
	}

	staircase := &Staircase{
		lengths:    lengths,
		distance:   distance,
		heights:    heights,
		transition: transition,
		steepness:  steepness,
	}

	return staircase, nil
}

// Evaluate computes the function at a point.
func (self *Staircase) Evaluate(point float64) float64 {
	l, d, h := self.lengths, self.distance, self.heights
	t, s := self.transition, self.steepness
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

	switch k {
	case 0:
		z := point / l[k]
		if z >= 1-t {
			x := (point - (1-t)*l[k]) / (t*l[k] + t*l[k+1])
			o := l[k] / (l[k] + l[k+1])
			return logistic(2*x-1, h[k], h[k+1], 2*o-1, s)
		} else {
			return h[k]
		}

	case n - 1:
		z := (point - d[k-1]) / l[k]
		if z <= t {
			x := (t*l[k-1] + point - d[k-1]) / (t*l[k-1] + t*l[k])
			o := l[k-1] / (l[k-1] + l[k])
			return logistic(2*x-1, h[k-1], h[k], 2*o-1, s)
		} else {
			return h[k]
		}

	default:
		z := (point - d[k-1]) / l[k]
		if z <= t {
			x := (t*l[k-1] + point - d[k-1]) / (t*l[k-1] + t*l[k])
			o := l[k-1] / (l[k-1] + l[k])
			return logistic(2*x-1, h[k-1], h[k], 2*o-1, s)
		} else if z >= (1 - t) {
			x := (point - d[k-1] - (1-t)*l[k]) / (t*l[k] + t*l[k+1])
			o := l[k] / (l[k] + l[k+1])
			return logistic(2*x-1, h[k], h[k+1], 2*o-1, s)
		} else {
			return h[k]
		}
	}
}

func logistic(x, a, b, offset, steepness float64) float64 {
	return (a*math.Exp(offset*steepness) + b*math.Exp(steepness*x)) /
		(math.Exp(offset*steepness) + math.Exp(steepness*x))
}
