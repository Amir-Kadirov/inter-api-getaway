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

// @Router		/v1/admin/create [post]
// @Summary		Creates a admin
// @Description	This api creates a admin and returns its id
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Param		admin body user_service.CreateAdmin true "admin"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateAdmin(c *gin.Context) {
	req := &us.CreateAdmin{}

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

	resp, err := h.grpcClient.AdminService().Create(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get create admin", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while creating admin")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/admin/getbyid/{id} [get]
// @Summary		Get by id a admin
// @Description	This api get by id a admin
// @Tags		Admin
// Accept		json
// @Produce		json
// @Param		id path string true "System User id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIdAdmin(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.grpcClient.AdminService().GetByID(c.Request.Context(), &us.AdminPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while get get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting by id")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/admin/getlist [get]
// @Summary		Get list a admin
// @Description	This api get list a admin
// @Tags		Admin
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListAdmin(c *gin.Context) {
	req := &us.GetListAdminRequest{}

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

	resp, err := h.grpcClient.AdminService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list admin")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/admin/update [PUT]
// @Summary		Update a admin
// @Description	This API updates a admin
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Param		admin body user_service.UpdateAdminRequest true "Admin object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateAdmin(c *gin.Context) {
	req := &us.UpdateAdminRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while update admin")
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

	resp, err := h.grpcClient.AdminService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update admin", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while ")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/admin/delete [delete]
// @Summary		delete a admin
// @Description	This api delete a admin
// @Tags		Admin
// Accept		json
// @Produce		json
// @Param		admin body user_service.AdminPrimaryKey true "admin"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteAdmin(c *gin.Context) {
	id := &us.AdminPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.AdminService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while delete admin")
		return
	}

	c.JSON(http.StatusOK, resp)
}
