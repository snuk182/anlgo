package anl

import (
	"math"
)

func Clamp(v, l, h float64) float64 {
	if v < l {
		v = l
	}
	if v > h {
		v = h
	}

	return v
}

func Lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}

/*func IsPowerOf2(n float64) bool {
    return !((n-1) & n);
}*/

func HermiteBlend(t float64) float64 {
	return (t * t * (3 - 2*t))
}

func QuinticBlend(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}

func ArrayDot(arr []float64, values ...float64) float64 {
	output := float64(0.)
	for i, v := range values {
		output += v * arr[i]
	}

	return output
}

func FastFloor(t float64) int32 {
	return (map[bool]int32{true: int32(t), false: int32(t) - 1})[t > 0]
}

func Bias(b, t float64) float64 {
	return math.Pow(t, math.Log(b)/math.Log(0.5))
}

func Gain(g, t float64) float64 {
	if t < 0.5 {
		return Bias(1.0-g, 2.0*t) / 2.0
	} else {
		return 1.0 - Bias(1.0-g, 2.0-2.0*t)/2.0
	}
}
