/*
Copyright © 2019 Adrien de Pierres <adrien@arccoza.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"k8s.io/klog"
	"google.golang.org/grpc"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"github.com/arccoza/kasm/pkg/kasm"
)

const (
	unixSockPath = "/tmp/kasm.sock"
)

func run(cmd *cobra.Command, args []string) {
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

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the KASM CRI service.",
	Long: ``,
	Run: run,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
