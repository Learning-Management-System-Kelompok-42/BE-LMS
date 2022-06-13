package main

import (
	"context"
	"fmt"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specializationCourse"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/userCourse"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/userModules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api"
	modules "github.com/Learning-Management-System-Kelompok-42/BE-LMS/app/module"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/certificate"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/faq"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/material"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/module"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/options"
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

	controllers := modules.RegisterModules(dbConnection)

	timeoutContext := time.Duration(cfg.App.Timeout) * time.Second

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "LMS API")
	})

	api.RegistrationPath(e, controllers)

	go func() {
		address := fmt.Sprintf(":%d", cfg.App.Port)
		if err := e.Start(address); err != nil {
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
	fmt.Println("jalan")
	// Set auto migration db
	cfg := config.GetConfig()
	dbConnection := util.NewConnectionDB(cfg)

	if err := dbConnection.PostgreSQL.AutoMigrate(
		&specialization.Specialization{},
		&company.Company{},
		&users.User{},
		&course.Course{},
		&certificate.Certificate{},
		&faq.Faq{},
		&material.Material{},
		&module.Module{},
		&quiz.Quiz{},
		&options.Option{},
		&requestCourse.RequestCourse{},
		&specializationCourse.SpecializationCourse{},
		&userCourse.UserCourse{},
		&userModules.UserModule{},
	); err != nil {
		panic(err)
	}
}
