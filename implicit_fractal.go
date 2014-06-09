package anl

import (
	"math"
)

type FractalType uint32

const (
	FBM FractalType = iota
	RidgedMulti
	Billow
	Multi
	HybridMulti
	DecarpentierSwiss
)

type ImplicitFractal struct {
	ImplicitModuleBase
	basis                                  [MaxSources]ImplicitBasisFunction
	source                                 [MaxSources]ImplicitModule
	exparray                               [MaxSources]float64
	correct                                [MaxSources][2]float64
	offset, gain, H, frequency, lacunarity float64
	numoctaves                             uint32
	ftype                                  FractalType
}

func (self *ImplicitFractal) SetNumOctaves(n uint32) {
	if n >= MaxSources {
		n = MaxSources - 1
		self.numoctaves = n
	}
}

func (self *ImplicitFractal) SetFrequency(f float64) {
	self.frequency = f
}

func (self *ImplicitFractal) SetLacunarity(l float64) {
	self.lacunarity = l
}

func (self *ImplicitFractal) SetGain(g float64) {
	self.gain = g
}

func (self *ImplicitFractal) SetOffset(o float64) {
	self.offset = o
}

func (self *ImplicitFractal) SetH(h float64) {
	self.H = h
}

func NewImplicitFractal(ftype FractalType, basistype BasisType, interptype InterpType) *ImplicitFractal {
	v := &ImplicitFractal{
		numoctaves: 8,
		frequency:  1.,
		lacunarity: 2.,
		ftype:      ftype,
	}
	
	v.SetAllSourceTypes(basistype, interptype);
	v.ResetAllSources();
	
	return v
}

func (self *ImplicitFractal) SetType(t FractalType) {
	self.ftype = t
	switch t {
	case FBM:
		self.H = 1.0
		self.gain = 0.5
		self.offset = 0
		self.FBmCalcWeights()

	case RidgedMulti:
		self.H = 0.9
		self.gain = 0.5
		self.offset = 1
		self.RidgedMultiCalcWeights()

	case Billow:
		self.H = 1
		self.gain = 0.5
		self.offset = 0
		self.BillowCalcWeights()

	case Multi:
		self.H = 1
		self.offset = 0
		self.gain = 0
		self.MultiCalcWeights()

	case HybridMulti:
		self.H = 0.25
		self.gain = 1
		self.offset = 0.7
		self.HybridMultiCalcWeights()

	case DecarpentierSwiss:
		self.H = 0.9
		self.gain = 0.6
		self.offset = 0.15
		self.DeCarpentierSwissCalcWeights()

	default:
		self.H = 1.0
		self.gain = 0
		self.offset = 0
		self.FBmCalcWeights()
	}
}

func (self *ImplicitFractal) SetAllSourceTypes(basis_type BasisType, interp InterpType) {
	for i := 0; i < int(MaxSources); i++ {
		self.basis[i].SetType(basis_type)
		self.basis[i].SetInterp(interp)
	}
}

func (self *ImplicitFractal) SetSourceType(which uint32, basis_type BasisType, interp InterpType) {
	if which >= MaxSources || which < 0 {
		return
	}

	self.basis[which].SetType(basis_type)
	self.basis[which].SetInterp(interp)
}

func (self *ImplicitFractal) OverrideSource(which uint32, b ImplicitModule) {
	if which >= MaxSources || which < 0 {
		return
	}
	self.source[which] = b
}

func (self *ImplicitFractal) ResetSource(which uint32) {
	if which >= MaxSources || which < 0 {
		return
	}
	self.source[which] = &self.basis[which]
}

func (self *ImplicitFractal) ResetAllSources() {
	for c := 0; c < int(MaxSources); c++ {
		self.source[c] = &self.basis[c]
	}
}

func (self *ImplicitFractal) SetSeed(seed uint32) {
	for c := 0; c < int(MaxSources); c++ {
		self.source[c].SetSeed(seed + uint32(c*300))
	}
}

func (self *ImplicitFractal) GetBasis(which uint32) *ImplicitBasisFunction {
	if which >= MaxSources || which < 0 {
		return nil
	}
	return &self.basis[which]
}

func (self *ImplicitFractal) Get2D(x, y float64) float64 {
	var v float64
	switch self.ftype {
	case FBM:
		v = self.FBmGet2D(x, y)
	case RidgedMulti:
		v = self.RidgedMultiGet2D(x, y)
	case Billow:
		v = self.BillowGet2D(x, y)
	case Multi:
		v = self.MultiGet2D(x, y)
	case HybridMulti:
		v = self.HybridMultiGet2D(x, y)
	case DecarpentierSwiss:
		v = self.DeCarpentierSwissGet2D(x, y)
	default:
		v = self.FBmGet2D(x, y)
	}
	//return clamp(v,-1.0,1.0);
	return v
}

func (self *ImplicitFractal) Get3D(x, y, z float64) float64 {
	var val float64
	switch self.ftype {
	case FBM:
		val = self.FBmGet3D(x, y, z)

	case RidgedMulti:
		val = self.RidgedMultiGet3D(x, y, z)

	case Billow:
		val = self.BillowGet3D(x, y, z)

	case Multi:
		val = self.MultiGet3D(x, y, z)

	case HybridMulti:
		val = self.HybridMultiGet3D(x, y, z)

	case DecarpentierSwiss:
		val = self.DeCarpentierSwissGet3D(x, y, z)

	default:
		val = self.FBmGet3D(x, y, z)

	}
	//return clamp(val,-1.0,1.0);
	return val
}

func (self *ImplicitFractal) Get4D(x, y, z, w float64) float64 {
	var val float64
	switch self.ftype {
	case FBM:
		val = self.FBmGet4D(x, y, z, w)

	case RidgedMulti:
		val = self.RidgedMultiGet4D(x, y, z, w)

	case Billow:
		val = self.BillowGet4D(x, y, z, w)

	case Multi:
		val = self.MultiGet4D(x, y, z, w)

	case HybridMulti:
		val = self.HybridMultiGet4D(x, y, z, w)

	case DecarpentierSwiss:
		val = self.DeCarpentierSwissGet4D(x, y, z, w)

	default:
		val = self.FBmGet4D(x, y, z, w)

	}
	return val
}

func (self *ImplicitFractal) Get6D(x, y, z, w, u, v float64) float64 {
	var val float64
	switch self.ftype {
	case FBM:
		val = self.FBmGet6D(x, y, z, w, u, v)

	case RidgedMulti:
		val = self.RidgedMultiGet6D(x, y, z, w, u, v)

	case Billow:
		val = self.BillowGet6D(x, y, z, w, u, v)

	case Multi:
		val = self.MultiGet6D(x, y, z, w, u, v)

	case HybridMulti:
		val = self.HybridMultiGet6D(x, y, z, w, u, v)

	case DecarpentierSwiss:
		val = self.DeCarpentierSwissGet6D(x, y, z, w, u, v)

	default:
		val = self.FBmGet6D(x, y, z, w, u, v)
	}

	return val
}

func (self *ImplicitFractal) FBmCalcWeights() {
	//std::cout << "Weights: ";
	for i := 0; i < int(MaxSources); i++ {
		self.exparray[i] = math.Pow(self.lacunarity, float64(-i)*self.H)
	}

	// Calculate scale/bias pairs by guessing at minimum and maximum values and remapping to [-1,1]
	minvalue, maxvalue := 0., 0.
	for i := 0; i < int(MaxSources); i++ {
		minvalue += -1.0 * self.exparray[i]
		maxvalue += 1.0 * self.exparray[i]

		A, B := -1., 1.
		scale := (B - A) / (maxvalue - minvalue)
		bias := A - minvalue*scale
		self.correct[i][0] = scale
		self.correct[i][1] = bias

		//std::cout << minvalue << " " << maxvalue << " " << scale << " " << bias << std::endl;
	}
}

func (self *ImplicitFractal) RidgedMultiCalcWeights() {
	for i := 0; i < int(MaxSources); i++ {
		self.exparray[i] = math.Pow(self.lacunarity, float64(-i)*self.H)
	}

	// Calculate scale/bias pairs by guessing at minimum and maximum values and remapping to [-1,1]
	minvalue, maxvalue := 0., 0.
	for i := 0; i < int(MaxSources); i++ {
		minvalue += (self.offset - 1.0) * (self.offset - 1.0) * self.exparray[i]
		maxvalue += (self.offset) * (self.offset) * self.exparray[i]

		A, B := -1., 1.
		scale := (B - A) / (maxvalue - minvalue)
		bias := A - minvalue*scale
		self.correct[i][0] = scale
		self.correct[i][1] = bias
	}

}

func (self *ImplicitFractal) DeCarpentierSwissCalcWeights() {
	for i := 0; i < int(MaxSources); i++ {
		self.exparray[i] = math.Pow(self.lacunarity, float64(-i)*self.H)
	}

	// Calculate scale/bias pairs by guessing at minimum and maximum values and remapping to [-1,1]
	minvalue, maxvalue := 0., 0.
	for i := 0; i < int(MaxSources); i++ {
		minvalue += (self.offset - 1.0) * (self.offset - 1.0) * self.exparray[i]
		maxvalue += (self.offset) * (self.offset) * self.exparray[i]

		A, B := -1., 1.
		scale := (B - A) / (maxvalue - minvalue)
		bias := A - minvalue*scale
		self.correct[i][0] = scale
		self.correct[i][1] = bias
	}

}

func (self *ImplicitFractal) BillowCalcWeights() {
	for i := 0; i < int(MaxSources); i++ {
		self.exparray[i] = math.Pow(self.lacunarity, float64(-i)*self.H)
	}

	// Calculate scale/bias pairs by guessing at minimum and maximum values and remapping to [-1,1]
	minvalue, maxvalue := 0., 0.
	for i := 0; i < int(MaxSources); i++ {
		minvalue += -1.0 * self.exparray[i]
		maxvalue += 1.0 * self.exparray[i]

		A, B := -1., 1.
		scale := (B - A) / (maxvalue - minvalue)
		bias := A - minvalue*scale
		self.correct[i][0] = scale
		self.correct[i][1] = bias
	}

}

func (self *ImplicitFractal) MultiCalcWeights() {
	for i := 0; i < int(MaxSources); i++ {
		self.exparray[i] = math.Pow(self.lacunarity, float64(-i)*self.H)
	}

	// Calculate scale/bias pairs by guessing at minimum and maximum values and remapping to [-1,1]
	minvalue, maxvalue := 1., 1.
	for i := 0; i < int(MaxSources); i++ {
		minvalue *= -1.0*self.exparray[i] + 1.0
		maxvalue *= 1.0*self.exparray[i] + 1.0

		A, B := -1., 1.
		scale := (B - A) / (maxvalue - minvalue)
		bias := A - minvalue*scale
		self.correct[i][0] = scale
		self.correct[i][1] = bias
	}

}

func (self *ImplicitFractal) HybridMultiCalcWeights() {
	for i := 0; i < int(MaxSources); i++ {
		self.exparray[i] = math.Pow(self.lacunarity, float64(-i)*self.H)
	}

	// Calculate scale/bias pairs by guessing at minimum and maximum values and remapping to [-1,1]
	A, B := -1., 1.

	minvalue := self.offset - 1.0
	maxvalue := self.offset + 1.0
	weightmin := self.gain * minvalue
	weightmax := self.gain * maxvalue

	scale := (B - A) / (maxvalue - minvalue)
	bias := A - minvalue*scale
	self.correct[0][0] = scale
	self.correct[0][1] = bias

	for i := 1; i < int(MaxSources); i++ {
		if weightmin > 1.0 {
			weightmin = 1.0
		}
		if weightmax > 1.0 {
			weightmax = 1.0
		}

		signal := (self.offset - 1.0) * self.exparray[i]
		minvalue += signal * weightmin
		weightmin *= self.gain * signal

		signal = (self.offset + 1.0) * self.exparray[i]
		maxvalue += signal * weightmax
		weightmax *= self.gain * signal

		scale = (B - A) / (maxvalue - minvalue)
		bias = A - minvalue*scale
		self.correct[i][0] = scale
		self.correct[i][1] = bias
	}
}

func (self *ImplicitFractal) FBmGet2D(x, y float64) float64 {
	sum := 0.
	amp := 1.

	x *= self.frequency
	y *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get2D(x, y)
		sum += n * amp
		amp *= self.gain

		x *= self.lacunarity
		y *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) FBmGet3D(x, y, z float64) float64 {
	sum := 0.
	amp := 1.

	x *= self.frequency
	y *= self.frequency
	z *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get3D(x, y, z)
		sum += n * amp
		amp *= self.gain

		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) FBmGet4D(x, y, z, w float64) float64 {
	sum := 0.
	amp := 1.

	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get4D(x, y, z, w)
		sum += n * amp
		amp *= self.gain

		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) FBmGet6D(x, y, z, w, u, v float64) float64 {
	sum := 0.
	amp := 1.

	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency
	u *= self.frequency
	v *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		//n := self.source[i].Get4D(x,y,z,w);
		n := self.source[i].Get6D(x, y, z, w, u, v)
		sum += n * amp
		amp *= self.gain

		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
		u *= self.lacunarity
		v *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) MultiGet2D(x, y float64) float64 {
	value := 1.
	x *= self.frequency
	y *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		value *= self.source[i].Get2D(x, y)*self.exparray[i] + 1.0
		x *= self.lacunarity
		y *= self.lacunarity

	}

	return value*self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1]
}

func (self *ImplicitFractal) MultiGet4D(x, y, z, w float64) float64 {
	value := 1.
	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		value *= self.source[i].Get4D(x, y, z, w)*self.exparray[i] + 1.0
		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
	}

	return value*self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1]
}

func (self *ImplicitFractal) MultiGet3D(x, y, z float64) float64 {
	value := 1.
	x *= self.frequency
	y *= self.frequency
	z *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		value *= self.source[i].Get3D(x, y, z)*self.exparray[i] + 1.0
		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
	}

	return value*self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1]
}

func (self *ImplicitFractal) MultiGet6D(x, y, z, w, u, v float64) float64 {
	value := 1.
	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency
	u *= self.frequency
	v *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		value *= self.source[i].Get6D(x, y, z, w, u, v)*self.exparray[i] + 1.0
		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
		u *= self.lacunarity
		v *= self.lacunarity
	}

	return value*self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1]
}

func (self *ImplicitFractal) BillowGet2D(x, y float64) float64 {
	sum := 0.
	amp := 1.

	x *= self.frequency
	y *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get2D(x, y)
		sum += (2.0*math.Abs(n) - 1.0) * amp
		amp *= self.gain

		x *= self.lacunarity
		y *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) BillowGet4D(x, y, z, w float64) float64 {
	sum := 0.
	amp := 1.

	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get4D(x, y, z, w)
		sum += (2.0*math.Abs(n) - 1.0) * amp
		amp *= self.gain

		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) BillowGet3D(x, y, z float64) float64 {
	sum := 0.
	amp := 1.

	x *= self.frequency
	y *= self.frequency
	z *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get3D(x, y, z)
		sum += (2.0*math.Abs(n) - 1.0) * amp
		amp *= self.gain

		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) BillowGet6D(x, y, z, w, u, v float64) float64 {
	sum := 0.
	amp := 1.

	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency
	u *= self.frequency
	v *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get6D(x, y, z, w, u, v)
		sum += (2.0*math.Abs(n) - 1.0) * amp
		amp *= self.gain

		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
		u *= self.lacunarity
		v *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) RidgedMultiGet2D(x, y float64) float64 {
	sum := 0.
	amp := 1.

	x *= self.frequency
	y *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get2D(x, y)
		n = 1.0 - math.Abs(n)
		sum += amp * n
		amp *= self.gain

		x *= self.lacunarity
		y *= self.lacunarity
	}
	return sum
	/*result,signal := 0.
	  x*=self.frequency;
	  y*=self.frequency;

	  for i:=0; i<self.numoctaves; i++ {
	      signal=self.source[i].Get2D(x,y);
	      signal=self.offset-math.Abs(signal);
	      signal *= signal;
	      result +=signal*exparray[i];

	      x*=self.lacunarity;
	      y*=self.lacunarity;

	  }

	  return result * self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1];*/
}

func (self *ImplicitFractal) RidgedMultiGet4D(x, y, z, w float64) float64 {
	result, signal := 0., 0.
	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		signal = self.source[i].Get4D(x, y, z, w)
		signal = self.offset - math.Abs(signal)
		signal *= signal
		result += signal * self.exparray[i]

		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
	}

	return result*self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1]
}

func (self *ImplicitFractal) RidgedMultiGet3D(x, y, z float64) float64 {
	sum := 0.
	amp := 1.

	x *= self.frequency
	y *= self.frequency
	z *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get3D(x, y, z)
		n = 1.0 - math.Abs(n)
		sum += amp * n
		amp *= self.gain

		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) RidgedMultiGet6D(x, y, z, w, u, v float64) float64 {
	result, signal := 0., 0.
	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency
	u *= self.frequency
	v *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		signal = self.source[i].Get6D(x, y, z, w, u, v)
		signal = self.offset - math.Abs(signal)
		signal *= signal
		result += signal * self.exparray[i]

		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
		u *= self.lacunarity
		v *= self.lacunarity
	}

	return result*self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1]
}

func (self *ImplicitFractal) HybridMultiGet2D(x, y float64) float64 {
	var value, signal, weight float64
	x *= self.frequency
	y *= self.frequency

	value = self.source[0].Get2D(x, y) + self.offset
	weight = self.gain * value
	x *= self.lacunarity
	y *= self.lacunarity

	for i := 1; i < int(self.numoctaves); i++ {
		if weight > 1.0 {
			weight = 1.0
		}
		signal = (self.source[i].Get2D(x, y) + self.offset) * self.exparray[i]
		value += weight * signal
		weight *= self.gain * signal
		x *= self.lacunarity
		y *= self.lacunarity
	}

	return value*self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1]
}

func (self *ImplicitFractal) HybridMultiGet3D(x, y, z float64) float64 {
	var value, signal, weight float64
	x *= self.frequency
	y *= self.frequency
	z *= self.frequency

	value = self.source[0].Get3D(x, y, z) + self.offset
	weight = self.gain * value
	x *= self.lacunarity
	y *= self.lacunarity
	z *= self.lacunarity

	for i := 1; i < int(self.numoctaves); i++ {
		if weight > 1.0 {
			weight = 1.0
		}
		signal = (self.source[i].Get3D(x, y, z) + self.offset) * self.exparray[i]
		value += weight * signal
		weight *= self.gain * signal
		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
	}

	return value*self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1]
}

func (self *ImplicitFractal) HybridMultiGet4D(x, y, z, w float64) float64 {
	var value, signal, weight float64
	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency

	value = self.source[0].Get4D(x, y, z, w) + self.offset
	weight = self.gain * value
	x *= self.lacunarity
	y *= self.lacunarity
	z *= self.lacunarity
	w *= self.lacunarity

	for i := 1; i < int(self.numoctaves); i++ {
		if weight > 1.0 {
			weight = 1.0
		}
		signal = (self.source[i].Get4D(x, y, z, w) + self.offset) * self.exparray[i]
		value += weight * signal
		weight *= self.gain * signal
		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
	}

	return value*self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1]
}

func (self *ImplicitFractal) HybridMultiGet6D(x, y, z, w, u, v float64) float64 {
	var value, signal, weight float64
	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency
	u *= self.frequency
	v *= self.frequency

	value = self.source[0].Get6D(x, y, z, w, u, v) + self.offset
	weight = self.gain * value
	x *= self.lacunarity
	y *= self.lacunarity
	z *= self.lacunarity
	w *= self.lacunarity
	u *= self.lacunarity
	v *= self.lacunarity

	for i := 1; i < int(self.numoctaves); i++ {
		if weight > 1.0 {
			weight = 1.0
		}
		signal = (self.source[i].Get6D(x, y, z, w, u, v) + self.offset) * self.exparray[i]
		value += weight * signal
		weight *= self.gain * signal
		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
		u *= self.lacunarity
		v *= self.lacunarity
	}

	return value*self.correct[self.numoctaves-1][0] + self.correct[self.numoctaves-1][1]
}

func (self *ImplicitFractal) DeCarpentierSwissGet2D(x, y float64) float64 {
	sum, amp := 0., 1.
	dx_sum, dy_sum := 0., 0.

	x *= self.frequency
	y *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get2D(x+self.offset*dx_sum, y+self.offset*dy_sum)
		dx := GetDx2(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum)
		dy := GetDy2(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum)
		sum += amp * (1.0 - math.Abs(n))
		dx_sum += amp * dx * -n
		dy_sum += amp * dy * -n
		amp *= self.gain * math.Max(0., math.Min(sum, 1.))
		x *= self.lacunarity
		y *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) DeCarpentierSwissGet4D(x, y, z, w float64) float64 {
	sum, amp := 0., 1.
	dx_sum, dy_sum, dz_sum, dw_sum := 0., 0., 0., 0.

	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get4D(x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum)
		dx := GetDx4(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum)
		dy := GetDy4(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum)
		dz := GetDz4(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum)
		dw := GetDw4(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum)
		sum += amp * (1.0 - math.Abs(n))
		dx_sum += amp * dx * -n
		dy_sum += amp * dy * -n
		dz_sum += amp * dz * -n
		dw_sum += amp * dw * -n
		amp *= self.gain * math.Max(0., math.Min(sum, 1.))
		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) DeCarpentierSwissGet3D(x, y, z float64) float64 {
	sum, amp := 0., 1.
	dx_sum, dy_sum, dz_sum := 0., 0., 0.

	x *= self.frequency
	y *= self.frequency
	z *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get3D(x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum)
		dx := GetDx3(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum)
		dy := GetDy3(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum)
		dz := GetDz3(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum)
		sum += amp * (1.0 - math.Abs(n))
		dx_sum += amp * dx * -n
		dy_sum += amp * dy * -n
		dz_sum += amp * dz * -n
		amp *= self.gain * math.Max(0., math.Min(sum, 1.))
		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
	}
	return sum
}

func (self *ImplicitFractal) DeCarpentierSwissGet6D(x, y, z, w, u, v float64) float64 {
	sum, amp := 0., 1.
	dx_sum, dy_sum, dz_sum, dw_sum, du_sum, dv_sum := 0., 0., 0., 0., 0., 0.

	x *= self.frequency
	y *= self.frequency
	z *= self.frequency
	w *= self.frequency
	u *= self.frequency
	v *= self.frequency

	for i := 0; i < int(self.numoctaves); i++ {
		n := self.source[i].Get6D(x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum, u+self.offset*du_sum, v+self.offset*dv_sum)
		dx := GetDx6(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dx_sum, w+self.offset*dw_sum, u+self.offset*du_sum, v+self.offset*dv_sum)
		dy := GetDy6(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum, u+self.offset*du_sum, v+self.offset*dv_sum)
		dz := GetDz6(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum, u+self.offset*du_sum, v+self.offset*dv_sum)
		dw := GetDw6(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum, u+self.offset*du_sum, v+self.offset*dv_sum)
		du := GetDu6(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum, u+self.offset*du_sum, v+self.offset*dv_sum)
		dv := GetDv6(self.source[i], x+self.offset*dx_sum, y+self.offset*dy_sum, z+self.offset*dz_sum, w+self.offset*dw_sum, u+self.offset*du_sum, v+self.offset*dv_sum)
		sum += amp * (1.0 - math.Abs(n))
		dx_sum += amp * dx * -n
		dy_sum += amp * dy * -n
		dz_sum += amp * dz * -n
		dw_sum += amp * dw * -n
		du_sum += amp * du * -n
		dv_sum += amp * dv * -n
		amp *= self.gain * math.Max(0., math.Min(sum, 1.))
		x *= self.lacunarity
		y *= self.lacunarity
		z *= self.lacunarity
		w *= self.lacunarity
		u *= self.lacunarity
		v *= self.lacunarity
	}
	return sum
}
