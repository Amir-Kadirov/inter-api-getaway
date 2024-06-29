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

// @Router		/v1/manager/create [post]
// @Summary		Creates a manager
// @Description	This api creates a manager and returns its id
// @Tags		Manager
// @Accept		json
// @Produce		json
// @Param		manager body user_service.CreateManager true "manager"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateManager(c *gin.Context) {
	req := &us.CreateManager{}

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

	resp, err := h.grpcClient.ManagerService().Create(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get create manager", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while create manager")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/manager/getbyid/{id} [get]
// @Summary		Get by id a manager
// @Description	This api get by id a manager
// @Tags		Manager
// @Produce		json
// @Param		id path string true "Manager id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIdManager(c *gin.Context) {
	id := c.Param("id")


	resp, err := h.grpcClient.ManagerService().GetByID(c.Request.Context(), &us.ManagerPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while  delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while delete")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/manager/getlist [get]
// @Summary		Get list a manager
// @Description	This api get list a manager
// @Tags		Manager
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListManager(c *gin.Context) {
	req := &us.GetListManagerRequest{}

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

	resp, err := h.grpcClient.ManagerService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list manager")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/manager/update [PUT]
// @Summary		Update a manager
// @Description	This API updates a manager
// @Tags		Manager
// @Accept		json
// @Produce		json
// @Param		manager body user_service.UpdateManagerRequest true "Manager object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateManager(c *gin.Context) {
	req := &us.UpdateManagerRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while update manager")
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

	resp, err := h.grpcClient.ManagerService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update manager", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while ")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/manager/delete [delete]
// @Summary		delete a manager
// @Description	This api delete a manager
// @Tags		Manager
// Accept		json
// @Produce		json
// @Param		manager body user_service.ManagerPrimaryKey true "manager"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteManager(c *gin.Context) {
	id := &us.ManagerPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.ManagerService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while delete manager")
		return
	}

	c.JSON(http.StatusOK, resp)
}
