package domain_test

import (
	"testing"

	"github.com/lpose/twitterGo/src/domain"
	"github.com/lpose/twitterGo/src/service"
)

func TestCanGetAPrintableTweet(t *testing.T) {

	// Initialization

	userManager := service.GetInstance()
	user := &domain.User{Nick: "Lucas"}
	userManager.Register(user)
	var id int
	tweet := domain.NewTweet(id, user, "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@Lucas: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	userManager := service.GetInstance()
	var id int
	user := &domain.User{Nick: "Lucas"}
	userManager.Register(user)
	tweet := domain.NewTweet(id, user, "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@Lucas: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
