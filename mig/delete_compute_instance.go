package mig

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/pegasus-cloud/nvml_client/protos"
)

// DeleteComputeInstance ...
func DeleteComputeInstance(in *pb.ComputeInstanceInfo) (*empty.Empty, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return pb.NewComputeInstanceControllerClient(use().conn).DeleteComputeInstance(ctx, in)
}
