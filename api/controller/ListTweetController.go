package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *TweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}
