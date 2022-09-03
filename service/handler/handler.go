package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hack2hire-2022/dtos"
	"hack2hire-2022/services"
	"net/http"
	"strconv"
)

type Handler struct {
	sampleService services.SampleService
}

func NewHandler(sampleService services.SampleService) Handler {
	return Handler{sampleService: sampleService}
}

func (h *Handler) Health(ctx *gin.Context) {
	zap.L().Info("health check request", zap.String("status", "running"))
	ctx.JSON(http.StatusOK, gin.H{
		"status": "running",
	})
}

func (h *Handler) GetMessage(ctx *gin.Context) {
	id, ok := ctx.Request.URL.Query()["id"]
	if !ok || len(id) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
	}
	IDint, err := strconv.ParseInt(id[0], 10, 64)
	message, err := h.sampleService.SayHello(IDint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.Response{
			Data: nil,
			Meta: dtos.Meta{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
	return
}
