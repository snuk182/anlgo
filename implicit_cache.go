package anl

import ()

type Cache struct {
	x, y, z, w, u, v, val float64
	valid                 bool
}

func NewCache() *Cache {
	return &Cache{valid: false}
}

type ImplicitCache struct {
	ImplicitModuleBase
	source               ScalarParameter
	c2, c3, c4, c6 Cache
}

func (this *ImplicitCache) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func (this *ImplicitCache) SetSeed(seed uint32){}

func (this *ImplicitCache) SetSourceModule(m ImplicitModule) {
	this.source.SetModule(m)
}
func (this *ImplicitCache) SetSource(v float64) {
	this.source.SetValue(v)
}
func (this *ImplicitCache) Get2D(x, y float64) float64 {
	if !this.c2.valid || this.c2.x != x || this.c2.y != y {
		this.c2.x = x
		this.c2.y = y
		this.c2.valid = true
		this.c2.val = this.source.Get2D(x, y)
	}
	return this.c2.val
}

func (this *ImplicitCache) Get3D(x, y, z float64) float64 {
	if !this.c3.valid || this.c3.x != x || this.c3.y != y || this.c3.z != z {
		this.c3.x = x
		this.c3.y = y
		this.c3.z = z
		this.c3.valid = true
		this.c3.val = this.source.Get3D(x, y, z)
	}
	return this.c3.val
}

func (this *ImplicitCache) Get4D(x, y, z, w float64) float64 {
	if !this.c4.valid || this.c4.x != x || this.c4.y != y || this.c4.z != z || this.c4.w != w {
		this.c4.x = x
		this.c4.y = y
		this.c4.z = z
		this.c4.w = w
		this.c4.valid = true
		this.c4.val = this.source.Get4D(x, y, z, w)
	}
	return this.c4.val
}

func (this *ImplicitCache) Get6D(x, y, z, w, u, v float64) float64 {
	if !this.c6.valid || this.c6.x != x || this.c6.y != y || this.c6.z != z || this.c6.w != w || this.c6.u != u || this.c6.v != v {
		this.c6.x = x
		this.c6.y = y
		this.c6.z = z
		this.c6.w = w
		this.c6.u = u
		this.c6.v = v
		this.c6.valid = true
		this.c6.val = this.source.Get6D(x, y, z, w, u, v)
	}
	return this.c6.val
}
