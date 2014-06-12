package anl

type ImplicitBias struct {
	ImplicitModuleBase
	source, bias ScalarParameter
}

func NewImplicitBias(b float64) (*ImplicitBias) {
	return &ImplicitBias{bias: NewScalarParameter(b)}
}

func (this *ImplicitBias) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func (this *ImplicitBias) SetSeed(seed uint32){}

func (this *ImplicitBias) SetSourceModule(b ImplicitModule) {
	this.source.SetModule(b)
}
func (this *ImplicitBias) SetSource(v float64) {
	this.source.SetValue(v)
}
func (this *ImplicitBias) SetBias(b float64) {
	this.bias.SetValue(b)
}
func (this *ImplicitBias) SetBiasModule(m ImplicitModule) {
	this.bias.SetModule(m)
}

func (this *ImplicitBias) Get2D(x,y float64) float64 {
        va := this.source.Get2D(x,y);
        return Bias(this.bias.Get2D(x,y), va);
    }

func (this *ImplicitBias) Get3D(x,y,z float64) float64 {
        va := this.source.Get3D(x,y,z);
        return Bias(this.bias.Get3D(x,y,z), va);
    }

func (this *ImplicitBias) Get4D(x,y,z,w float64) float64 {
        va := this.source.Get4D(x,y,z,w);
        return Bias(this.bias.Get4D(x,y,z,w), va);
    }

func (this *ImplicitBias) Get6D(x,y,z,w,u,v float64) float64 {
        va := this.source.Get6D(x,y,z,w, u, v);
        return Bias(this.bias.Get6D(x,y,z,w,u,v), va);
    }