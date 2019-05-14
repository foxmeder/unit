package unit

import (
	"math"
	"testing"
)

type tItem struct {
	from, to string
	value    float64
	expect   float64
}

func TestConversion(t *testing.T) {
	var r float64
	for _, v := range []tItem{
		// mass
		tItem{"mass.kg", "mass.g", 1.0, 1000.0},         // 千克 克
		tItem{"mass.kg", "mass.jin", 1.0, 2.0},          // 千克 斤
		tItem{"mass.jin", "mass.g", 1.0, 500.0},         // 斤 克
		tItem{"mass.lb", "mass.ct", 1.0, 2267.96185},    // 磅 克拉
		tItem{"mass.lt", "mass.jin", 1.0, 2032.0938176}, // 长吨 斤
		// length
		tItem{"length.km", "length.m", 1.0, 1e3},                  // 千米 米
		tItem{"length.in", "length.ft", 1.0, 0.08333333333333334}, // 英寸 英尺
		tItem{"length.fur", "length.mi", 100.0, 12.5},             // 弗隆 英里
		tItem{"length.li", "length.m", 1.0, 500.0},                // 里 米
		tItem{"length.m", "length.cun", 1.0, 30.0},                // 米 寸
		// angle
		tItem{"angle.rad", "angle.degree", 1.0, 57.295779513082321},     // 弧度 度
		tItem{"angle.degree", "angle.mrad", 1.0, 17.453292519943295},    // 度 毫弧度
		tItem{"angle.min", "angle.rad", 1.0, 0.0002908882086657216},     // 分 弧度
		tItem{"angle.degree", "angle.rad", 180.0, math.Pi},              // 度 弧度
		tItem{"angle.right_angle", "angle.rad", 1.0, 1.570796326794897}, // 直角 弧度
		// speed
		tItem{"speed.m_s", "speed.km_h", 10.0, 35.99997120002304},    // 米/秒 千米/时
		tItem{"speed.m_s", "speed.mph", 10.0, 22.369362920544024},    // 米/秒 英里/时
		tItem{"speed.m_s", "speed.c", 10000000, 0.03335640951981521}, // 米/秒 光速
		tItem{"speed.mach", "speed.mph", 1.0, 761.2294201861131},     // 马赫 英里/时
		tItem{"speed.mach", "speed.km_h", 1.0, 1225.0790199367839},   // 马赫 千米/时
		// area
		tItem{"area.sq_m", "area.sq_ft", 1.0, 10.763910416709724}, // 平方米 平方英尺
		tItem{"area.acre", "area.sq_mi", 1.0, 0.0015625},          // 亩 平方英里
		tItem{"area.sq_m", "area.fen", 1.0, 0.015},                // 平方米 分
		tItem{"area.sq_m", "area.sq_cun", 1.0, 900.0},             // 平方米 平方寸
		tItem{"area.sq_mi", "area.sq_m", 1.0, 2589988.110336},     // 平方英里 平方米
		// volumn
		tItem{"volumn.l", "volumn.uk_gal", 1.0, 0.21996924829908776}, // 升 英式加仑
		tItem{"volumn.l", "volumn.us_gal", 1.0, 0.26417205235814845}, // 升 美式加仑
		tItem{"volumn.l", "volumn.us_oz", 1.0, 33.814022701843},      // 升 英式液体盎司
		tItem{"volumn.ac_in", "volumn.hl", 1.0, 12334.818375475199},  // 亩英尺 公石
		tItem{"volumn.ac_in", "volumn.cu_in", 1.0, 75271680.0},       // 亩英尺 立方英寸
		// temperature
		tItem{"temperature.c", "temperature.k", 100, 373.15},            // 摄氏度 开氏度
		tItem{"temperature.f", "temperature.k", 100, 310.9277777777778}, // 华氏度 开氏度
		tItem{"temperature.re", "temperature.k", 100, 398.15},           // 列氏度 开氏度
		tItem{"temperature.k", "temperature.r", 100, 180},               // 开氏度 兰氏度
		tItem{"temperature.k", "temperature.c", 100, -173.15},           // 亩英尺 立方英寸
		// pressure
		tItem{"pressure.kpa", "pressure.atm", 10, 0.09869232667160129},     // 帕 大气压
		tItem{"pressure.pa", "pressure.psi", 10000, 1.4503774389728312},    // 帕 磅力/平方英寸
		tItem{"pressure.atm", "pressure.mmhg", 1, 760.0},                   // 大气压 mmhg
		tItem{"pressure.inhg", "pressure.pa", 1, 3386.3881578947367},       // inhg 帕
		tItem{"pressure.mpa", "pressure.kgf_sq_cm", 1, 10.197162129779283}, // 兆帕 公斤力/平方厘米
		// power
		tItem{"power.w", "power.ps", 1000, 1.3596216173039042},     // 瓦 米制马力(ps)
		tItem{"power.w", "power.hp", 1000, 1.3410220899228118},     // 瓦 英制马力(hp)
		tItem{"power.w", "power.kcal_s", 1000, 0.2390057361376673}, // 瓦 千卡/秒(kcal/s)
		tItem{"power.w", "power.ftlb_s", 1000, 737.5621494575464},  // 瓦 英尺·磅/秒(ft·lb/s)
		tItem{"power.w", "power.btu_s", 1000, 0.9478171226670133},  // 瓦 英热单位/秒(Btu/s)
		// duration
		tItem{"duration.s", "duration.m", 1, 0.016666666666666666},                  // 秒 分钟
		tItem{"duration.s", "duration.h", 1, 0.0002777777777777778},                 // 秒 小时
		tItem{"duration.s", "duration.d", 1, 1.1574074074074073e-5},                 // 秒 天
		tItem{"duration.s", "duration.wk", 1, 1.6534391534391535e-6},                // 秒 星期
		tItem{"duration.s", "duration.yr", 1, 3.1709791983764586504312531709792e-8}, // 秒 年
		// datasize
		tItem{"datasize.bit", "datasize.byte", 1, 0.125},         // bit 字节
		tItem{"datasize.bit", "datasize.kb", 1, 0.0001220703125}, // bit KB
	} {
		r, _ = UnitConv(v.from, v.to, v.value)
		assertFloatEqual(t, r, v.expect)
	}
	// volumn
	// temperature
	// pressure
	// power
	// duration
	// datasize
}
