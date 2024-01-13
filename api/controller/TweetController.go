package controllers

import (
	"github.com/diogenesOliver/goAPI/api/entities"
)

type TweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *TweetController {
	return &TweetController{}
}
