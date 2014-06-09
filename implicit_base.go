package anl

const MaxSources uint32 = 20 

type ImplicitModule interface {
	SetSeed(seed uint32)

	Get2D(x, y float64) float64
	Get3D(x, y, z float64) float64
	Get4D(x, y, z, w float64) float64
	Get6D(x, y, z, w, u, v float64) float64
	
	Spacing() float64
}

type ImplicitModuleBase struct {
	spacing float64
}

func (self *ImplicitModuleBase) Spacing() float64 {
	return self.spacing 
}

func GetDx2(module ImplicitModule, x, y float64) float64 {
	return (module.Get2D(x-module.Spacing(), y) - module.Get2D(x+module.Spacing(), y)) / module.Spacing()
}

func GetDy2(module ImplicitModule, x, y float64) float64 {
	return (module.Get2D(x, y-module.Spacing()) - module.Get2D(x, y+module.Spacing())) / module.Spacing()
}

func GetDx3(module ImplicitModule, x, y, z float64) float64 {
	return (module.Get3D(x-module.Spacing(), y, z) - module.Get3D(x+module.Spacing(), y, z)) / module.Spacing()
}

func GetDy3(module ImplicitModule, x, y, z float64) float64 {
	return (module.Get3D(x, y-module.Spacing(), z) - module.Get3D(x, y+module.Spacing(), z)) / module.Spacing()
}

func GetDz3(module ImplicitModule, x, y, z float64) float64 {
	return (module.Get3D(x, y, z-module.Spacing()) - module.Get3D(x, y, z+module.Spacing())) / module.Spacing()
}

func GetDx4(module ImplicitModule, x, y, z, w float64) float64 {
	return (module.Get4D(x-module.Spacing(), y, z, w) - module.Get4D(x+module.Spacing(), y, z, w)) / module.Spacing()
}

func GetDy4(module ImplicitModule, x, y, z, w float64) float64 {
	return (module.Get4D(x, y-module.Spacing(), z, w) - module.Get4D(x, y+module.Spacing(), z, w)) / module.Spacing()
}

func GetDz4(module ImplicitModule, x, y, z, w float64) float64 {
	return (module.Get4D(x, y, z-module.Spacing(), w) - module.Get4D(x, y, z+module.Spacing(), w)) / module.Spacing()
}

func GetDw4(module ImplicitModule, x, y, z, w float64) float64 {
	return (module.Get4D(x, y, z, w-module.Spacing()) - module.Get4D(x, y, z, w+module.Spacing())) / module.Spacing()
}

func GetDx6(module ImplicitModule, x, y, z, w, u, v float64) float64 {
	return (module.Get6D(x-module.Spacing(), y, z, w, u, v) - module.Get6D(x+module.Spacing(), y, z, w, u, v)) / module.Spacing()
}

func GetDy6(module ImplicitModule, x, y, z, w, u, v float64) float64 {
	return (module.Get6D(x, y-module.Spacing(), z, w, u, v) - module.Get6D(x, y+module.Spacing(), z, w, u, v)) / module.Spacing()
}

func GetDz6(module ImplicitModule, x, y, z, w, u, v float64) float64 {
	return (module.Get6D(x, y, z-module.Spacing(), w, u, v) - module.Get6D(x, y, z+module.Spacing(), w, u, v)) / module.Spacing()
}

func GetDw6(module ImplicitModule, x, y, z, w, u, v float64) float64 {
	return (module.Get6D(x, y, z, w-module.Spacing(), u, v) - module.Get6D(x, y, z, w+module.Spacing(), u, v)) / module.Spacing()
}

func GetDu6(module ImplicitModule, x, y, z, w, u, v float64) float64 {
	return (module.Get6D(x, y, z, w, u-module.Spacing(), v) - module.Get6D(x, y, z, w, u+module.Spacing(), v)) / module.Spacing()
}

func GetDv6(module ImplicitModule, x, y, z, w, u, v float64) float64 {
	return (module.Get6D(x, y, z, w, u, v-module.Spacing()) - module.Get6D(x, y, z, w, u, v+module.Spacing())) / module.Spacing()
}

type ScalarParameter struct {
	val    float64
	source ImplicitModule
}

func NewScalarParameter(v float64) ScalarParameter {
	return ScalarParameter{val: v, source: nil}
}

func (self *ScalarParameter) SetValue(v float64) {
	self.source = nil
	self.val = v
}

func (self *ScalarParameter) SetModule(m ImplicitModule) {
	self.source = m
}

func (self *ScalarParameter) Get2D(x, y float64) float64 {
	if self.source != nil {
		return self.source.Get2D(x, y)
	} else {
		return self.val
	}
}

func (self *ScalarParameter) Get3D(x, y, z float64) float64 {
	if self.source != nil {
		return self.source.Get3D(x, y, z)
	} else {
		return self.val
	}
}

func (self *ScalarParameter) Get4D(x, y, z, w float64) float64 {
	if self.source != nil {
		return self.source.Get4D(x, y, z, w)
	} else {
		return self.val
	}
}

func (self *ScalarParameter) Get6D(x, y, z, w, u, v float64) float64 {
	if self.source != nil {
		return self.source.Get6D(x, y, z, w, u, v)
	} else {
		return self.val
	}
}
