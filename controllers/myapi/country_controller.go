package myapi

import (
	"../../servicies/myapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	paramCountryID = "countryID"
)

func GetCountry(c *gin.Context) {
	countryID := c.Param(paramCountryID)
	country, apiError := myapi.GetCountryFromAPI(countryID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, country)
}
