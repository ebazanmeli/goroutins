package main

import "github.com/gin-gonic/gin"
import "./controllers/myapi"

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/user/:userID", myapi.GetUser)
	router.GET("/country/:countryID", myapi.GetCountry)
	router.GET("/site/:siteID", myapi.GetSite)
	router.GET("/result/:userID", myapi.GetResult)
	router.Run(port)
}
