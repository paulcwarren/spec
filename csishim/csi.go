package csishim

import (
	"google.golang.org/grpc"

	"github.com/paulcwarren/spec"
)

type Csi interface {
	NewNodeClient(cc *grpc.ClientConn) csi.NodeClient
}
