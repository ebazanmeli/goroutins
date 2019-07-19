package myapi

import (
	"../../utils/"
	"../../utils/apierrors"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Site struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	CountryID          string `json:"country_id"`
	SaleFeesMode       string `json:"sale_fees_mode"`
	MercadopagoVersion int64  `json:"mercadopago_version"`
	DefaultCurrencyId  string `json:"default_currency_id"`
	ImmediatePayment   string `json:"immediate_payment"`
	PaymentMethodIds   []interface {
	} `json:"payment_method_ids"`
	Settings struct {
		IdentificationTypes      []interface{} `json:"identification_types"`
		TaxpayerTypes            []interface{} `json:"taxpayer_types"`
		IdentificationTypesRules interface{}
	} `json:"settings"`
	Currencies []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	}
	Categories []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
}

func (site *Site) GetSite() *apierrors.ApiError {
	if site.ID == "" {
		return &apierrors.ApiError{
			Message: "El id site está vacío",
			Status:  http.StatusBadRequest,
		}
	}

	url := utils.UrlSite + site.ID
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

	if err = json.Unmarshal(data, &site); err != nil {
		if err != nil {
			return &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	return nil
}
