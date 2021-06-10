package main

//import "C"
import (
	CC "./Consul"
	TokenService "./TokenService"
	pb "./pb"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
)


var ServeIP = "127.0.0.1"
var ServePort = 1234
var metricPort = 1235

func main() {
	flag.Parse()
	ServeAddress := fmt.Sprintf("%v:%v", ServeIP, ServePort)
	lis, err := net.Listen("tcp", ServeAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}


	//新建gRPCserver, 注册gRPC拦截器供prometheus使用
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	TokenService := new(TokenService.TokenService)
	//注册服务
	pb.RegisterTokenServiceServer(grpcServer, TokenService)
	// 在所有的拦截器注册之后，确保所有的prometheus指标都被初始化。
	grpc_prometheus.Register(grpcServer)
	// 注册prometheus metrics处理器
	//http.HandleFunc("/metrics", promhttp.Handler)
	http.Handle("/metrics", promhttp.Handler())
	//提供http服务

	//consul注册
	register:= CC.NewConsulRegister()
	register.Name = "TokenService"
	register.Port = ServePort
	register.IP = ServeIP
	if err := register.Register(); err !=nil{
		panic(err)
	}

	grpc_health_v1.RegisterHealthServer(grpcServer, &CC.HealthImpl{Status:grpc_health_v1.HealthCheckResponse_SERVING})


	go func() {
		MetricsAdders := fmt.Sprintf("%v:%v", ServeIP, metricPort)
		err := http.ListenAndServe(MetricsAdders, nil)
		if err != nil {
			fmt.Println("error: ", err.Error())
		}
	}()


	//开始服务

	grpcServer.Serve(lis)



}
