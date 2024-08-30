package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tarot-app/models"
	"tarot-app/repositories"
)

type TarotSpreadController struct {
	Repo *repositories.TarotSpreadRepository
}

func NewTarotSpreadController(repo *repositories.TarotSpreadRepository) *TarotSpreadController {
	return &TarotSpreadController{Repo: repo}
}

func (ctrl *TarotSpreadController) CreateSpread(c *gin.Context) {
	var spread models.TarotSpread
	if err := c.ShouldBindJSON(&spread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdSpread := ctrl.Repo.Create(spread)
	c.JSON(http.StatusOK, createdSpread)
}

func (ctrl *TarotSpreadController) GetSpread(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	spread, err := ctrl.Repo.Read(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, spread)
}

func (ctrl *TarotSpreadController) UpdateSpread(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var spread models.TarotSpread
	if err := c.ShouldBindJSON(&spread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedSpread, err := ctrl.Repo.Update(id, spread)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedSpread)
}

func (ctrl *TarotSpreadController) DeleteSpread(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.Repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "spread deleted"})
}

func (ctrl *TarotSpreadController) ListSpreads(c *gin.Context) {
	spreads := ctrl.Repo.List()
	
	c.JSON(http.StatusOK, spreads)
}
