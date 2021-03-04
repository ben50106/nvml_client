package mig

import (
	"context"
	"time"

	pb "github.com/pegasus-cloud/nvml_client/protos"
)

// ListGpuInstanceByProfile ...
func ListGpuInstanceByProfile(in *pb.GpuInstanceInfo) (*pb.GpuInstanceInfos, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return pb.NewGpuInstanceControllerClient(use().conn).ListGpuInstanceByProfile(ctx, in)
}
