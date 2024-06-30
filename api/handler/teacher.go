package handler

import (
	us "backend_course/branch_api_gateway/genproto/user_service"
	"backend_course/branch_api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/teacher/create [post]
// @Summary		Creates a Teacher
// @Description	This api creates a Teacher and returns its id
// @Tags		Teacher
// @Accept		json
// @Produce		json
// @Param		Teacher body user_service.CreateTeacher true "Teacher"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateTeacher(c *gin.Context) {
	req := &us.CreateTeacher{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if !validator.ValidatePhone(req.Phone) {
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating phone"), "wrong phone")
		return
	}

	if !validator.ValidateGmail(req.Email) {
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating gmail"), "wrong gmail")
		return
	}

	resp, err := h.grpcClient.TeacherService().Create(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get create Teacher", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while create Teacher")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/teacher/getbyid/{id} [get]
// @Summary		Get by id a Teacher
// @Description	This api get by id a Teacher
// @Tags		Teacher
// @Produce		json
// @Param		id path string true "Teacher id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIdTeacher(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.TeacherService().GetByID(c.Request.Context(), &us.TeacherPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting by id")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/teacher/getlist [get]
// @Summary		Get list a Teacher
// @Description	This api get list a Teacher
// @Tags		Teacher
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListTeacher(c *gin.Context) {
	req := &us.GetListTeacherRequest{}

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

	resp, err := h.grpcClient.TeacherService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list Teacher")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/teacher/update [put]
// @Summary		Update a Teacher
// @Description	This API updates a Teacher
// @Tags		Teacher
// @Accept		json
// @Produce		json
// @Param		Teacher body user_service.UpdateTeacherRequest true "Teacher object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateTeacher(c *gin.Context) {
	req := &us.UpdateTeacherRequest{}

	if !validator.ValidatePhone(req.Phone) {
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating phone"), "wrong phone")
		return
	}

	if !validator.ValidateGmail(req.Email) {
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating gmail"), "wrong gmail")
		return
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while update Teacher")
		return
	}

	resp, err := h.grpcClient.TeacherService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update Teacher", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while ")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/teacher/delete [delete]
// @Summary		delete a Teacher
// @Description	This api delete a Teacher
// @Tags		Teacher
// Accept		json
// @Produce		json
// @Param		Teacher body user_service.TeacherPrimaryKey true "Teacher"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteTeacher(c *gin.Context) {
	id := &us.TeacherPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.TeacherService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while delete Teacher")
		return
	}

	c.JSON(http.StatusOK, resp)
}