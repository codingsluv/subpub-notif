package controller

import (
	"fmt"
	"log"

	"github.com/codingsluv/pubsub-go/configs"
	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	mssg := fmt.Sprintf("User denga ip %v mengakses website", c.IP())
	err := configs.RDS.Publish(
		configs.RDS_CTX,
		configs.REDIS_CHANNEL_NOTIFICATION,
		mssg,
	).Err()

	if err != nil {
		log.Println("Error :", err)
	}

	return c.Render("index", fiber.Map{})
}
