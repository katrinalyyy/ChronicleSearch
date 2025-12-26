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
		chronicles := api.Group("/chronicle_resources")
		{
			chronicles.GET("", h.GetChronicleResourcesAPI)
			chronicles.GET("/:id_chronicle_resource", h.GetChronicleResourceAPI)
			chronicles.POST("", h.CreateChronicleResourceAPI)
			chronicles.PUT("/:id_chronicle_resource", h.UpdateChronicleResourceAPI)
			chronicles.DELETE("/:id_chronicle_resource", h.DeleteChronicleResourceAPI)
			chronicles.POST("/:id_chronicle_resource/image", h.UploadChronicleResourceImageAPI)
			chronicles.POST("/:id_chronicle_resource/add_to_chronicle_request", h.AddChronicleToRequestAPI)
		}

		requests := api.Group("/ChronicleRequestList")
		{
			requests.GET("/chronicle_draft", h.GetDraftRequestInfoAPI)
			requests.GET("", h.GetRequestChronicleResearchAPI)
			requests.GET("/:id_chronicle_request", h.GetRequestWithChroniclesAPI)
			requests.PUT("/:id_chronicle_request", h.UpdateRequestChronicleResearchAPI)
			requests.PUT("/:id_chronicle_request/chronicle_request-form", h.FormRequestChronicleResearchAPI)
			requests.PUT("/:id_chronicle_request/chronicle_complete-or-reject", h.CompleteOrRejectRequestChronicleResearchAPI)
			requests.DELETE("/:id_chronicle_request", h.DeleteRequestChronicleResearchAPI)
		}

		chronicleResearch := api.Group("/chronicle_research")
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
