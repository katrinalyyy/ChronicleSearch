package handler

import (
	"Lab1/intermal/app/config"
	"Lab1/intermal/app/redis"
	"Lab1/intermal/app/repository"
	"Lab1/intermal/app/role"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	Repository  *repository.Repository
	Config      *config.Config
	RedisClient *redis.Client
	App         AuthMiddleware
}

type AuthMiddleware interface {
	WithAuthCheck(assignedRoles ...interface{}) gin.HandlerFunc
}

func NewHandler(r *repository.Repository, cfg *config.Config, redisClient *redis.Client, app AuthMiddleware) *Handler {
	return &Handler{
		Repository:  r,
		Config:      cfg,
		RedisClient: redisClient,
		App:         app,
	}
}

func (h *Handler) RegisterAPI(router *gin.Engine) {
	router.POST("/sign_up", h.Register)
	router.POST("/login", h.Login)
	router.POST("/logout", h.Logout)

	api := router.Group("/api")
	{
		// Chronicle Resources - разделяем по уровням доступа
		chronicles := api.Group("/chronicle_resources")
		{
			// GET методы - доступны всем (без авторизации)
			chronicles.GET("", h.GetChronicleResourcesAPI)
			chronicles.GET("/:id_chronicle_resource", h.GetChronicleResourceAPI)

			// Методы создания/изменения - требуют авторизации (Buyer, Manager, Admin)
			chronicles.POST("", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.CreateChronicleResourceAPI)
			chronicles.PUT("/:id_chronicle_resource", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.UpdateChronicleResourceAPI)
			chronicles.DELETE("/:id_chronicle_resource", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.DeleteChronicleResourceAPI)
			chronicles.POST("/:id_chronicle_resource/image", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.UploadChronicleResourceImageAPI)
			chronicles.POST("/:id_chronicle_resource/add_to_chronicle_request", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.AddChronicleToRequestAPI)
		}

		// Requests - требуют авторизации
		requests := api.Group("/ChronicleRequestList")
		{
			requests.GET("/chronicle_draft", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.GetDraftRequestInfoAPI)
			requests.GET("", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.GetRequestChronicleResearchAPI)
			requests.GET("/:id_chronicle_request", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.GetRequestWithChroniclesAPI)
			requests.PUT("/:id_chronicle_request", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.UpdateRequestChronicleResearchAPI)
			requests.PUT("/:id_chronicle_request/chronicle_request-form", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.FormRequestChronicleResearchAPI)

			// Завершение/отклонение - только для Admin
			requests.PUT("/:id_chronicle_request/chronicle_complete-or-reject", h.App.WithAuthCheck(role.Admin), h.CompleteOrRejectRequestChronicleResearchAPI)

			requests.DELETE("/:id_chronicle_request", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.DeleteRequestChronicleResearchAPI)
		}

		// Chronicle Research - требуют авторизации
		chronicleResearch := api.Group("/chronicle_research")
		{
			chronicleResearch.PUT("/:id/chronicles/:chronicle_id", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.UpdateChronicleResearchInRequestAPI)
			chronicleResearch.DELETE("/:id/chronicles/:chronicle_id", h.App.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.DeleteChronicleResearchFromRequestAPI)
		}
	}
}

func (h *Handler) RegisterStatic(router *gin.Engine) {
	// 	router.LoadHTMLGlob("templates/*")
	// 	router.Static("/static", "./resources")
}

func (h *Handler) RegisterSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (h *Handler) errorHandler(ctx *gin.Context, errorStatusCode int, err error) {
	logrus.Error(err.Error())
	ctx.JSON(errorStatusCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})
}
