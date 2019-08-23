package service

import (
	"errors"

	"github.com/lpose/twitterGo/src/domain"
)

type TweetManager struct {
	miTweet     *domain.TextTweet
	misTweets   []*domain.TextTweet
	quoteTweets []*domain.QuoteTweet
	imageTweets []*domain.ImageTweet
}

func NewTweetManager() *TweetManager {
	tweetManager := new(TweetManager)
	tweetManager.misTweets = make([]*domain.TextTweet, 0)
	return tweetManager
}

func (manager *TweetManager) PublishTweet(tweet domain.Tweet) (int, error) {
	userManager := GetInstance()
	if userManager.GetUser() == nil {
		return -1, errors.New("No hay usuario registrado")
	}
	if tweet.User.Name == "" {
		return -1, errors.New("user is required")
	}

	if tweet.Text == "" {
		return -1, errors.New("text is required")
	}

	if len(tweet.Text) > 140 {
		return -1, errors.New("Text must not exced 140 characters")
	}

	manager.miTweet = tweet
	manager.misTweets = append(manager.misTweets, tweet)

	return tweet.GetId(), nil
}

func (manager *TweetManager) GetTweet() *domain.TextTweet {
	return manager.miTweet
}

func (manager *TweetManager) GetTweets() []*domain.TextTweet {
	return manager.misTweets
}

func (manager *TweetManager) GetTweetById(id int) *domain.TextTweet {
	for _, tweet := range manager.misTweets {
		if tweet.GetId() == id {
			return tweet
		}
	}
	return nil
}

func (manager *TweetManager) CountTweetsByUser(user *domain.User) int {
	count := 0
	for _, tweet := range manager.misTweets {
		if tweet.User.GetId() == user.GetId() {
			count++
		}

	}
	return count
}

func (manager *TweetManager) GetTweetsByUser(user *domain.User) []*domain.TextTweet {
	tweets := make([]*domain.TextTweet, 0)
	for _, tweet := range manager.misTweets {
		if tweet.User.GetId() == user.GetId() {
			tweets = append(tweets, tweet)
		}
	}
	return tweets
}
