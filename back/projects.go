package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Project struct {
	ID              int               `json:"id"`
	Name            string            `json:"name"`
	Slug            string            `json:"slug"`
	Parent          Parent            `json:"parent"`
	Children        []interface{}     `json:"children"`
	Attachments     []interface{}     `json:"attachments"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	Exam            bool              `json:"exam"`
	GitID           int               `json:"git_id"`
	Repository      string            `json:"repository"`
	Recommendation  string            `json:"recommendation"`
	Cursus          []Cursus          `json:"cursus"`
	Campus          []Campus          `json:"campus"`
	Videos          []interface{}     `json:"videos"`
	ProjectSessions []ProjectSessions `json:"project_sessions"`
}
type Parent struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	URL  string `json:"url"`
}
type Cursus struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
}
type Language struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Identifier string    `json:"identifier"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type Campus struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	TimeZone           string      `json:"time_zone"`
	Language           Language    `json:"language"`
	UsersCount         int         `json:"users_count"`
	VogsphereID        interface{} `json:"vogsphere_id"`
	Country            string      `json:"country"`
	Address            string      `json:"address"`
	Zip                string      `json:"zip"`
	City               string      `json:"city"`
	Website            string      `json:"website"`
	Facebook           string      `json:"facebook"`
	Twitter            string      `json:"twitter"`
	Active             bool        `json:"active"`
	EmailExtension     string      `json:"email_extension"`
	DefaultHiddenPhone bool        `json:"default_hidden_phone"`
}
type Scales struct {
	ID               int  `json:"id"`
	CorrectionNumber int  `json:"correction_number"`
	IsPrimary        bool `json:"is_primary"`
}
type Uploads struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type ProjectSessions struct {
	ID               int           `json:"id"`
	Solo             bool          `json:"solo"`
	BeginAt          interface{}   `json:"begin_at"`
	EndAt            interface{}   `json:"end_at"`
	EstimateTime     string        `json:"estimate_time"`
	Difficulty       int           `json:"difficulty"`
	Objectives       []interface{} `json:"objectives"`
	Description      string        `json:"description"`
	DurationDays     interface{}   `json:"duration_days"`
	TerminatingAfter int           `json:"terminating_after"`
	ProjectID        int           `json:"project_id"`
	CampusID         interface{}   `json:"campus_id"`
	CursusID         interface{}   `json:"cursus_id"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	MaxPeople        interface{}   `json:"max_people"`
	IsSubscriptable  bool          `json:"is_subscriptable"`
	Scales           []Scales      `json:"scales"`
	Uploads          []Uploads     `json:"uploads"`
	TeamBehaviour    string        `json:"team_behaviour"`
	Commit           string        `json:"commit"`
}

// ApiGet makes a GET request to a given url with a given token and returns
// the response
func apiGet(authToken string, url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req.Header.Add("Authorization", "Bearer "+authToken)
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
func getProjects(authToken string) ([]Project, error) {

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
	return allProjects, nil
}

// func main() {
// 	// baseUrl := "https://api.intra.42.fr/v2/cursus/21/projects"
// 	token := "62b7a7138e019c1638d52c1f18beb41280e618140b632937f92a699985b9aa1b"
// 	// resp, _ := getPage(token, baseUrl, 1)
// 	// fmt.Println(resp)
// 	// fmt.Println(resp.Header["X-Total"])

// 	projects, err := getProjects(token)

// 	if err != nil {
// 		fmt.Println("Error! ", err)
// 	}
// 	// for _, project := range projects {
// 	// 	// fmt.Println(project.ID, project.Name)
// 	// }
// 	file, _ := json.MarshalIndent(projects, "", " ")

// 	_ = ioutil.WriteFile("test.json", file, 0644)

// }
