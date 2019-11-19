package main

import (
	// "fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"k8s.io/klog"
	"google.golang.org/grpc"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"

	"github.com/arccoza/kasm/pkg/kasm"
)

const (
	unixSockPath = "/tmp/kasm.sock"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	if err := syscall.Unlink(unixSockPath); err != nil && !os.IsNotExist(err) {
		klog.Fatalf("Failed to remove %q: %v", unixSockPath, err)
	}
	defer os.Remove(unixSockPath)

	sock, err := net.Listen("unix", unixSockPath)
	if err != nil {
		klog.Fatalf("Failed to listen on %q: %v", unixSockPath, err)
	}
	defer sock.Close()

	grpcServer := grpc.NewServer()

	imgService := &kasm.Service{}
	runService := &kasm.Service{}

	runtimeapi.RegisterImageServiceServer(grpcServer, imgService)
	runtimeapi.RegisterRuntimeServiceServer(grpcServer, runService)

	klog.Infof("Starting to serve on %q", unixSockPath)
	go grpcServer.Serve(sock)

	<-sigCh
	klog.Infof("kasm service exiting...")
}
