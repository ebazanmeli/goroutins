package myapi

import (
	"../../domain/myapi"
	"../../utils/apierrors"
)

var user *myapi.User

func GetResultFromAPI(userID int64) (*myapi.Result, *apierrors.ApiError) {

	getUser(userID)

	countryChannel := make(chan *myapi.Country)
	siteChannel := make(chan *myapi.Site)
	if user != nil {
		go getCountry(user, countryChannel)
		go getSite(user, siteChannel)
	}
	country, site := <-countryChannel, <-siteChannel

	result := &myapi.Result{
		User:    user,
		Country: country,
		Site:    site,
	}

	return result, nil
}

func getUser(userID int64) *apierrors.ApiError {
	var err *apierrors.ApiError
	user, err = GetUserFromAPI(userID)
	if err != nil {
		return err
	}
	return nil
}

func getCountry(user *myapi.User, countryChannel chan *myapi.Country) *apierrors.ApiError {
	var err *apierrors.ApiError
	var country *myapi.Country
	country, err = GetCountryFromAPI(user.CountryID)
	if err != nil {
		return err
	}
	countryChannel <- country
	return nil
}

func getSite(user *myapi.User, siteChannel chan *myapi.Site) *apierrors.ApiError {
	var err *apierrors.ApiError
	var site *myapi.Site
	site, err = GetSiteFromAPI(user.SiteID)
	if err != nil {
		return err
	}
	siteChannel <- site
	return nil
}
