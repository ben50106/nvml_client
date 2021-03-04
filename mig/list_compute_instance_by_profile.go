package mig

import (
	"context"
	"time"

	pb "github.com/pegasus-cloud/nvml_client/protos"
)

// ListComputeInstanceByProfile ...
func ListComputeInstanceByProfile(in *pb.ComputeInstanceInfo) (*pb.ComputeInstanceInfos, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return pb.NewComputeInstanceControllerClient(use().conn).ListComputeInstanceByProfile(ctx, in)
}
