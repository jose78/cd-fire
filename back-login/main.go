package main

import (
	"fmt"
	"log"
	"os"

	//"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/shareed2k/goth_fiber"
)

var (
	clientID     = os.Getenv("fire_clientId")
	clientSecret = os.Getenv("fire_clientSecret")
	domain       = os.Getenv("fire_frontUrl")
)

func main() {
	app := fiber.New()

	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	// Optionally, you can override the session store here:
	goth_fiber.SessionStore = session.New(session.Config{
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:_gothic_session",
		CookieDomain:   "",
		CookiePath:     "",
		CookieSecure:   false,
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
		KeyGenerator:   utils.UUIDv4,
	})
	goth.UseProviders(
		github.New(clientID, clientSecret, "http://localhost:8080/auth/callback/github"),
	)
	// e36f57adcb81d0361b69e8150b48a5689ec955de
	app.Get("/login/:provider", goth_fiber.BeginAuthHandler)
	app.Get("/auth/callback/:provider", func(ctx *fiber.Ctx) error {
		fmt.Println(ctx)
		user, err := goth_fiber.CompleteUserAuth(ctx)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.SendString(user.Email)
	})
	app.Get("/logout", func(ctx *fiber.Ctx) error {
		if err := goth_fiber.Logout(ctx); err != nil {
			log.Fatal(err)
		}

		return ctx.SendString("logout")
	})

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
