package handlers

import (
	"bus_depot/internal/errs"
	"errors"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	} else if errors.Is(err, errs.ErrValidationFailed) ||
		errors.Is(err, errs.ErrUserAlreadyExists) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrUserNotFound) ||
		errors.Is(err, errs.ErrUserNotFound) ||
		errors.Is(err, errs.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrValidationFailed) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrNotFound) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("something went wrong: %s", err.Error()),
		})
	}
}
