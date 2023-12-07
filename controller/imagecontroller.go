package controller

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz") // rune is used to store the unicode characters
// rune is similar to byte but rune is used to store the unicode characters

func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["image"]
	fileName := ""
	for _, file := range files {
		fileName = randLetter(5) + "-" + file.Filename
		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return err
		}
	}

	fmt.Println(fileName)
	return c.JSON(fiber.Map{
		"url": "http://https://go-blog-api.onrender.com/api/uploads/" + fileName,
	})
}
