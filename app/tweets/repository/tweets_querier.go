package repository

import (
	"context"

	"github.com/ssidletsky/esportal-twitter/app/tweets/entities"
)

// TweetsQuerier describes required methods to be implemented by repository
type TweetsQuerier interface {
	// FindTweets retrieves a list of latest tweets posted by users followed by *userID*.
	// *fromTweetID* allows to skip already seen tweets.
	// *limit* limits the list.
	FindTweets(ctx context.Context, userID, fromTweetID uint32, limit int) ([]entities.Tweet, error)
}
