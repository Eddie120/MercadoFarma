package core

type City string

const (
	Cali         City = "cali"
	Medellin     City = "medellin"
	Barranquilla City = "barranquilla"
)

var CitiesAllowed = map[City]bool{
	Cali:         true,
	Medellin:     true,
	Barranquilla: true,
}
