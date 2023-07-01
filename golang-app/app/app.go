package app

import (
	"example/employee-app/config"
	"example/employee-app/db"
	"example/employee-app/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	db   *gorm.DB
	repo *repository.EmployeRepository
}

func Run(cfg *config.Config) {
	router := gin.Default()
	app := new(App)

	app.db = db.Connect(cfg)
	app.repo = repository.NewEmployeeRepository(app.db)
	if err := app.repo.Migrate(); err != nil {
		panic(err)
	}

	app.UseCors(router)
	app.Routes(router)
	router.Run(cfg.Port)
}
