package kasm

import (
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

type Service struct {
	runtimeapi.UnimplementedImageServiceServer
	runtimeapi.UnimplementedRuntimeServiceServer
}