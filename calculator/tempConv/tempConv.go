package tempConv

type Convert interface {
	Convert2C(arg float64) float64
	Convert2F(arg float64) float64
	Convert2K(arg float64) float64
}

type Celcius struct {
	temperature float64
}

type Fahrenheit struct {
	temperature float64
}

type Kelvin struct {
	temperature float64
}

func (c Celcius) Convert2C(arg float64) float64 {
	return c.temperature
}

func (c Celcius) Convert2F(arg float64) float64 {
	return c.temperature*9/5 + 32
}

func (c Celcius) Convert2K(arg float64) float64 {
	return c.temperature + 273.15
}

func (c Fahrenheit) Convert2C(arg float64) float64 {
	return (c.temperature - 32) * 5 / 9
}

func (c Fahrenheit) Convert2F(arg float64) float64 {
	return c.temperature
}

func (c Fahrenheit) Convert2K(arg float64) float64 {
	return (c.temperature + 459.67) * 5 / 9
}

func (c Kelvin) Convert2C(arg float64) float64 {
	return c.temperature - 273.15
}

func (c Kelvin) Convert2F(arg float64) float64 {
	return c.temperature*9/5 - 459.67
}

func (c Kelvin) Convert2K(arg float64) float64 {
	return c.temperature
}
