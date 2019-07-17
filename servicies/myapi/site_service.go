package myapi

import "../../domain/myapi"
import "../../utils/apierrors"

func GetSiteFromAPI(siteID string) (*myapi.Site, *apierrors.ApiError) {
	site := &myapi.Site{
		ID: siteID,
	}
	err := site.GetSite()
	if err != nil {
		return nil, err
	}
	return site, nil
}
