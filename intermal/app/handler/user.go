package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"Lab1/intermal/app/ds"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUserAPI(ctx *gin.Context) {
	var user ds.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
		return
	}

	if user.Email == "" || user.Password == "" {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("email and password are required"))
		return
	}

	_, err := h.Repository.GetUserByEmail(user.Email)
	if err == nil {
		h.errorHandler(ctx, http.StatusConflict, fmt.Errorf("user with this email already exists"))
		return
	}

	createdUser, err := h.Repository.CreateUser(user)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	createdUser.Password = ""
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"data":    createdUser,
		"message": "User registered successfully",
	})
}

func (h *Handler) AuthenticateUserAPI(ctx *gin.Context) {
	var authData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&authData); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
		return
	}

	user, err := h.Repository.CheckCredentials(authData.Email, authData.Password)
	if err != nil {
		h.errorHandler(ctx, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		return
	}

	user.Password = ""
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    user,
		"message": "Authentication successful",
	})
}

func (h *Handler) GetUserProfileAPI(ctx *gin.Context) {
	userIDStr := ctx.Query("user_id")
	if userIDStr == "" {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("user_id is required"))
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid user_id format"))
		return
	}

	user, err := h.Repository.GetUserByID(uint(userID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.errorHandler(ctx, http.StatusNotFound, err)
		} else {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	user.Password = ""
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func (h *Handler) UpdateUserProfileAPI(ctx *gin.Context) {
	userIDStr := ctx.Query("user_id")
	if userIDStr == "" {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("user_id is required"))
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid user_id format"))
		return
	}

	var user ds.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
		return
	}

	user.ID = 0
	user.Email = ""
	user.Password = ""

	err = h.Repository.UpdateUser(uint(userID), user)
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
		"message": "User profile updated successfully",
	})
}

func (h *Handler) LogoutUserAPI(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User logged out successfully",
	})
}
