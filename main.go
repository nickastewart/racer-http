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
	var locationRepository repository.LocationRepository = repository.NewLocationRepository(queries)
	var eventRepository repository.EventRepository = repository.NewEventRepository(queries)
	var eventResultRepository repository.EventResultRepository = repository.NewEventResultRepository(queries)

	authController := controllers.NewAuthController(userRepository)
	fileUploadController := controllers.NewFileUploadController(userRepository, eventRepository, locationRepository, eventResultRepository)
	eventController := controllers.NewEventsController(userRepository, eventRepository, locationRepository, eventResultRepository)

	if err != nil {
		log.Panic(err)
	}

	router := gin.Default()

	router.POST("auth/user", authController.CreateUser)
	router.GET("auth/user", authController.Login)
	router.POST("/login", authController.LoginForm)

	// TODO: delete this when a get profile functionlity is implemented
	router.GET("user", authController.CheckAuth, authController.GetUser)

	router.POST("upload", authController.CheckAuth, fileUploadController.UploadFile)
	router.GET("events", authController.CheckAuth, eventController.GetEventsByUser)

	// TODO: Add add friend endpoint, need to check auth

	//router.POST("friend", authCOntroller.checkAuth, friendController.AddFriend)

	// TODO: Add get events endpoint that includes friends, need to check auth
	// TODO: Add endpoint to remove friends

	router.HTMLRender = &TemplRender{}
	router.GET("/", authController.CheckAccessToken, func(c *gin.Context) {
		c.HTML(http.StatusOK, "Home Page", templates.Home())
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Home Page", templates.Login())
	})
	router.Run()
}
