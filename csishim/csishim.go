package csishim

import (
	"google.golang.org/grpc"

	spec "github.com/paulcwarren/spec"
)

type CsiShim struct{}

func (c *CsiShim) NewNodeClient(conn *grpc.ClientConn) spec.NodeClient {
	return spec.NewNodeClient(conn)
}

func (c *CsiShim) NewControllerClient(conn *grpc.ClientConn) spec.ControllerClient {
	return spec.NewControllerClient(conn)
}
