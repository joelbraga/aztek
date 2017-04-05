package api

import "github.com/gin-gonic/gin"

func(api *ApiResource) HttpResponse(status string, message string, data interface{}) map[string]interface{}{
	return gin.H{
		"status": status,
		"message": message,
		"result": data,
	}
}


