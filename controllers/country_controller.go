package controllers

import (
	"../domain"
	"../servicies"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	paramCountryID = "countryID"
)

func GetCountry(c *gin.Context) {
	rafaga := make(chan time.Time, 3)
	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			for i := 0; i < 3; i++ {
				rafaga <- t
			}
		}
	}()

	rafagaRequest := make(chan *domain.Country, 15)
	for i := 1; i <= 15; i++ {
		countryID := c.Param(paramCountryID)
		country, apiError := servicies.GetCountryFromAPI(countryID)
		if apiError != nil {
			c.JSON(apiError.Status, apiError)
			return
		}
		rafagaRequest <- country
	}
	close(rafagaRequest)
	for req := range rafagaRequest {
		<-rafaga
		c.JSON(http.StatusOK, req)
		fmt.Println(req)
	}
}
