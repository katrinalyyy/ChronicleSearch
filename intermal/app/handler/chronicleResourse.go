package handler

import (
	"Lab1/intermal/app/ds"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetChronicleResources(ctx *gin.Context) {
	var chronicleResources []ds.ChronicleResource
	var err error

	searchChronicle := ctx.Query("chronicle")
	if searchChronicle == "" {
		chronicleResources, err = h.Repository.GetChronicleResources()
		if err != nil {
			logrus.Error(err)
		}
	} else {
		chronicleResources, err = h.Repository.GetChronicleResourcesByTitle(searchChronicle)
		if err != nil {
			logrus.Error(err)
		}
	}

	draftRequest, chronicleResearch, err := h.Repository.GetDraftRequestChronicleResearchInfo()
	var draftRequestID uint = 0
	var chronicleResearchCount int = 0
	if err == nil {
		draftRequestID = draftRequest.ID
		chronicleResearchCount = len(chronicleResearch)
	}

	ctx.HTML(http.StatusOK, "chronicleResources.html", gin.H{
		"chronicleResources":        chronicleResources,
		"chronicleQuery":            searchChronicle,
		"chronicleApplicationCount": chronicleResearchCount,
		"draftRequestID":            draftRequestID,
	})
}

func (h *Handler) GetChronicleResource(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Error(err)
	}

	chronicleResource, err := h.Repository.GetChronicleResource(id)
	if err != nil {
		logrus.Error(err)
	}

	draftRequest, chronicleResearch, err := h.Repository.GetDraftRequestChronicleResearchInfo()
	var chronicleResearchCount int = 0
	if err == nil {
		chronicleResearchCount = len(chronicleResearch)
	}

	ctx.HTML(http.StatusOK, "chronicleDetailedResource.html", gin.H{
		"chronicleResource":         chronicleResource,
		"chronicleApplicationCount": chronicleResearchCount,
		"draftRequestID":            draftRequest.ID,
	})
}

func (h *Handler) AddChronicleToRequest(ctx *gin.Context) {
	chronicleIDStr := ctx.Param("id")
	chronicleID, err := strconv.Atoi(chronicleIDStr)
	if err != nil {
		logrus.Error("Error converting chronicle ID:", err)
	}

	draftRequest, _, err := h.Repository.GetDraftRequestChronicleResearchInfo()
	if err != nil {
		_, err := h.Repository.CreateRequestChronicleResearch(uint(chronicleID))
		if err != nil {
			logrus.Error("Error creating new draft request:", err)
		}
	} else {
		err = h.Repository.AddChronicleResearchToRequest(draftRequest.ID, uint(chronicleID))
		if err != nil {
			logrus.Error("Error adding chronicle to existing request:", err)
		}
	}

	ctx.Redirect(http.StatusFound, "/chronicles")
}
