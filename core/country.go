package core

type Country string

const Colombia Country = "colombia"

var CountriesAllowed = map[Country]bool{
	Colombia: true,
}
