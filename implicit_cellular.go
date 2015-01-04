package anl

import ()

type ImplicitCellular struct {
	ImplicitModuleBase
	generator    *CellularGenerator
	coefficients [4]float64
}

func NewImplicitCellularEmpty() *ImplicitCellular {
	this := new(ImplicitCellular)
	this.SetCoefficients(1, 0, 0, 0)
	return this
}
func NewImplicitCellular(a, b, c, d float64) *ImplicitCellular {
	this := new(ImplicitCellular)
	this.SetCoefficients(a, b, c, d)
	return this
}

func (this *ImplicitCellular) SetCoefficients(a, b, c, d float64) {
	this.coefficients[0] = a
	this.coefficients[1] = b
	this.coefficients[2] = c
	this.coefficients[3] = d
}

func (this *ImplicitCellular) SetCellularSource(m *CellularGenerator) {
	this.generator = m
}

func (this *ImplicitCellular) Get2D(x, y float64) float64 {
	if this.generator == nil {
		return 0.0
	}

	c := this.generator.Get2D(x, y)
	return c.f[0]*this.coefficients[0] + c.f[1]*this.coefficients[1] + c.f[2]*this.coefficients[2] + c.f[3]*this.coefficients[3]
}

func (this *ImplicitCellular) Get3D(x, y, z float64) float64 {
	if this.generator == nil {
		return 0.0
	}

	c := this.generator.Get3D(x, y, z)
	return c.f[0]*this.coefficients[0] + c.f[1]*this.coefficients[1] + c.f[2]*this.coefficients[2] + c.f[3]*this.coefficients[3]
}

func (this *ImplicitCellular) Get4D(x, y, z, w float64) float64 {
	if this.generator == nil {
		return 0.0
	}

	c := this.generator.Get4D(x, y, z, w)
	return c.f[0]*this.coefficients[0] + c.f[1]*this.coefficients[1] + c.f[2]*this.coefficients[2] + c.f[3]*this.coefficients[3]
}

func (this *ImplicitCellular) Get6D(x, y, z, w, u, v float64) float64 {
	if this.generator == nil {
		return 0.0
	}

	c := this.generator.Get6D(x, y, z, w, u, v)
	return c.f[0]*this.coefficients[0] + c.f[1]*this.coefficients[1] + c.f[2]*this.coefficients[2] + c.f[3]*this.coefficients[3]
}

func (this *ImplicitCellular) SetSeed(seed uint32) {
	if this.generator != nil {
		this.generator.SetSeed(seed)
	}
}
