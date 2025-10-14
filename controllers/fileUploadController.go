package controllers

import (
	"fmt"
	"racer_http/repository"
	"racer_http/sqlite/entities"

	"context"
	"github.com/gin-gonic/gin"
	racer "github.com/nickastewart/racer-parser"
	"github.com/nickastewart/racer-parser/model"
	"net/http"
)

type FileUploadController struct {
	UserRepository     repository.UserRepository
	LocationRepository repository.LocationRepository
	EventRepository    repository.EventRepository
}

func NewFileUploadController(userRepository repository.UserRepository,
	eventRepository repository.EventRepository,
	locationRepository repository.LocationRepository) *FileUploadController {
	return &FileUploadController{
		UserRepository:     userRepository,
		EventRepository:    eventRepository,
		LocationRepository: locationRepository,
	}
}

func (controller *FileUploadController) UploadFile(c *gin.Context) {
	ctx := context.Background()
	u, exists := c.Get("currentUser")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not authenticated"})
		return
	}

	user := u.(entities.GetUserByIdRow)

	form, err := c.MultipartForm()
	multipartFile := form.File["file"]

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Delete temp logging
	fmt.Printf("User %s, uploaded file %s \n", user.FirstName, multipartFile[0].Filename)

	file, err := multipartFile[0].Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	event := racer.ParseFile(file)

	// TODO: Save event to the DB

	locationEntity, err := controller.processLocation(ctx, &event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventEntity, err := controller.processEvent(ctx, &event, &locationEntity)
	fmt.Println(eventEntity)
	fmt.Println(locationEntity)

}

func (controller *FileUploadController) processLocation(ctx context.Context, event *model.Event) (entities.Location, error) {

	location, err := controller.LocationRepository.GetLocationByName(ctx, event.Location)

	if err != nil {
		return location, err
	}

	if location.ID == 0 {
		createdLocation, err := controller.LocationRepository.CreateLocation(ctx, event.Location)
		if err != nil {
			return location, err
		}
		location = createdLocation
	}
	return location, err
}

func (controller *FileUploadController) processEvent(ctx context.Context, event *model.Event, location *entities.Location) (entities.Event, error) {

	getEventParams := entities.GetEventByLocationAndTypeAndDateParams{
		LocationID: location.ID,
		Type:       event.RaceType,
		Date:       event.Date,
	}

	eventEntity, err := controller.EventRepository.GetEventByLocationAndTypeAndDate(ctx, getEventParams)

	if err != nil {
		return eventEntity, err
	}

	if eventEntity.ID == 0 {
		createEventParams := entities.CreateEventParams{
			LocationID:   location.ID,
			Type:         event.RaceType,
			Date:         event.Date,
			TotalDrivers: int64(len(event.DriverTimes)),
		}

		savedEvent, err := controller.EventRepository.CreateEvent(ctx, createEventParams)

		if err != nil {
			return eventEntity, err
		}
		eventEntity = savedEvent
	}

	return eventEntity, err
}
