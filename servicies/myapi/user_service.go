package myapi

import "../../domain/myapi"
import "../../utils/apierrors"

func GetUserFromAPI(userID int64) (*myapi.User, *apierrors.ApiError) {
	user := &myapi.User{
		ID: userID,
	}
	err := user.GetUser()
	if err != nil {
		return nil, err
	}
	return user, nil
}
