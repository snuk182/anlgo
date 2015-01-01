package anl

import ()

type ImplicitSelect struct {
	ImplicitModuleBase
	low, high, control, threshold, falloff ScalarParameter
}

func NewImplicitSelect() *ImplicitSelect {
	return &ImplicitSelect{
		low: NewScalarParameter(0),
		high: NewScalarParameter(0),
		control: NewScalarParameter(0),
		threshhold: NewScalarParameter(0),
		falloff: NewScalarParameter(0),
	}
}

func (this *ImplicitSelect) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func (this *ImplicitSelect) SetSeed(seed uint32){}

func (this *ImplicitSelect) SetLowSourceModule(b ImplicitModule) {
	this.low.SetModule(b)
}
func (this *ImplicitSelect) SetHighSourceModule(b ImplicitModule) {
	this.high.SetModule(b)
}
func (this *ImplicitSelect) SetControlSourceModule(b ImplicitModule) {
	this.control.SetModule(b)
}

func (this *ImplicitSelect) SetLowSourceValue(b float64) {
	this.low.SetValue(b)
}
func (this *ImplicitSelect) SetHighSourceValue(b float64) {
	this.high.SetValue(b)
}
func (this *ImplicitSelect) SetControlSourceValue(b float64) {
	this.control.SetValue(b)
}

func (this *ImplicitSelect) SetThresholdValue(t float64) {
	//this.threshold=t;
	this.threshold.SetValue(t)
}
func (this *ImplicitSelect) SetFalloffValue(f float64) {
	//this.falloff=f;
	this.falloff.SetValue(f)
}
func (this *ImplicitSelect) SetThresholdModule(m ImplicitModule) {
	this.threshold.SetModule(m)
}
func (this *ImplicitSelect) SetFalloffModule(m ImplicitModule) {
	this.falloff.SetModule(m)
}

func (this *ImplicitSelect) Get2D(x, y float64) float64 {
	control := this.control.Get2D(x, y)
	falloff := this.falloff.Get2D(x, y)
	threshold := this.threshold.Get2D(x, y)

	if falloff > 0. {
		switch {
		case control < (threshold - falloff):
			return this.low.Get2D(x, y)
		case control > (threshold + falloff):
			return this.high.Get2D(x, y)
		default:
			lower := threshold - falloff
			upper := threshold + falloff
			blend := QuinticBlend((control - lower) / (upper - lower))
			return Lerp(blend, this.low.Get2D(x, y), this.high.Get2D(x, y))
		}
	} else {
		if control < threshold {
			return this.low.Get2D(x, y)
		} else {
			return this.high.Get2D(x, y)
		}
	}
}

func (this *ImplicitSelect) Get3D(x, y, z float64) float64 {
	control := this.control.Get3D(x, y, z)
	falloff := this.falloff.Get3D(x, y, z)
	threshold := this.threshold.Get3D(x, y, z)

	if falloff > 0. {
		switch {
		case control < (threshold - falloff):
			return this.low.Get3D(x, y, z)
		case control > (threshold + falloff):
			return this.high.Get3D(x, y, z)
		default:
			lower := threshold - falloff
			upper := threshold + falloff
			blend := QuinticBlend((control - lower) / (upper - lower))
			return Lerp(blend, this.low.Get3D(x, y, z), this.high.Get3D(x, y, z))
		}
	} else {
		if control < threshold {
			return this.low.Get3D(x, y, z)
		} else {
			return this.high.Get3D(x, y, z)
		}
	}
}

func (this *ImplicitSelect) Get4D(x, y, z, w float64) float64 {
	control := this.control.Get4D(x, y, z, w)
	falloff := this.falloff.Get4D(x, y, z, w)
	threshold := this.threshold.Get4D(x, y, z, w)

	if falloff > 0. {
		switch {
		case control < (threshold - falloff):
			return this.low.Get4D(x, y, z, w)
		case control > (threshold + falloff):
			return this.high.Get4D(x, y, z, w)
		default:
			lower := threshold - falloff
			upper := threshold + falloff
			blend := QuinticBlend((control - lower) / (upper - lower))
			return Lerp(blend, this.low.Get4D(x, y, z, w), this.high.Get4D(x, y, z, w))
		}
	} else {
		if control < threshold {
			return this.low.Get4D(x, y, z, w)
		} else {
			return this.high.Get4D(x, y, z, w)
		}
	}
}

func (this *ImplicitSelect) Get6D(x, y, z, w, u, v float64) float64 {
	control := this.control.Get6D(x, y, z, w, u, v)
	falloff := this.falloff.Get6D(x, y, z, w, u, v)
	threshold := this.threshold.Get6D(x, y, z, w, u, v)

	if falloff > 0. {
		switch {
		case control < (threshold - falloff):
			return this.low.Get6D(x, y, z, w, u, v)
		case control > (threshold + falloff):
			return this.high.Get6D(x, y, z, w, u, v)
		default:
			lower := threshold - falloff
			upper := threshold + falloff
			blend := QuinticBlend((control - lower) / (upper - lower))
			return Lerp(blend, this.low.Get6D(x, y, z, w, u, v), this.high.Get6D(x, y, z, w, u, v))
		}
	} else {
		if control < threshold {
			return this.low.Get6D(x, y, z, w, u, v)
		} else {
			return this.high.Get6D(x, y, z, w, u, v)
		}
	}
}
