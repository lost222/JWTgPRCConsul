package Consul

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/health/grpc_health_v1"
	"time"
)

var consulAddress = "127.0.0.1:8500"
var ServerIP = "127.0.0.1"


type HealthImpl struct {
	Status   grpc_health_v1.HealthCheckResponse_ServingStatus
	Reason string
}

func (h *HealthImpl) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return nil
}

func (h *HealthImpl) OffLine(reason string) {
	h.Status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	h.Reason = reason
	fmt.Println(reason)
}
func (h *HealthImpl) OnLine(reason string) {
	h.Status = grpc_health_v1.HealthCheckResponse_SERVING
	h.Reason = reason
	fmt.Println(reason)
}

func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status:h.Status,
	}, nil
}


//consul
// NewConsulRegister create a new consul register
func NewConsulRegister() *ConsulRegister {
	return &ConsulRegister{
		IP:                        ServerIP, //ServerIP
		Name:                           "unknown",
		Tag:                            []string{},
		Port:                           8082,
		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
		Interval:                       time.Duration(10) * time.Second,
	}
}

// ConsulRegister consul service register
type ConsulRegister struct {
	IP                        string
	Name                           string
	Tag                            []string
	Port                           int
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
}

// Register register service
func (r *ConsulRegister) Register() error {
	config := api.DefaultConfig()
	config.Address = consulAddress
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}
	agent := client.Agent()


	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", r.Name, r.IP, r.Port), // 服务节点的名称
		Name:    r.Name,  										  // 服务名称
		Tags:    r.Tag,                                          // tag，可以为空
		Port:    r.Port,                                         // 服务端口
		Address: r.IP, 											// 服务 IP
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval: r.Interval.String(), // 健康检查间隔
			GRPC:     fmt.Sprintf("%v:%v/%v", r.IP, r.Port, r.Name), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(), // 注销时间，相当于过期时间
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		return err
	}

	return nil
}



