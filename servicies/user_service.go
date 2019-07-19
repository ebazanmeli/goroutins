package servicies

import "../domain"
import "../utils/apierrors"

func GetUserFromAPI(userID int64) (*domain.User, *apierrors.ApiError) {
	user := &domain.User{
		ID: userID,
	}
	err := user.GetUser()
	if err != nil {
		return nil, err
	}
	return user, nil
}
