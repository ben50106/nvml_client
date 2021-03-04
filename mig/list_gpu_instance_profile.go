package mig

import (
	"context"
	"time"

	pb "github.com/pegasus-cloud/nvml_client/protos"
)

// ListGpuInstanceProfile ...
func ListGpuInstanceProfile(in *pb.GpuInstanceInfo) (*pb.GpuInstanceProfileInfos, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return pb.NewGpuInstanceControllerClient(use().conn).ListGpuInstanceProfile(ctx, in)
}
