package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"Lab1/intermal/app/ds"
	"Lab1/intermal/app/repository"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDraftRequestInfoAPI(ctx *gin.Context) {
	requestID, count, err := h.Repository.GetDraftRequestInfo()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"request_id": requestID,
		"count":      count,
	})
}

func (h *Handler) GetRequestChronicleResearchAPI(ctx *gin.Context) {
	var startDate, endDate *time.Time

	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &parsed
		}
	}

	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &parsed
		}
	}

	requests, err := h.Repository.GetRequestChronicleResearch("", startDate, endDate)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   requests,
	})
}

func (h *Handler) GetRequestWithChroniclesAPI(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid ID format"))
		return
	}

	request, chronicles, err := h.Repository.GetRequestWithChronicles(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.errorHandler(ctx, http.StatusNotFound, err)
		} else {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"request":    request,
		"chronicles": chronicles,
	})
}

func (h *Handler) UpdateRequestChronicleResearchAPI(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid ID format"))
		return
	}

	var request ds.RequestChronicleResearch
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
		return
	}

	err = h.Repository.UpdateRequestChronicleResearch(uint(id), request)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.errorHandler(ctx, http.StatusNotFound, err)
		} else {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Request updated successfully",
	})
}

func (h *Handler) FormRequestChronicleResearchAPI(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid ID format"))
		return
	}

	err = h.Repository.FormRequestChronicleResearch(uint(id), repository.GetFixedCreatorID())
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "cannot be formed") ||
			strings.Contains(err.Error(), "доступен только черновик") || strings.Contains(err.Error(), "заявка пуста") ||
			strings.Contains(err.Error(), "необходимо заполнить") {
			h.errorHandler(ctx, http.StatusBadRequest, err)
		} else {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Request formed successfully",
	})
}

func (h *Handler) CompleteOrRejectRequestChronicleResearchAPI(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid ID format"))
		return
	}

	var requestBody struct {
		Action string `json:"action" binding:"required"` // "complete" или "reject"
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
		return
	}

	switch requestBody.Action {
	case "complete":
		err = h.Repository.CompleteRequestChronicleResearch(uint(id), repository.GetFixedModeratorID())
		if err != nil {
			if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "cannot be completed") {
				h.errorHandler(ctx, http.StatusBadRequest, err)
			} else {
				h.errorHandler(ctx, http.StatusInternalServerError, err)
			}
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Request completed successfully",
		})
	case "reject":
		err = h.Repository.RejectRequestChronicleResearch(uint(id), repository.GetFixedModeratorID())
		if err != nil {
			if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "cannot be rejected") {
				h.errorHandler(ctx, http.StatusBadRequest, err)
			} else {
				h.errorHandler(ctx, http.StatusInternalServerError, err)
			}
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Request rejected successfully",
		})
	default:
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid action. Use 'complete' or 'reject'"))
	}
}

func (h *Handler) DeleteRequestChronicleResearchAPI(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid ID format"))
		return
	}

	err = h.Repository.DeleteRequestChronicleResearch(uint(id), repository.GetFixedCreatorID())
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "cannot be deleted") {
			h.errorHandler(ctx, http.StatusBadRequest, err)
		} else {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Request deleted successfully",
	})
}
