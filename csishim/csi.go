package csishim

import (
	csi "github.com/paulcwarren/spec"
	"google.golang.org/grpc"
)

//go:generate counterfeiter -o csi_fake/fake_csi.go . Csi
type Csi interface {
	NewNodeClient(cc *grpc.ClientConn) csi.NodeClient
	NewControllerClient(cc *grpc.ClientConn) csi.ControllerClient
}
