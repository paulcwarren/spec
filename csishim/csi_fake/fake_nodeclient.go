// Code generated by counterfeiter. DO NOT EDIT.
package csi_fake

import (
	"sync"

	csi "github.com/paulcwarren/spec"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type FakeNodeClient struct {
	NodePublishVolumeStub        func(ctx context.Context, in *csi.NodePublishVolumeRequest, opts ...grpc.CallOption) (*csi.NodePublishVolumeResponse, error)
	nodePublishVolumeMutex       sync.RWMutex
	nodePublishVolumeArgsForCall []struct {
		ctx  context.Context
		in   *csi.NodePublishVolumeRequest
		opts []grpc.CallOption
	}
	nodePublishVolumeReturns struct {
		result1 *csi.NodePublishVolumeResponse
		result2 error
	}
	nodePublishVolumeReturnsOnCall map[int]struct {
		result1 *csi.NodePublishVolumeResponse
		result2 error
	}
	NodeUnpublishVolumeStub        func(ctx context.Context, in *csi.NodeUnpublishVolumeRequest, opts ...grpc.CallOption) (*csi.NodeUnpublishVolumeResponse, error)
	nodeUnpublishVolumeMutex       sync.RWMutex
	nodeUnpublishVolumeArgsForCall []struct {
		ctx  context.Context
		in   *csi.NodeUnpublishVolumeRequest
		opts []grpc.CallOption
	}
	nodeUnpublishVolumeReturns struct {
		result1 *csi.NodeUnpublishVolumeResponse
		result2 error
	}
	nodeUnpublishVolumeReturnsOnCall map[int]struct {
		result1 *csi.NodeUnpublishVolumeResponse
		result2 error
	}
	GetNodeIDStub        func(ctx context.Context, in *csi.GetNodeIDRequest, opts ...grpc.CallOption) (*csi.GetNodeIDResponse, error)
	getNodeIDMutex       sync.RWMutex
	getNodeIDArgsForCall []struct {
		ctx  context.Context
		in   *csi.GetNodeIDRequest
		opts []grpc.CallOption
	}
	getNodeIDReturns struct {
		result1 *csi.GetNodeIDResponse
		result2 error
	}
	getNodeIDReturnsOnCall map[int]struct {
		result1 *csi.GetNodeIDResponse
		result2 error
	}
	ProbeNodeStub        func(ctx context.Context, in *csi.ProbeNodeRequest, opts ...grpc.CallOption) (*csi.ProbeNodeResponse, error)
	probeNodeMutex       sync.RWMutex
	probeNodeArgsForCall []struct {
		ctx  context.Context
		in   *csi.ProbeNodeRequest
		opts []grpc.CallOption
	}
	probeNodeReturns struct {
		result1 *csi.ProbeNodeResponse
		result2 error
	}
	probeNodeReturnsOnCall map[int]struct {
		result1 *csi.ProbeNodeResponse
		result2 error
	}
	NodeGetCapabilitiesStub        func(ctx context.Context, in *csi.NodeGetCapabilitiesRequest, opts ...grpc.CallOption) (*csi.NodeGetCapabilitiesResponse, error)
	nodeGetCapabilitiesMutex       sync.RWMutex
	nodeGetCapabilitiesArgsForCall []struct {
		ctx  context.Context
		in   *csi.NodeGetCapabilitiesRequest
		opts []grpc.CallOption
	}
	nodeGetCapabilitiesReturns struct {
		result1 *csi.NodeGetCapabilitiesResponse
		result2 error
	}
	nodeGetCapabilitiesReturnsOnCall map[int]struct {
		result1 *csi.NodeGetCapabilitiesResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNodeClient) NodePublishVolume(ctx context.Context, in *csi.NodePublishVolumeRequest, opts ...grpc.CallOption) (*csi.NodePublishVolumeResponse, error) {
	fake.nodePublishVolumeMutex.Lock()
	ret, specificReturn := fake.nodePublishVolumeReturnsOnCall[len(fake.nodePublishVolumeArgsForCall)]
	fake.nodePublishVolumeArgsForCall = append(fake.nodePublishVolumeArgsForCall, struct {
		ctx  context.Context
		in   *csi.NodePublishVolumeRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("NodePublishVolume", []interface{}{ctx, in, opts})
	fake.nodePublishVolumeMutex.Unlock()
	if fake.NodePublishVolumeStub != nil {
		return fake.NodePublishVolumeStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.nodePublishVolumeReturns.result1, fake.nodePublishVolumeReturns.result2
}

func (fake *FakeNodeClient) NodePublishVolumeCallCount() int {
	fake.nodePublishVolumeMutex.RLock()
	defer fake.nodePublishVolumeMutex.RUnlock()
	return len(fake.nodePublishVolumeArgsForCall)
}

func (fake *FakeNodeClient) NodePublishVolumeArgsForCall(i int) (context.Context, *csi.NodePublishVolumeRequest, []grpc.CallOption) {
	fake.nodePublishVolumeMutex.RLock()
	defer fake.nodePublishVolumeMutex.RUnlock()
	return fake.nodePublishVolumeArgsForCall[i].ctx, fake.nodePublishVolumeArgsForCall[i].in, fake.nodePublishVolumeArgsForCall[i].opts
}

func (fake *FakeNodeClient) NodePublishVolumeReturns(result1 *csi.NodePublishVolumeResponse, result2 error) {
	fake.NodePublishVolumeStub = nil
	fake.nodePublishVolumeReturns = struct {
		result1 *csi.NodePublishVolumeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) NodePublishVolumeReturnsOnCall(i int, result1 *csi.NodePublishVolumeResponse, result2 error) {
	fake.NodePublishVolumeStub = nil
	if fake.nodePublishVolumeReturnsOnCall == nil {
		fake.nodePublishVolumeReturnsOnCall = make(map[int]struct {
			result1 *csi.NodePublishVolumeResponse
			result2 error
		})
	}
	fake.nodePublishVolumeReturnsOnCall[i] = struct {
		result1 *csi.NodePublishVolumeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) NodeUnpublishVolume(ctx context.Context, in *csi.NodeUnpublishVolumeRequest, opts ...grpc.CallOption) (*csi.NodeUnpublishVolumeResponse, error) {
	fake.nodeUnpublishVolumeMutex.Lock()
	ret, specificReturn := fake.nodeUnpublishVolumeReturnsOnCall[len(fake.nodeUnpublishVolumeArgsForCall)]
	fake.nodeUnpublishVolumeArgsForCall = append(fake.nodeUnpublishVolumeArgsForCall, struct {
		ctx  context.Context
		in   *csi.NodeUnpublishVolumeRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("NodeUnpublishVolume", []interface{}{ctx, in, opts})
	fake.nodeUnpublishVolumeMutex.Unlock()
	if fake.NodeUnpublishVolumeStub != nil {
		return fake.NodeUnpublishVolumeStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.nodeUnpublishVolumeReturns.result1, fake.nodeUnpublishVolumeReturns.result2
}

func (fake *FakeNodeClient) NodeUnpublishVolumeCallCount() int {
	fake.nodeUnpublishVolumeMutex.RLock()
	defer fake.nodeUnpublishVolumeMutex.RUnlock()
	return len(fake.nodeUnpublishVolumeArgsForCall)
}

func (fake *FakeNodeClient) NodeUnpublishVolumeArgsForCall(i int) (context.Context, *csi.NodeUnpublishVolumeRequest, []grpc.CallOption) {
	fake.nodeUnpublishVolumeMutex.RLock()
	defer fake.nodeUnpublishVolumeMutex.RUnlock()
	return fake.nodeUnpublishVolumeArgsForCall[i].ctx, fake.nodeUnpublishVolumeArgsForCall[i].in, fake.nodeUnpublishVolumeArgsForCall[i].opts
}

func (fake *FakeNodeClient) NodeUnpublishVolumeReturns(result1 *csi.NodeUnpublishVolumeResponse, result2 error) {
	fake.NodeUnpublishVolumeStub = nil
	fake.nodeUnpublishVolumeReturns = struct {
		result1 *csi.NodeUnpublishVolumeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) NodeUnpublishVolumeReturnsOnCall(i int, result1 *csi.NodeUnpublishVolumeResponse, result2 error) {
	fake.NodeUnpublishVolumeStub = nil
	if fake.nodeUnpublishVolumeReturnsOnCall == nil {
		fake.nodeUnpublishVolumeReturnsOnCall = make(map[int]struct {
			result1 *csi.NodeUnpublishVolumeResponse
			result2 error
		})
	}
	fake.nodeUnpublishVolumeReturnsOnCall[i] = struct {
		result1 *csi.NodeUnpublishVolumeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) GetNodeID(ctx context.Context, in *csi.GetNodeIDRequest, opts ...grpc.CallOption) (*csi.GetNodeIDResponse, error) {
	fake.getNodeIDMutex.Lock()
	ret, specificReturn := fake.getNodeIDReturnsOnCall[len(fake.getNodeIDArgsForCall)]
	fake.getNodeIDArgsForCall = append(fake.getNodeIDArgsForCall, struct {
		ctx  context.Context
		in   *csi.GetNodeIDRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("GetNodeID", []interface{}{ctx, in, opts})
	fake.getNodeIDMutex.Unlock()
	if fake.GetNodeIDStub != nil {
		return fake.GetNodeIDStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getNodeIDReturns.result1, fake.getNodeIDReturns.result2
}

func (fake *FakeNodeClient) GetNodeIDCallCount() int {
	fake.getNodeIDMutex.RLock()
	defer fake.getNodeIDMutex.RUnlock()
	return len(fake.getNodeIDArgsForCall)
}

func (fake *FakeNodeClient) GetNodeIDArgsForCall(i int) (context.Context, *csi.GetNodeIDRequest, []grpc.CallOption) {
	fake.getNodeIDMutex.RLock()
	defer fake.getNodeIDMutex.RUnlock()
	return fake.getNodeIDArgsForCall[i].ctx, fake.getNodeIDArgsForCall[i].in, fake.getNodeIDArgsForCall[i].opts
}

func (fake *FakeNodeClient) GetNodeIDReturns(result1 *csi.GetNodeIDResponse, result2 error) {
	fake.GetNodeIDStub = nil
	fake.getNodeIDReturns = struct {
		result1 *csi.GetNodeIDResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) GetNodeIDReturnsOnCall(i int, result1 *csi.GetNodeIDResponse, result2 error) {
	fake.GetNodeIDStub = nil
	if fake.getNodeIDReturnsOnCall == nil {
		fake.getNodeIDReturnsOnCall = make(map[int]struct {
			result1 *csi.GetNodeIDResponse
			result2 error
		})
	}
	fake.getNodeIDReturnsOnCall[i] = struct {
		result1 *csi.GetNodeIDResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) ProbeNode(ctx context.Context, in *csi.ProbeNodeRequest, opts ...grpc.CallOption) (*csi.ProbeNodeResponse, error) {
	fake.probeNodeMutex.Lock()
	ret, specificReturn := fake.probeNodeReturnsOnCall[len(fake.probeNodeArgsForCall)]
	fake.probeNodeArgsForCall = append(fake.probeNodeArgsForCall, struct {
		ctx  context.Context
		in   *csi.ProbeNodeRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("ProbeNode", []interface{}{ctx, in, opts})
	fake.probeNodeMutex.Unlock()
	if fake.ProbeNodeStub != nil {
		return fake.ProbeNodeStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.probeNodeReturns.result1, fake.probeNodeReturns.result2
}

func (fake *FakeNodeClient) ProbeNodeCallCount() int {
	fake.probeNodeMutex.RLock()
	defer fake.probeNodeMutex.RUnlock()
	return len(fake.probeNodeArgsForCall)
}

func (fake *FakeNodeClient) ProbeNodeArgsForCall(i int) (context.Context, *csi.ProbeNodeRequest, []grpc.CallOption) {
	fake.probeNodeMutex.RLock()
	defer fake.probeNodeMutex.RUnlock()
	return fake.probeNodeArgsForCall[i].ctx, fake.probeNodeArgsForCall[i].in, fake.probeNodeArgsForCall[i].opts
}

func (fake *FakeNodeClient) ProbeNodeReturns(result1 *csi.ProbeNodeResponse, result2 error) {
	fake.ProbeNodeStub = nil
	fake.probeNodeReturns = struct {
		result1 *csi.ProbeNodeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) ProbeNodeReturnsOnCall(i int, result1 *csi.ProbeNodeResponse, result2 error) {
	fake.ProbeNodeStub = nil
	if fake.probeNodeReturnsOnCall == nil {
		fake.probeNodeReturnsOnCall = make(map[int]struct {
			result1 *csi.ProbeNodeResponse
			result2 error
		})
	}
	fake.probeNodeReturnsOnCall[i] = struct {
		result1 *csi.ProbeNodeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) NodeGetCapabilities(ctx context.Context, in *csi.NodeGetCapabilitiesRequest, opts ...grpc.CallOption) (*csi.NodeGetCapabilitiesResponse, error) {
	fake.nodeGetCapabilitiesMutex.Lock()
	ret, specificReturn := fake.nodeGetCapabilitiesReturnsOnCall[len(fake.nodeGetCapabilitiesArgsForCall)]
	fake.nodeGetCapabilitiesArgsForCall = append(fake.nodeGetCapabilitiesArgsForCall, struct {
		ctx  context.Context
		in   *csi.NodeGetCapabilitiesRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("NodeGetCapabilities", []interface{}{ctx, in, opts})
	fake.nodeGetCapabilitiesMutex.Unlock()
	if fake.NodeGetCapabilitiesStub != nil {
		return fake.NodeGetCapabilitiesStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.nodeGetCapabilitiesReturns.result1, fake.nodeGetCapabilitiesReturns.result2
}

func (fake *FakeNodeClient) NodeGetCapabilitiesCallCount() int {
	fake.nodeGetCapabilitiesMutex.RLock()
	defer fake.nodeGetCapabilitiesMutex.RUnlock()
	return len(fake.nodeGetCapabilitiesArgsForCall)
}

func (fake *FakeNodeClient) NodeGetCapabilitiesArgsForCall(i int) (context.Context, *csi.NodeGetCapabilitiesRequest, []grpc.CallOption) {
	fake.nodeGetCapabilitiesMutex.RLock()
	defer fake.nodeGetCapabilitiesMutex.RUnlock()
	return fake.nodeGetCapabilitiesArgsForCall[i].ctx, fake.nodeGetCapabilitiesArgsForCall[i].in, fake.nodeGetCapabilitiesArgsForCall[i].opts
}

func (fake *FakeNodeClient) NodeGetCapabilitiesReturns(result1 *csi.NodeGetCapabilitiesResponse, result2 error) {
	fake.NodeGetCapabilitiesStub = nil
	fake.nodeGetCapabilitiesReturns = struct {
		result1 *csi.NodeGetCapabilitiesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) NodeGetCapabilitiesReturnsOnCall(i int, result1 *csi.NodeGetCapabilitiesResponse, result2 error) {
	fake.NodeGetCapabilitiesStub = nil
	if fake.nodeGetCapabilitiesReturnsOnCall == nil {
		fake.nodeGetCapabilitiesReturnsOnCall = make(map[int]struct {
			result1 *csi.NodeGetCapabilitiesResponse
			result2 error
		})
	}
	fake.nodeGetCapabilitiesReturnsOnCall[i] = struct {
		result1 *csi.NodeGetCapabilitiesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nodePublishVolumeMutex.RLock()
	defer fake.nodePublishVolumeMutex.RUnlock()
	fake.nodeUnpublishVolumeMutex.RLock()
	defer fake.nodeUnpublishVolumeMutex.RUnlock()
	fake.getNodeIDMutex.RLock()
	defer fake.getNodeIDMutex.RUnlock()
	fake.probeNodeMutex.RLock()
	defer fake.probeNodeMutex.RUnlock()
	fake.nodeGetCapabilitiesMutex.RLock()
	defer fake.nodeGetCapabilitiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNodeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ csi.NodeClient = new(FakeNodeClient)
