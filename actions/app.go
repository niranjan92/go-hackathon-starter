package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
	"github.com/markbates/goth/gothic"
	"github.com/niranjan92/go_hackathon_starter/models"

	"log"
	"net/http"
	_ "net/http/pprof"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_go_hackathon_starter_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
			go func() {
				log.Println(http.ListenAndServe(":8080", nil))
			}()
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		app.GET("/logout", AuthDestroy)
		app.GET("/login", LoginHandler)
		app.DELETE("/users/", UserDestroy)
		app.GET("/", HomeHandler)
		app.GET("/react", ReactHandler)
		app.GET("/api-examples", ApiExampleHandler)
		app.GET("/dashboard", DashboardHandler)
		app.GET("/account/profile", ProfileHandler)
		app.POST("/account/profile", UpdateProfileHandler)

		app.ServeFiles("/assets", assetsBox)

		app.Use(SetCurrentUser)
		app.Use(Authorize)
		app.Middleware.Skip(Authorize, HomeHandler)
		app.Middleware.Skip(Authorize, LoginHandler)
		app.Middleware.Skip(Authorize, ApiExampleHandler)

		auth := app.Group("/auth")
		bah := buffalo.WrapHandlerFunc(gothic.BeginAuthHandler)
		auth.GET("/{provider}", bah)
		auth.GET("/{provider}/callback", AuthCallback)
		auth.Middleware.Skip(Authorize, bah, AuthCallback)

		api := app.Group("/api")
		api.GET("/upload", GetUploadHandler)
		api.POST("/upload", PostUploadHandler)
		api.GET("/github", GithubHandler)
		api.GET("/twitter", TwitterHandler)
		api.GET("/scraping", ScrapingHandler)

		cr := &ContactsResource{&buffalo.BaseResource{}}
		app.Middleware.Skip(Authorize, cr.New)
		app.Middleware.Skip(Authorize, cr.Create)
		app.Resource("/contacts", cr)
		app.Resource("/widgets", WidgetsResource{})
	}

	return app
}
