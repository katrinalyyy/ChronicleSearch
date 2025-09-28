package handler

import (
	"Lab1/intermal/app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ChronicleController struct {
	ChronicleModel *repository.ChronicleModel
}

func NewChronicleController(r *repository.ChronicleModel) *ChronicleController {
	return &ChronicleController{
		ChronicleModel: r,
	}
}

func (h *ChronicleController) GetChronicleResources(ctx *gin.Context) {
	var chronicleResources []repository.ChronicleResource
	var err error

	searchChronicleQuery := ctx.Query("chronicle") // получаем значение из поля поиска
	if searchChronicleQuery == "" {                // если поле поиска пусто, то просто получаем из репозитория все записи
		chronicleResources, err = h.ChronicleModel.GetChronicleResources()
		if err != nil {
			logrus.Error(err)
		}
	} else {
		chronicleResources, err = h.ChronicleModel.GetChronicleResourcesByTitle(searchChronicleQuery) // в ином случае ищем заказ по заголовку
		if err != nil {
			logrus.Error(err)
		}
	}

	requestID := 1
	_, chronicleResearch, err := h.ChronicleModel.GetChronicleResearchForRequest(requestID)
	if err != nil {
		logrus.Error(err)
	}
	chronicleResearchCount := len(chronicleResearch)

	ctx.HTML(http.StatusOK, "chronicleResources.html", gin.H{
		"chronicleResources":        chronicleResources,
		"chronicleQuery":            searchChronicleQuery,
		"chronicleApplicationCount": chronicleResearchCount,
	})
}

func (h *ChronicleController) GetChronicleResource(ctx *gin.Context) {
	idStr := ctx.Param("id") // получаем id заказа из урла (то есть из /order/:id)
	// через двоеточие мы указываем параметры, которые потом сможем считать через функцию выше
	id, err := strconv.Atoi(idStr) // так как функция выше возвращает нам строку, нужно ее преобразовать в int
	if err != nil {
		logrus.Error(err)
	}

	chronicleResource, err := h.ChronicleModel.GetChronicleResource(id)
	if err != nil {
		logrus.Error(err)
	}

	requestID := 1
	_, chronicleResearch, err := h.ChronicleModel.GetChronicleResearchForRequest(requestID)
	if err != nil {
		logrus.Error(err)
	}
	chronicleResearchCount := len(chronicleResearch)

	ctx.HTML(http.StatusOK, "chronicleDetailedResource.html", gin.H{
		"chronicleResource":         chronicleResource,
		"chronicleApplicationCount": chronicleResearchCount,
	})
}

func (h *ChronicleController) GetChronicleApplication(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Error(err)
	}
	requestChronicleResearch, chronicleResearch, err := h.ChronicleModel.GetChronicleResearchForRequest(id)
	if err != nil {
		logrus.Error(err)
	}

	var resources []repository.ChronicleResource
	resources, err = h.ChronicleModel.GetChronicleResources()
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "requestChronicleResearch.html", gin.H{
		"requestChronicleResearch": requestChronicleResearch,
		"research":                 chronicleResearch,
		"resources":                resources,
	})
}
