package anl

import ()

type ImplicitScaleDomain struct {
	ImplicitModuleBase
	source, sx, sy, sz, sw, su, sv ScalarParameter
}

func (this *ImplicitScaleDomain) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func (this *ImplicitScaleDomain) SetSeed(seed uint32){}

func NewImplicitScaleDomainEmpty() *ImplicitScaleDomain {
	return NewImplicitScaleDomain(1, 1, 1, 1, 1, 1)
}

func NewImplicitScaleDomain(x, y, z, w, u, v float64) *ImplicitScaleDomain {
	return &ImplicitScaleDomain{
		sx: NewScalarParameter(x),
		sy: NewScalarParameter(y),
		sz: NewScalarParameter(z),
		sw: NewScalarParameter(w),
		su: NewScalarParameter(u),
		sv: NewScalarParameter(v),
	}
}

func (this *ImplicitScaleDomain) SetScale(x, y, z, w, u, v float64) {
	this.sx.SetValue(x)
	this.sy.SetValue(y)
	this.sz.SetValue(z)
	this.sw.SetValue(w)
	this.su.SetValue(u)
	this.sv.SetValue(v)
}

func (this *ImplicitScaleDomain) SetXScaleValue(x float64)   { this.sx.SetValue(x) }
func (this *ImplicitScaleDomain) SetYScaleValue(x float64)   { this.sy.SetValue(x) }
func (this *ImplicitScaleDomain) SetZScaleValue(x float64)   { this.sz.SetValue(x) }
func (this *ImplicitScaleDomain) SetWScaleValue(x float64)   { this.sw.SetValue(x) }
func (this *ImplicitScaleDomain) SetUScaleValue(x float64)   { this.su.SetValue(x) }
func (this *ImplicitScaleDomain) SetVScaleValue(x float64)   { this.sv.SetValue(x) }
func (this *ImplicitScaleDomain) SetXScale(x ImplicitModule) { this.sx.SetModule(x) }
func (this *ImplicitScaleDomain) SetYScale(y ImplicitModule) { this.sy.SetModule(y) }
func (this *ImplicitScaleDomain) SetZScale(z ImplicitModule) { this.sz.SetModule(z) }
func (this *ImplicitScaleDomain) SetWScale(w ImplicitModule) { this.sw.SetModule(w) }
func (this *ImplicitScaleDomain) SetUScale(u ImplicitModule) { this.su.SetModule(u) }
func (this *ImplicitScaleDomain) SetVScale(v ImplicitModule) { this.sv.SetModule(v) }

func (this *ImplicitScaleDomain) SetSourceModule(m ImplicitModule) {
	this.source.SetModule(m)
}

func (this *ImplicitScaleDomain) SetSourceValue(v float64) {
	this.source.SetValue(v)
}

func (this *ImplicitScaleDomain) Get2D(x, y float64) float64 {
	return this.source.Get2D(x*this.sx.Get2D(x, y), y*this.sy.Get2D(x, y))
}

func (this *ImplicitScaleDomain) Get3D(x, y, z float64) float64 {
	return this.source.Get3D(x*this.sx.Get3D(x, y, z), y*this.sy.Get3D(x, y, z), z*this.sz.Get3D(x, y, z))
}

func (this *ImplicitScaleDomain) Get4D(x, y, z, w float64) float64 {
	return this.source.Get4D(x*this.sx.Get4D(x, y, z, w), y*this.sy.Get4D(x, y, z, w), z*this.sz.Get4D(x, y, z, w), w*this.sw.Get4D(x, y, z, w))
}

func (this *ImplicitScaleDomain) Get6D(x, y, z, w, u, v float64) float64 {
	return this.source.Get6D(x*this.sx.Get6D(x, y, z, w, u, v), y*this.sy.Get6D(x, y, z, w, u, v), z*this.sz.Get6D(x, y, z, w, u, v), w*this.sw.Get6D(x, y, z, w, u, v), u*this.su.Get6D(x, y, z, w, u, v), v*this.sv.Get6D(x, y, z, w, u, v))
}
