package unit

// Speed represents a unit of speed (in meters per second, m/s)
type Speed Unit

// Unit converts the Speed to a Unit
func (s Speed) Unit() Unit {
	return Unit(s)
}

// ...
const (
	MetersPerSecond     Speed = 1e0                         // 米/秒(m/s)
	InchesPerSecond           = MetersPerSecond * 0.0254    // 英寸/秒(in/s)
	KilometersPerHour         = MetersPerSecond * 0.277778  // 千米/时(km/h)
	FeetPerSecond             = MetersPerSecond * 0.3048    // 英尺/秒
	MilesPerHour              = MetersPerSecond * 0.44704   // 英里/时(mile/h)
	Knot                      = MetersPerSecond * 0.514444  // 节
	Mach                      = MetersPerSecond * 340.3     // 马赫(mach)
	SpeedOfLight              = MetersPerSecond * 299792458 // 光速(c)
	KilometersPerSecond       = MetersPerSecond * 1e3       // 千米/秒(km/s)
)

// MetersPerSecond returns the speed in m/s
func (s Speed) MetersPerSecond() float64 {
	return float64(s)
}

// KilometersPerHour returns the speed in km/h
func (s Speed) KilometersPerHour() float64 {
	return float64(s / KilometersPerHour)
}

// FeetPerSecond returns the speed in ft/s
func (s Speed) FeetPerSecond() float64 {
	return float64(s / FeetPerSecond)
}

// MilesPerHour returns the speed in mph
func (s Speed) MilesPerHour() float64 {
	return float64(s / MilesPerHour)
}

// Knots returns the speed in knots
func (s Speed) Knots() float64 {
	return float64(s / Knot)
}

// SpeedOfLight returns the speed in c
func (s Speed) SpeedOfLight() float64 {
	return float64(s / SpeedOfLight)
}
