package anl

type ImplicitAutoCorrect struct {
	ImplicitModuleBase
	source          ImplicitModule
	low, high       float64
	scale2, offset2 float64
	scale3, offset3 float64
	scale4, offset4 float64
	scale6, offset6 float64
}

func NewImplicitAutoCorrect(low, high float64) *ImplicitAutoCorrect {
	return &ImplicitAutoCorrect{
		low:  low,
		high: high,
	}
}

func NewImplicitAutoCorrectEmpty() *ImplicitAutoCorrect {
	return NewImplicitAutoCorrect(-1., 1.)
}

func (this *ImplicitAutoCorrect) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func (this *ImplicitAutoCorrect) SetSource(m ImplicitModule) {
	this.source = m
	this.Calculate()
}

func (this *ImplicitAutoCorrect) SetRange(low, high float64) {
	this.low = low
	this.high = high
	this.Calculate()
}

func (this *ImplicitAutoCorrect) Calculate() {
	if this.source == nil {
		return
	}
	var mn, mx float64
	lcg := NewLCG()
	//lcg.setSeedTime();

	// Calculate 2D
	mn = 10000.0
	mx = -10000.0
	for c := 0; c < 10000; c++ {
		nx := Get01(lcg)*4.0 - 2.0
		ny := Get01(lcg)*4.0 - 2.0

		v := this.source.Get2D(nx, ny)
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}
	this.scale2 = (this.high - this.low) / (mx - mn)
	this.offset2 = this.low - mn*this.scale2

	// Calculate 3D
	mn = 10000.0
	mx = -10000.0
	for c := 0; c < 10000; c++ {
		nx := Get01(lcg)*4.0 - 2.0
		ny := Get01(lcg)*4.0 - 2.0
		nz := Get01(lcg)*4.0 - 2.0

		v := this.source.Get3D(nx, ny, nz)
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}
	this.scale3 = (this.high - this.low) / (mx - mn)
	this.offset3 = this.low - mn*this.scale3

	// Calculate 4D
	mn = 10000.0
	mx = -10000.0
	for c := 0; c < 10000; c++ {
		nx := Get01(lcg)*4.0 - 2.0
		ny := Get01(lcg)*4.0 - 2.0
		nz := Get01(lcg)*4.0 - 2.0
		nw := Get01(lcg)*4.0 - 2.0

		v := this.source.Get4D(nx, ny, nz, nw)
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}
	this.scale4 = (this.high - this.low) / (mx - mn)
	this.offset4 = this.low - mn*this.scale4

	// Calculate 6D
	mn = 10000.0
	mx = -10000.0
	for c := 0; c < 10000; c++ {
		nx := Get01(lcg)*4.0 - 2.0
		ny := Get01(lcg)*4.0 - 2.0
		nz := Get01(lcg)*4.0 - 2.0
		nw := Get01(lcg)*4.0 - 2.0
		nu := Get01(lcg)*4.0 - 2.0
		nv := Get01(lcg)*4.0 - 2.0

		v := this.source.Get6D(nx, ny, nz, nw, nu, nv)
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}
	this.scale6 = (this.high - this.low) / (mx - mn)
	this.offset6 = this.low - mn*this.scale6
}

func (this *ImplicitAutoCorrect) Get2D(x, y float64) float64 {
	if this.source == nil {
		return 0.
	}

	v := this.source.Get2D(x, y)
	return Clamp(v*this.scale2+this.offset2, this.low, this.high)
}

func (this *ImplicitAutoCorrect) Get3D(x, y, z float64) float64 {
	if this.source == nil {
		return 0.
	}

	v := this.source.Get3D(x, y, z)
	return Clamp(v*this.scale3+this.offset3, this.low, this.high)
}
func (this *ImplicitAutoCorrect) Get4D(x, y, z, w float64) float64 {
	if this.source == nil {
		return 0.
	}

	v := this.source.Get4D(x, y, z, w)
	return Clamp(v*this.scale4+this.offset4, this.low, this.high)
}

func (this *ImplicitAutoCorrect) Get6D(x, y, z, w, u, v float64) float64 {
	if this.source == nil {
		return 0.
	}

	val := this.source.Get6D(x, y, z, w, u, v)
	return Clamp(val*this.scale6+this.offset6, this.low, this.high)
}
