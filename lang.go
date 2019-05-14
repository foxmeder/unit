package unit

// 语言map
type langMap map[string]langList

// 参数与语言的map
type langList map[string]string

// 英文维度和语言列表的map
type langHash map[string]langMap

// 重量(mass)
var massLang = langMap{
	"kg":    langList{"zh-cn": "千克(kg)"},
	"g":     langList{"zh-cn": "克(g)"},
	"mg":    langList{"zh-cn": "毫克(mg)"},
	"μg":    langList{"zh-cn": "微克(μg)"},
	"t":     langList{"zh-cn": "吨(t)"},
	"q":     langList{"zh-cn": "公担(q)"},
	"lb":    langList{"zh-cn": "磅(lb)"},
	"oz":    langList{"zh-cn": "盎司(oz)"},
	"ct":    langList{"zh-cn": "克拉(ct)"},
	"gr":    langList{"zh-cn": "格令(gr)"},
	"lt":    langList{"zh-cn": "长吨(lt)"},
	"st":    langList{"zh-cn": "短吨(st)"},
	"cwt":   langList{"zh-cn": "英担"},
	"scwt":  langList{"zh-cn": "美担"},
	"ust":   langList{"zh-cn": "英石(st)"},
	"dr":    langList{"zh-cn": "打兰(dr)"},
	"dan":   langList{"zh-cn": "担"},
	"jin":   langList{"zh-cn": "斤"},
	"liang": langList{"zh-cn": "两"},
	"qian":  langList{"zh-cn": "钱"},
}

// 长度(length)
var lengthLang = langMap{
	"km":    langList{"zh-cn": "千米(km)"},
	"m":     langList{"zh-cn": "米(m)"},
	"dm":    langList{"zh-cn": "分米(dm)"},
	"cm":    langList{"zh-cn": "厘米(cm)"},
	"mm":    langList{"zh-cn": "毫米(mm)"},
	"um":    langList{"zh-cn": "微米(um)"},
	"nm":    langList{"zh-cn": "纳米(nm)"},
	"pm":    langList{"zh-cn": "皮米(pm)"},
	"ly":    langList{"zh-cn": "光年(ly)"},
	"au":    langList{"zh-cn": "天文单位(AU)"},
	"in":    langList{"zh-cn": "英寸(in)"},
	"ft":    langList{"zh-cn": "英尺(ft)"},
	"yd":    langList{"zh-cn": "码(yd)"},
	"mi":    langList{"zh-cn": "英里(mi)"},
	"nmi":   langList{"zh-cn": "海里(nmi)"},
	"fm":    langList{"zh-cn": "英寻(fm)"},
	"fur":   langList{"zh-cn": "弗隆(fur)"},
	"li":    langList{"zh-cn": "里"},
	"zhang": langList{"zh-cn": "丈"},
	"chi":   langList{"zh-cn": "尺"},
	"cun":   langList{"zh-cn": "寸"},
	"fen":   langList{"zh-cn": "分"},
	"li2":   langList{"zh-cn": "厘"},
	"hao":   langList{"zh-cn": "毫"},
}

// 角度(angle)
var angleLang = langMap{
	"circle":      langList{"zh-cn": "圆周"},
	"right_angle": langList{"zh-cn": "直角"},
	"degree":      langList{"zh-cn": "度(°)"},
	"min":         langList{"zh-cn": "分( ′)"},
	"sec":         langList{"zh-cn": "秒(')"},
	"rad":         langList{"zh-cn": "弧度(rad)"},
	"mrad":        langList{"zh-cn": "毫弧度(mrad)"},
}

// 速度(speed)
var speedLang = langMap{
	"m_s":  langList{"zh-cn": "米/秒(m/s)"},
	"km_s": langList{"zh-cn": "千米/秒(km/s)"},
	"km_h": langList{"zh-cn": "千米/时(km/h)"},
	"c":    langList{"zh-cn": "光速(c)"},
	"mach": langList{"zh-cn": "马赫(Mach)"},
	"mph":  langList{"zh-cn": "英里/时(mile/h)"},
	"in_s": langList{"zh-cn": "英寸/秒(in/s)"},
}

// 面积(area)
var areaLang = langMap{
	"sq_km":  langList{"zh-cn": "平方千米(km²)"},
	"ha":     langList{"zh-cn": "公顷(ha)"},
	"are":    langList{"zh-cn": "公亩(are)"},
	"sq_m":   langList{"zh-cn": "平方米(㎡)"},
	"sq_dm":  langList{"zh-cn": "平方分米(dm²)"},
	"sq_cm":  langList{"zh-cn": "平方厘米(cm²)"},
	"sq_mm":  langList{"zh-cn": "平方毫米(mm²)"},
	"acre":   langList{"zh-cn": "英亩(acre)"},
	"sq_mi":  langList{"zh-cn": "平方英里(sq.mi)"},
	"sq_yd":  langList{"zh-cn": "平方码(sq.yd)"},
	"sq_ft":  langList{"zh-cn": "平方英尺(sq.ft)"},
	"sq_in":  langList{"zh-cn": "平方英寸(sq.in)"},
	"qing":   langList{"zh-cn": "顷"},
	"mu":     langList{"zh-cn": "亩"},
	"fen":    langList{"zh-cn": "分"},
	"sq_chi": langList{"zh-cn": "平方尺"},
	"sq_cun": langList{"zh-cn": "平方寸"},
}

// 体积(volumn)
var volumnLang = langMap{
	"cu_m":   langList{"zh-cn": "立方米(m³)"},
	"cu_dm":  langList{"zh-cn": "立方分米(dm³)"},
	"cu_cm":  langList{"zh-cn": "立方厘米(cm³)"},
	"cu_mm":  langList{"zh-cn": "立方毫米(mm³)"},
	"l":      langList{"zh-cn": "升(l)"},
	"dl":     langList{"zh-cn": "分升(dl)"},
	"ml":     langList{"zh-cn": "毫升(ml)"},
	"cl":     langList{"zh-cn": "厘升(cl)"},
	"hl":     langList{"zh-cn": "公石(hl)"},
	"ul":     langList{"zh-cn": "微升(ul)"},
	"cu_ft":  langList{"zh-cn": "立方英尺(cu ft)"},
	"cu_in":  langList{"zh-cn": "立方英寸(cu in)"},
	"cu_yd":  langList{"zh-cn": "立方码(cu yd)"},
	"ac_in":  langList{"zh-cn": "亩英尺"},
	"uk_gal": langList{"zh-cn": "英制加仑(uk gal)"},
	"us_gal": langList{"zh-cn": "美制加仑(us gal)"},
	"uk_oz":  langList{"zh-cn": "英制液体盎司(oz)"},
	"us_oz":  langList{"zh-cn": "美制液体盎司(oz)"},
}

// 温度(temperature)
var temperatureLang = langMap{
	"c":  langList{"zh-cn": "摄氏度(℃)"},
	"f":  langList{"zh-cn": "华氏度(℉)"},
	"k":  langList{"zh-cn": "开氏度(K)"},
	"r":  langList{"zh-cn": "兰氏度(°R)"},
	"re": langList{"zh-cn": "列氏度(°Re)"},
}

// 压力(pressure)
var pressureLang = langMap{
	"pa":        langList{"zh-cn": "帕斯卡(Pa)"},
	"mpa":       langList{"zh-cn": "兆帕(MPa)"},
	"kpa":       langList{"zh-cn": "千帕(kpa)"},
	"hpa":       langList{"zh-cn": "百帕(hpa)"},
	"atm":       langList{"zh-cn": "标准大气压(atm)"},
	"mmhg":      langList{"zh-cn": "毫米汞柱(mmHg)"},
	"inhg":      langList{"zh-cn": "英寸汞柱(in Hg)"},
	"bar":       langList{"zh-cn": "巴(bar)"},
	"mbar":      langList{"zh-cn": "毫巴(mbar)"},
	"psf":       langList{"zh-cn": "磅力/平方英尺(psf)"},
	"psi":       langList{"zh-cn": "磅力/平方英寸(psi)"},
	"mmh2o":     langList{"zh-cn": "毫米水柱"},
	"kgf_sq_cm": langList{"zh-cn": "公斤力/平方厘米(kgf/cm²)"},
	"kgf_sq_m":  langList{"zh-cn": "公斤力/平方米(kgf/㎡)"},
}

// 功率(power)
var powerLang = langMap{
	"w":      langList{"zh-cn": "瓦(W)"},
	"kw":     langList{"zh-cn": "千瓦(kW)"},
	"hp":     langList{"zh-cn": "英制马力(hp)"},
	"ps":     langList{"zh-cn": "米制马力(ps)"},
	"kgm_s":  langList{"zh-cn": "公斤·米/秒(kg·m/s)"},
	"kcal_s": langList{"zh-cn": "千卡/秒(kcal/s)"},
	"btu_s":  langList{"zh-cn": "英热单位/秒(Btu/s)"},
	"ftlb_s": langList{"zh-cn": "英尺·磅/秒(ft·lb/s)"},
	"j_s":    langList{"zh-cn": "焦耳/秒(J/s)"},
	"nm_s":   langList{"zh-cn": "牛顿·米/秒(N·m/s)"},
}

// 时间(duration)
var durationLang = langMap{
	"yr":   langList{"zh-cn": "年(yr)"},
	"week": langList{"zh-cn": "周(week)"},
	"d":    langList{"zh-cn": "天(d)"},
	"h":    langList{"zh-cn": "时(h)"},
	"m":    langList{"zh-cn": "分(min)"},
	"s":    langList{"zh-cn": "秒(s)"},
	"ms":   langList{"zh-cn": "毫秒(ms)"},
	"μs":   langList{"zh-cn": "微秒(μs)"},
	"ns":   langList{"zh-cn": "纳秒(ns)"},
}

// 数据存储(datasize)
var datasizeLang = langMap{
	"bit":  langList{"zh-cn": "比特(bit)"},
	"byte": langList{"zh-cn": "字节(byte)"},
	"kb":   langList{"zh-cn": "千字节(kb)"},
	"mb":   langList{"zh-cn": "兆字节(mb)"},
	"gb":   langList{"zh-cn": "千兆字节(gb)"},
	"tb":   langList{"zh-cn": "太字节(tb)"},
	"pb":   langList{"zh-cn": "拍字节(pb)"},
	"eb":   langList{"zh-cn": "艾字节(eb)"},
}

// 功,能(energy)
var energyLang = langMap{
	"j":     langList{"zh-cn": "焦耳(J)"},
	"kj":    langList{"zh-cn": "千焦(kJ)"},
	"w_h":   langList{"zh-cn": "瓦·时(W·h)"},
	"kw_h":  langList{"zh-cn": "千瓦·时(kW·h),度"},
	"cal":   langList{"zh-cn": "卡(cal)"},
	"kcal":  langList{"zh-cn": "千卡(kcal)"},
	"ft_lb": langList{"zh-cn": "英尺·磅(ft·lb)"},
	"kg_m":  langList{"zh-cn": "公斤·米(kg·m)"},
	"hp_h":  langList{"zh-cn": "英制马力·时(hp·h)"},
	"ps_h":  langList{"zh-cn": "米制马力·时(ps·h)"},
	"btu":   langList{"zh-cn": "英热单位(btu)"},
}

// 力(force)
var forceLang = langMap{
	"n":   langList{"zh-cn": "牛(N)"},
	"kn":  langList{"zh-cn": "千牛(kN)"},
	"kgf": langList{"zh-cn": "千克力(kgf)"},
	"gf":  langList{"zh-cn": "克力(gf)"},
	"lbf": langList{"zh-cn": "磅力(lbf)"},
	"kip": langList{"zh-cn": "千磅力(kip)"},
	"tf":  langList{"zh-cn": "公吨力(tf)"},
	"dyn": langList{"zh-cn": "达因(dyn)"},
}

// 密度(density)
var densityLang = langMap{
	"kg_cu_m":  langList{"zh-cn": "千克/立方米(kg/m³)"},
	"kg_cu_dm": langList{"zh-cn": "千克/立方分米(kg/dm³)"},
	"kg_cu_cm": langList{"zh-cn": "千克/立方厘米(kg/cm³)"},
	"g_cu_m":   langList{"zh-cn": "克/立方米(g/m³)"},
	"g_cu_dm":  langList{"zh-cn": "克/立方分米(g/dm³)"},
	"g_cu_cm":  langList{"zh-cn": "克/立方厘米(g/cm³)"},
}

// 英文到单位翻译的map
var dime2lang = langHash{
	"mass":        massLang,
	"length":      lengthLang,
	"angle":       angleLang,
	"speed":       speedLang,
	"area":        areaLang,
	"volumn":      volumnLang,
	"temperature": temperatureLang,
	"pressure":    pressureLang,
	"power":       powerLang,
	"duration":    durationLang,
	"datasize":    datasizeLang,
	"energy":      energyLang,
	"force":       forceLang,
	"density":     densityLang,
}
