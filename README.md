# About

[![GoDoc](https://godoc.org/github.com/foxmeder/unit?status.svg)](https://godoc.org/github.com/foxmeder/unit)

Conversion of unit library for golang

forked from [martinlindhe/unit](https://github.com/martinlindhe/unit)

## new feature

Add more unit including common unit from China
new function `UnitConv` to do conversion conveniently

## Installation

```
go get -u github.com/foxmeder/unit
```


## General use

Basic usage:
```go
ft := 1 * unit.Foot
fmt.Println(ft.Feet(), "feet is", ft.Meters(), "meters")
```

To use your own data type, you need to convert to the base unit first (eg Length, Speed etc):
```go
type MyUnit int

n := MyUnit(2)
ft := Length(n) * Foot
fmt.Println(ft.Feet(), "feet is", ft.Meters(), "meters")
```

Conversion func usage:
`from` and `to` use format of `dimension`.`name`,see [dimension list](dimension.md) for details
For example:use `pressure.atm` for Atmosphere

```go
from := "pressure.atm"
to := "pressure.mmhg"
result,err := unit.UnitConv(from, to, 100)
```

## Temperature

Cannot be used to scale directly like the other units.
Instead, use the From* functions to create a Temperature type:

```go
f := unit.FromFahrenheit(100)

fmt.Println("100 fahrenheit in celsius = ", f.Celsius())
```


## Future work

Please note the resulting precision is limited to the float64 type.
Big decimal version is being tracked in https://github.com/martinlindhe/unit/issues/3


## License

Under [MIT](LICENSE)
