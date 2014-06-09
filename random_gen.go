package anl

import (
	"time"
)

type BasePRNG struct{}

type PRNG interface {
	Get() uint32
	SetSeed(seed uint32)
}

func SetSeedTime(prng PRNG) {
	prng.SetSeed(uint32(time.Now().Unix()))
}

func GetTarget(prng PRNG, t uint32) uint32 {
	v := Get01(prng)
	return uint32(v * float64(t))
}

func GetRange(prng PRNG, low, high uint32) uint32 {
	if high < low {
		high, low = low, high
	}
	rg := float64((high - low) + 1)
	val := float64(low) + Get01(prng)*rg
	return uint32(val)
}

func Get01(prng PRNG) float64 {
	return (float64(prng.Get()) / float64(4294967295))
}

type LCG struct {
	BasePRNG
	m_state uint32
}

func NewLCG() *LCG {
	v := new(LCG)
	v.SetSeed(10000)

	return v
}

func (self *LCG) SetSeed(seed uint32) {
	self.m_state = seed
}

func (self *LCG) Get() uint32 {
	self.m_state = 69069*self.m_state + 362437
	return self.m_state
}

// Setup a static, global LCG for seeding other generators.
var lcg *LCG = NewLCG()

// The following generators are based on generators created by George Marsaglia
// They use the static lcg created above for seeding, to initialize various
// state and tables. Seeding them is a bit more involved than an LCG.
type Xorshift struct {
	BasePRNG
	m_x, m_y, m_z, m_w, m_v uint32
}

func NewXorshift() *Xorshift {
	v := new(Xorshift)
	v.SetSeed(10000)
	return v
}

func (self *Xorshift) SetSeed(s uint32) {
	lcg.SetSeed(s)
	self.m_x = lcg.Get()
	self.m_y = lcg.Get()
	self.m_z = lcg.Get()
	self.m_w = lcg.Get()
	self.m_v = lcg.Get()
}

func (self *Xorshift) Get() uint32 {
	t := (self.m_x ^ (self.m_x >> 7))
	self.m_x = self.m_y
	self.m_y = self.m_z
	self.m_z = self.m_w
	self.m_w = self.m_v
	self.m_v = (self.m_v ^ (self.m_v << 6)) ^ (t ^ (t << 13))
	return (self.m_y + self.m_y + 1) * self.m_v
}

type MWC256 struct {
	BasePRNG
	m_Q [256]uint32
	c   uint32
}

var iMWC256 uint32 = 0

func NewMWC256() *MWC256 {
	v := new(MWC256)
	v.SetSeed(10000)
	return v
}

func (self *MWC256) SetSeed(s uint32) {
	lcg.SetSeed(s)
	for i := 0; i < 256; i++ {
		self.m_Q[i] = lcg.Get()
	}
	self.c = GetTarget(lcg, 809430660)
}

func (self *MWC256) Get() uint32 {
	var t, a uint32
	t, a = 809430660, 809430660

	t = a*self.m_Q[iMWC256] + self.c
	iMWC256++
	self.c = (t >> 32)
	self.m_Q[iMWC256] = t
	return self.m_Q[iMWC256]
}

var iCMWC4096 uint32 = 2095

type CMWC4096 struct {
	BasePRNG
	m_Q [256]uint32
	c   uint32
}

func NewCMWC4096() *CMWC4096 {
	v := new(CMWC4096)
	v.SetSeed(10000)
	return v
}

func (self *CMWC4096) SetSeed(s uint32) {
	lcg.SetSeed(s) // Seed the global random source

	// Seed the table
	for i := 0; i < 4096; i++ {
		self.m_Q[i] = lcg.Get()
	}

	self.c = GetTarget(lcg, 18781)
}

func (self *CMWC4096) Get() uint32 {
	var t, a, b uint32
	t, a, b = 18782, 18782, 294967295
	r := uint32(b - 1)

	iCMWC4096 = (iCMWC4096 + 1) & 4095
	t = a*self.m_Q[iCMWC4096] + self.c
	self.c = (t >> 32)
	t = (t & b) + self.c
	if t > r {
		self.c++
		t = t - b
	}
	self.m_Q[iCMWC4096] = uint32(r - t)
	return self.m_Q[iCMWC4096]
}

type KISS struct {
	BasePRNG
	z, w, jsr, jcong uint32
}

func NewKISS() *KISS {
	v := new(KISS)
	v.SetSeed(10000)
	return v
}

func (self *KISS) SetSeed(s uint32) {
	lcg.SetSeed(s)
	self.z = lcg.Get()
	self.w = lcg.Get()
	self.jsr = lcg.Get()
	self.jcong = lcg.Get()
}

func (self *KISS) Get() uint32 {
	self.z = 36969*(self.z&65535) + (self.z >> 16)
	self.w = 18000*(self.w&65535) + (self.w >> 16)
	mwc := (self.z << 16) + self.w

	self.jcong = 69069*self.jcong + 1234567

	self.jsr ^= (self.jsr << 17)
	self.jsr ^= (self.jsr >> 13)
	self.jsr ^= (self.jsr << 5)

	return (mwc ^ self.jcong) + self.jsr
}
