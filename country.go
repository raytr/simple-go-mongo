package main

type CountryNames struct {
	CountriesName []string `json:"countries"`
}

type Country struct {
	CountryName string  `json:"country_name"`
	Population  uint64  `json:"population"`
	Ranking     int     `json:"ranking"`
	WorldShare  float64 `json:"world_share"`
}
