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
	"github.com/niranjan92/go-hackathon-starter/apiExamples"
	"github.com/niranjan92/go-hackathon-starter/models"
	"github.com/niranjan92/go-hackathon-starter/render"
	// used for performance profiling
	// _ "net/http/pprof"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

// var BASE_URL = envy.Get("BASE_URL", "http://localhost:3000")
var app *buffalo.App

// T is a translator used for internationalization
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

			// Uncomment lines below to enable pprof debugging
			// go func() {
			// 	log.Println(http.ListenAndServe(":8080", nil))
			// }()
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
		app.ServeFiles("/assets", render.AssetsBox)

		app.Use(SetCurrentUser)
		app.Use(Authorize)

		configureRoutes(app)

	}

	return app
}

func configureRoutes(app *buffalo.App) {
	app.GET("/", HomeHandler)
	app.GET("/login", LoginHandler)
	app.GET("/api-examples", apiExamples.APIExampleHandler)
	app.Middleware.Skip(Authorize, HomeHandler, LoginHandler, apiExamples.APIExampleHandler)

	cr := &ContactsResource{&buffalo.BaseResource{}}
	app.Middleware.Skip(Authorize, cr.New)
	app.Middleware.Skip(Authorize, cr.Create)
	app.Resource("/contacts", cr)
	app.Resource("/widgets", WidgetsResource{}) //TODO: remove widget and only use contact

	// configure login routes
	auth := app.Group("/auth")
	bah := buffalo.WrapHandlerFunc(gothic.BeginAuthHandler)
	auth.GET("/{provider}", bah)
	auth.GET("/{provider}/callback", AuthCallback)
	auth.Middleware.Skip(Authorize, bah, AuthCallback)

	// account management routes
	app.DELETE("/users/", UserDestroy)
	app.GET("/logout", AuthDestroy)
	app.GET("/account/profile", ProfileHandler)
	app.POST("/account/profile", UpdateProfileHandler)

	// routes for api examples
	api := app.Group("/api")
	api.GET("/upload", apiExamples.GetUploadHandler)
	api.POST("/upload", apiExamples.PostUploadHandler)
	api.GET("/github", apiExamples.GithubHandler)
	api.GET("/twitter", apiExamples.TwitterHandler)
	api.GET("/scraping", apiExamples.ScrapingHandler)
	api.Middleware.Skip(Authorize, apiExamples.GetUploadHandler, apiExamples.PostUploadHandler, apiExamples.GithubHandler,
		apiExamples.TwitterHandler, apiExamples.ScrapingHandler)

	app.GET("/react", ReactHandler) //TODO:

}
