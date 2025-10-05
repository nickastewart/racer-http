package controllers

import (
	"fmt"
	"racer_http/repository"
	"racer_http/sqlite/entities"
	
	racer "github.com/nickastewart/racer-parser"
	"github.com/gin-gonic/gin"
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
	multipartFile, _ := form.File["file"]

	fmt.Printf("User %s, uploaded file %s \n", user.FirstName, multipartFile[0].Filename)

	file, _ := multipartFile[0].Open()
	
	fmt.Print(racer.ParseFile(file))
}
