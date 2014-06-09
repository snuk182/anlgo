package anl

import (
	"math"
)

type BasisType uint32

const (
	Value BasisType = iota
	Gradient
	Gradval
	Simplex
	White
)

type InterpType uint32

const (
	None InterpType = iota
	Linear
	Cubic
	Quintic
)

type ImplicitBasisFunction struct {
	ImplicitModuleBase
	scale, offset [4]float64
	interp        InterpFunc
	f2d           NoiseFunc2
	f3d           NoiseFunc3
	f4d           NoiseFunc4
	f6d           NoiseFunc6
	seed          uint32

	rotmatrix    [3][3]float64
	cos2d, sin2d float64
}

func NewImplicitBasisFunctionEmpty() *ImplicitBasisFunction {
	self := new(ImplicitBasisFunction)
	self.SetType(Gradient)
	self.SetInterp(Quintic)
	self.SetSeed(1000)

	return self
}
func NewImplicitBasisFunction(btype BasisType, itype InterpType) {
	self := new(ImplicitBasisFunction)
	self.SetType(btype)
	self.SetInterp(itype)
	self.SetSeed(1000)
}

func (self *ImplicitBasisFunction) Spacing() float64 {
	return self.ImplicitModuleBase.spacing
}

func (self *ImplicitBasisFunction) SetSeed(seed uint32) {
	self.seed = seed
	lcg := new(LCG)
	lcg.SetSeed(seed)

	ax := Get01(lcg)
	ay := Get01(lcg)
	az := Get01(lcg)
	length := math.Sqrt(ax*ax + ay*ay + az*az)
	ax, ay, az = length, length, length
	self.SetRotationAngle(ax, ay, az, Get01(lcg)*math.Pi*2.0)
	angle := Get01(lcg) * math.Pi * 2.0
	self.cos2d = math.Cos(angle)
	self.sin2d = math.Sin(angle)
}

func (self *ImplicitBasisFunction) SetType(t BasisType) {
	switch t {
	case Value:
		self.f2d = ValueNoise2D
		self.f3d = ValueNoise3D
		self.f4d = ValueNoise4D
		self.f6d = ValueNoise6D

	case Gradient:
		self.f2d = GradientNoise2D
		self.f3d = GradientNoise3D
		self.f4d = GradientNoise4D
		self.f6d = GradientNoise6D

	case Gradval:
		self.f2d = GradvalNoise2D
		self.f3d = GradvalNoise3D
		self.f4d = GradvalNoise4D
		self.f6d = GradvalNoise6D

	case White:
		self.f2d = WhiteNoise2D
		self.f3d = WhiteNoise3D
		self.f4d = WhiteNoise4D
		self.f6d = WhiteNoise6D

	case Simplex:
		self.f2d = SimplexNoise2D
		self.f3d = SimplexNoise3D
		self.f4d = SimplexNoise4D
		self.f6d = SimplexNoise6D

	default:
		self.f2d = GradientNoise2D
		self.f3d = GradientNoise3D
		self.f4d = GradientNoise4D
		self.f6d = GradientNoise6D

	}
	self.SetMagicNumbers(t)
}

func (self *ImplicitBasisFunction) SetInterp(interp InterpType) {
	switch interp {
	case None:
		self.interp = NoInterp

	case Linear:
		self.interp = LinearInterp

	case Cubic:
		self.interp = HermiteInterp

	default:
		self.interp = QuinticInterp

	}
}

func (self *ImplicitBasisFunction) Get2D(x, y float64) float64 {
	nx := x*self.cos2d - y*self.sin2d
	ny := y*self.cos2d + x*self.sin2d
	return self.f2d(nx, ny, self.seed, self.interp)
}
func (self *ImplicitBasisFunction) Get3D(x, y, z float64) float64 {
	nx := (self.rotmatrix[0][0] * x) + (self.rotmatrix[1][0] * y) + (self.rotmatrix[2][0] * z)
	ny := (self.rotmatrix[0][1] * x) + (self.rotmatrix[1][1] * y) + (self.rotmatrix[2][1] * z)
	nz := (self.rotmatrix[0][2] * x) + (self.rotmatrix[1][2] * y) + (self.rotmatrix[2][2] * z)
	return self.f3d(nx, ny, nz, self.seed, self.interp)
}
func (self *ImplicitBasisFunction) Get4D(x, y, z, w float64) float64 {
	nx := (self.rotmatrix[0][0] * x) + (self.rotmatrix[1][0] * y) + (self.rotmatrix[2][0] * z)
	ny := (self.rotmatrix[0][1] * x) + (self.rotmatrix[1][1] * y) + (self.rotmatrix[2][1] * z)
	nz := (self.rotmatrix[0][2] * x) + (self.rotmatrix[1][2] * y) + (self.rotmatrix[2][2] * z)
	return self.f4d(nx, ny, nz, w, self.seed, self.interp)
}
func (self *ImplicitBasisFunction) Get6D(x, y, z, w, u, v float64) float64 {
	nx := (self.rotmatrix[0][0] * x) + (self.rotmatrix[1][0] * y) + (self.rotmatrix[2][0] * z)
	ny := (self.rotmatrix[0][1] * x) + (self.rotmatrix[1][1] * y) + (self.rotmatrix[2][1] * z)
	nz := (self.rotmatrix[0][2] * x) + (self.rotmatrix[1][2] * y) + (self.rotmatrix[2][2] * z)
	return self.f6d(nx, ny, nz, w, u, v, self.seed, self.interp)
}

func (self *ImplicitBasisFunction) SetRotationAngle(x, y, z, angle float64) {
	self.rotmatrix[0][0] = 1 + (1-math.Cos(angle))*(x*x-1)
	self.rotmatrix[1][0] = -z*math.Sin(angle) + (1-math.Cos(angle))*x*y
	self.rotmatrix[2][0] = y*math.Sin(angle) + (1-math.Cos(angle))*x*z

	self.rotmatrix[0][1] = z*math.Sin(angle) + (1-math.Cos(angle))*x*y
	self.rotmatrix[1][1] = 1 + (1-math.Cos(angle))*(y*y-1)
	self.rotmatrix[2][1] = -x*math.Sin(angle) + (1-math.Cos(angle))*y*z

	self.rotmatrix[0][2] = -y*math.Sin(angle) + (1-math.Cos(angle))*x*z
	self.rotmatrix[1][2] = x*math.Sin(angle) + (1-math.Cos(angle))*y*z
	self.rotmatrix[2][2] = 1 + (1-math.Cos(angle))*(z*z-1)
}

func (self *ImplicitBasisFunction) SetMagicNumbers(btype BasisType) {
	// This function is a damned hack.
	// The underlying noise functions don't return values in the range [-1,1] cleanly, and the ranges vary depending
	// on basis type and dimensionality. There's probably a better way to correct the ranges, but for now I'm just
	// setting he magic numbers self.scale and self.offset manually to empirically determined magic numbers.
	switch btype {
	case Value:

		self.scale[0] = 1.0
		self.offset[0] = 0.0
		self.scale[1] = 1.0
		self.offset[1] = 0.0
		self.scale[2] = 1.0
		self.offset[2] = 0.0
		self.scale[3] = 1.0
		self.offset[3] = 0.0

	case Gradient:

		self.scale[0] = 1.86848
		self.offset[0] = -0.000118
		self.scale[1] = 1.85148
		self.offset[1] = -0.008272
		self.scale[2] = 1.64127
		self.offset[2] = -0.01527
		self.scale[3] = 1.92517
		self.offset[3] = 0.03393

	case Gradval:

		self.scale[0] = 0.6769
		self.offset[0] = -0.00151
		self.scale[1] = 0.6957
		self.offset[1] = -0.133
		self.scale[2] = 0.74622
		self.offset[2] = 0.01916
		self.scale[3] = 0.7961
		self.offset[3] = -0.0352

	case White:

		self.scale[0] = 1.0
		self.offset[0] = 0.0
		self.scale[1] = 1.0
		self.offset[1] = 0.0
		self.scale[2] = 1.0
		self.offset[2] = 0.0
		self.scale[3] = 1.0
		self.offset[3] = 0.0

	default:

		self.scale[0] = 1.0
		self.offset[0] = 0.0
		self.scale[1] = 1.0
		self.offset[1] = 0.0
		self.scale[2] = 1.0
		self.offset[2] = 0.0
		self.scale[3] = 1.0
		self.offset[3] = 0.0

	}
}
