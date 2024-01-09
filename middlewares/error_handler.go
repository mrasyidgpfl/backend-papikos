package middlewares

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/httperror"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	if len(c.Errors) == 0 {
		return
	}

	firstError := c.Errors[0].Err
	appError, isAppError := firstError.(httperror.AppError)
	if isAppError {
		c.JSON(appError.StatusCode, appError)
		return
	}
	serverErr := httperror.InternalServerError(firstError.Error())
	c.JSON(serverErr.StatusCode, serverErr)
}
