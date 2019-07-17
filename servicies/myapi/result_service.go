package myapi

import (
	"../../domain/myapi"
	"../../utils/apierrors"
	"sync"
)

var user *myapi.User
var country *myapi.Country
var site *myapi.Site
var wg sync.WaitGroup

func GetResultFromAPI(userID int64) (*myapi.Result, *apierrors.ApiError) {

	wg.Add(1)

	go getUser(userID)
	wg.Wait()

	if user != nil {
		wg.Add(2)
		go getCountry(user)
		go getSite(user)
	}
	wg.Wait()

	result := &myapi.Result{
		User:    user,
		Country: country,
		Site:    site,
	}

	return result, nil
}

func getUser(userID int64) *apierrors.ApiError {
	defer wg.Done()

	var err *apierrors.ApiError
	user, err = GetUserFromAPI(userID)
	if err != nil {
		return err

	}
	return nil
}

func getCountry(user *myapi.User) *apierrors.ApiError {
	defer wg.Done()

	var err *apierrors.ApiError
	country, err = GetCountryFromAPI(user.CountryID)
	if err != nil {
		return err

	}
	return nil
}

func getSite(user *myapi.User) *apierrors.ApiError {
	defer wg.Done()

	var err *apierrors.ApiError
	site, err = GetSiteFromAPI(user.SiteID)
	if err != nil {
		return err

	}
	return nil
}
