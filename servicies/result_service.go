package servicies

import (
	"../domain"
	"../utils/apierrors"
)

var user *domain.User

func GetResultFromAPI(userID int64) (*domain.Result, *apierrors.ApiError) {

	getUser(userID)

	countryChannel := make(chan *domain.Country)
	siteChannel := make(chan *domain.Site)
	if user != nil {
		go getCountry(user, countryChannel)
		go getSite(user, siteChannel)
	}
	country, site := <-countryChannel, <-siteChannel

	result := &domain.Result{
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

func getCountry(user *domain.User, countryChannel chan *domain.Country) *apierrors.ApiError {
	var err *apierrors.ApiError
	var country *domain.Country
	country, err = GetCountryFromAPI(user.CountryID)
	if err != nil {
		return err
	}
	countryChannel <- country
	return nil
}

func getSite(user *domain.User, siteChannel chan *domain.Site) *apierrors.ApiError {
	var err *apierrors.ApiError
	var site *domain.Site
	site, err = GetSiteFromAPI(user.SiteID)
	if err != nil {
		return err
	}
	siteChannel <- site
	return nil
}
