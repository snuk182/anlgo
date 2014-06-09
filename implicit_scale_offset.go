package anl

import ()

type ImplicitScaleOffset struct {
	ImplicitBasisFunction
	source, scale, offset ScalarParameter
}

func NewImplicitScaleOffset(scale, offset float64) (*ImplicitScaleOffset) {
	return &ImplicitScaleOffset{scale: NewScalarParameter(scale), offset: NewScalarParameter(offset)}
}

func (this *ImplicitScaleOffset) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func (this *ImplicitScaleOffset) SetSourceModule(b ImplicitModule) {
	this.source.SetModule(b)
}
func (this *ImplicitScaleOffset) SetSourceValue(v float64) {
	this.source.SetValue(v)
}

func (this *ImplicitScaleOffset) SetScaleValue(scale float64) {
	this.scale.SetValue(scale)
}
func (this *ImplicitScaleOffset) SetOffsetValue(offset float64) {
	this.offset.SetValue(offset)
}
func (this *ImplicitScaleOffset) SetScaleModule(scale ImplicitModule) {
	this.scale.SetModule(scale)
}
func (this *ImplicitScaleOffset) SetOffsetModule(offset ImplicitModule) {
	this.offset.SetModule(offset)
}

func (this *ImplicitScaleOffset) Get2D(x, y float64) float64 {
	return this.source.Get2D(x, y)*this.scale.Get2D(x, y) + this.offset.Get2D(x, y)
}

func (this *ImplicitScaleOffset) Get3D(x, y, z float64) float64 {
	return this.source.Get3D(x, y, z)*this.scale.Get3D(x, y, z) + this.offset.Get3D(x, y, z)
}

func (this *ImplicitScaleOffset) Get4D(x,y,z,w float64) float64 {
	return this.source.Get4D(x, y, z, w)*this.scale.Get4D(x, y, z, w) + this.offset.Get4D(x, y, z, w)
}

func (this *ImplicitScaleOffset) Get6D(x, y, z, w, u, v float64) float64 {
	return this.source.Get6D(x, y, z, w, u, v)*this.scale.Get6D(x, y, z, w, u, v) + this.offset.Get6D(x, y, z, w, u, v)
}
