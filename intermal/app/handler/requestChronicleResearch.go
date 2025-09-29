package handler

import (
	"Lab1/intermal/app/ds"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetChronicleResearch(ctx *gin.Context) {
	idRequestStr := ctx.Param("id")
	idRequest, err := strconv.Atoi(idRequestStr)
	if err != nil {
		logrus.Error(err)
	}

	requestChronicleResearch, chronicleResearch, err := h.Repository.GetChronicleResearchForRequest(idRequest)
	if err != nil || requestChronicleResearch.Status == ds.RequestStatusDeleted {
		ctx.HTML(http.StatusOK, "requestChronicleResearch.html", gin.H{
			"error": "Заявка не найдена",
		})
		return
	}

	var resources []ds.ChronicleResource
	resources, err = h.Repository.GetChronicleResources()
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "requestChronicleResearch.html", gin.H{
		"requestChronicleResearch": requestChronicleResearch,
		"research":                 chronicleResearch,
		"resources":                resources,
	})
}

func (h *Handler) DeleteRequestChronicleResearch(ctx *gin.Context) {
	idRequestStr := ctx.Param("id")
	id, err := strconv.Atoi(idRequestStr)
	if err != nil {
		logrus.Errorf("Error converting request ID: %v", err)
	}

	err = h.Repository.DeleteRequestChronicleResearch(id)
	if err != nil {
		logrus.Errorf("Error deleting request: %v", err)
	}

	ctx.Redirect(http.StatusFound, "/chronicles")
}
