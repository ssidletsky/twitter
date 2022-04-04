package usecases

import (
	"context"

	"github.com/ssidletsky/esportal-twitter/app/tweets/entities"
	"github.com/ssidletsky/esportal-twitter/app/tweets/repository"
)

// Tweet interface
type Tweet interface {
	FindTweets(ctx context.Context, userID, fromTweetID uint32, limit int) ([]entities.Tweet, error)
}

// NewTweet configures and returns TweetCases
func NewTweet(r repository.TweetsQuerier) Tweet {
	if r == nil {
		panic("invalid tweets querier")
	}
	tq := tweetUseCases{
		tq: r,
	}
	return &tq
}

type tweetUseCases struct {
	tq repository.TweetsQuerier
}

func (c *tweetUseCases) FindTweets(ctx context.Context, userID, fromTweetID uint32, limit int) ([]entities.Tweet, error) {
	return c.tq.FindTweets(ctx, userID, fromTweetID, limit)
}
