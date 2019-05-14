package unit

import (
	"errors"
	"strings"
)

var (
	UnsupportUnit error = errors.New("unit is not supported")
	UnsupportDime error = errors.New("Dime is not supported")
	NotSameDime   error = errors.New("not the same dimesion")
	NotFloat64    error = errors.New("unit convert to float64 failed")
)

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
	case "datasize":
		hash = datasizeHash
	case "energy":
		hash = energyHash
	case "force":
		hash = forceHash
	case "density":
		hash = densityHash
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
