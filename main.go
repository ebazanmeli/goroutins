package main

import "github.com/gin-gonic/gin"
import "./controllers"

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/user/:userID", controllers.GetUser)
	router.GET("/country/:countryID", controllers.GetCountry)
	router.GET("/site/:siteID", controllers.GetSite)
	router.GET("/result/:userID", controllers.GetResult)
	router.GET("/resultRL/:userID", controllers.GetResultRateLimited)
	router.Run(port)
}
