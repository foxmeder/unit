package unit

// Density represents a unit of density
type Density Unit

// Unit converts the Density to a Unit
func (d Density) Unit() Unit {
	return Unit(d)
}

const (
	GramPerCubicMeter                  = KilogramPerCubicMeter * 1e-3 // 克/立方米(g/m³)
	GramPerCubicDecimeter              = KilogramPerCubicMeter        // 克/立方分米(g/dm³)
	KilogramPerCubicMeter      Density = 1                            // 千克/立方米(kg/m³)
	GramPerCubicCentimeter             = KilogramPerCubicMeter * 1e3  // 克/立方厘米(g/cm³)
	KilogramPerCubicDecimeter          = KilogramPerCubicMeter * 1e3  // 千克/立方分米(kg/dm³)
	KilogramPerCubicCentimeter         = KilogramPerCubicMeter * 1e6  // 千克/立方厘米(kg/cm³)
)
