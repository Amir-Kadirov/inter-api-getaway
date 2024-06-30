package handler

import (
	us "backend_course/branch_api_gateway/genproto/schedule_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/task/create [post]
// @Summary		Creates a Task
// @Description	This api creates a Task and returns its id
// @Tags		Task
// @Accept		json
// @Produce		json
// @Param		Task body schedule_service.CreateTask true "Task"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateTask(c *gin.Context) {
	var req us.CreateTask

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.TaskService().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while create task")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/task/getbyid/{id} [get]
// @Summary		Get by id a Task
// @Description	This api get by id a Task
// @Tags		Task
// @Produce		json
// @Param		id path string true "Task id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIDTask(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.TaskService().GetByID(c.Request.Context(), &us.TaskPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get by id task")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/task/getlist [get]
// @Summary		Get list a task
// @Description	This api get list a task
// @Tags		Task
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListTask(c *gin.Context) {
	req := &us.GetListTaskRequest{}

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

	resp, err := h.grpcClient.TaskService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list task")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/task/updatetask [PUT]
// @Summary		Update a Task
// @Description	This API updates a Task
// @Tags		Task
// @Accept		json
// @Produce		json
// @Param		Task body schedule_service.UpdateTask true "Task object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateTask(c *gin.Context) {
	req := &us.UpdateTask{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding binding body")
		return
	}

	resp, err := h.grpcClient.TaskService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update Task", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while updating task")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/task/delete [delete]
// @Summary		delete a Task
// @Description	This api delete a Task
// @Tags		Task
// Accept		json
// @Produce		json
// @Param		Task body schedule_service.TaskPrimaryKey true "Task"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteTask(c *gin.Context) {
	id := &us.TaskPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.TaskService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting task")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/task/dotask [post]
// @Summary		Creates a Task
// @Description	This api do task
// @Tags		Task
// @Accept		json
// @Produce		json
// @Param		Task body schedule_service.DoTaskReq true "Task"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DoTask(c *gin.Context) {
	var req us.DoTaskReq

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.TaskService().DoTask(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while do task task")
		return
	}

	c.JSON(http.StatusOK, resp)
}
