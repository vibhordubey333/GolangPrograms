/*
	Purpose of interface:
	Interface allows us to encapsulate the logic within user defined data type.
*/
package main

import (
	"fmt"
)

type Metal struct {
	mass   float64
	volume float64
}

type Gas struct {
	pressure      float64
	temprature    float64
	molecularMass float64
}

//Return density of metal.
func (m *Metal) Density() float64 {
	return m.mass / m.volume
}

//Return density if gas.
func (g *Gas) Density() float64 {
	return (g.molecularMass * g.pressure) / (0.0821 * (g.temprature + 273))

}

/*
	Here Dense interface is having Density() method which has been implemented by
	Metal and Gas structs.
*/
type Dense interface {
	Density() float64
}

/*
	IsDenser() takes interface as an arguement. Based on the
	passed object it calls respective Density() method to compute .
*/
func IsDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

func main() {
	gold := Metal{478, 24}
	silver := Metal{100, 10}

	result := IsDenser(&gold, &silver)
	if result {
		fmt.Println("Gold has higher density than silver.")
	} else {
		fmt.Println("Silver has higher density than gold.")
	}

	oxygen := Gas{
		pressure:      5,
		temprature:    27,
		molecularMass: 32,
	}

	hydrogen := Gas{
		pressure:      1,
		temprature:    0,
		molecularMass: 2,
	}

	result = IsDenser(&oxygen, &hydrogen)

	if result {
		fmt.Println("Oxygen has higher density than hydrogen")
	} else {
		fmt.Println("Hygrogen has higher density than oxygen")
	}
}
