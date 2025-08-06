package main

import (
	"fmt"

	"github.com/CosmeticsShiraz/Backend/bootstrap"
	"github.com/CosmeticsShiraz/Backend/internal/domain/entity"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/routes"
	"github.com/CosmeticsShiraz/Backend/wire"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()
	ginEngine := gin.New()

	config := bootstrap.Run()

	app, err := wire.InitializeApplication(config)
	if err != nil {
		panic(err)
	}

	app.Database.DB.GetDB().AutoMigrate(
		&entity.Address{},
		&entity.City{},
		&entity.Permission{},
		&entity.Province{},
		&entity.Role{},
		&entity.User{},
		&entity.Media{},
		&entity.News{},
		&entity.Like{},
		&entity.Product{},
		&entity.Category{},
		&entity.Brand{},
		&entity.Picture{},
	)

	app.Seeds.AddressSeeder.SeedProvincesAndCities()
	app.Seeds.RoleSeeder.SeedRoles()

	routes.Run(ginEngine, app)

	ginEngine.Run(fmt.Sprintf(":%v", config.Env.Server.Port))
}
