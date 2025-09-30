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

func (h *Handler) RegisterAPI(router *gin.Engine) {
	api := router.Group("/api")
	{
		chronicles := api.Group("/chronicle-resources")
		{
			chronicles.GET("", h.GetChronicleResourcesAPI)
			chronicles.GET("/:id", h.GetChronicleResourceAPI)
			chronicles.POST("", h.CreateChronicleResourceAPI)
			chronicles.PUT("/:id", h.UpdateChronicleResourceAPI)
			chronicles.DELETE("/:id", h.DeleteChronicleResourceAPI)
			chronicles.POST("/:id/image", h.UploadChronicleResourceImageAPI)
			chronicles.POST("/:id/add-to-request", h.AddChronicleToRequestAPI)
		}

		requests := api.Group("/requests")
		{
			requests.GET("/draft-info", h.GetDraftRequestInfoAPI)
			requests.GET("", h.GetRequestChronicleResearchAPI)
			requests.GET("/:id", h.GetRequestWithChroniclesAPI)
			requests.PUT("/:id", h.UpdateRequestChronicleResearchAPI)
			requests.PUT("/:id/form", h.FormRequestChronicleResearchAPI)
			requests.PUT("/:id/complete-or-reject", h.CompleteOrRejectRequestChronicleResearchAPI)
			requests.DELETE("/:id", h.DeleteRequestChronicleResearchAPI)
		}

		chronicleResearch := api.Group("/chronicle-research")
		{
			chronicleResearch.PUT("/:id/chronicles/:chronicle_id", h.UpdateChronicleResearchInRequestAPI)
			chronicleResearch.DELETE("/:id/chronicles/:chronicle_id", h.DeleteChronicleResearchFromRequestAPI)
		}

		users := api.Group("/users")
		{
			users.POST("/register", h.RegisterUserAPI)
			users.POST("/auth", h.AuthenticateUserAPI)
			users.POST("/logout", h.LogoutUserAPI)
			users.GET("/profile", h.GetUserProfileAPI)
			users.PUT("/profile", h.UpdateUserProfileAPI)
		}

	}
}

func (h *Handler) RegisterStatic(router *gin.Engine) {
	// 	router.LoadHTMLGlob("templates/*")
	// 	router.Static("/static", "./resources")
}

func (h *Handler) errorHandler(ctx *gin.Context, errorStatusCode int, err error) {
	logrus.Error(err.Error())
	ctx.JSON(errorStatusCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})
}
