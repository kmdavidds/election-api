package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmdavidds/election-api/entity"
	"github.com/kmdavidds/election-api/model"
)

func (r *Rest) CreatePost(ctx *gin.Context) {
	param := model.PostCreate{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error":   err,
		})
		return
	}

	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed get login user",
		})
		return
	}

	param.UserID = user.(entity.User).ID

	err = r.usecase.PostUsecase.CreatePost(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}