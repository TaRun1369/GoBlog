package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/TaRun1369/GoBlog/database"
	"github.com/TaRun1369/GoBlog/models"
	"github.com/TaRun1369/GoBlog/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)

}

func Register(c *fiber.Ctx) error { // fiber.Ctx is used to get the context of the request
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
	err := database.DB.Create(&user)            // Create is used to create the record in the database
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{ // fiber.Map is used to create a map
		"user":    user, // user is the user struct
		"message": "Account Created Successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string                  // map is a key value pair
	if err := c.BodyParser(&data); err != nil { // BodyParser is used to parse the data from the body
		// data parsing is needed to get the data from the body as data is in the form of json
		fmt.Println("Unable to parse json or body")
	}
	var user models.User

	// check if email already exists
	database.DB.Where("email=?", strings.TrimSpace(data["email"])).First(&user) // First is used to get the first record from the database
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email does not exists, kindly create/register your account",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id))) //Itoa is used to convert the int to string
	//strconv is used to convert the data types in golang from one to another it is not only for string to int it can do many more things such as int to string, string to int, int to float etc
	// GenerateJwt is used to generate the jwt token
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	// c.Cookie is used to set the cookie
	// cookie is used to store the data in the browser
	cookie := fiber.Cookie{
		Name:  "jwt", // name of the cookie
		Value: token, // value of the cookie as token because we are storing the token in the cookie
		// use of token in cookie is that we can access the token from the frontend
		// and we can use the token for authentication
		Expires:  time.Now().Add(time.Hour * 24), // expires after 24 hours
		HTTPOnly: true,                           // http only is used to make the cookie accessible only by the http protocol
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Login Successful",
		"user":    user,
	})
}

type Claims struct {
	
// here we create struct for payload 
	jwt.StandardClaims
}
// type is used to create a new datatype 
// Claims is the type of the struct