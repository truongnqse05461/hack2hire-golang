package handler

import (
	"hack2hire-2022/dtos"
	"hack2hire-2022/model"
	"hack2hire-2022/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	bookingService services.BookingService
}

func NewHandler(bookingService services.BookingService) Handler {
	return Handler{bookingService: bookingService}
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
	message, err := h.bookingService.SayHello(id)
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

func (h *Handler) SaveBookings(ctx *gin.Context) {
	showId := ctx.Param("id")
	if showId == "" {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	var req dtos.BookingReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.L().Error("parse request failed", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	err := h.bookingService.Save(model.Bookings{
		User: model.User{
			Name:        req.User.Name,
			PhoneNumber: req.User.PhoneNumber,
		},
		SeatCodes: req.SeatCodes,
	})
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
		"message": "success",
	})
}
