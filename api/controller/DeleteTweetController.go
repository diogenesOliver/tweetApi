package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *TweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, tweet := range t.tweets {
		if tweet.ID == id {
			t.tweets = append(t.tweets[0:idx], t.tweets[idx+1:]...)
			ctx.JSON(http.StatusOK, gin.H{
				"OK": "Tweet Deleted",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet not found",
	})
}
