package main

import (
	"fmt"
	"go_mvc/controllers"
	"go_mvc/initializers"
	"go_mvc/middlewares"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
  initializers.SyncDB()
}

func main() {
	fmt.Println("Booting Go_MVC")
  // load Templates
  engine := html.New("./views", ".tmpl")

	// Setup Global `app` variable
	app := fiber.New(fiber.Config{
    Views: engine,
  })

  // Configure app
  app.Static("/", "./public")

  // CORS
  // app.Use(cors.New(cors.Config(middlewares.SetupCORS)))

	// Routes
	app.Get("/", controllers.PostIndex)

	// Listen Server
	app.Listen(":" + os.Getenv("PORT"))
}
