package core

type City string

const (
	cali         City = "cali"
	medellin     City = "medellin"
	barranquilla City = "barranquilla"
)

var CitiesAllowed = map[City]bool{
	cali:         true,
	medellin:     true,
	barranquilla: true,
}
