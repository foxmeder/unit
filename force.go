package unit

// Force represents a SI unit of force (in newtons, N)
type Force Unit

// Unit converts the Force to a Unit
func (v Force) Unit() Unit {
	return Unit(v)
}

// ...
const (
	// SI
	Newton Force = 1e0 // 牛 N

	// non-SI
	Dyne          = Newton * 1e-5     // 达因 dyn
	KilogramForce = Newton * 9.80665  // 千克力 kgf
	PoundForce    = Newton * 4.448222 // 磅力 lbf
	Poundal       = Newton * 0.138255

	// aliases
	Kilopond = KilogramForce

	// misc
	KiloNewton     = Newton * 1e3         // 千牛 kN
	GramForce      = KilogramForce * 1e-3 // 克力 gf
	TonneForce     = KilogramForce * 1e3  // 公吨力 tf
	KiloPoundForce = PoundForce * 1e3     // 千磅力 kip
)

// Newtons returns the force in N
func (f Force) Newtons() float64 {
	return float64(f)
}

// Dynes returns the force in dyn
func (f Force) Dynes() float64 {
	return float64(f / Dyne)
}

// KilogramForce returns the force in kp
func (f Force) KilogramForce() float64 {
	return float64(f / KilogramForce)
}

// PoundForce returns the force in lbf
func (f Force) PoundForce() float64 {
	return float64(f / PoundForce)
}

// Poundals returns the force in pdl
func (f Force) Poundals() float64 {
	return float64(f / Poundal)
}
