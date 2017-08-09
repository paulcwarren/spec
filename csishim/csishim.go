package csishim

import (
	"google.golang.org/grpc"

	"github.com/paulcwarren/spec"
)

type CsiShim struct{}

func (csi *CsiShim) NewNodeClient(conn *grpc.ClientConn) csi.NodeClient {
	return csi.NewNodeClient(conn)
}
