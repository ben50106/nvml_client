package mig

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/pegasus-cloud/nvml_client/protos"
)

// CreateComputeInstance ...
func CreateComputeInstance(in *pb.ComputeInstanceInfo) (*empty.Empty, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return pb.NewComputeInstanceControllerClient(use().conn).CreateComputeInstance(ctx, in)
}
