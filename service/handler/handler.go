package handler

import (
	"hack2hire-2022/dtos"
	"hack2hire-2022/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		zap.L().Error("parse request failed", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	message, err := h.sampleService.SayHello(id)
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
