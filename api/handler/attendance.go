package handler

import (
	us "backend_course/branch_api_gateway/genproto/schedule_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/attendance/create [post]
// @Summary		Creates a Attendance
// @Description	This api creates a Attendance and returns its id
// @Tags		Attendance
// @Accept		json
// @Produce		json
// @Param		Attendance body schedule_service.CreateAttendance true "Attendance"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateAttendance(c *gin.Context) {
	var req us.CreateAttendance

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.AttendanceService().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while create attendance")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/attendance/getbyid/{id} [get]
// @Summary		Get by id a Attendance
// @Description	This api get by id a Attendance
// @Tags		Attendance
// @Produce		json
// @Param		id path string true "Student id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIDAttendance(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.AttendanceService().GetByID(c.Request.Context(), &us.AttendancePrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get by id attendance")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/attendance/getlist [get]
// @Summary		Get list a attendance
// @Description	This api get list a attendance
// @Tags		Attendance
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search student by id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListAttendance(c *gin.Context) {
	req := &us.GetListAttendanceRequest{}

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

	resp, err := h.grpcClient.AttendanceService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list attendance")
		return
	}

	c.JSON(http.StatusOK, resp)
}
