package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/userCourse"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/userModules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api"
	modules "github.com/Learning-Management-System-Kelompok-42/BE-LMS/app/module"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/certificate"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/faq"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/quiz"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/requestCourse"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	cfg := config.GetConfig()

	dbConnection := util.NewConnectionDB(cfg)

	controllers := modules.RegisterModules(dbConnection, cfg)

	timeoutContext := time.Duration(cfg.App.Timeout) * time.Second

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "LMS API")
	})

	api.RegistrationPath(e, controllers, cfg)

	// port := os.Getenv("PORT")

	go func() {
		address := fmt.Sprintf(":%d", cfg.App.Port)
		// if err := e.Start(address); err != nil {
		// 	log.Info("Shutting down the server")
		// }
		// run server with https with file localhost.crt and localhost.key
		if err := e.StartTLS(address, "localhost.crt", "localhost.key"); err != nil {
			log.Info("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// close the database connection
	defer dbConnection.Close()

	// timout 30 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Set auto migration db
	cfg := config.GetConfig()
	dbConnection := util.NewConnectionDB(cfg)

	if err := dbConnection.PostgreSQL.AutoMigrate(
		&company.Company{},
		&specialization.Specialization{},
		&users.User{},
		&course.Course{},
		&certificate.Certificate{},
		&faq.Faq{},
		&quiz.Quiz{},
		&module.Module{},
		&requestCourse.RequestCourse{},
		&specialization.SpecializationCourse{},
		&userCourse.UserCourse{},
		&userModules.UserModule{},
	); err != nil {
		panic(err)
	}
}
