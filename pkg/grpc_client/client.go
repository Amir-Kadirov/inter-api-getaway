package grpc_client

import (
	"backend_course/branch_api_gateway/config"
	sc "backend_course/branch_api_gateway/genproto/schedule_service"
	pc "backend_course/branch_api_gateway/genproto/user_service"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClientI ...
type GrpcClientI interface {
	BranchBranch() pc.BranchServiceClient
	StudentService() sc.StudentServiceClient
	TeacherService() pc.TeacherServiceClient
	SupportTeacherService() pc.SupportTeacherServiceClient
	AdminService() pc.AdminServiceClient
	ManagerService() pc.ManagerServiceClient
	GroupService() sc.GroupServiceClient
	LessonService() sc.LessonServiceClient
	TaskService() sc.TaskServiceClient
	JournalService() sc.JournalServiceClient
	ScheduleService() sc.ScheduleServiceClient
	EventService() sc.EventServiceClient
	GroupManyService() sc.GroupManyServiceClient
	AttendanceService() sc.AttendanceServiceClient
}

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
			"group_service": sc.NewGroupServiceClient(connSchedule),
			"task_service": sc.NewTaskServiceClient(connSchedule),
			"lesson_service": sc.NewLessonServiceClient(connSchedule),
			"journal_service": sc.NewJournalServiceClient(connSchedule),
			"schedule_service": sc.NewScheduleServiceClient(connSchedule),
			"event_service": sc.NewEventServiceClient(connSchedule),
			"group_many_service": sc.NewGroupManyServiceClient(connSchedule),
			"attendance_service": sc.NewAttendanceServiceClient(connSchedule),
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

func (g *GrpcClient) GroupService() sc.GroupServiceClient {
	return g.connections["group_service"].(sc.GroupServiceClient)
}

func (g *GrpcClient) TaskService() sc.TaskServiceClient {
	return g.connections["task_service"].(sc.TaskServiceClient)
}

func (g *GrpcClient) LessonService() sc.LessonServiceClient {
	return g.connections["lesson_service"].(sc.LessonServiceClient)
}

func (g *GrpcClient) JournalService() sc.JournalServiceClient {
	return g.connections["journal_service"].(sc.JournalServiceClient)
}

func (g *GrpcClient) ScheduleService() sc.ScheduleServiceClient {
	return g.connections["schedule_service"].(sc.ScheduleServiceClient)
}

func (g *GrpcClient) EventService() sc.EventServiceClient {
	return g.connections["event_service"].(sc.EventServiceClient)
}

func (g *GrpcClient) GroupManyService() sc.GroupManyServiceClient {
	return g.connections["group_many_service"].(sc.GroupManyServiceClient)
}

func (g *GrpcClient) AttendanceService() sc.AttendanceServiceClient {
	return g.connections["attendance_service"].(sc.AttendanceServiceClient)
}