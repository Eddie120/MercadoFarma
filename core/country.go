package core

type Country string

const colombia Country = "colombia"

var CountriesAllowed = map[Country]bool{
	colombia: true,
}
