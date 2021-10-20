package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

const redirect_url string = "http://127.0.0.1:3001"
const demo_redirect_url string = "http://127.0.0.1:3000/demoAuth"

type AuthResponse struct { // AuthResponse has other fields, but for now we only use one
	AccessToken string `json:"access_token"`
	// TokenType    string `json:"token_type"`
	// ExpiresIn    int    `json:"expires_in"`
	// RefreshToken string `json:"refresh_token"`
	// Scope        string `json:"scope"`
	// CreatedAt    int    `json:"created_at"`
}

// getAuthToken exchanges a 42 redirect code for a 42 API auth token
func getAuthToken(CLIENT_ID string, CLIENT_SECRET string, code string, redirect_url string) (string, error) {
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

	fmt.Println(resp)

	// Read response as JSON
	defer resp.Body.Close()

	authResp := new(AuthResponse)
	err = json.NewDecoder(resp.Body).Decode(authResp)

	if err != nil {
		return "", err
	}

	return authResp.AccessToken, nil
}

// getAuthToken exchanges a 42 redirect code for a 42 API auth token
func getAuthTokenServer(CLIENT_ID string, CLIENT_SECRET string) (string, error) {
	// Hit 42 api
	resp, err := http.PostForm("https://api.intra.42.fr/oauth/token",
		url.Values{
			"grant_type":    {"client_credentials"},
			"client_id":     {CLIENT_ID},
			"client_secret": {CLIENT_SECRET}})

	if err != nil {
		return "", err
	}

	fmt.Println(resp)

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

	db := InitialMigration()

	app := fiber.New()
	app.Use(cors.New()) // Or extend your config for customization
	// Default encrypted cookie middleware config
	app.Use(encryptcookie.New(encryptcookie.Config{ // this re-creates keys each time
		Key: encryptcookie.GenerateKey()})) // later we should use a random, but stable value
	token, err := getAuthTokenServer(CLIENT_ID, CLIENT_SECRET)

	if err != nil {
		return
	}

	r := db.Find(&Project{})
	if r.RowsAffected <= 0 { // If we don't have any projects loaded, load them
		log.Println("No projects present in the database, downloading")
		initProjects(token)
	} else {
		log.Println("Projects already present in the database, skipping init")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML) // to display as html

		return c.SendString(`
			<h1>Hello, World! Backend speaking</h1>
			<a href="/login">42 login demo</a>`)
	})

	// Redirect to 42 api
	app.Get("/login", func(c *fiber.Ctx) error {
		url := fmt.Sprint("https://api.intra.42.fr/oauth/authorize?client_id=", CLIENT_ID, "&redirect_uri=", demo_redirect_url, "&response_type=code&scope=public&state=1234")
		return c.Redirect(url)
	})

	// Demo backend-only auth for easy testing
	app.Get("/demoAuth", func(c *fiber.Ctx) error {
		// Capture 42 redirect with auth code in Body
		code := c.Query("code")
		fmt.Println(code)

		// Exchange the code for an auth token
		authToken, err := getAuthToken(CLIENT_ID, CLIENT_SECRET, code, demo_redirect_url)

		fmt.Println("Auth token: " + authToken)

		if err != nil {
			return c.SendString("42 api call failed")
		}

		// Save token in encrypted cookie
		c.Cookie(&fiber.Cookie{
			Name:  "42session",
			Value: authToken})

		// c.Set(fiber.HeaderContentType, fiber.MIMETextHTML) // to display as html
		return c.SendString(fmt.Sprint("Your token is: ", authToken))
	})

	// Get code and trade it for auth token
	app.Post("/auth", func(c *fiber.Ctx) error {
		// Capture 42 redirect with auth code in Body
		payload := struct {
			Code string `json:"code"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		// Exchange the code for an auth token
		authToken, err := getAuthToken(CLIENT_ID, CLIENT_SECRET, payload.Code, redirect_url)

		if err != nil {
			return c.SendString("42 api call failed")
		}

		myData, err := GetUserData(db, authToken)
		if err != nil {
			c.SendString(fmt.Sprint(err))
		}

		rtn, err := json.Marshal(myData)
		if err != nil {
			return c.SendString(fmt.Sprint(err))
		}
		return c.SendString(string(rtn))
	})

	app.Get("/projects", func(c *fiber.Ctx) error {
		projects := GetProjects(c)

		return projects
	})

	app.Get("/testUserData", func(c *fiber.Ctx) error {

		token := c.Cookies("42session")

		if token == "" {
			return c.SendString("Please first get a token through demo auth")
		}

		myData, err := GetUserData(db, token)
		if err != nil {
			c.SendString(fmt.Sprint(err))
		}
		
		rtn, err := json.Marshal(&myData)
		if err != nil {
			return c.SendString(fmt.Sprint(err))
		}
		fmt.Println(fmt.Sprint(rtn))
		return c.SendString(fmt.Sprint(rtn))
	})

	app.Listen(":3000")
}
