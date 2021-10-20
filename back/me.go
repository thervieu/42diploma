package main

import (
	"encoding/json"
	"time"
	"strings"
	"gorm.io/gorm"
)

// The response of /me from 42 api
type MeResponse struct {
	ID              int               `json:"id"`
	Email           string            `json:"email"`
	Login           string            `json:"login"`
	FirstName       string            `json:"first_name"`
	LastName        string            `json:"last_name"`
	UsualFullName   string            `json:"usual_full_name"`
	UsualFirstName  interface{}       `json:"usual_first_name"`
	URL             string            `json:"url"`
	Phone           string            `json:"phone"`
	Displayname     string            `json:"displayname"`
	ImageURL        string            `json:"image_url"`
	Staff           bool              `json:"staff?"`
	CorrectionPoint int               `json:"correction_point"`
	PoolMonth       string            `json:"pool_month"`
	PoolYear        string            `json:"pool_year"`
	Location        interface{}       `json:"location"`
	Wallet          int               `json:"wallet"`
	AnonymizeDate   time.Time         `json:"anonymize_date"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	Alumni          bool              `json:"alumni"`
	IsLaunched      bool              `json:"is_launched?"`
	Groups          []interface{}     `json:"groups"`
	CursusUsers     []CursusUsers     `json:"cursus_users"`
	ProjectsUsers   []ProjectsUsers   `json:"projects_users"`
	LanguagesUsers  []LanguagesUsers  `json:"languages_users"`
	Achievements    []Achievements    `json:"achievements"`
	Titles          []Titles          `json:"titles"`
	TitlesUsers     []TitlesUsers     `json:"titles_users"`
	Partnerships    []interface{}     `json:"partnerships"`
	Patroned        []Patroned        `json:"patroned"`
	Patroning       []interface{}     `json:"patroning"`
	ExpertisesUsers []ExpertisesUsers `json:"expertises_users"`
	Roles           []interface{}     `json:"roles"`
	Campus          []Campus          `json:"campus"`
	CampusUsers     []CampusUsers     `json:"campus_users"`
}
type Skills struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Level float64 `json:"level"`
}
type User struct {
	ID              int         `json:"id"`
	Login           string      `json:"login"`
	Email           string      `json:"email"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	UsualFirstName  interface{} `json:"usual_first_name"`
	UsualFullName   string      `json:"usual_full_name"`
	Displayname     string      `json:"displayname"`
	Staff           bool        `json:"staff?"`
	CorrectionPoint int         `json:"correction_point"`
	PoolMonth       string      `json:"pool_month"`
	PoolYear        string      `json:"pool_year"`
	Location        interface{} `json:"location"`
	Wallet          int         `json:"wallet"`
	URL             string      `json:"url"`
	AnonymizeDate   time.Time   `json:"anonymize_date"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}
type Cursus struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
}
type CursusUsers struct {
	Grade        interface{} `json:"grade"`
	Level        float64     `json:"level"`
	Skills       []Skills    `json:"skills"`
	BlackholedAt interface{} `json:"blackholed_at"`
	ID           int         `json:"id"`
	BeginAt      time.Time   `json:"begin_at"`
	EndAt        time.Time   `json:"end_at"`
	CursusID     int         `json:"cursus_id"`
	HasCoalition bool        `json:"has_coalition"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	User         User        `json:"user"`
	Cursus       Cursus      `json:"cursus"`
}
type ProjectUsersProject struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ParentID int    `json:"parent_id"`
}
type ProjectsUsers struct {
	ID            int                 `json:"id"`
	Occurrence    int                 `json:"occurrence"`
	FinalMark     interface{}         `json:"final_mark"`
	Status        string              `json:"status"`
	Validated     interface{}         `json:"validated?"`
	CurrentTeamID int                 `json:"current_team_id"`
	Project       ProjectUsersProject `json:"project"`
	CursusIds     []int               `json:"cursus_ids"`
	MarkedAt      interface{}         `json:"marked_at"`
	Marked        bool                `json:"marked"`
	RetriableAt   interface{}         `json:"retriable_at"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
}
type LanguagesUsers struct {
	ID         int       `json:"id"`
	LanguageID int       `json:"language_id"`
	UserID     int       `json:"user_id"`
	Position   int       `json:"position"`
	CreatedAt  time.Time `json:"created_at"`
}
type Achievements struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Tier         string      `json:"tier"`
	Kind         string      `json:"kind"`
	Visible      bool        `json:"visible"`
	Image        string      `json:"image"`
	NbrOfSuccess interface{} `json:"nbr_of_success"`
	UsersURL     string      `json:"users_url"`
}
type Titles struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type TitlesUsers struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	TitleID   int       `json:"title_id"`
	Selected  bool      `json:"selected"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Patroned struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	GodfatherID int       `json:"godfather_id"`
	Ongoing     bool      `json:"ongoing"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type ExpertisesUsers struct {
	ID          int       `json:"id"`
	ExpertiseID int       `json:"expertise_id"`
	Interested  bool      `json:"interested"`
	Value       int       `json:"value"`
	ContactMe   bool      `json:"contact_me"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int       `json:"user_id"`
}
type Language struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Identifier string    `json:"identifier"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type Campus struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	TimeZone           string   `json:"time_zone"`
	Language           Language `json:"language"`
	UsersCount         int      `json:"users_count"`
	VogsphereID        int      `json:"vogsphere_id"`
	Country            string   `json:"country"`
	Address            string   `json:"address"`
	Zip                string   `json:"zip"`
	City               string   `json:"city"`
	Website            string   `json:"website"`
	Facebook           string   `json:"facebook"`
	Twitter            string   `json:"twitter"`
	Active             bool     `json:"active"`
	EmailExtension     string   `json:"email_extension"`
	DefaultHiddenPhone bool     `json:"default_hidden_phone"`
}
type CampusUsers struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	CampusID  int       `json:"campus_id"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectUserData struct {
	ProjectID   int 		`json:"project_id"`
	ID			int			`json:"id"`
	Slug        string		`json:"slug"`
	XP			int			`json:"xp"`
}
// The data that is necessary to us
type UserData struct {
	Login        string				`json:"login"`
	Level        float64			`json:"level"`
	Projects     []ProjectUserData	`json:"projects";gorm:"foreignkey:ProjectID"`
}

// projectsDone returns a slice with IDs of the projects done by user
func ProjectsDone(projUs []ProjectsUsers) []int {
	ret := make([]int, len(projUs))

	for i, pu := range projUs {
		ret[i] = pu.Project.ID
	}
	return ret
}

// ProjectsWXP returns a slice w/ Project not done by user and their xp
// determined by 
func ProjectsWXP(db *gorm.DB, pNotDone []Project) []ProjectUserData {
	var rtnProjects []ProjectUserData
	for o := range pNotDone {
		xp := 0
		first := true
		var allPS []ProjectSessions
		db.Model(&pNotDone[o]).Association("ProjectSessions").Find(&allPS)

		var last_updated time.Time
		for ps := 0 ; ps < len(allPS) ; ps++ {
				if first {
					last_updated = allPS[ps].UpdatedAt
					xp = allPS[ps].Difficulty
					first = false
				} else {
					if allPS[ps].UpdatedAt.After(last_updated) {
						last_updated = allPS[ps].UpdatedAt
						xp = allPS[ps].Difficulty
					}
				}
		}
		if xp != 0 {
			var pSave ProjectUserData
			pSave.ID = pNotDone[o].ID
			pSave.Slug = pNotDone[o].Slug
			pSave.XP = xp

			rtnProjects = append(rtnProjects, pSave)
		}
	}

	var finalProjects []ProjectUserData
	for o := range rtnProjects {
		if !strings.Contains(rtnProjects[o].Slug, "piscine") {
			finalProjects = append(finalProjects, rtnProjects[o])
		}
	}

	return finalProjects
}

// GetUserData returns the data we will need about the user, given their
// authToken
func GetUserData(db *gorm.DB, authToken string) (UserData, error) {
	// Ask 42 api for our data
	resp, err := apiGet(authToken, "https://api.intra.42.fr/v2/me")
	if err != nil {
		return UserData{}, err
	}

	// Parse this data
	var me MeResponse
	err = json.NewDecoder(resp.Body).Decode(&me)
	if err != nil {
		return UserData{}, err
	}

	projectsDone := ProjectsDone(me.ProjectsUsers)
	
	// find projects not in done
	var projectNotDone []Project
	db.Not(map[string]interface{}{"id": projectsDone}).Find(&projectNotDone)
	
	rtnProjects := ProjectsWXP(db, projectNotDone)

	return UserData{
		Login:    me.Login,
		Level:    me.CursusUsers[1].Level,
		Projects: rtnProjects,
	}, nil
}
