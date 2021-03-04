package mig

import (
	"context"
	"time"

	pb "github.com/pegasus-cloud/nvml_client/protos"
)

// ListComputeInstanceProfile ...
func ListComputeInstanceProfile(in *pb.GpuInstanceInfo) (*pb.ComputeInstanceProfileInfos, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return pb.NewComputeInstanceControllerClient(use().conn).ListComputeInstanceProfile(ctx, in)
}
