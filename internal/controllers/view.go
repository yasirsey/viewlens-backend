package controllers

import (
	"net/http"
	"viewlens/internal/services"

	"github.com/gin-gonic/gin"
)

type ViewController struct {
	service *services.ViewService
}

func NewViewController(service *services.ViewService) *ViewController {
	return &ViewController{service}
}

func (vc *ViewController) ListViews(c *gin.Context) {
	views, err := vc.service.ListViews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"views": views})
}

func (vc *ViewController) GetViewDetails(c *gin.Context) {
	viewName := c.Param("viewName")
	columns, data, err := vc.service.GetViewDetails(viewName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"columns": columns, "data": data})
}
