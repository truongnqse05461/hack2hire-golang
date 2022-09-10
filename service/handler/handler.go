package handler

import (
	"hack2hire-2022/dtos"
	"hack2hire-2022/services"
	"hack2hire-2022/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	OK = "OK"
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

}

func (h *Handler) GetShows(ctx *gin.Context) {
	shows, err := h.bookingService.GetShows(ctx)
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
	ctx.JSON(http.StatusOK, dtos.Response{
		Data: dtos.ShowRes{Total: len(shows), Shows: shows},
		Meta: dtos.Meta{
			Message:    OK,
			StatusCode: http.StatusOK,
		},
	})

}

func (h *Handler) GetSeats(ctx *gin.Context) {
	showId := ctx.Param("show_id")
	if showId == "" {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Data: nil,
			Meta: dtos.Meta{
				Message:    "bad request",
				StatusCode: http.StatusBadRequest,
			},
		})
		return
	}
	seats, err := h.bookingService.GetSeats(ctx, showId)
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == utils.ErrShowNotFound {
			code = http.StatusNotFound
		}
		ctx.JSON(code, dtos.Response{
			Data: nil,
			Meta: dtos.Meta{
				Message:    err.Error(),
				StatusCode: code,
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, dtos.Response{
		Data: dtos.SeatRes{Total: len(seats), Seats: seats},
		Meta: dtos.Meta{
			Message:    OK,
			StatusCode: http.StatusOK,
		},
	})

}

func (h *Handler) GetReservations(ctx *gin.Context) {
	showId := ctx.Param("show_id")
	if showId == "" {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Data: nil,
			Meta: dtos.Meta{
				Message:    "bad request",
				StatusCode: http.StatusBadRequest,
			},
		})
		return
	}
	phoneNum := ctx.Query("phone_number")
	seatCodes := ctx.QueryArray("seat_codes")

	data, err := h.bookingService.GetReservations(ctx, showId, phoneNum, seatCodes...)
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == utils.ErrShowNotFound {
			code = http.StatusNotFound
		}
		ctx.JSON(code, dtos.Response{
			Data: nil,
			Meta: dtos.Meta{
				Message:    err.Error(),
				StatusCode: code,
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, dtos.Response{
		Data: dtos.ReservationRes{Total: len(data), Reservations: data},
		Meta: dtos.Meta{
			Message:    OK,
			StatusCode: http.StatusOK,
		},
	})

}

func (h *Handler) SaveReservation(ctx *gin.Context) {
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
	// err := h.bookingService.Save(model.Reservation{
	// 	User: model.User{
	// 		Name:        req.User.Name,
	// 		PhoneNumber: req.User.PhoneNumber,
	// 	},
	// 	SeatCodes: req.SeatCodes,
	// })
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, dtos.Response{
	// 		Data: nil,
	// 		Meta: dtos.Meta{
	// 			Message:    err.Error(),
	// 			StatusCode: http.StatusInternalServerError,
	// 		},
	// 	})
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *Handler) SaveShows(ctx *gin.Context) {
	var req dtos.SaveShowsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.L().Error("parse request failed", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	if err := h.bookingService.SaveShows(ctx, req.Shows...); err != nil {
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
