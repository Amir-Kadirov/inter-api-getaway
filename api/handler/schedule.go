package handler

import (
	us "backend_course/branch_api_gateway/genproto/schedule_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/schedule/create [post]
// @Summary		Creates a Schedule
// @Description	This api creates a Schedule and returns its id
// @Tags		Schedule
// @Accept		json
// @Produce		json
// @Param		Schedule body schedule_service.CreateSchedule true "Schedule"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateSchedule(c *gin.Context) {
	var req us.CreateSchedule

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.ScheduleService().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while create schedule")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/schedule/getbyid/{id} [get]
// @Summary		Get by id a Schedule
// @Description	This api get by id a Schedule
// @Tags		Schedule
// @Produce		json
// @Param		id path string true "Schedule id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIDSchedule(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.ScheduleService().GetByID(c.Request.Context(), &us.SchedulePrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get by id schedule")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/schedule/getlist [get]
// @Summary		Get list a schedule
// @Description	This api get list a schedule
// @Tags		Schedule
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListSchedule(c *gin.Context) {
	req := &us.GetListScheduleRequest{}

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

	resp, err := h.grpcClient.ScheduleService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list schedule")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/schedule/updateschedule [PUT]
// @Summary		Update a Schedule
// @Description	This API updates a Schedule
// @Tags		Schedule
// @Accept		json
// @Produce		json
// @Param		Schedule body schedule_service.UpdateScheduleRequest true "Schedule object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateSchedule(c *gin.Context) {
	req := &us.UpdateScheduleRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding binding body")
		return
	}

	resp, err := h.grpcClient.ScheduleService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update Schedule", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while updating schedule")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/schedule/delete [delete]
// @Summary		delete a Schedule
// @Description	This api delete a Schedule
// @Tags		Schedule
// Accept		json
// @Produce		json
// @Param		Schedule body schedule_service.SchedulePrimaryKey true "Schedule"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteSchedule(c *gin.Context) {
	id := &us.SchedulePrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.ScheduleService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting schedule")
		return
	}

	c.JSON(http.StatusOK, resp)
}
