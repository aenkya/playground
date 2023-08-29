package prototype

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

type Pair[K comparable, V any] struct {
	key   K
	value V
}

func Creational(p Person, newData ...Pair[string, string]) Person {

	p2 := p

	for _, d := range newData {
		switch d.key {
		case "name":
			p2.Name = d.value
		case "streetaddress":
			p2.Address.StreetAddress = d.value
		case "city":
			p2.Address.City = d.value
		case "country":
			p2.Address.Country = d.value
		default:
			panic(fmt.Sprintf("key %s not recognized", d.key))
		}
	}

	return p2
}
