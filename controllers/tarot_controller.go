package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tarot-app/database"
	"tarot-app/repositories"
	"tarot-app/services"
)

type TarotController struct {
	Service *services.TarotService
}

func NewTarotController(service *services.TarotService) *TarotController {
	return &TarotController{Service: service}
}

func (ctrl *TarotController) GetAllCards(c *gin.Context) {
	cards, err := ctrl.Service.GetAllCards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve cards"})
		return
	}
	c.JSON(http.StatusOK, cards)
}

func (ctrl *TarotController) GetCardByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	card, err := ctrl.Service.GetCardByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}
	c.JSON(http.StatusOK, card)
}

func RegisterRoutes(r *gin.Engine) {
	db := database.DB
	repo := repositories.NewTarotRepository(db)
	service := services.NewTarotService(repo)
	controller := NewTarotController(service)

	r.GET("/tarot", controller.GetAllCards)
	r.GET("/tarot/:id", controller.GetCardByID)
}
