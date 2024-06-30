package handler

import (
	us "backend_course/branch_api_gateway/genproto/schedule_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/lesson/create [post]
// @Summary		Creates a Lesson
// @Description	This api creates a Lesson and returns its id
// @Tags		Lesson
// @Accept		json
// @Produce		json
// @Param		Lesson body schedule_service.CreateLesson true "Lesson"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateLesson(c *gin.Context) {
	var req us.CreateLesson

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.LessonService().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while create lesson")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/lesson/getbyid/{id} [get]
// @Summary		Get by id a Lesson
// @Description	This api get by id a Lesson
// @Tags		Lesson
// @Produce		json
// @Param		id path string true "Lesson id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIDLesson(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.LessonService().GetByID(c.Request.Context(), &us.LessonPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get by id lesson")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/lesson/getlist [get]
// @Summary		Get list a lesson
// @Description	This api get list a lesson
// @Tags		Lesson
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListLesson(c *gin.Context) {
	req := &us.GetListLessonRequest{}

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

	resp, err := h.grpcClient.LessonService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list lesson")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/lesson/updatelesson [PUT]
// @Summary		Update a Lesson
// @Description	This API updates a Lesson
// @Tags		Lesson
// @Accept		json
// @Produce		json
// @Param		Lesson body schedule_service.UpdateLessonRequest true "Lesson object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateLesson(c *gin.Context) {
	req := &us.UpdateLessonRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding binding body")
		return
	}

	resp, err := h.grpcClient.LessonService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update Lesson", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while updating lesson")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/lesson/delete [delete]
// @Summary		delete a Lesson
// @Description	This api delete a Lesson
// @Tags		Lesson
// Accept		json
// @Produce		json
// @Param		Lesson body schedule_service.LessonPrimaryKey true "Lesson"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteLesson(c *gin.Context) {
	id := &us.LessonPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.LessonService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting lesson")
		return
	}

	c.JSON(http.StatusOK, resp)
}
