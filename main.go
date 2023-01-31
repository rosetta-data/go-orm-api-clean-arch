package main

import (
	"log"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
	"github.com/juliocesarscheidt/go-orm-api/infra/repository"
	"github.com/juliocesarscheidt/go-orm-api/infra/router"
	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
)

func main() {
	connectionString := utils.GetDbConnectionString()
	db, _ := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		utils.Logger.Errorf("Err %v", err)
	}

	// create repositories
	userRepository := repository.UserRepository{Db: db}
	// create router and its routes
	r := router.GetRouter()
	router.InjectRoutes(r, userRepository)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	utils.Logger.Info("Server listening on 0.0.0.0:8000")
	log.Fatal(srv.ListenAndServe())
}
