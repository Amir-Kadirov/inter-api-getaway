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

// @Router		/v1/supportteacher/create [post]
// @Summary		Creates a supportteacher
// @Description	This api creates a supportteacher and returns its id
// @Tags		SupportTeacher
// @Accept		json
// @Produce		json
// @Param		supportteacher body user_service.CreateSupportTeacher true "supportteacher"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateSupportTeacher(c *gin.Context) {
	req := &us.CreateSupportTeacher{}

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

	resp, err := h.grpcClient.SupportTeacherService().Create(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get create supportteacher", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while creating supportteacher")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/supportteacher/getbyid/{id} [get]
// @Summary		Get by id a supportteacher
// @Description	This api get by id a supportteacher
// @Tags		SupportTeacher
// Accept		json
// @Produce		json
// @Param		id path string true "System User id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIdSupportTeacher(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.grpcClient.SupportTeacherService().GetByID(c.Request.Context(), &us.SupportTeacherPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while get get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting by id")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/supportteacher/getlist [get]
// @Summary		Get list a supportteacher
// @Description	This api get list a supportteacher
// @Tags		SupportTeacher
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListSupportTeacher(c *gin.Context) {
	req := &us.GetListSupportTeacherRequest{}

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

	resp, err := h.grpcClient.SupportTeacherService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list supportteacher")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/supportteacher/update [PUT]
// @Summary		Update a supportteacher
// @Description	This API updates a supportteacher
// @Tags		SupportTeacher
// @Accept		json
// @Produce		json
// @Param		supportteacher body user_service.UpdateSupportTeacherRequest true "SupportTeacher object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateSupportTeacher(c *gin.Context) {
	req := &us.UpdateSupportTeacherRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while update supportteacher")
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

	resp, err := h.grpcClient.SupportTeacherService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update supportteacher", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while ")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/supportteacher/delete [delete]
// @Summary		delete a supportteacher
// @Description	This api delete a supportteacher
// @Tags		SupportTeacher
// Accept		json
// @Produce		json
// @Param		supportteacher body user_service.SupportTeacherPrimaryKey true "supportteacher"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteSupportTeacher(c *gin.Context) {
	id := &us.SupportTeacherPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SupportTeacherService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while delete supportteacher")
		return
	}

	c.JSON(http.StatusOK, resp)
}
