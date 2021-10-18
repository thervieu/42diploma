package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type Project struct {
	ID              int               `json:"id"`
	Name            string            `json:"name"`
	Slug            string            `json:"slug"`
}

func InitialMigration() {
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&Project{})
}

func GetProjects(c *fiber.Ctx) error {
	var projects []Project
	DB.Find(&projects)
	return c.JSON(projects)
}

func SaveProjects(projects []Project) []Project {
	DB.Create(&projects)
	return projects
}

// ApiGet makes a GET request to a given url with a given token and returns
// the response
func apiGet(authToken string, url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req.Header.Add("Authorization", "Bearer " + authToken)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getPage(authToken string, baseUrl string, page int) (*http.Response, error) {
	return apiGet(authToken, baseUrl+"?page="+strconv.Itoa(page))
}

// Get all projects from 42 api
func initProjects(authToken string) ([]Project, error) {

	baseUrl := "https://api.intra.42.fr/v2/cursus/21/projects"

	var allProjects []Project
	var p []Project

	i := 1

	// Temporary value until we get the first response with the actual value
	var totalElements int64 = 1

	// Get new pages while we don't reach the total number of projects
	for len(allProjects) < int(totalElements) {

		fmt.Println(len(allProjects))

		// Get a page
		resp, err := getPage(authToken, baseUrl, i)
		if err != nil {
			return nil, err
		}
		i++

		// Get the number of total pages from the page
		totalElements, err = strconv.ParseInt(resp.Header["X-Total"][0], 10, 32)
		if err != nil {
			return nil, err
		}

		// Add the new projects to our list
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			return nil, err
		}

		allProjects = append(allProjects, p...)
	}

	allKeys := make(map[Project]bool)
	list := []Project{}
	for _, item := range allProjects {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	fmt.Println(allProjects)
	return SaveProjects(list), nil
}
