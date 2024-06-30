package handler

// import (
// 	us "backend_course/branch_api_gateway/genproto/auth_service"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // @Router		/v1/teacher-auth/loginbypas [post]
// // @Summary		Login by Password a Teacher
// // @Description	This api logging a Teacher
// // @Tags		Teacher-Auth
// // @Accept		json
// // @Produce		json
// // @Param		Teacher body auth_service.TeacherLoginRequest true "Teacher"
// // @Success		200  {object}  models.ResponseSuccess
// // @Failure		400  {object}  models.ResponseError
// // @Failure		404  {object}  models.ResponseError
// // @Failure		500  {object}  models.ResponseError
// func (h *handler) LoginTeacherByPas(c *gin.Context) {
// 	var req us.TeacherLoginRequest	

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
// 		return
// 	}

// 	resp, err := h.service.Teacher().TeacherLoginByPassword(c.Request.Context(),&req)
// 	if err != nil {
// 		handleGrpcErrWithDescription(c, h.log, err, "error while create teacher-auth")
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp)
// }
