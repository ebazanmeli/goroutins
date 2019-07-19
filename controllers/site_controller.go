package controllers

import (
	"../servicies"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	paramSiteID = "siteID"
)

func GetSite(c *gin.Context) {
	siteID := c.Param(paramSiteID)
	site, apiError := servicies.GetSiteFromAPI(siteID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, site)
}
