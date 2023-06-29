package app

import (
	"example/employee-app/repository"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	db   *gorm.DB
	repo *repository.EmployeRepository
}

func Run() {
	router := gin.Default()
	app := new(App)

	dsn := "root@tcp(127.0.0.1:3306)/employee-app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	app.db = db
	app.repo = repository.NewEmployeeRepository(db)
	if err := app.repo.Migrate(); err != nil {
		panic(err)
	}
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	app.Routes(router)

	router.Run("localhost:8080")
}
