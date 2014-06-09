package anl

import ()

type ImplicitGradient struct {
	ImplicitModuleBase
	gx1, gy1, gz1, gw1, gu1, gv1 float64
	gx2, gy2, gz2, gw2, gu2, gv2 float64
	x, y, z, w, u, v             float64
	vlen                                   float64
}

func (this *ImplicitGradient) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func (this *ImplicitGradient) SetSeed(seed uint32){}

func NewImplicitGradient() *ImplicitGradient {
	v := new(ImplicitGradient)
	v.SetGradient(0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0)
	return v
}

func (this *ImplicitGradient) SetGradient(x1, x2, y1, y2, z1, z2, w1, w2, u1, u2, v1, v2 float64) {
	this.gx1 = x1
	this.gx2 = x2
	this.gy1 = y1
	this.gy2 = y2
	this.gz1 = z1
	this.gz2 = z2
	this.gw1 = w1
	this.gw2 = w2
	this.gu1 = u1
	this.gu2 = u2
	this.gv1 = v1
	this.gv2 = v2

	this.x = x2 - x1
	this.y = y2 - y1
	this.z = z2 - z1
	this.w = w2 - w1
	this.u = u2 - u1
	this.v = v2 - v1

	this.vlen = (this.x*this.x + this.y*this.y + this.z*this.z + this.w*this.w + this.u*this.u + this.v*this.v)
}

func (this *ImplicitGradient) Get2D(x, y float64) float64 {
	// Subtract from (1) and take dotprod
	dx := x - this.gx1
	dy := y - this.gy1
	dp := dx*this.x + dy*this.y
	dp /= this.vlen
	//dp=clamp(dp/this.vlen,0.0,1.0);
	//return lerp(dp,1.0,-1.0);
	return dp
}

func (this *ImplicitGradient) Get3D(x, y, z float64) float64 {
	dx := x - this.gx1
	dy := y - this.gy1
	dz := z - this.gz1
	dp := dx*this.x + dy*this.y + dz*this.z
	dp /= this.vlen
	//dp=clamp(dp/this.vlen,0.0,1.0);
	//return lerp(dp,1.0,-1.0);
	return dp
}

func (this *ImplicitGradient) Get4D(x, y, z, w float64) float64 {
	dx := x - this.gx1
	dy := y - this.gy1
	dz := z - this.gz1
	dw := w - this.gw1
	dp := dx*this.x + dy*this.y + dz*this.z + dw*this.w
	dp /= this.vlen
	//dp=clamp(dp/this.vlen,0.0,1.0);
	//return lerp(dp,1.0,-1.0);
	return dp
}

func (this *ImplicitGradient) Get6D(x, y, z, w, u, v float64) float64 {
	dx := x - this.gx1
	dy := y - this.gy1
	dz := z - this.gz1
	dw := w - this.gw1
	du := u - this.gu1
	dv := v - this.gv1
	dp := dx*this.x + dy*this.y + dz*this.z + dw*this.w + du*this.u + dv*this.v
	dp /= this.vlen
	//dp=clamp(dp/this.vlen,0.0,1.0);
	//return lerp(clamp(dp,0.0,1.0),1.0,-1.0);
	return dp
}
