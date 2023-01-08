package controller

import (
	"net/http"
	"userservice/helper"
	"userservice/model"

	"github.com/gin-gonic/gin"
)

func ApplicationAddNew(context *gin.Context) {
	var input model.Applications
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unique, _ := helper.GenerateAppId(3)
	input.Id = unique

	savedEntry, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}
