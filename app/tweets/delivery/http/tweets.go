package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/ssidletsky/esportal-twitter/app/tweets/usecases"
)

// TweetsController implements restful methods for "/tweets" endpoint
type TweetsController struct {
	uc usecases.Tweet
}

// RegisterTweetsHandler registers tweets handlers to the fiber app
func RegisterTweetsController(app *fiber.App, uc usecases.Tweet) {
	tc := TweetsController{uc}
	app.Get("/tweets", tc.Index)
}

// Index provides tweets for currently authenticated user posted by user's followers
func (tc *TweetsController) Index(c *fiber.Ctx) error {
	ctx := c.UserContext()

	// using directly user-id here as there is no authorization
	userIDParam := c.Get("x-user-id")
	userID, err := strconv.ParseInt(userIDParam, 10, 0)
	if userID == 0 || err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "unauthorized",
		})
	}

	fromTweetIDParam := c.Query("from_tweet_id", "0")
	fromTweetID, err := strconv.ParseInt(fromTweetIDParam, 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "invalid from_tweet_id param",
		})
	}

	perPageString := c.Query("per_page", "100")
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "invalid per_page param",
		})
	}

	tweets, err := tc.uc.FindTweets(ctx, 1, uint32(fromTweetID), perPage)
	if err != nil {
		log.Errorf("FindTweets error: %w", err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": fiber.ErrInternalServerError,
		})
	}
	return c.JSON(&fiber.Map{
		"tweets": tweets,
	})
}
