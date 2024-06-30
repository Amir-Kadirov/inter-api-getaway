package handler

import (
	us "backend_course/branch_api_gateway/genproto/schedule_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/journal/create [post]
// @Summary		Creates a journal
// @Description	This api creates a journal and returns its id
// @Tags		Journal
// @Accept		json
// @Produce		json
// @Param		journal body schedule_service.CreateJournal true "journal"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateJournal(c *gin.Context) {
	req := &us.CreateJournal{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.JournalService().Create(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get create journal", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while creating journal")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/journal/getbyid/{id} [get]
// @Summary		Get by id a journal
// @Description	This api get by id a journal
// @Tags		Journal
// Accept		json
// @Produce		json
// @Param		id path string true "System User id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIdJournal(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.grpcClient.JournalService().GetByID(c.Request.Context(), &us.JournalPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("error while get get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting by id")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/journal/getlist [get]
// @Summary		Get list a journal
// @Description	This api get list a journal
// @Tags		Journal
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListJournal(c *gin.Context) {
	req := &us.GetListJournalRequest{}

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

	resp, err := h.grpcClient.JournalService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list journal")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/journal/update [PUT]
// @Summary		Update a journal
// @Description	This API updates a journal
// @Tags		Journal
// @Accept		json
// @Produce		json
// @Param		journal body schedule_service.UpdateJournalRequest true "Journal object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateJournal(c *gin.Context) {
	req := &us.UpdateJournalRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while update journal")
		return
	}

	resp, err := h.grpcClient.JournalService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update journal", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while ")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/journal/delete [delete]
// @Summary		delete a journal
// @Description	This api delete a journal
// @Tags		Journal
// Accept		json
// @Produce		json
// @Param		journal body schedule_service.JournalPrimaryKey true "journal"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteJournal(c *gin.Context) {
	id := &us.JournalPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.JournalService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while delete journal")
		return
	}

	c.JSON(http.StatusOK, resp)
}
