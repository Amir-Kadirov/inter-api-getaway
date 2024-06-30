package handler

import (
	"backend_course/branch_api_gateway/genproto/schedule_service"
	us "backend_course/branch_api_gateway/genproto/schedule_service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/groupmany/create [post]
// @Summary		Creates a GroupMany
// @Description	This api creates a GroupMany and returns its id
// @Tags		GroupMany
// @Accept		json
// @Produce		json
// @Param		GroupMany body schedule_service.CreateGroupMany true "GroupMany"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateGroupMany(c *gin.Context) {
	var req us.CreateGroupMany

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	_,err:=h.grpcClient.GroupService().GetByID(c.Request.Context(),&schedule_service.GroupPrimaryKey{Id: req.GroupId})
	if err!=nil {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong group id"), "wrong group id")
		return
	}

	_,err=h.grpcClient.ScheduleService().GetByID(c.Request.Context(),&schedule_service.SchedulePrimaryKey{Id: req.ScheduleId})
	if err!=nil {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong schedule id"), "wrong schedule id")
		return
	}

	_,err=h.grpcClient.JournalService().GetByID(c.Request.Context(),&schedule_service.JournalPrimaryKey{Id: req.JournalId})
	if err!=nil {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong journal id"), "wrong journal id")
		return
	}

	resp, err := h.grpcClient.GroupManyService().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while create groupmany")
		return
	}

	c.JSON(http.StatusOK, resp)
}