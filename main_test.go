package staircase

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestEvaluate(t *testing.T) {
	lengths := []float64{0.2, 0.3, 0.5}
	heights := []float64{1, -2, 3}
	transition := 0.4

	staircase := New(lengths, heights, transition)

	x := []float64{
		0.00, 0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09, 0.10, 0.11,
		0.12, 0.13, 0.14, 0.15, 0.16, 0.17, 0.18, 0.19, 0.20, 0.21, 0.22, 0.23,
		0.24, 0.25, 0.26, 0.27, 0.28, 0.29, 0.30, 0.31, 0.32, 0.33, 0.34, 0.35,
		0.36, 0.37, 0.38, 0.39, 0.40, 0.41, 0.42, 0.43, 0.44, 0.45, 0.46, 0.47,
		0.48, 0.49, 0.50, 0.51, 0.52, 0.53, 0.54, 0.55, 0.56, 0.57, 0.58, 0.59,
		0.60, 0.61, 0.62, 0.63, 0.64, 0.65, 0.66, 0.67, 0.68, 0.69, 0.70, 0.71,
		0.72, 0.73, 0.74, 0.75, 0.76, 0.77, 0.78, 0.79, 0.80, 0.81, 0.82, 0.83,
		0.84, 0.85, 0.86, 0.87, 0.88, 0.89, 0.90, 0.91, 0.92, 0.93, 0.94, 0.95,
		0.96, 0.97, 0.98, 0.99, 1.00,
	}

	y := []float64{
		+1.0000000000000000e+00, +1.0000000000000000e+00,
		+1.0000000000000000e+00, +1.0000000000000000e+00,
		+1.0000000000000000e+00, +1.0000000000000000e+00,
		+1.0000000000000000e+00, +1.0000000000000000e+00,
		+1.0000000000000000e+00, +1.0000000000000000e+00,
		+1.0000000000000000e+00, +1.0000000000000000e+00,
		+1.0000000000000000e+00, +9.7824999999999995e-01,
		+9.1599999999999993e-01, +8.1774999999999998e-01,
		+6.8799999999999972e-01, +5.3124999999999989e-01,
		+3.5199999999999998e-01, +1.5474999999999983e-01,
		-5.6000000000000272e-02, -2.7575000000000038e-01,
		-5.0000000000000067e-01, -7.2425000000000028e-01,
		-9.4400000000000039e-01, -1.1547500000000002e+00,
		-1.3520000000000001e+00, -1.5312500000000007e+00,
		-1.6880000000000004e+00, -1.8177500000000004e+00,
		-1.9160000000000006e+00, -1.9782500000000007e+00,
		-2.0000000000000000e+00, -2.0000000000000000e+00,
		-2.0000000000000000e+00, -2.0000000000000000e+00,
		-2.0000000000000000e+00, -2.0000000000000000e+00,
		-2.0000000000000000e+00, -1.9856567382812500e+00,
		-1.9438476562500000e+00, -1.8764038085937498e+00,
		-1.7851562500000000e+00, -1.6719360351562500e+00,
		-1.5385742187500000e+00, -1.3869018554687500e+00,
		-1.2187499999999998e+00, -1.0359497070312498e+00,
		-8.4033203125000044e-01, -6.3372802734375022e-01,
		-4.1796875000000000e-01, -1.9488525390625000e-01,
		+3.3691406250000888e-02, +2.6593017578125067e-01,
		+5.0000000000000111e-01, +7.3406982421875111e-01,
		+9.6630859375000067e-01, +1.1948852539062522e+00,
		+1.4179687500000022e+00, +1.6337280273437493e+00,
		+1.8403320312499993e+00, +2.0359497070312491e+00,
		+2.2187500000000000e+00, +2.3869018554687500e+00,
		+2.5385742187500000e+00, +2.6719360351562478e+00,
		+2.7851562499999987e+00, +2.8764038085937487e+00,
		+2.9438476562499987e+00, +2.9856567382812500e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00, +3.0000000000000000e+00,
		+3.0000000000000000e+00,
	}

	for i := range x {
		assert.EqualWithin(staircase.Evaluate(x[i]), y[i], 1e-14, t)
	}
}
