package actions

import (
	"fmt"
	"os"

	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/joho/godotenv"
	"github.com/markbates/going/defaults"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/twitter"
	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
	"github.com/niranjan92/go-hackathon-starter/models"
	"github.com/pkg/errors"
)

func init() {
	gothic.Store = App().SessionStore
	err := godotenv.Load()
	if err != nil {
		log.Println("You need to add .env file to enable authentications")
	}
	var providers []goth.Provider

	if os.Getenv("TWITTER_CONSUMER_KEY") != "" {
		providers = append(providers,
			twitter.New(
				os.Getenv("TWITTER_CONSUMER_KEY"),
				os.Getenv("TWITTER_CONSUMER_SECRET"),
				"http://localhost:3000/auth/twitter/callback"))
	}

	if os.Getenv("FACEBOOK_KEY") != "" {
		providers = append(providers,
			facebook.New(
				os.Getenv("FACEBOOK_KEY"),
				os.Getenv("FACEBOOK_SECRET"),
				fmt.Sprintf("%s%s", "http://localhost:3000", "/auth/facebook/callback")))
	}

	if os.Getenv("GITHUB_KEY") != "" {
		providers = append(providers,
			github.New(
				os.Getenv("GITHUB_KEY"),
				os.Getenv("GITHUB_SECRET"),
				fmt.Sprintf("%s%s", "http://localhost:3000", "/auth/github/callback")))
	}

	goth.UseProviders(providers[:]...)

}

func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			if err := tx.Find(u, uid); err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}

func AuthCallback(c buffalo.Context) error {
	gu, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	tx := c.Value("tx").(*pop.Connection)

	//TODO: save data to user profiles after creating user
	// a user should not be created if there is already an user profile
	// for the given user for given email. Everything should get tied to
	// and email
	q := tx.Where("provider = ? and provider_id = ?", gu.Provider, gu.UserID)
	exists, err := q.Exists("users")
	if err != nil {
		return errors.WithStack(err)
	}
	u := &models.User{}
	if exists {
		if err = q.First(u); err != nil {
			return errors.WithStack(err)
		}
	}
	u.Provider = gu.Provider
	u.ProviderID = gu.UserID
	u.Name = defaults.String(gu.Name, gu.NickName)
	u.Email = nulls.NewString(gu.Email)
	u.Gravatar = nulls.NewString(gu.AvatarURL)
	if err = tx.Save(u); err != nil {
		return errors.WithStack(err)
	}

	c.Session().Set("current_user_id", u.ID)
	if err = c.Session().Save(); err != nil {
		return errors.WithStack(err)
	}

	c.Flash().Add("success", "You have been logged in")
	return c.Redirect(302, "/")

}

func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You have been logged out")
	return c.Redirect(302, "/")
}
