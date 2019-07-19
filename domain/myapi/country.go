package myapi

import (
	"../../utils"
	"../../utils/apierrors"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Country struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Locale             string `json:"locale"`
	Currency           string `json:"currency"`
	DecimalSeparator   string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone           string `json:"time_zone"`
	GeoInformation     struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"states"`
}

func (country *Country) GetCountry() *apierrors.ApiError {
	if country.ID == "" {
		return &apierrors.ApiError{
			Message: "El id country está vacío",
			Status:  http.StatusBadRequest,
		}
	}

	url := utils.UrlCountry + country.ID
	res, err := http.Get(url)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err = json.Unmarshal(data, &country); err != nil {
		if err != nil {
			return &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	return nil
}
