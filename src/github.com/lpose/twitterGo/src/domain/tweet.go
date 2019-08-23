package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	GetId() int
	GetName() string
	GetText() string
	GetQuote() Tweet
	GetUrlImage() string
	GetDate() *time.Time
}

type TextTweet struct {
	id   int
	User *User
	Text string
	Date *time.Time
}

type ImageTweet struct {
	TextTweet
	urlImage string
}

type QuoteTweet struct {
	TextTweet
	quoteTweet Tweet
}

func NewTweetText(id int, user *User, text string) *TextTweet {
	date := time.Now()

	tweet := TextTweet{
		id,
		user,
		text,
		&date,
	}

	return &tweet
}

func NewTweetImage(id int, user *User, text string, url string) *ImageTweet {
	date := time.Now()

	tweet := ImageTweet{
		TextTweet{
			id,
			user,
			text,
			&date,
		},
		url,
	}

	return &tweet
}

func NewTweetQuote(id int, user *User, text string, quote Tweet) *QuoteTweet {
	date := time.Now()

	tweet := QuoteTweet{
		TextTweet{
			id,
			user,
			text,
			&date,
		},
		quote,
	}

	return &tweet
}

func (t *TextTweet) GetId() int {
	return t.id
}

func (t *TextTweet) String() string {
	return fmt.Sprintf("@%v: %v", t.User.Nick, t.Text)

}

func (t *TextTweet) GetName() string {
	return t.User.Name
}

func (t *TextTweet) GetText() string {
	return t.Text
}

func (t *ImageTweet) GetUrlImage() string {
	return t.urlImage
}

func (t *QuoteTweet) GetQuote() Tweet {
	return t.quoteTweet
}

func (t *TextTweet) GetDate() *time.Time {
	return t.Date
}
