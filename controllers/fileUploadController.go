package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"racer_http/repository"
	"racer_http/sqlite/entities"
)

type FileUploadController struct {
	UserRepository repository.UserRepository
}

func NewFileUploadController(userRepository repository.UserRepository) *FileUploadController {
	return &FileUploadController{
		UserRepository: userRepository,
	}
}

func (controller *FileUploadController) UploadFile(c *gin.Context) {
	u, _ := c.Get("currentUser")

	user, _ := u.(entities.GetUserByIdRow)
	form, _ := c.MultipartForm()
	file, _ := form.File["file"]

	fmt.Printf("User %s, uploaded file %s", user.FirstName, file[0].Filename)
}
