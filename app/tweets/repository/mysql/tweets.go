package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/huandu/go-sqlbuilder"

	"github.com/ssidletsky/esportal-twitter/app/tweets/entities"
	"github.com/ssidletsky/esportal-twitter/app/tweets/repository"
)

// NewTweetsRepository provides tweets repository.
func NewRepository() repository.TweetsQuerier {
	if Conn == nil {
		panic("conn is not initialized")
	}
	r := repo{
		conn: Conn,
	}
	return &r
}

// repository implements repository.TweetsQuerier interface using MySQL database.
type repo struct {
	conn *sql.DB
}

// FindTweets retrieves a list of the latest tweets posted by users followed by *userID*.
// *fromTweetID* allows to skip already seen tweets.
// *limit* limits the list.
func (r *repo) FindTweets(ctx context.Context, userID, fromTweetID uint32, limit int) ([]entities.Tweet, error) {
	if userID == 0 {
		return nil, errors.New("invalid userID")
	}

	sb := sqlbuilder.NewSelectBuilder()
	whereExpr := make([]string, 0, 2)
	whereExpr = append(whereExpr, sb.Equal("followers.followed_user_id", userID))
	if fromTweetID > 0 {
		whereExpr = append(whereExpr, sb.LessThan("tweets.id", fromTweetID))
	}
	sb.Select("tweets.id", "users.username", "users.first_name as first_name", "users.last_name", "tweets.text", "tweets.publication_date").
		From("tweets").
		Join("users", "users.id=tweets.author_user_id").
		Join("followers", "followers.follower_user_id=tweets.author_user_id").
		Where(whereExpr...).
		OrderBy("tweets.id").Desc().
		Limit(limit)

	query, args := sb.Build()
	rows, err := r.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var tweets []entities.Tweet
	for rows.Next() {
		var tweet entities.Tweet
		err := rows.Scan(
			&tweet.ID,
			&tweet.Username,
			&tweet.FirstName,
			&tweet.LastName,
			&tweet.Text,
			&tweet.PublicationDate,
		)
		if err != nil {
			return nil, fmt.Errorf("rows scan error: %w", err)
		}
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
