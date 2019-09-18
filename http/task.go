package http

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"todo/models"
)

func fileHeaderToByte(file *multipart.FileHeader) ([]byte, error) {
	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}

	byteFile, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return nil, err
	}

	return byteFile, nil
}

func newTask(c *gin.Context) (*models.Task, error) {
	claims := jwt.ExtractClaims(c)
	data := models.NewTask(claims[identityKey].(string))

	if err := c.Bind(data); err != nil {
		return nil, err
	}

	if attachmentRequest, err := c.FormFile("attachment"); err == nil {
		file, err := fileHeaderToByte(attachmentRequest)
		if err != nil {
			return nil, err
		}

		data.AttachmentFile = file
	}

	return data, nil
}

func (s *Server) allTasks(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	tasks, err := s.taskService.FindAll(&models.TaskFilter{UserID: claims[identityKey].(string)})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

func (s *Server) getTask(c *gin.Context) {
	id := c.Param("id")
	tasks, err := s.taskService.Get(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

func (s *Server) createTask(c *gin.Context) {
	data, err := newTask(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := s.taskService.Create(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func (s *Server) updateTask(c *gin.Context) {
	data, err := newTask(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := s.taskService.Update(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, task)
	}
}
