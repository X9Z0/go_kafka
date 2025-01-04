package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v3"
)

type Comment struct {
	Text string `json:"text" form:"text"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comments", createComment)
	app.Listen(":3000")
}

func PushCommentToQueue(topic string, message []byte) error {
	bokersUrl := 
}

func createComment(c *fiber.Ctx) error {
	cmt := new(Comment)
	if err := c.BodyParser(cmt); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err

		cmtInBytes, err := json.Marshal(cmt)
		ushCommentToQueue("comments", cmtInBytes)
		if err != nil {
			log.Println(err)
		}
		err = c.JSON(&fiber.Map{
			"success": true,
			"message": "comment pushed successfully",
			"comment": cmt,
		})

		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": "Error creating product",
			})
			return err
		}
	}

}
