package csishim

import (
	"google.golang.org/grpc"

	spec "github.com/paulcwarren/spec"
)

type CsiShim struct{}

func (csi *CsiShim) NewNodeClient(conn *grpc.ClientConn) spec.NodeClient {
	return csi.NewNodeClient(conn)
}

func (csi *CsiShim) NewControllerClient(conn *grpc.ClientConn) spec.ControllerClient {
	return csi.NewControllerClient(conn)
}
