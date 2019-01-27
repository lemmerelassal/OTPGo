package main

import (
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	backendp "otpgo/pkg/proto/backend"
	backendf "otpgo/pkg/svc/backend"
)

const (
	port = ":4711"
)

var (
	logger = logrus.NewEntry(logrus.StandardLogger())
)

func main() {

	// setup logging
	logrus.SetOutput(os.Stdout)
	grpc.EnableTracing = true
	grpc_logrus.ReplaceGrpcLogger(logger)

	logger.Infoln("Starting Server...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logger),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(logger),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_recovery.StreamServerInterceptor(),
		))

	if err != nil {
		logger.Fatalln("Instantiating Datastore failed", err)
	}

	backendp.RegisterOTPGOBackendServer(s, backendf.New())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logger.Errorln("failed to serve: ", err)
	}
}
