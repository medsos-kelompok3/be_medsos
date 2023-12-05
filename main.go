package main

import (
	"be_medsos/config"
	ph "be_medsos/features/posting/handler"
	pr "be_medsos/features/posting/repository"
	ps "be_medsos/features/posting/service"
	uh "be_medsos/features/user/handler"
	ur "be_medsos/features/user/repository"
	us "be_medsos/features/user/service"
	ek "be_medsos/helper/enkrip"
	"be_medsos/routes"
	cld "be_medsos/utils/cld"
	"be_medsos/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	if cfg == nil {
		e.Logger.Fatal("tidak bisa start server kesalahan database")
	}
	cld, ctx, param := cld.InitCloudnr(*cfg)

	db, err := database.InitMySql(*cfg)
	if err != nil {
		e.Logger.Fatal("tidak bisa start bro", err.Error())
	}

	db.AutoMigrate(&ur.UserModel{}, &pr.PostingModel{})

	ekrip := ek.New()
	userRepo := ur.New(db)
	userService := us.New(userRepo, ekrip)
	userHandler := uh.New(userService, cld, ctx, param)

	postingRepo := pr.New(db)
	postingService := ps.New(postingRepo)
	postingHandler := ph.New(postingService, cld, ctx, param)

	routes.InitRoute(e, userHandler, postingHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
