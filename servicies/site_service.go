package servicies

import "../domain"
import "../utils/apierrors"

func GetSiteFromAPI(siteID string) (*domain.Site, *apierrors.ApiError) {
	site := &domain.Site{
		ID: siteID,
	}
	err := site.GetSite()
	if err != nil {
		return nil, err
	}
	return site, nil
}
