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

	// Shipper endpoints
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}