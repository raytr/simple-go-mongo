package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	loadData()

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("all_country", currentWorldPopulation)
	e.GET("country_names", allCountryNamePublic)

	log.Fatal(e.Start(":8080"))
}

func allCountryNamePublic(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, CountriesName)
}

func currentWorldPopulation(ctx echo.Context) error {
	a := CountryRangingMapping
	fmt.Println(a)
	return ctx.JSON(http.StatusOK, Countries)
}

func top20(ctx echo.Context) error {
	top20 := make([]Country, 0, 20)
	for i := 1; i <= 20; i++ {
		top20 = append(top20, CountryRangingMapping[i])
	}
	return ctx.JSON(http.StatusOK, top20)
}
