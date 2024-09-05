package main

import (
	"fmt"

	"github.com/codingsluv/pubsub-go/configs"
	"github.com/codingsluv/pubsub-go/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	configs.InitRedisClient()

	err := configs.RDS.Ping(configs.RDS_CTX).Err()
	if err != nil {
		fmt.Println("gagal connect redis")
		panic(err)
	} else {
		fmt.Println("sukses connect redis")
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", controller.HomePage)
	app.Get("/notifikasi", controller.Notification)
	app.Listen(":8080")
}
