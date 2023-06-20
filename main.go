package main

import (
	"GolangApiFiber/database"
	"GolangApiFiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
	"log"
)

var pplTable = ksql.NewTable("peoples", "id")

func main() {
	ctx := context.Background()

	db, err := kpgx.New(ctx, dbURL, ksql.Config{})
	if err != nil {
		log.Fatalf("unable to start database: %s", err)
	}

	app := fiber.New()
	database.ConnectDb()
	log.Println("Aplicação iniciada")
	routes.SetupRoutes(app)
	log.Println("Rotas configuradas")
	app.Listen(":3000")
}
