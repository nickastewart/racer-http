package main

import (
	"database/sql"
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	_ "modernc.org/sqlite"
	"net/http"
	"racer_http/controllers"
	"racer_http/repository"
	"racer_http/sqlite/entities"
	"racer_http/templates"
)

func main() {
	// TODO: Add testing to parser
	// TODO: Start using Templ

	db, err := sql.Open("sqlite", "/Users/nickstewart/sqlite/racer.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	queries := entities.New(db)
	var userRepository repository.UserRepository = repository.NewUserRepository(queries)
	//var locationRepository repository.LocationRepository = repository.NewLocationRepository(queries)
	//var eventRepository repository.EventRepository = repository.NewEventRepository(queries)
	//var eventResultRepository repository.EventResultRepository = repository.NewEventResultRepository(queries)

	authController := controllers.NewAuthController(userRepository)
	//	fileUploadController := controllers.NewFileUploadController(userRepository, eventRepository, locationRepository, eventResultRepository)
	//	eventController := controllers.NewEventsController(userRepository, eventRepository, locationRepository, eventResultRepository)

	if err != nil {
		log.Panic(err)
	}

	router := gin.Default()

	router.POST("/signup", authController.Signup)
	router.POST("/login", authController.LoginForm)

	// TODO: Add add friend endpoint, need to check auth

	//router.POST("friend", authCOntroller.checkAuth, friendController.AddFriend)

	// TODO: Add get events endpoint that includes friends, need to check auth
	// TODO: Add endpoint to remove friends

	router.HTMLRender = &TemplRender{}
	router.GET("/", authController.CheckAccessToken, func(c *gin.Context) {
		c.HTML(http.StatusOK, "", templates.Home())
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", templates.Login())
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", templates.Signup())
	})

	router.Run()
}
