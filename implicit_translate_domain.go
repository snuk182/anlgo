package anl

type ImplicitTranslateDomain struct {
	ImplicitModuleBase
	source, ax, ay, az, aw, au, av ScalarParameter
}

func NewImplicitTranslateDomain() *ImplicitTranslateDomain {
	return &ImplicitTranslateDomain{
		source: NewScalarParameter(0),
		ax: NewScalarParameter(0),
		ay: NewScalarParameter(0),
		az: NewScalarParameter(0),
		aw: NewScalarParameter(0),
		au: NewScalarParameter(0),
		av: NewScalarParameter(0),
	}
}

func (this *ImplicitTranslateDomain) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func (this *ImplicitTranslateDomain) SetSeed(seed uint32){}

func (this *ImplicitTranslateDomain) SetXAxisSourceModule(m ImplicitModule) {
	this.ax.SetModule(m)
}
func (this *ImplicitTranslateDomain) SetYAxisSourceModule(m ImplicitModule) {
	this.ay.SetModule(m)
}
func (this *ImplicitTranslateDomain) SetZAxisSourceModule(m ImplicitModule) {
	this.az.SetModule(m)
}
func (this *ImplicitTranslateDomain) SetWAxisSourceModule(m ImplicitModule) {
	this.aw.SetModule(m)
}
func (this *ImplicitTranslateDomain) SetUAxisSourceModule(m ImplicitModule) {
	this.au.SetModule(m)
}
func (this *ImplicitTranslateDomain) SetVAxisSourceModule(m ImplicitModule) {
	this.av.SetModule(m)
}
func (this *ImplicitTranslateDomain) SetXAxisSourceValue(v float64) {
	this.ax.SetValue(v)
}
func (this *ImplicitTranslateDomain) SetYAxisSourceValue(v float64) {
	this.ay.SetValue(v)
}
func (this *ImplicitTranslateDomain) SetZAxisSourceValue(v float64) {
	this.az.SetValue(v)
}
func (this *ImplicitTranslateDomain) SetWAxisSourceValue(v float64) {
	this.aw.SetValue(v)
}
func (this *ImplicitTranslateDomain) SetUAxisSourceValue(v float64) {
	this.au.SetValue(v)
}
func (this *ImplicitTranslateDomain) SetVAxisSourceValue(v float64) {
	this.av.SetValue(v)
}
func (this *ImplicitTranslateDomain) SetSourceModule(m ImplicitModule) {
	this.source.SetModule(m)
}
func (this *ImplicitTranslateDomain) SetSourceValue(v float64) {
	this.source.SetValue(v)
}

func (this *ImplicitTranslateDomain) Get2D(x, y float64) float64 {
	return this.source.Get2D(x+this.ax.Get2D(x, y), y+this.ay.Get2D(x, y))
}
func (this *ImplicitTranslateDomain) Get3D(x, y, z float64) float64 {
	return this.source.Get3D(x+this.ax.Get3D(x, y, z), y+this.ay.Get3D(x, y, z), z+this.az.Get3D(x, y, z))
}
func (this *ImplicitTranslateDomain) Get4D(x, y, z, w float64) float64 {
	return this.source.Get4D(x+this.ax.Get4D(x, y, z, w), y+this.ay.Get4D(x, y, z, w), z+this.az.Get4D(x, y, z, w), w+this.aw.Get4D(x, y, z, w))
}
func (this *ImplicitTranslateDomain) Get6D(x, y, z, w, u, v float64) float64 {
	return this.source.Get6D(x+this.ax.Get6D(x, y, z, w, u, v), y+this.ay.Get6D(x, y, z, w, u, v), z+this.az.Get6D(x, y, z, w, u, v),
		w+this.aw.Get6D(x, y, z, w, u, v), u+this.au.Get6D(x, y, z, w, u, v), v+this.av.Get6D(x, y, z, w, u, v))
}
