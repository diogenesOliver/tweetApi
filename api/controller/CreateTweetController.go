package controllers

import (
	"net/http"

	"github.com/diogenesOliver/goAPI/api/entities"
	"github.com/gin-gonic/gin"
)

func (t *TweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()
	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}

	t.tweets = append(t.tweets, *tweet)
	ctx.JSON(http.StatusOK, t.tweets)
}
