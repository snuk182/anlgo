package anl

import ()

type ImplicitGradient struct {
	ImplicitModuleBase
	m_gx1, m_gy1, m_gz1, m_gw1, m_gu1, m_gv1 float64
	m_gx2, m_gy2, m_gz2, m_gw2, m_gu2, m_gv2 float64
	m_x, m_y, m_z, m_w, m_u, m_v             float64
	m_vlen                                   float64
}

func (this *ImplicitGradient) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func NewImplicitGradient() *ImplicitGradient {
	v := new(ImplicitGradient)
	v.SetGradient(0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0)
	return v
}

func (this *ImplicitGradient) SetGradient(x1, x2, y1, y2, z1, z2, w1, w2, u1, u2, v1, v2 float64) {
	this.m_gx1 = x1
	this.m_gx2 = x2
	this.m_gy1 = y1
	this.m_gy2 = y2
	this.m_gz1 = z1
	this.m_gz2 = z2
	this.m_gw1 = w1
	this.m_gw2 = w2
	this.m_gu1 = u1
	this.m_gu2 = u2
	this.m_gv1 = v1
	this.m_gv2 = v2

	this.m_x = x2 - x1
	this.m_y = y2 - y1
	this.m_z = z2 - z1
	this.m_w = w2 - w1
	this.m_u = u2 - u1
	this.m_v = v2 - v1

	this.m_vlen = (this.m_x*this.m_x + this.m_y*this.m_y + this.m_z*this.m_z + this.m_w*this.m_w + this.m_u*this.m_u + this.m_v*this.m_v)
}

func (this *ImplicitGradient) Get2D(x, y float64) float64 {
	// Subtract from (1) and take dotprod
	dx := x - this.m_gx1
	dy := y - this.m_gy1
	dp := dx*this.m_x + dy*this.m_y
	dp /= this.m_vlen
	//dp=clamp(dp/this.m_vlen,0.0,1.0);
	//return lerp(dp,1.0,-1.0);
	return dp
}

func (this *ImplicitGradient) Get3D(x, y, z float64) float64 {
	dx := x - this.m_gx1
	dy := y - this.m_gy1
	dz := z - this.m_gz1
	dp := dx*this.m_x + dy*this.m_y + dz*this.m_z
	dp /= this.m_vlen
	//dp=clamp(dp/this.m_vlen,0.0,1.0);
	//return lerp(dp,1.0,-1.0);
	return dp
}

func (this *ImplicitGradient) Get4D(x, y, z, w float64) float64 {
	dx := x - this.m_gx1
	dy := y - this.m_gy1
	dz := z - this.m_gz1
	dw := w - this.m_gw1
	dp := dx*this.m_x + dy*this.m_y + dz*this.m_z + dw*this.m_w
	dp /= this.m_vlen
	//dp=clamp(dp/this.m_vlen,0.0,1.0);
	//return lerp(dp,1.0,-1.0);
	return dp
}

func (this *ImplicitGradient) Get6D(x, y, z, w, u, v float64) float64 {
	dx := x - this.m_gx1
	dy := y - this.m_gy1
	dz := z - this.m_gz1
	dw := w - this.m_gw1
	du := u - this.m_gu1
	dv := v - this.m_gv1
	dp := dx*this.m_x + dy*this.m_y + dz*this.m_z + dw*this.m_w + du*this.m_u + dv*this.m_v
	dp /= this.m_vlen
	//dp=clamp(dp/this.m_vlen,0.0,1.0);
	//return lerp(clamp(dp,0.0,1.0),1.0,-1.0);
	return dp
}
