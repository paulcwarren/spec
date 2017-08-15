package csishim

import (
	csi "github.com/paulcwarren/spec"
	"google.golang.org/grpc"
	"code.cloudfoundry.org/goshims/grpcshim"
)

//go:generate counterfeiter -o csi_fake/fake_csi.go . Csi
type Csi interface {
	NewNodeClient(cc grpcshim.ClientConn) csi.NodeClient
	NewControllerClient(cc *grpc.ClientConn) csi.ControllerClient
}
