package anl

import ()

type CombinerType uint32

const (
	Add CombinerType = iota
	Mult
	Max
	Min
	Avg
)

type ImplicitCombiner struct {
	ImplicitModuleBase
	sources []ImplicitModule
	ctype    CombinerType
}

func (this *ImplicitCombiner) Spacing() float64 {
	return this.ImplicitModuleBase.spacing
}

func NewImplicitCombiner(t CombinerType) *ImplicitCombiner {
	v := &ImplicitCombiner{ctype: t}
	v.ClearAllSources()
	return v
}

func (this *ImplicitCombiner) SetType(t CombinerType) {
	this.ctype = t
}

func (this *ImplicitCombiner) ClearAllSources() {
	for c := 0; c < int(MaxSources); c++ {
		this.sources[c] = nil
	}
}

func (this *ImplicitCombiner) SetSource(which uint32, b ImplicitModule) {
	if which < 0 || which >= MaxSources {
		return
	}
	this.sources[which] = b
}

func (this *ImplicitCombiner) Get2D(x, y float64) float64 {
	switch this.ctype {
	case Add:
		return this.AddGet2D(x, y)
	case Mult:
		return this.MultGet2D(x, y)
	case Max:
		return this.MaxGet2D(x, y)
	case Min:
		return this.MinGet2D(x, y)
	case Avg:
		return this.AvgGet2D(x, y)
	default:
		return 0.0
	}
}

func (this *ImplicitCombiner) Get3D(x, y, z float64) float64 {

	switch this.ctype {
	case Add:
		return this.AddGet3D(x, y, z)
	case Mult:
		return this.MultGet3D(x, y, z)
	case Max:
		return this.MaxGet3D(x, y, z)
	case Min:
		return this.MinGet3D(x, y, z)
	case Avg:
		return this.AvgGet3D(x, y, z)
	default:
		return 0.0
	}
}

func (this *ImplicitCombiner) Get4D(x, y, z, w float64) float64 {
	switch this.ctype {
	case Add:
		return this.AddGet4D(x, y, z, w)
	case Mult:
		return this.MultGet4D(x, y, z, w)
	case Max:
		return this.MaxGet4D(x, y, z, w)
	case Min:
		return this.MinGet4D(x, y, z, w)
	case Avg:
		return this.AvgGet4D(x, y, z, w)
	default:
		return 0.0
	}
}

func (this *ImplicitCombiner) Get6D(x, y, z, w, u, v float64) float64 {
	switch this.ctype {
	case Add:
		return this.AddGet6D(x, y, z, w, u, v)
	case Mult:
		return this.MultGet6D(x, y, z, w, u, v)
	case Max:
		return this.MaxGet6D(x, y, z, w, u, v)
	case Min:
		return this.MinGet6D(x, y, z, w, u, v)
	case Avg:
		return this.AvgGet6D(x, y, z, w, u, v)
	default:
		return 0.0
	}
}

func (this *ImplicitCombiner) AddGet2D(x, y float64) float64 {
	value := 0.0
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value += this.sources[c].Get2D(x, y)
		}
	}
	return value
}
func (this *ImplicitCombiner) AddGet3D(x, y, z float64) float64 {
	value := 0.0
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value += this.sources[c].Get3D(x, y, z)
		}
	}
	return value
}
func (this *ImplicitCombiner) AddGet4D(x, y, z, w float64) float64 {
	value := 0.0
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value += this.sources[c].Get4D(x, y, z, w)
		}
	}
	return value
}
func (this *ImplicitCombiner) AddGet6D(x, y, z, w, u, v float64) float64 {
	value := 0.0
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value += this.sources[c].Get6D(x, y, z, w, u, v)
		}
	}
	return value
}

func (this *ImplicitCombiner) MultGet2D(x, y float64) float64 {
	value := 1.0
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value *= this.sources[c].Get2D(x, y)
		}
	}
	return value
}
func (this *ImplicitCombiner) MultGet3D(x, y, z float64) float64 {
	value := 1.0
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value *= this.sources[c].Get3D(x, y, z)
		}
	}
	return value
}
func (this *ImplicitCombiner) MultGet4D(x, y, z, w float64) float64 {
	value := 1.0
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value *= this.sources[c].Get4D(x, y, z, w)
		}
	}
	return value
}
func (this *ImplicitCombiner) MultGet6D(x, y, z, w, u, v float64) float64 {
	value := 1.0
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value *= this.sources[c].Get6D(x, y, z, w, u, v)
		}
	}
	return value
}

func (this *ImplicitCombiner) MinGet2D(x, y float64) float64 {
	c := 0
	for c < int(MaxSources) && this.sources[c] == nil {
		c++
	}
	if c == int(MaxSources) {
		return 0.0
	}
	mn := this.sources[c].Get2D(x, y)

	for d := c; d < int(MaxSources); d++ {
		if this.sources[d] != nil {
			v := this.sources[d].Get2D(x, y)
			if v < mn {
				mn = v
			}
		}
	}

	return mn
}

func (this *ImplicitCombiner) MinGet3D(x, y, z float64) float64 {
	c := 0
	for c < int(MaxSources) && this.sources[c] == nil {
		c++
	}
	if c == int(MaxSources) {
		return 0.0
	}
	mn := this.sources[c].Get3D(x, y, z)

	for d := c; d < int(MaxSources); d++ {
		if this.sources[d] != nil {
			v := this.sources[d].Get3D(x, y, z)
			if v < mn {
				mn = v
			}
		}
	}

	return mn
}

func (this *ImplicitCombiner) MinGet4D(x, y, z, w float64) float64 {
	c := 0
	for c < int(MaxSources) && this.sources[c] == nil {
		c++
	}
	if c == int(MaxSources) {
		return 0.0
	}
	mn := this.sources[c].Get4D(x, y, z, w)

	for d := c; d < int(MaxSources); d++ {
		if this.sources[d] != nil {
			v := this.sources[d].Get4D(x, y, z, w)
			if v < mn {
				mn = v
			}
		}
	}

	return mn
}

func (this *ImplicitCombiner) MinGet6D(x, y, z, w, u, v float64) float64 {
	c := 0
	for c < int(MaxSources) && this.sources[c] == nil {
		c++
	}
	if c == int(MaxSources) {
		return 0.0
	}
	mn := this.sources[c].Get6D(x, y, z, w, u, v)

	for d := c; d < int(MaxSources); d++ {
		if this.sources[d] != nil {
			v := this.sources[d].Get6D(x, y, z, w, u, v)
			if v < mn {
				mn = v
			}
		}
	}

	return mn
}

func (this *ImplicitCombiner) MaxGet2D(x, y float64) float64 {
	c := 0
	for c < int(MaxSources) && this.sources[c] == nil {
		c++
	}
	if c == int(MaxSources) {
		return 0.0
	}
	mn := this.sources[c].Get2D(x, y)

	for d := c; d < int(MaxSources); d++ {
		if this.sources[d] != nil {
			v := this.sources[d].Get2D(x, y)
			if v > mn {
				mn = v
			}
		}
	}

	return mn
}

func (this *ImplicitCombiner) MaxGet3D(x, y, z float64) float64 {
	c := 0
	for c < int(MaxSources) && this.sources[c] == nil {
		c++
	}
	if c == int(MaxSources) {
		return 0.0
	}
	mn := this.sources[c].Get3D(x, y, z)

	for d := c; d < int(MaxSources); d++ {
		if this.sources[d] != nil {
			v := this.sources[d].Get3D(x, y, z)
			if v > mn {
				mn = v
			}
		}
	}

	return mn
}

func (this *ImplicitCombiner) MaxGet4D(x, y, z, w float64) float64 {
	c := 0
	for c < int(MaxSources) && this.sources[c] == nil {
		c++
	}
	if c == int(MaxSources) {
		return 0.0
	}
	mn := this.sources[c].Get4D(x, y, z, w)

	for d := c; d < int(MaxSources); d++ {
		if this.sources[d] != nil {
			v := this.sources[d].Get4D(x, y, z, w)
			if v > mn {
				mn = v
			}
		}
	}

	return mn
}

func (this *ImplicitCombiner) MaxGet6D(x, y, z, w, u, v float64) float64 {
	c := 0
	for c < int(MaxSources) && this.sources[c] == nil {
		c++
	}
	if c == int(MaxSources) {
		return 0.0
	}
	mn := this.sources[c].Get6D(x, y, z, w, u, v)

	for d := c; d < int(MaxSources); d++ {
		if this.sources[d] != nil {
			v := this.sources[d].Get6D(x, y, z, w, u, v)
			if v > mn {
				mn = v
			}
		}
	}

	return mn
}

func (this *ImplicitCombiner) AvgGet2D(x, y float64) float64 {
	count := 0.
	value := 0.
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value += this.sources[c].Get2D(x, y)
			count += 1.0
		}
	}
	if count == 0.0 {
		return 0.0
	}
	return value / count
}

func (this *ImplicitCombiner) AvgGet3D(x, y, z float64) float64 {
	count := 0.
	value := 0.
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value += this.sources[c].Get3D(x, y, z)
			count += 1.0
		}
	}
	if count == 0.0 {
		return 0.0
	}
	return value / count
}

func (this *ImplicitCombiner) AvgGet4D(x, y, z, w float64) float64 {
	count := 0.
	value := 0.
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value += this.sources[c].Get4D(x, y, z, w)
			count += 1.0
		}
	}
	if count == 0.0 {
		return 0.0
	}
	return value / count
}

func (this *ImplicitCombiner) AvgGet6D(x, y, z, w, u, v float64) float64 {
	count := 0.
	value := 0.
	for c := 0; c < int(MaxSources); c++ {
		if this.sources[c] != nil {
			value += this.sources[c].Get6D(x, y, z, w, u, v)
			count += 1.0
		}
	}
	if count == 0.0 {
		return 0.0
	}
	return value / count
}
