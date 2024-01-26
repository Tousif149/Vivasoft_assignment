package containers

import (
	"book-crud/pkg/config"
	"book-crud/pkg/connection"
	"book-crud/pkg/controllers"
	"book-crud/pkg/repositories"
	"book-crud/pkg/routes"
	"book-crud/pkg/services"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {

	//config initialization
	config.SetConfig()

	//database initializations
	db := connection.GetDB()

	// repository initialization
	bookRepo := repositories.BookDBInstance(db)
	authorRepo := repositories.AuthorDBInstance(db)
	userRepo := repositories.UserDBInstance(db)

	//service initialization
	bookService := services.BookServiceInstance(bookRepo)
	authorService := services.AuthorServiceInstance(authorRepo)
	userService := services.AuthServiceInstance(userRepo)
	//controller initialization
	bookCtr := controllers.NewBookController(bookService)
	authorCtr := controllers.NewAuthorController(authorService)
	userCtr := controllers.NewAuthController(userService)
	//route initialization
	b := routes.BookRoutes(e, bookCtr)
	a := routes.AuthorRoutes(e, authorCtr)
	r := routes.AuthRoutes(e, userCtr)

	b.InitBookRoute()
	a.InitAuthorRoute()
	r.InitAuthRoutes()
	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))

}
