package controller

import (
	// "errors"
	"fmt"
	"math"
	"strconv"

	"github.com/TaRun1369/GoBlog/database"
	"github.com/TaRun1369/GoBlog/models"
	"github.com/TaRun1369/GoBlog/util"
	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil { // BodyParser is used to parse the data from the body
		fmt.Println("Unable to parse json or body")
	}
	if err := database.DB.Create(&blogpost).Error; err != nil {
		// Create is used to create the record in the database // error such as email ki jaghah emails dala postman ke body mein toh wo error waise error ke case mein ye
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Congrats!, Your post is liveeeeeee",
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1")) // Query is used to get the query from the url query is like ?page=1 so we are getting the page from the url
	limit := 5                                    // 5 blog he lenge blog ki list mein se
	offset := (page - 1) * limit                  // offset is used to skip the records
	// offset meaning agar page 1 hai toh 0 se 5 tak records lenge agar page 2 hai toh 5 se 10 tak records lenge
	var total int64           // total is used to get the total number of records
	var getblog []models.Blog // getblog is used to get the blog from the database
	// list of blogs
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog) // Preload is used to get the user details from the user table
	database.DB.Model(&models.Blog{}).Count(&total)                        // Model is used to get the model from the database
	return c.JSON(fiber.Map{
		"data": getblog, //
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

// func DetailPost(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id")) // Params is used to get the params from the url
// 	var blogpost models.Blog
// 	database.DB.Where("id=?", id).Preload("User").First(&blogpost) // Preload is inbuilt function in gorm to get the user details from the user table
// 	return c.JSON(fiber.Map{
// 		"data": blogpost,
// 	})
// }

func DetailPost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id")) // Params is used to get the params from the url
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var blogpost models.Blog
	result := database.DB.Where("id=?", id).Preload("User").First(&blogpost) // Preload is inbuilt function in gorm to get the user details from the user table
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}
	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body or json")
	}

	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"message": "post updated successfully",
	})
}

// func UniquePost(c *fiber.Ctx) error{
// 	cookie:=c.Cookies("jwt")
// 	id,_:=util.ParseJwt(cookie)
// 	var blog models.Blog
// 	database.DB.Model(&blog).Where("userid=?",id).Preload("User").First(&blog)
// 	return c.JSON(blog)
// }

// ye work nhi karta pata nhi kyuuuuuuuuuu
func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, err := util.ParseJwt(cookie)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid JWT token",
		})
	}

	var blog models.Blog
	result := database.DB.Model(&blog).Where("userid=?", id).Preload("User").First(&blog)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	return c.JSON(blog)
}

// func DeletePost(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))
// 	blog := models.Blog{
// 		Id: uint(id),
// 	}
// 	deleteQuery := database.DB.Delete(&blog)
// 	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Post not found",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Post deleted successfully",
// 	})

// }

func DeletePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	blog := models.Blog{
		Id: uint(id),
	}
	deleteQuery := database.DB.Delete(&blog)
	if deleteQuery.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}
