package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const redirect_url string = "http://127.0.0.1:3001"

type AuthResponse struct { // AuthResponse has other fields, but for now we only use one
	AccessToken string `json:"access_token"`
	// TokenType    string `json:"token_type"`
	// ExpiresIn    int    `json:"expires_in"`
	// RefreshToken string `json:"refresh_token"`
	// Scope        string `json:"scope"`
	// CreatedAt    int    `json:"created_at"`
}

func getAuthToken(CLIENT_ID string, CLIENT_SECRET string, code string) (string, error) {
	// Hit 42 api
	resp, err := http.PostForm("https://api.intra.42.fr/oauth/token",
		url.Values{
			"grant_type":    {"authorization_code"},
			"client_id":     {CLIENT_ID},
			"client_secret": {CLIENT_SECRET},
			"code":          {code},
			"redirect_uri":  {redirect_url},
			"state":         {"1234"}})

	if err != nil {
		return "", err
	}

	// Read response as JSON
	defer resp.Body.Close()

	authResp := new(AuthResponse)
	err = json.NewDecoder(resp.Body).Decode(authResp)

	if err != nil {
		return "", err
	}

	return authResp.AccessToken, nil
}

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

	app.Use(cors.New())// Or extend your config for customization
	// Default encrypted cookie middleware config
	app.Use(encryptcookie.New(encryptcookie.Config{ // this re-creates keys each time
		Key: encryptcookie.GenerateKey()})) // later we should use a random, but stable value

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML) // to display as html

		return c.SendString(`
			<h1>Hello, World! Backend speaking</h1>
			<a href="/login">42 login demo</a>`)
	})

	// Redirect to 42 api
	app.Get("/login", func(c *fiber.Ctx) error {
		url := fmt.Sprint("https://api.intra.42.fr/oauth/authorize?client_id=", CLIENT_ID, "&redirect_uri=", redirect_url, "&response_type=code&scope=public&state=1234")
		return c.Redirect(url)
	})

	// Get code and trade it for auth token
	app.Post("/auth", func(c *fiber.Ctx) error {
		// Capture 42 redirect with auth code in Body
		payload := struct {
			Code  string `json:"code"`
		}{}
	
		if err := c.BodyParser(&payload); err != nil {
			return err
		}
	
		// Exchange the code for an auth token
		authToken, err := getAuthToken(CLIENT_ID, CLIENT_SECRET, payload.Code)

		if err != nil {
			return c.SendString("42 api call failed")
		}

		// Save token in encrypted cookie
		c.Cookie(&fiber.Cookie{
			Name:  "42session",
			Value: authToken})

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML) // to display as html

		// Get 'me' from 42 api
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.intra.42.fr/v2/me", nil)

		if err != nil {
			return c.SendString("Couldn't create request")
		}

		req.Header.Add("Authorization", fmt.Sprint("Bearer ", authToken))

		resp, err := client.Do(req)

		if err != nil {
			return c.SendString("42 api request failed")
		}

		// Read response
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return c.SendString("Couldn't read 42 api response")
		}

		// Show the primitive json
		return c.SendString(string(body))
	})

	app.Listen(":3000")
}
