package unit

import (
	"errors"
	"strings"
)

type unitMap map[string]Uniter

var (
	UnsupportUnit error = errors.New("unit is not supported")
	UnsupportDime error = errors.New("Dime is not supported")
	NotSameDime   error = errors.New("not the same dimesion")
	NotFloat64    error = errors.New("unit convert to float64 failed")
)

func (m unitMap) getElem(name string) (float64, error) {
	unit, ok := m[name]
	if !ok {
		return 0, UnsupportUnit
	}
	return float64(unit.Unit()), nil
}

// 重量
var massHash unitMap = unitMap{
	"kg":    Kilogram,           // 千克 公斤
	"g":     Gram,               // 克(g)
	"mg":    Milligram,          // 毫克(mg)
	"μg":    Microgram,          // 微克(μg)
	"t":     Tonne,              // 吨(t)
	"q":     Quintal,            // 公担(q)
	"lb":    AvoirdupoisPound,   // 磅(lb)
	"oz":    AvoirdupoisOunce,   // 盎司(oz)
	"ct":    Carat,              // 克拉(ct)
	"gr":    TroyGrain,          // 格令(gr)
	"lt":    LongTon,            // 长吨(lt)
	"st":    ShortTon,           // 短吨(st)
	"cwt":   LongHundredweight,  // 英担
	"scwt":  ShortHundredweight, // 美担
	"ust":   UsStone,            // 英石(st)
	"dr":    AvoirdupoisDram,    // 打兰(dr)
	"dan":   Dan,                // 担
	"jin":   Jin,                // 斤
	"liang": Liang,              // 两
	"qian":  Qian,               // 钱
}

// 长度
var lengthHash unitMap = unitMap{
	"km":    Kilometer,        // 千米(km)
	"m":     Meter,            // 米(m)
	"dm":    Decimeter,        // 分米(dm)
	"cm":    Centimeter,       // 厘米(cm)
	"mm":    Millimeter,       // 毫米(mm)
	"um":    Micrometer,       // 微米(um)
	"nm":    Nanometer,        // 纳米(nm)
	"pm":    Picometer,        // 皮米(pm)
	"ly":    LightYear,        // 光年(ly)
	"au":    AstronomicalUnit, // 天文单位(AU)
	"in":    Inch,             // 英寸(in)
	"ft":    Foot,             // 英尺(ft)
	"yd":    Yard,             // 码(yd)
	"mi":    Mile,             // 英里(mi)
	"nmi":   NauticalMile,     // 海里(nmi)
	"fm":    Fathom,           // 英寻(fm)
	"fur":   Furlong,          // 弗隆(fur)
	"li":    Li,               // 里
	"zhang": Zhang,            // 丈
	"chi":   Chi,              // 尺
	"cun":   Cun,              // 寸
	"fen":   CnFen,            // 分
	"li2":   Li2,              // 厘
	"hao":   Hao,              // 毫
}

// 角度
var angleHash unitMap = unitMap{
	"circle":      Circle,      // 圆周
	"right_angle": RightAngle,  // 直角
	"degree":      Degree,      // 度(°)
	"min":         Arcminute,   // 分( ′)
	"sec":         Arcsecond,   // 秒(')
	"rad":         Radian,      // 弧度(rad)
	"mrad":        Milliradian, // 毫弧度(mrad)
}

// 速度
var speedHash unitMap = unitMap{
	"m_s":  MetersPerSecond,     // 米/秒(m/s)
	"km_s": KilometersPerSecond, // 千米/秒(km/s)
	"km_h": KilometersPerHour,   // 千米/时(km/h)
	"c":    SpeedOfLight,        // 光速(c)
	"mach": Mach,                // 马赫(Mach)
	"mph":  MilesPerHour,        // 英里/时(mile/h)
	"in_s": InchesPerSecond,     // 英寸/秒(in/s)
}

// 面积
var areaHash unitMap = unitMap{
	"sq_km":  SquareKilometer,  // 平方千米(km²)
	"ha":     SquareHectometer, // 公顷(ha)
	"are":    SquareDecameter,  // 公亩(are)
	"sq_m":   SquareMeter,      // 平方米(㎡)
	"sq_dm":  SquareDecimeter,  // 平方分米(dm²)
	"sq_cm":  SquareCentimeter, // 平方厘米(cm²)
	"sq_mm":  SquareMillimeter, // 平方毫米(mm²)
	"acre":   Acre,             // 英亩(acre)
	"sq_mi":  SquareMile,       // 平方英里(sq.mi)
	"sq_yd":  SquareYard,       // 平方码(sq.yd)
	"sq_ft":  SquareFoot,       // 平方英尺(sq.ft)
	"sq_in":  SquareInch,       // 平方英寸(sq.in)
	"qing":   Qing,             // 顷
	"mu":     Mu,               // 亩
	"fen":    Fen,              // 分
	"sq_chi": SquareChi,        // 平方尺
	"sq_cun": SquareCun,        // 平方寸
}

// 体积
var volumnHash unitMap = unitMap{
	"cu_m":   CubicMeter,         // 立方米(m³)
	"cu_dm":  CubicDecimeter,     // 立方分米(dm³)
	"cu_cm":  CubicCentimeter,    // 立方厘米(cm³)
	"cu_mm":  CubicMillimeter,    // 立方毫米(mm³)
	"l":      Liter,              // 升(l)
	"dl":     Deciliter,          // 分升(dl)
	"ml":     Milliliter,         // 毫升(ml)
	"cl":     Centiliter,         // 厘升(cl)
	"hl":     Hectoliter,         // 公石(hl)
	"ul":     Microliter,         // 微升(ul)
	"cu_ft":  CubicFoot,          // 立方英尺(cu ft)
	"cu_in":  CubicInch,          // 立方英寸(cu in)
	"cu_yd":  CubicYard,          // 立方码(cu yd)
	"ac_in":  AcreInch,           // 亩英尺
	"uk_gal": ImperialGallon,     // 英制加仑(uk gal)
	"us_gal": USLiquidGallon,     // 美制加仑(us gal)
	"uk_oz":  ImperialFluidOunce, // 英制液体盎司(oz)
	"us_oz":  USFluidOunce,       // 美制液体盎司(oz)
}

// 压力
var pressureHash unitMap = unitMap{
	"pa":        Pascal,                           // 帕斯卡(Pa)
	"mpa":       Megapascal,                       // 兆帕(MPa)
	"kpa":       Kilopascal,                       // 千帕(kpa)
	"hpa":       Hectopascal,                      // 百帕(hpa)
	"atm":       Atmosphere,                       // 标准大气压(atm)
	"mmhg":      MilliHg,                          // 毫米汞柱(mmHg)
	"inhg":      InchHg,                           // 英寸汞柱(in Hg)
	"bar":       Bar,                              // 巴(bar)
	"mbar":      Millibar,                         // 毫巴(mbar)
	"psf":       PoundsPerSquareFoot,              // 磅力/平方英尺(psf)
	"psi":       PoundsPerSquareInch,              // 磅力/平方英寸(psi)
	"mmh2o":     MilliH2o,                         // 毫米水柱
	"kgf_cu_m":  KilogramForcePerSquareMeter,      // 公斤力/平方米(kgf/㎡)
	"kgf_cu_cm": KilogramForcePerSquareCentimeter, // 公斤力/平方厘米(kgf/cm²)
}

// 功率
var powerHash unitMap = unitMap{
	"w":      Watt,                   // 瓦(W)
	"kw":     Kilowatt,               // 千瓦(kW)
	"hp":     BritishHorsepower,      // 英制马力(hp)
	"ps":     MetricHorsepower,       // 米制马力(ps)
	"kgm_s":  KilogramMeterPerSecond, // 公斤·米/秒(kg·m/s)
	"kcal_s": KilocaloriePerSecond,   // 千卡/秒(kcal/s)
	"btu_s":  BTUPerSecond,           // 英热单位/秒(Btu/s)
	"ftlb_s": FootPoundPerSecond,     // 英尺·磅/秒(ft·lb/s)
	"j_s":    JoulePerSecond,         // 焦耳/秒(J/s)
	"nm_s":   NewtonMeterPerSecond,   // 牛顿·米/秒(N·m/s)
}

// 时间
var durationHash unitMap = unitMap{
	"yr": Year,        // 年
	"wk": Week,        // 周
	"d":  Day,         // 天
	"h":  Hour,        // 时
	"m":  Minute,      // 分
	"s":  Second,      // 秒
	"ms": Millisecond, //毫秒
	"μs": Microsecond, // 微秒(μs)
	"ns": Nanosecond,  // 纳秒(ns)
}

// 数据存储
var datasizeHash unitMap = unitMap{
	"bit":  Bit,      // 比特(bit)
	"byte": Byte,     // 字节(byte)
	"kb":   Kibibyte, // 千字节(kb)
	"mb":   Mebibyte,
	"gb":   Gibibyte,
	"tb":   Tebibyte, // 太字节(tb)
	"pb":   Pebibyte, // 拍字节(pb)
	"eb":   Exbibyte, // 艾字节(eb)
}

// UnitConv 把单位从from转换成to
func UnitConv(from, to string, value float64) (float64, error) {
	f := strings.Split(from, ".")
	if len(f) != 2 {
		return 0, UnsupportUnit
	}
	t := strings.Split(to, ".")
	if len(t) != 2 {
		return 0, UnsupportUnit
	}
	if f[0] != t[0] {
		return 0, NotSameDime
	}
	var hash unitMap
	switch f[0] {
	case "mass":
		hash = massHash
	case "length":
		hash = lengthHash
	case "angle":
		hash = angleHash
	case "speed":
		hash = speedHash
	case "area":
		hash = areaHash
	case "volumn":
		hash = volumnHash
	case "pressure":
		hash = pressureHash
	case "power":
		hash = powerHash
	case "duration":
		hash = durationHash
	case "datastorage":
		hash = datasizeHash
	case "temperature":
		return temperatureConv(f[1], t[1], value)
	default:
		return 0, UnsupportDime
	}
	return divConv(hash, f[1], t[1], value)
}

// 除法转换
func divConv(mapHash unitMap, from, to string, value float64) (float64, error) {
	left, err := mapHash.getElem(from)
	if err != nil {
		return 0, err
	}
	right, err := mapHash.getElem(to)
	if err != nil {
		return 0, err
	}
	return value * (left / right), nil
}

// 温度转换
func temperatureConv(from, to string, value float64) (float64, error) {
	var (
		funit  Temperature
		tValue float64
	)
	switch from {
	case "c":
		funit = FromCelsius(value)
	case "f":
		funit = FromFahrenheit(value)
	case "k":
		funit = FromKelvin(value)
	case "r":
		funit = FromRankine(value)
	case "re":
		funit = FromReaumur(value)
	default:
		return 0, UnsupportUnit
	}
	switch to {
	case "c":
		tValue = funit.Celsius()
	case "f":
		tValue = funit.Fahrenheit()
	case "k":
		tValue = funit.Kelvin()
	case "r":
		tValue = funit.Rankine()
	case "re":
		tValue = funit.Reaumur()
	default:
		return 0, UnsupportUnit
	}
	return tValue, nil
}
