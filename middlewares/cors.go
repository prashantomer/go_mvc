package middlewares

import "github.com/gofiber/fiber/v2/middleware/cors"

  func SetupCORS()  {
  {
    AllowOrigins: "https://gofiber.io, https://gofiber.net",
    AllowHeaders: "Origin, Content-Type, Accept"
  }
}
