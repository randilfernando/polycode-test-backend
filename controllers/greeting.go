package controllers

import (
	"net/http"

	"github.com/cloudimpl/next-coder-sdk/api"
	"github.com/cloudimpl/next-coder-sdk/apicontext"
	"github.com/cloudimpl/next-coder-sdk/polycode"
	"github.com/gin-gonic/gin"
	"portal/register/model"
)

func greeting(c *gin.Context) {
	request := model.HelloRequest{}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiCtx, err := apicontext.FromContext(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// The service call remains the same, but the request/response types are updated
	resp, err := apiCtx.Service("greeting-service").Get().
		RequestReply(polycode.TaskOptions{}, "Greeting", request).GetAny()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

var Greeting = api.FromWorkflow(greeting)