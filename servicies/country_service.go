package servicies

import "../domain"
import "../utils/apierrors"

func GetCountryFromAPI(countryID string) (*domain.Country, *apierrors.ApiError) {
	country := &domain.Country{
		ID: countryID,
	}
	err := country.GetCountry()
	if err != nil {
		return nil, err
	}
	return country, nil
}
