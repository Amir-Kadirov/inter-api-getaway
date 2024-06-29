package grpc_client

import (
	pc "backend_course/branch_api_gateway/genproto/user_service"
	sc "backend_course/branch_api_gateway/genproto/schedule_service"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"backend_course/branch_api_gateway/config"
)

// GrpcClientI ...
type GrpcClientI interface {
	BranchBranch() pc.BranchServiceClient
	StudentService() sc.StudentServiceClient
	TeacherService() pc.TeacherServiceClient
	SupportTeacherService() pc.SupportTeacherServiceClient
	AdminService() pc.AdminServiceClient
	ManagerService() pc.ManagerServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.UserBranchHost, cfg.UserBranchPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",
			cfg.UserBranchHost, cfg.UserBranchPort, err)
	}

	connSchedule, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.ScheduleBranchHost, cfg.ScheduleBranchPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("schedule service dial host: %s port:%s err: %s",
			cfg.ScheduleBranchHost, cfg.ScheduleBranchPort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"branch_service":      pc.NewBranchServiceClient(connUser),
			"student_service":     sc.NewStudentServiceClient(connSchedule),
			"teacher_service":     pc.NewTeacherServiceClient(connUser),
			"support_teacher_service": pc.NewSupportTeacherServiceClient(connUser),	
			"administration_service": pc.NewAdminServiceClient(connUser),
			"manager_service": pc.NewManagerServiceClient(connUser),
		},
	}, nil
}

func (g *GrpcClient) BranchBranch() pc.BranchServiceClient {
	return g.connections["branch_service"].(pc.BranchServiceClient)
}

func (g *GrpcClient) StudentService() sc.StudentServiceClient {
	return g.connections["student_service"].(sc.StudentServiceClient)
}

func (g *GrpcClient) TeacherService() pc.TeacherServiceClient {
	return g.connections["teacher_service"].(pc.TeacherServiceClient)
}

func (g *GrpcClient) SupportTeacherService() pc.SupportTeacherServiceClient {
	return g.connections["support_teacher_service"].(pc.SupportTeacherServiceClient)
}

func (g *GrpcClient) AdminService() pc.AdminServiceClient {
	return g.connections["administration_service"].(pc.AdminServiceClient)
}

func (g *GrpcClient) ManagerService() pc.ManagerServiceClient {
	return g.connections["manager_service"].(pc.ManagerServiceClient)
}