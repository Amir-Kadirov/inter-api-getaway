package handler

import (
	us "backend_course/branch_api_gateway/genproto/schedule_service"
	"backend_course/branch_api_gateway/pkg/validator"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/event/create [post]
// @Summary		Creates a Event
// @Description	This api creates a Event and returns its id
// @Tags		Event
// @Accept		json
// @Produce		json
// @Param		Event body schedule_service.CreateEvent true "Event"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateEvent(c *gin.Context) {
	var req us.CreateEvent

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if err := validator.ValidateDaySunDay(req.Day); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while validating event day")
		return
	}

	resp, err := h.grpcClient.EventService().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while create event")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/event/getbyid/{id} [get]
// @Summary		Get by id a Event
// @Description	This api get by id a Event
// @Tags		Event
// @Produce		json
// @Param		id path string true "Event id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIDEvent(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.EventService().GetByID(c.Request.Context(), &us.EventPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get by id event")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/event/getlist [get]
// @Summary		Get list a event
// @Description	This api get list a event
// @Tags		Event
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListEvent(c *gin.Context) {
	req := &us.GetListEventRequest{}

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	search := c.Query("search")

	if limitStr != "" {
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			fmt.Errorf("error while parse limit", err)
			handleGrpcErrWithDescription(c, h.log, err, "error while parse limit")
			return
		}
		req.Limit = limit
	}

	if offsetStr != "" {
		offset, err := strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			fmt.Errorf("error while parse offset", err)
			handleGrpcErrWithDescription(c, h.log, err, "error while parse offset")
			return
		}
		req.Offset = offset
	}

	if search != "" {
		req.Search = search
	}

	resp, err := h.grpcClient.EventService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list event")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/event/updateevent [PUT]
// @Summary		Update a Event
// @Description	This API updates a Event
// @Tags		Event
// @Accept		json
// @Produce		json
// @Param		Event body schedule_service.UpdateEventRequest true "Event object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateEvent(c *gin.Context) {
	req := &us.UpdateEventRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding binding body")
		return
	}

	resp, err := h.grpcClient.EventService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update Event", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while updating event")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/event/delete [delete]
// @Summary		delete a Event
// @Description	This api delete a Event
// @Tags		Event
// Accept		json
// @Produce		json
// @Param		Event body schedule_service.EventPrimaryKey true "Event"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteEvent(c *gin.Context) {
	id := &us.EventPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.EventService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting event")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/event/registerevent [post]
// @Summary		Registers to Event
// @Description	This api register event
// @Tags		Event
// @Accept		json
// @Produce		json
// @Param		Register-Event body schedule_service.RegisterEv true "Event"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) RegisterEvent(c *gin.Context) {
	var req us.RegisterEv

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.EventService().RegisterEvent(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while register event")
		return
	}

	c.JSON(http.StatusOK, resp)
}
