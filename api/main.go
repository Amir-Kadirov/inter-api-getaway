package api

import (
	"net/http"

	_ "backend_course/branch_api_gateway/api/docs" //for swagger
	"backend_course/branch_api_gateway/api/handler"
	"backend_course/branch_api_gateway/config"
	"backend_course/branch_api_gateway/pkg/grpc_client"
	"backend_course/branch_api_gateway/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")
	// config.AllowOrigins = cnf.Cfg.AllowOrigins
	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})


	// User Service
	r.POST("/v1/branch/create", handler.CreateBranch)
	r.GET("/v1/branch/getbyid/:id", handler.GetByID)
	r.GET("/v1/branch/getlist", handler.GetListBranch)
	r.PUT("/v1/branch/updatebranch", handler.Update)
	r.DELETE("/v1/branch/delete", handler.Delete)

	r.POST("/v1/Student/create", handler.CreateStudent)
	r.GET("/v1/Student/getbyid/:id", handler.GetByIdStudent)
	r.GET("/v1/Student/getlist", handler.GetListStudent)
	r.PUT("/v1/Student/update", handler.UpdateStudent)
	r.DELETE("/v1/Student/delete", handler.DeleteStudent)

	r.POST("/v1/admin/create", handler.CreateAdmin)
	r.GET("/v1/admin/getbyid/:id", handler.GetByIdAdmin)
	r.GET("/v1/admin/getlist", handler.GetListAdmin)
	r.PUT("/v1/admin/update", handler.UpdateAdmin)
	r.DELETE("/v1/admin/delete", handler.DeleteAdmin)
	r.GET("/v1/admin/report",handler.AdminReport)

	r.POST("/v1/manager/create", handler.CreateManager)
	r.GET("/v1/manager/getbyid/:id", handler.GetByIdManager)
	r.GET("/v1/manager/getlist", handler.GetListManager)
	r.PUT("/v1/manager/update", handler.UpdateManager)
	r.DELETE("/v1/manager/delete", handler.DeleteManager)

	r.POST("/v1/teacher/create", handler.CreateTeacher)
	r.GET("/v1/teacher/getbyid/:id", handler.GetByIdTeacher)
	r.GET("/v1/teacher/getlist", handler.GetListTeacher)
	r.PUT("/v1/teacher/update", handler.UpdateTeacher)
	r.DELETE("/v1/teacher/delete", handler.DeleteTeacher)

	r.POST("/v1/supportteacher/create", handler.CreateSupportTeacher)
	r.GET("/v1/supportteacher/getbyid/:id", handler.GetByIdSupportTeacher)
	r.GET("/v1/supportteacher/getlist", handler.GetListSupportTeacher)
	r.PUT("/v1/supportteacher/update", handler.UpdateSupportTeacher)
	r.DELETE("/v1/supportteacher/delete", handler.DeleteSupportTeacher)

	// Schedule Service
	r.POST("/v1/event/create", handler.CreateEvent)
	r.GET("/v1/event/getbyid/:id", handler.GetByIDEvent)
	r.GET("/v1/event/getlist", handler.GetListEvent)
	r.PUT("/v1/event/update", handler.UpdateEvent)
	r.DELETE("/v1/event/delete", handler.DeleteEvent)
	r.POST("/v1/event/registerevent",handler.RegisterEvent)

	r.POST("/v1/group/create", handler.CreateGroup)
	r.GET("/v1/group/getbyid/:id", handler.GetByIDGroup)
	r.GET("/v1/group/getlist", handler.GetListGroup)
	r.PUT("/v1/group/update", handler.UpdateGroup)
	r.DELETE("/v1/group/delete", handler.DeleteGroup)

	r.POST("/v1/journal/create", handler.CreateJournal)
	r.GET("/v1/journal/getbyid/:id", handler.GetByIdJournal)
	r.GET("/v1/journal/getlist", handler.GetListJournal)
	r.PUT("/v1/journal/update", handler.UpdateJournal)
	r.DELETE("/v1/journal/delete", handler.DeleteJournal)

	r.POST("/v1/lesson/create", handler.CreateLesson)
	r.GET("/v1/lesson/getbyid/:id", handler.GetByIDLesson)
	r.GET("/v1/lesson/getlist", handler.GetListLesson)
	r.PUT("/v1/lesson/update", handler.UpdateLesson)
	r.DELETE("/v1/lesson/delete", handler.DeleteLesson)

	r.POST("/v1/schedule/create", handler.CreateSchedule)
	r.GET("/v1/schedule/getbyid/:id", handler.GetByIDSchedule)
	r.GET("/v1/schedule/getlist", handler.GetListSchedule)
	r.PUT("/v1/schedule/update", handler.UpdateSchedule)
	r.DELETE("/v1/schedule/delete", handler.DeleteSchedule)

	r.POST("/v1/task/create", handler.CreateTask)
	r.GET("/v1/task/getbyid/:id", handler.GetByIDTask)
	r.GET("/v1/task/getlist", handler.GetListTask)
	r.PUT("/v1/task/update", handler.UpdateTask)
	r.DELETE("/v1/task/delete", handler.DeleteTask)
	r.POST("/v1/task/dotask",handler.DoTask)

	r.POST("/v1/attendance/create", handler.CreateAttendance)
	r.GET("/v1/attendance/getbyid/:id", handler.GetByIDAttendance)
	r.GET("/v1/attendance/getlist", handler.GetListAttendance)


	r.POST("/v1/groupmany/create", handler.CreateGroupMany)

	r.GET("/v1/report/schedulemonth",handler.ScheduleMonth)
	r.GET("/v1/report/teacher",handler.TeacherReport)
	r.GET("/v1/report/admin",handler.AdminReport)
	r.GET("/v1/report/student",handler.StudentReport)
	r.GET("/v1/report/supportteacher",handler.SupportTeacherReport)

	// Shipper endpoints
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}