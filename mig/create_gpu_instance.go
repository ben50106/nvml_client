package mig

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/pegasus-cloud/nvml_client/protos"
)

// CreateGpuInstance ...
func CreateGpuInstance(in *pb.GpuInstanceInfo) (*empty.Empty, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return pb.NewGpuInstanceControllerClient(use().conn).CreateGpuInstance(ctx, in)
}
