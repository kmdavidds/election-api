package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmdavidds/election-api/entity"
)

func (m *middleware) OnlyCandidates(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user",
		})
		ctx.Abort()
		return
	}

	if !user.(entity.User).IsCandidate {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user is not a candidate",
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}
