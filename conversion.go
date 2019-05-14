package unit

import (
	"errors"
	"strings"
)

type unitItem struct {
	UnitName string `json:"unit"`
	UnitDime string `json:"type"`
	UnitCode string `json:"code"`
}

var (
	UnsupportUnit error = errors.New("unit is not supported")
	UnsupportDime error = errors.New("Dime is not supported")
	UnsupportLang error = errors.New("Lang is not supported")
	NotSameDime   error = errors.New("not the same dimesion")
	NotFloat64    error = errors.New("unit convert to float64 failed")
)

// UnitList 返回支持的单位列表
func UnitList(dime string) ([]unitItem, error) {
	lang := "zh-cn"
	if dime == "" {
		// 返回所有单位
		return allLang2list(lang)
	}
	langH, ok := dime2lang[dime]
	if !ok {
		return nil, UnsupportDime
	}
	// 返回单位列表
	return lang2list(dime, langH, lang)
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
	hash, ok := dime2unit[f[0]]
	if ok {
		return divConv(hash, f[1], t[1], value)
	}
	switch f[0] {
	case "temperature":
		return temperatureConv(f[1], t[1], value)
	default:
		return 0, UnsupportDime
	}
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

// 全部支持的单位列表
func allLang2list(lang string) ([]unitItem, error) {
	// lang = "zh-cn"
	var ret []unitItem
	for dime, langH := range dime2lang {
		tmp, err := lang2list(dime, langH, lang)
		if err != nil {
			return nil, err
		}
		ret = append(ret, tmp...)
	}
	return ret, nil
}

// 单个langMap转化为单位列表
func lang2list(dime string, langH langMap, lang string) ([]unitItem, error) {
	// lang = "zh-cn"
	len := len(langH)
	ret := make([]unitItem, len)
	i := 0
	for code, langTr := range langH {
		unitName, ok := langTr[lang]
		if !ok {
			return nil, UnsupportLang
		}
		ret[i] = unitItem{
			UnitName: unitName,
			UnitDime: dime,
			UnitCode: code,
		}
		i++
	}
	return ret, nil
}
