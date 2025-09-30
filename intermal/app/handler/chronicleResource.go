package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"Lab1/intermal/app/ds"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetChronicleResourcesAPI(ctx *gin.Context) {
	title := ctx.Query("title")
	author := ctx.Query("author")
	location := ctx.Query("location")

	resources, err := h.Repository.GetChronicleResources(title, author, location)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   resources,
	})
}

func (h *Handler) GetChronicleResourceAPI(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid ID format"))
		return
	}

	resource, err := h.Repository.GetChronicleResource(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.errorHandler(ctx, http.StatusNotFound, err)
		} else {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   resource,
	})
}

func (h *Handler) CreateChronicleResourceAPI(ctx *gin.Context) {
	var resource ds.ChronicleResource
	if err := ctx.ShouldBindJSON(&resource); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
		return
	}

	resource.ID = 0

	createdResource, err := h.Repository.CreateChronicleResource(resource)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   createdResource,
	})
}

func (h *Handler) UpdateChronicleResourceAPI(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid ID format"))
		return
	}

	var resource ds.ChronicleResource
	if err := ctx.ShouldBindJSON(&resource); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
		return
	}

	err = h.Repository.UpdateChronicleResource(uint(id), resource)
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
		"message": "Chronicle resource updated successfully",
	})
}

func (h *Handler) DeleteChronicleResourceAPI(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid ID format"))
		return
	}

	err = h.Repository.DeleteChronicleResource(uint(id))
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
		"message": "Chronicle resource deleted successfully",
	})
}

func (h *Handler) UploadChronicleResourceImageAPI(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid ID format"))
		return
	}

	_, err = h.Repository.GetChronicleResource(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.errorHandler(ctx, http.StatusNotFound, err)
		} else {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("no image file provided"))
		return
	}

	src, err := file.Open()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}
	defer src.Close()

	fileName := h.Repository.GenerateImageFileName(file.Filename)

	err = h.Repository.UploadFileToMinIO(
		context.Background(),
		fileName,
		src,
		file.Size,
		file.Header.Get("Content-Type"),
	)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	imagePath := "http://127.0.0.1:9000/chronicles/" + fileName
	err = h.Repository.UpdateChronicleResourceImage(uint(id), imagePath)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"message":    "Image uploaded successfully",
		"image_path": imagePath,
	})
}

func (h *Handler) AddChronicleToRequestAPI(ctx *gin.Context) {
	idStr := ctx.Param("id")
	chronicleID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid ID format"))
		return
	}

	_, err = h.Repository.GetChronicleResource(uint(chronicleID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.errorHandler(ctx, http.StatusNotFound, err)
		} else {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	draftRequest, _, err := h.Repository.GetDraftRequestChronicleResearchInfo()
	if err != nil {
		draftRequest, err = h.Repository.CreateRequestChronicleResearchWithChronicle(uint(chronicleID))
		if err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, fmt.Errorf("failed to create draft request: %v", err))
			return
		}
	} else {
		err = h.Repository.AddChronicleToRequest(draftRequest.ID, uint(chronicleID), "")
		if err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, fmt.Errorf("failed to add chronicle to request: %v", err))
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"message":    "Chronicle added to draft request successfully",
		"request_id": draftRequest.ID,
	})
}
