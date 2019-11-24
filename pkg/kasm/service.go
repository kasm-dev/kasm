package kasm

import (
	"context"

	// "k8s.io/klog"
	// codes "google.golang.org/grpc/codes"
	// status "google.golang.org/grpc/status"
	cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

type Service struct {
	cri.UnimplementedImageServiceServer
	cri.UnimplementedRuntimeServiceServer
}

var (
	version = cri.VersionResponse{
		Version:           "0.1.0",
		RuntimeName:       "kasm",
		RuntimeVersion:    "0.1.0",
		RuntimeApiVersion: "0.1.0",
	}
)

func (s *Service) Version(ctx context.Context, req *cri.VersionRequest) (*cri.VersionResponse, error) {
	return &version, nil
}

func (s *Service) RunPodSandbox(ctx context.Context, req *cri.RunPodSandboxRequest) (*cri.RunPodSandboxResponse, error) {
	return &cri.RunPodSandboxResponse{PodSandboxId:"kasm_podid_1"}, nil
}

// func (s *Service) StartContainer(ctx context.Context, req *cri.StartContainerRequest) (*cri.StartContainerResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method StartContainer not implemented")
// }

// func (s *Service) StopContainer(ctx context.Context, req *cri.StopContainerRequest) (*cri.StopContainerResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method StopContainer not implemented")
// }

// func (s *Service) PortForward(ctx context.Context, req *cri.PortForwardRequest) (*cri.PortForwardResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method PortForward not implemented")
// }
