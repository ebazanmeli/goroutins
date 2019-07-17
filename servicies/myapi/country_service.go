package myapi

import "../../domain/myapi"
import "../../utils/apierrors"

func GetCountryFromAPI(countryID string) (*myapi.Country, *apierrors.ApiError) {
	country := &myapi.Country{
		ID: countryID,
	}
	err := country.GetCountry()
	if err != nil {
		return nil, err
	}
	return country, nil
}
