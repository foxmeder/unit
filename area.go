package unit

// Area represents a SI unit of area (in square meters, m²)
type Area Unit

// Unit converts the Area to a Unit
func (a Area) Unit() Unit {
	return Unit(a)
}

// ...
const (
	// SI
	SquareYoctometer      = SquareMeter * 1e-48
	SquareZeptometer      = SquareMeter * 1e-42
	SquareAttometer       = SquareMeter * 1e-36
	SquareFemtometer      = SquareMeter * 1e-30
	SquarePicometer       = SquareMeter * 1e-24
	SquareNanometer       = SquareMeter * 1e-18
	SquareMicrometer      = SquareMeter * 1e-12
	SquareMillimeter      = SquareMeter * 1e-6 // 平方毫米(mm²)
	SquareCentimeter      = SquareMeter * 1e-4 // 平方厘米(cm²)
	SquareDecimeter       = SquareMeter * 1e-2 // 平方分米(dm²)
	SquareMeter      Area = 1e0                // 平方米(㎡)
	SquareDecameter       = SquareMeter * 1e2  // 公亩(are)
	SquareHectometer      = SquareMeter * 1e4  // 公顷(ha)
	SquareKilometer       = SquareMeter * 1e6  // 平方千米(km²)
	SquareMegameter       = SquareMeter * 1e12
	SquareGigameter       = SquareMeter * 1e18
	SquareTerameter       = SquareMeter * 1e24
	SquarePetameter       = SquareMeter * 1e30
	SquareExameter        = SquareMeter * 1e36
	SquareZettameter      = SquareMeter * 1e42
	SquareYottameter      = SquareMeter * 1e48

	// US
	SquareInch = SquareMeter * 0.00064516 // 平方英寸(sq.in)
	SquareFoot = SquareInch * 144         // 平方英尺(sq.ft)
	SquareYard = SquareFoot * 9           // 平方码(sq.yd)
	Acre       = SquareYard * 4840        // 英亩(acre)
	SquareMile = Acre * 640               // 平方英里(sq.mi)

	// imperial
	SquareRod = SquareFoot * 272.25
	Rood      = SquareYard * 1210

	// aliases
	Centiare    = SquareMeter
	Are         = SquareDecameter  // 公亩(are)
	Hectare     = SquareHectometer // 公顷(ha)
	SquarePerch = SquareRod

	// CN
	SquareCun = SquareChi * 1e-2 // 平方寸
	SquareChi = SquareMeter / 9  // 平方尺
	Fen       = SquareChi * 600  // 分
	Mu        = Fen * 10         // 亩
	Qing      = Mu * 100         // 顷
)

// SquareYoctometers returns the area in ym²
func (a Area) SquareYoctometers() float64 {
	return float64(a / SquareYoctometer)
}

// SquareZeptometers returns the area in zm²
func (a Area) SquareZeptometers() float64 {
	return float64(a / SquareZeptometer)
}

// SquareAttometers returns the area in am²
func (a Area) SquareAttometers() float64 {
	return float64(a / SquareAttometer)
}

// SquareFemtometers returns the area in fm²
func (a Area) SquareFemtometers() float64 {
	return float64(a / SquareFemtometer)
}

// SquarePicometers returns the area in pm²
func (a Area) SquarePicometers() float64 {
	return float64(a / SquarePicometer)
}

// SquareNanometers returns the area in nm²
func (a Area) SquareNanometers() float64 {
	return float64(a / SquareNanometer)
}

// SquareMicrometers returns the area in µm²
func (a Area) SquareMicrometers() float64 {
	return float64(a / SquareMicrometer)
}

// SquareMillimeters returns the area in mm²
func (a Area) SquareMillimeters() float64 {
	return float64(a / SquareMillimeter)
}

// SquareCentimeters returns the area in cm²
func (a Area) SquareCentimeters() float64 {
	return float64(a / SquareCentimeter)
}

// SquareDecimeters returns the area in dm²
func (a Area) SquareDecimeters() float64 {
	return float64(a / SquareDecimeter)
}

// SquareMeters returns the area in m²
func (a Area) SquareMeters() float64 {
	return float64(a / SquareMeter)
}

// Centiares returns the area in ca² (m²)
func (a Area) Centiares() float64 {
	return float64(a / Centiare)
}

// Ares returns the area in a²
func (a Area) Ares() float64 {
	return float64(a / Are)
}

// SquareDecameter returns the area in dam²
func (a Area) SquareDecameter() float64 {
	return float64(a / SquareDecameter)
}

// SquareHectometer returns the area in hm² (hectare)
func (a Area) SquareHectometer() float64 {
	return float64(a / SquareHectometer)
}

// Hectares returns the area in ha
func (a Area) Hectares() float64 {
	return float64(a / Hectare)
}

// SquareKilometers returns the area in km²
func (a Area) SquareKilometers() float64 {
	return float64(a / SquareKilometer)
}

// SquareMegameters returns the area in Mm²
func (a Area) SquareMegameters() float64 {
	return float64(a / SquareMegameter)
}

// SquareGigameters returns the area in Gm²
func (a Area) SquareGigameters() float64 {
	return float64(a / SquareGigameter)
}

// SquareTerameters returns the area in Tm²
func (a Area) SquareTerameters() float64 {
	return float64(a / SquareTerameter)
}

// SquarePetameters returns the area in Pm²
func (a Area) SquarePetameters() float64 {
	return float64(a / SquarePetameter)
}

// SquareExameters returns the area in Em²
func (a Area) SquareExameters() float64 {
	return float64(a / SquareExameter)
}

// SquareZettameters returns the area in Zm²
func (a Area) SquareZettameters() float64 {
	return float64(a / SquareZettameter)
}

// SquareYottameters returns the area in Ym²
func (a Area) SquareYottameters() float64 {
	return float64(a / SquareYottameter)
}

// SquareInches returns the area in in²
func (a Area) SquareInches() float64 {
	return float64(a / SquareInch)
}

// SquareFeet returns the area in ft²
func (a Area) SquareFeet() float64 {
	return float64(a / SquareFoot)
}

// SquareYards returns the area in yd
func (a Area) SquareYards() float64 {
	return float64(a / SquareYard)
}

// Acres returns the area in ac
func (a Area) Acres() float64 {
	return float64(a / Acre)
}

// SquareMiles returns the area in mi²
func (a Area) SquareMiles() float64 {
	return float64(a / SquareMile)
}

// SquareRods returns the area in square rods
func (a Area) SquareRods() float64 {
	return float64(a / SquareRod)
}

// Roods returns the area in roods
func (a Area) Roods() float64 {
	return float64(a / Rood)
}

// SquareCuns returns the area in square cuns
// 平方寸
func (a Area) SquareCuns() float64 {
	return float64(a / SquareCun)
}

// SquareChis returns the area in square chis
// 平方尺
func (a Area) SquareChis() float64 {
	return float64(a / SquareChi)
}

// Fens returns the area in fen
// 分
func (a Area) Fens() float64 {
	return float64(a / Fen)
}

// Mus returns the area in mu
// 亩
func (a Area) Mus() float64 {
	return float64(a / Mu)
}

// Qings returns the area in qing
// 顷
func (a Area) Qings() float64 {
	return float64(a / Qing)
}
