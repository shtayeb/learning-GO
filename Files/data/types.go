package data

type integer = int
type json = map[string]string

var x integer

type distance float32 // miles
type distanceKm float64

// Method for a type
func (miles distance) ToKm() distanceKm {
	return distanceKm(1.6 * miles)
}

// Method for a type
func (km distanceKm) ToMiles() distance {
	return distance(km / 1.6)
}

func test() {
	d := distance(4.5)
	// using the type method
	km := d.ToKm()

	print(km.ToMiles())
}