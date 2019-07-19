package controllers

import (
	"../servicies"
	"../utils/apierrors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetResult(c *gin.Context) {
	userID := c.Param(paramUserID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiError := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}
	result, apiError := servicies.GetResultFromAPI(id)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, result)
}
