package handler

import (
	us "backend_course/branch_api_gateway/genproto/schedule_service"
	"backend_course/branch_api_gateway/genproto/user_service"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/group/create [post]
// @Summary		Creates a Group
// @Description	This api creates a Group and returns its id
// @Tags		Group
// @Accept		json
// @Produce		json
// @Param		Group body schedule_service.CreateGroup true "Group"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateGroup(c *gin.Context) {
	var req us.CreateGroup


	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	_,err:=h.grpcClient.TeacherService().GetByID(c.Request.Context(),&user_service.TeacherPrimaryKey{Id: req.TeacherId})
	if err!=nil {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong teacher id"), "wrong teacher id")
		return
	}

	_,err=h.grpcClient.SupportTeacherService().GetByID(c.Request.Context(),&user_service.SupportTeacherPrimaryKey{Id: req.SupportTeacherId})
	if err!=nil {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong support teacher id"), "wrong support teacher id")
		return
	}

	_,err=h.grpcClient.BranchBranch().GetByID(c.Request.Context(),&user_service.BranchPrimaryKey{Id: req.Branchid})
	if err!=nil {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong branch id"), "wrong branch id")
		return
	}
	
	resp, err := h.grpcClient.GroupService().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while create group")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/group/getbyid/{id} [get]
// @Summary		Get by id a Group
// @Description	This api get by id a Group
// @Tags		Group
// @Produce		json
// @Param		id path string true "Group id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIDGroup(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.GroupService().GetByID(c.Request.Context(), &us.GroupPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get by id group")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/group/getlist [get]
// @Summary		Get list a group
// @Description	This api get list a group
// @Tags		Group
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListGroup(c *gin.Context) {
	req := &us.GetListGroupRequest{}

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

	resp, err := h.grpcClient.GroupService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list group")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/group/updategroup [PUT]
// @Summary		Update a Group
// @Description	This API updates a Group
// @Tags		Group
// @Accept		json
// @Produce		json
// @Param		Group body schedule_service.UpdateGroupRequest true "Group object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateGroup(c *gin.Context) {
	req := &us.UpdateGroupRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding binding body")
		return
	}

	resp, err := h.grpcClient.GroupService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update Group", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while updating group")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/group/delete [delete]
// @Summary		delete a Group
// @Description	This api delete a Group
// @Tags		Group
// Accept		json
// @Produce		json
// @Param		Group body schedule_service.GroupPrimaryKey true "Group"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteGroup(c *gin.Context) {
	id := &us.GroupPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.GroupService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting group")
		return
	}

	c.JSON(http.StatusOK, resp)
}
