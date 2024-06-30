package handler

import (
	"backend_course/branch_api_gateway/genproto/schedule_service"
	us "backend_course/branch_api_gateway/genproto/schedule_service"
	"backend_course/branch_api_gateway/genproto/user_service"
	"backend_course/branch_api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/Student/create [post]
// @Summary		Creates a Student
// @Description	This api creates a Student and returns its id
// @Tags		Student
// @Accept		json
// @Produce		json
// @Param		Student body schedule_service.CreateStudent true "Student"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateStudent(c *gin.Context) {
	req := &us.CreateStudent{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	_, err := h.grpcClient.GroupService().GetByID(c.Request.Context(), &schedule_service.GroupPrimaryKey{Id: req.GroupId})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong group id"), "wrong group id")
		return
	}

	_, err = h.grpcClient.BranchBranch().GetByID(c.Request.Context(), &user_service.BranchPrimaryKey{Id: req.Branchid})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong branch id"), "wrong branch id")
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

	resp, err := h.grpcClient.StudentService().Create(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get create Student", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while creating Student")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router      /v1/Student/getbyid/{id} [get]
// @Summary     Get by id a Student
// @Description This api get by id a Student
// @Tags        Student
// @Produce     json
// @Param       id path string true "Student Id"
// @Success     200  {object}  models.ResponseSuccess
// @Failure     400  {object}  models.ResponseError
// @Failure     404  {object}  models.ResponseError
// @Failure     500  {object}  models.ResponseError
func (h *handler) GetByIdStudent(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.StudentService().GetByID(c.Request.Context(), &us.StudentPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("Error while getting Student by id: %v", err)
		handleGrpcErrWithDescription(c, h.log, err, "Error while getting Student by id")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/Student/getlist [get]
// @Summary		Get list a Student
// @Description	This api get list a Student
// @Tags		Student
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListStudent(c *gin.Context) {
	req := &us.GetListStudentRequest{}

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

	resp, err := h.grpcClient.StudentService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list Student")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/Student/update [PUT]
// @Summary		Update a Student
// @Description	This API updates a Student
// @Tags		Student
// @Accept		json
// @Produce		json
// @Param		Student body schedule_service.UpdateStudentRequest true "Student object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateStudent(c *gin.Context) {
	req := &us.UpdateStudentRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.StudentService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update Student", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while updating Student")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/Student/delete [delete]
// @Summary		delete a Student
// @Description	This api delete a Student
// @Tags		Student
// Accept		json
// @Produce		json
// @Param		Student body schedule_service.StudentPrimaryKey true "Student"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteStudent(c *gin.Context) {
	id := &us.StudentPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.StudentService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting Student")
		return
	}

	c.JSON(http.StatusOK, resp)
}
