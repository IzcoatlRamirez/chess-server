package main

import (
	// "chess-server/database"
	// "chess-server/routes"
	"chess-server/treebasedmodel"
	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
// 	database.Init()
// 	app := fiber.New()
//   //por ahora puede recibir de cualquier sitio
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins:     "*",
// 		AllowMethods:     "GET,POST,DELETE",
// 		AllowCredentials: false,
// 	}))

	// routes.UserRoutes(app)
	// routes.MatchRoutes(app)

	// app.Listen(":3000")



	treebasedmodel.GetBestMove()

   
}
