package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

var CountriesName []string
var CountryRangingMapping map[int]Country

func loadData() {
	allCountryName()
	getAllWorldPopulation()

}

func getAllWorldPopulation() {

	//prepare model
	type respModel struct {
		Body Country `json:"body"`
	}
	//Countries := make([]Country, 0, len(CountriesName))

	//load data
	for _, name := range CountriesName {
		url := "https://world-population.p.rapidapi.com/population?country_name=" + name
		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-rapidapi-host", "world-population.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", "c007516201msh3e36c96f39107f2p1581b5jsn95b947e33086")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		//fmt.Println(res)
		//fmt.Println(string(body))

		resp := respModel{}
		if err := json.Unmarshal(body, &resp); err != nil {
			continue
		}

		CountryRangingMapping[resp.Body.Ranking] = Country{
			CountryName: resp.Body.CountryName,
			Population:  resp.Body.Population,
			Ranking:     resp.Body.Ranking,
			WorldShare:  resp.Body.WorldShare,
		}
	}

	log.Infoln(CountryRangingMapping)
}

func allCountryName() {

	type respModel struct {
		Body CountryNames `json:"body"`
	}

	url := "https://world-population.p.rapidapi.com/allcountriesname"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "world-population.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "c007516201msh3e36c96f39107f2p1581b5jsn95b947e33086")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	resp := respModel{}
	json.Unmarshal(body, &resp)

	CountriesName = resp.Body.CountriesName
}
