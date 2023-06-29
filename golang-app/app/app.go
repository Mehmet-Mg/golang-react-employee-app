package app

import (
	"example/employee-app/config"
	"example/employee-app/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	db   *gorm.DB
	repo *repository.EmployeRepository
}

func Run(cfg *config.Config) {
	router := gin.Default()
	app := new(App)

	db, err := gorm.Open(mysql.Open(cfg.ConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	app.db = db
	app.repo = repository.NewEmployeeRepository(db)
	if err := app.repo.Migrate(); err != nil {
		panic(err)
	}

	app.UseCors(router)
	app.Routes(router)
	router.Run(cfg.Port)
}
