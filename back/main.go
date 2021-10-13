package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
)

const redirect_url string = "http://127.0.0.1:3000/auth"
const success_redirect_url string = "http://127.0.0.1:3000/success"

func main() {
	CLIENT_ID := os.Getenv("CLIENT_ID")
	if CLIENT_ID == "" {
		log.Fatal("Please set the CLIENT_ID env variable to your 42 API client ID")
	}

	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")
	if CLIENT_SECRET == "" {
		log.Fatal("Please set the CLIENT_SECRET env variable to your 42 API client ID")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML) // to display as html
		return c.SendString(`
			<h1>Hello, World!</h1>
			<a href="/login">42 login demo</a>`)
	})

	// Redirect to 42 api
	app.Get("/login", func(c *fiber.Ctx) error {
		url := fmt.Sprint("https://api.intra.42.fr/oauth/authorize?client_id=", CLIENT_ID, "&redirect_uri=", redirect_url, "&response_type=code&scope=public&state=1234")
		return c.Redirect(url)
	})

	// Get code and trade it for auth token
	app.Get("/auth", func(c *fiber.Ctx) error {
		code := c.Query("code")
		if code == "" {
			return c.SendString("Please get back with the code")
		}

		resp, err := http.PostForm("https://api.intra.42.fr/oauth/token",
			url.Values{"grant_type": {"authorization_code"}, "client_id": {CLIENT_ID}, "client_secret": {CLIENT_SECRET}, "code": {code}, "redirect_uri": {success_redirect_url}, "state": {"1234"}})

		if err != nil {
			fmt.Print("err: ", err)
			c.SendString("42 api call failed")
		}
		fmt.Print(resp)

		return c.SendString(fmt.Sprint("Your code is: ", code))
	})

	// A route for success
	app.Get("/success", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML) // to display as html
		return c.SendString("Yay!!")
	})

	app.Listen(":3000")
}
