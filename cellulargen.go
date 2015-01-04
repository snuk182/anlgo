package anl

import (

)

type CellularCache struct {
    f,d [4]float64
    x,y,z,w,u,v float64
    valid bool
};

func NewCellularCache() CellularCache {
	return CellularCache{
		valid: false,
	}
}

type CellularGenerator struct {
    seed uint32
    cache2, cache3, cache4, cache6 CellularCache
};

func NewCellularGenerator() *CellularGenerator {
	return &CellularGenerator{
		seed: 1000,
	}
}

func (this *CellularGenerator) Get2D(x,y float64) CellularCache {
    if(!this.cache2.valid || x!=this.cache2.x || y!=this.cache2.y) {
        CellularFunction2D(x,y,this.seed,this.cache2.f[:],this.cache2.d[:]);
        this.cache2.x=x;
        this.cache2.y=y;
        this.cache2.valid=true;
    }
    return this.cache2;
}

func (this *CellularGenerator) Get3D(x,y,z float64) CellularCache {
    if(!this.cache3.valid || x!=this.cache3.x || y!=this.cache3.y || z!=this.cache3.z) {
        CellularFunction3D(x,y,z,this.seed,this.cache3.f[:],this.cache3.d[:]);
        this.cache3.x=x;
        this.cache3.y=y;
        this.cache3.z=z;
        this.cache3.valid=true;
    }
    return this.cache3;
}

func (this *CellularGenerator) Get4D(x,y,z,w float64) CellularCache {
    if(!this.cache4.valid || x!=this.cache4.x || y!=this.cache4.y || z!=this.cache4.z || w!=this.cache4.w) {
        CellularFunction4D(x,y,z,w,this.seed,this.cache4.f[:],this.cache4.d[:]);
        this.cache4.x=x;
        this.cache4.y=y;
        this.cache4.z=z;
        this.cache4.w=w;
        this.cache4.valid=true;
    }
    return this.cache4;
}

func (this *CellularGenerator) Get6D(x,y,z,w,u,v float64) CellularCache {
    if(this.cache6.valid || x!=this.cache6.x || y!=this.cache6.y || z!=this.cache6.z || w!=this.cache6.w || u!=this.cache6.u || v!=this.cache6.v) {
        CellularFunction6D(x,y,z,w,u,v,this.seed,this.cache6.f[:],this.cache6.d[:]);
        this.cache6.x=x;
        this.cache6.y=y;
        this.cache6.z=z;
        this.cache6.w=w;
        this.cache6.u=u;
        this.cache6.v=v;
        this.cache6.valid=true;
    }

    return this.cache6;
}

func (this *CellularGenerator) setSeed(seed uint32) {
    this.seed=seed;
    this.cache2.valid=false;
    this.cache3.valid=false;
    this.cache4.valid=false;
    this.cache6.valid=false;
}