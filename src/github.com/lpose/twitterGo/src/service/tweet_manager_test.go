package service_test

import (
	"testing"

	"github.com/lpose/twitterGo/src/domain"
	"github.com/lpose/twitterGo/src/service"
)

func isValidTweet(t *testing.T, publishedTweet *domain.TextTweet, id int, user *domain.User, text string) bool {
	if publishedTweet.User == user &&
		publishedTweet.Text == text &&
		publishedTweet.GetId() == id {
		return true
	}
	return false
}

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	userManager := service.GetInstance()
	tweetManager := service.NewTweetManager()
	var id int
	var tweet *domain.TextTweet
	user := &domain.User{Name: "lucas"}
	userManager.Register(user)
	text := "This is my first tweet"
	tweet = domain.NewTweet(id, user, text)

	// Operation
	tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %v: %s \nbut is %v: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	userManager := service.GetInstance()
	tweetManager := service.NewTweetManager()
	var tweet *domain.TextTweet
	var id int
	user := &domain.User{Name: ""}
	userManager.Register(user)
	text := "This is my first tweet"
	tweet = domain.NewTweet(id, user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	userManager := service.GetInstance()
	tweetManager := service.NewTweetManager()
	var tweet *domain.TextTweet
	var id int
	user := &domain.User{Name: "lucas"}
	userManager.Register(user)
	var text string
	tweet = domain.NewTweet(id, user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("text is required")
	}

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}
func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization
	userManager := service.GetInstance()
	tweetManager := service.NewTweetManager()
	var tweet *domain.TextTweet
	var id int
	user := &domain.User{Name: "lucas"}
	userManager.Register(user)
	text := "Este tweet tiene mas de 140 caracr probaraaaa el test TweetWhichExceeding140CharactersIsNotPublished modificando el tweet manager para que controle la longitud del texto"
	tweet = domain.NewTweet(id, user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Text must not exced 140 characters")
	}

	if err != nil && err.Error() != "Text must not exced 140 characters" {
		t.Error("Text must not exced 140 characters", err.Error())

	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	userManager := service.GetInstance()
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet *domain.TextTweet // Fill the tweets with data
	var id int
	user := &domain.User{Name: "lucas"}
	userManager.Register(user)
	text := "This is my first tweet"
	tweet = domain.NewTweet(id, user, text)

	userSecond := &domain.User{Name: "Juan"}
	userManager.Register(userSecond)
	textSecond := "This is my second tweet"
	secondTweet = domain.NewTweet(id, userSecond, textSecond)

	// Operation
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	if !isValidTweet(t, firstPublishedTweet, id, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, id, userSecond, textSecond) {
		return
	}
	// Same for secondPublishedTweet
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	userManager := service.GetInstance()
	tweetManager := service.NewTweetManager()
	var tweet *domain.TextTweet
	id := 1

	user := &domain.User{Name: "lucas"}
	user.SetId(id)
	userManager.Register(user)
	text := "This is my first tweet"

	tweet = domain.NewTweet(id, user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	userManager := service.GetInstance()
	tweetManager := service.NewTweetManager()
	var tweet, secondTweet, thirdTweet *domain.TextTweet

	user := &domain.User{Name: "lucas"}
	userManager.Register(user)
	id := 1
	user.SetId(1)
	secondUser := &domain.User{Name: "jose"}
	userManager.Register(secondUser)
	secondUser.SetId(2)

	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(id, user, text)
	secondTweet = domain.NewTweet(id, user, secondText)
	thirdTweet = domain.NewTweet(id, secondUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)
	// Operation
	count := tweetManager.CountTweetsByUser(user)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	userManager := service.GetInstance()
	tweetManager := service.NewTweetManager()
	var id int
	var tweet, secondTweet, thirdTweet *domain.TextTweet

	user := &domain.User{Name: "lucas"}
	userManager.Register(user)
	user.SetId(1)
	secondUser := &domain.User{Name: "Jorge"}
	userManager.Register(secondUser)
	secondUser.SetId(2)
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(id, user, text)
	secondTweet = domain.NewTweet(id, user, secondText)
	thirdTweet = domain.NewTweet(id, secondUser, text)
	// publish the 3 tweets
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)
	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */
		t.Errorf("Fallo")
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	isValidTweet(t, firstPublishedTweet, id, user, text)
	isValidTweet(t, secondPublishedTweet, id, user, secondText)
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet
}
