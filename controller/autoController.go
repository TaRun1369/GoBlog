package controller

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/TaRun1369/GoBlog/database"
	"github.com/TaRun1369/GoBlog/models"
	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)

}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{} // map is a key value pair
	// interface ki jaghah string bhi le sakta agar saare attributes string hote yaha password is in  byte so we use interface
	var userData models.User
	if err := c.BodyParser(&data); err != nil { // BodyParser is used to parse the data from the body
		fmt.Println("Unable to parse json or body")
	}

	// check if password is lesser than 6 characters
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwrod must be greater than 6 characters",
		})
	}
	// Trimspace is used to removes the spaces from the string
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			// c.JSON is used to send the json response
			// fiber.Map is used to create a map
			"message": "Invalid Email",
		})
	}

	// check if email already exists
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData) // First is used to get the first record from the database
	// Where is inbuilt function in gorm for querying the database for finding the record
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exists",
		})
	}
	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string)) // SetPassword is used to set the password
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":user,
		"message":"Account Created Successfully",
	})
}
