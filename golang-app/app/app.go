package app

import (
	"example/employee-app/config"
	"example/employee-app/db"
	"example/employee-app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	db      *gorm.DB
	empRepo *repositories.GormEmployeeRepository
}

func Run(cfg *config.Config) {
	router := gin.Default()
	app := new(App)

	app.db = db.Connect(cfg)
	db.Migrate(app.db)

	app.empRepo = repositories.NewGormEmployeeRepository(app.db)

	app.UseCors(router)
	app.Routes(router)
	router.Run(cfg.Port)
}
