package csishim

import (
	"google.golang.org/grpc"

	spec "github.com/paulcwarren/spec"
	"code.cloudfoundry.org/goshims/grpcshim"
)

type CsiShim struct{}

func (c *CsiShim) NewNodeClient(conn grpcshim.ClientConn) spec.NodeClient {
	return spec.NewNodeClient(conn.(*grpcshim.ClientConnShim).ClientConn)
}

func (c *CsiShim) NewControllerClient(conn *grpc.ClientConn) spec.ControllerClient {
	return spec.NewControllerClient(conn)
}
