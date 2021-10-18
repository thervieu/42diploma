package project

import (
	"fmt"	
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)
var DB *gorm.DB
var err error

const DNS = "root:admin@tcp(127.0.0.1:3000)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type Project struct {
	gorm.Model
	Id string 'json:"id"'
	Name string 'json:"name"'
	Slug string 'json:"slug"'
	Xp string 'json:"xp"'
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&Project{})
}

func GetProjects(c *fiber.Ctx) error {
	var projects []Project
	DB.Find(&projects)
	return c.JSON(&projects)
}

func SaveProjects(c *fiber.Ctx) error {
	projects := new([]Projects)
	if err := c.BodyParser(projects); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Create(&projects)
	return c.JSON(&projects)
}
