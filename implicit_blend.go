package anl

import ()

type ImplicitBlend struct {
	ImplicitModuleBase
	low, high, control ScalarParameter
}

func (this *ImplicitBlend) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func (this *ImplicitBlend) SetSeed(seed uint32){}

func (this *ImplicitBlend) SetLowSourceModule(b ImplicitModule) {
	this.low.SetModule(b)
}

func (this *ImplicitBlend) SetHighSourceModule(b ImplicitModule) {
	this.high.SetModule(b)
}

func (this *ImplicitBlend) SetControlSourceModule(b ImplicitModule) {
	this.control.SetModule(b)
}

func (this *ImplicitBlend) SetLowSourceValue(v float64) {
	this.low.SetValue(v)
}

func (this *ImplicitBlend) SetHighSourceValue(v float64) {
	this.high.SetValue(v)
}

func (this *ImplicitBlend) SetControlSourceValue(v float64) {
	this.control.SetValue(v)
}

func (this *ImplicitBlend) Get2D(x, y float64) float64 {
	v1 := this.low.Get2D(x, y)
	v2 := this.high.Get2D(x, y)
	blend := this.control.Get2D(x, y)
	blend = (blend + 1.0) * 0.5

	return Lerp(blend, v1, v2)
}

func (this *ImplicitBlend) Get3D(x, y, z float64) float64 {
	v1 := this.low.Get3D(x, y, z)
	v2 := this.high.Get3D(x, y, z)
	blend := this.control.Get3D(x, y, z)
	return Lerp(blend, v1, v2)
}

func (this *ImplicitBlend) Get4D(x, y, z, w float64) float64 {
	v1 := this.low.Get4D(x, y, z, w)
	v2 := this.high.Get4D(x, y, z, w)
	blend := this.control.Get4D(x, y, z, w)
	return Lerp(blend, v1, v2)
}

func (this *ImplicitBlend) Get6D(x, y, z, w, u, v float64) float64 {
	v1 := this.low.Get6D(x, y, z, w, u, v)
	v2 := this.high.Get6D(x, y, z, w, u, v)
	blend := this.control.Get6D(x, y, z, w, u, v)
	return Lerp(blend, v1, v2)
}
