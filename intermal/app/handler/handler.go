package handler

import (
	"Lab1/intermal/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{
		Repository: r,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET("/chronicles", h.GetChronicleResources)
	router.GET("/chronicle/:id", h.GetChronicleResource)
	router.GET("/chronicle-research/:id", h.GetChronicleResearch)

	router.POST("/chronicle-research/:id/delete-request", h.DeleteRequestChronicleResearch)

	router.POST("/chronicle/:id/add-to-research", h.AddChronicleToRequest)
}

func (h *Handler) RegisterStatic(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./resources")
}

func (h *Handler) errorHandler(ctx *gin.Context, errorStatusCode int, err error) {
	logrus.Error(err.Error())
	ctx.JSON(errorStatusCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})
}
