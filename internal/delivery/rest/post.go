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
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *Rest) DeletePost(ctx *gin.Context) {
	param := model.PostDelete{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error":   err,
		})
		return
	}

	post, err := r.usecase.PostUsecase.GetPost(model.PostParam{
		ID: param.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get post",
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

	if post.UserID != user.(entity.User).ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "access not permitted",
		})
		return
	}

	err = r.usecase.PostUsecase.DeletePost(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
