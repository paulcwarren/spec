// This file was generated by counterfeiter
package csi_fake

import (
	"sync"

	csi "github.com/paulcwarren/spec"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type FakeControllerClient struct {
	CreateVolumeStub        func(ctx context.Context, in *csi.CreateVolumeRequest, opts ...grpc.CallOption) (*csi.CreateVolumeResponse, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		ctx  context.Context
		in   *csi.CreateVolumeRequest
		opts []grpc.CallOption
	}
	createVolumeReturns struct {
		result1 *csi.CreateVolumeResponse
		result2 error
	}
	DeleteVolumeStub        func(ctx context.Context, in *csi.DeleteVolumeRequest, opts ...grpc.CallOption) (*csi.DeleteVolumeResponse, error)
	deleteVolumeMutex       sync.RWMutex
	deleteVolumeArgsForCall []struct {
		ctx  context.Context
		in   *csi.DeleteVolumeRequest
		opts []grpc.CallOption
	}
	deleteVolumeReturns struct {
		result1 *csi.DeleteVolumeResponse
		result2 error
	}
	ControllerPublishVolumeStub        func(ctx context.Context, in *csi.ControllerPublishVolumeRequest, opts ...grpc.CallOption) (*csi.ControllerPublishVolumeResponse, error)
	controllerPublishVolumeMutex       sync.RWMutex
	controllerPublishVolumeArgsForCall []struct {
		ctx  context.Context
		in   *csi.ControllerPublishVolumeRequest
		opts []grpc.CallOption
	}
	controllerPublishVolumeReturns struct {
		result1 *csi.ControllerPublishVolumeResponse
		result2 error
	}
	ControllerUnpublishVolumeStub        func(ctx context.Context, in *csi.ControllerUnpublishVolumeRequest, opts ...grpc.CallOption) (*csi.ControllerUnpublishVolumeResponse, error)
	controllerUnpublishVolumeMutex       sync.RWMutex
	controllerUnpublishVolumeArgsForCall []struct {
		ctx  context.Context
		in   *csi.ControllerUnpublishVolumeRequest
		opts []grpc.CallOption
	}
	controllerUnpublishVolumeReturns struct {
		result1 *csi.ControllerUnpublishVolumeResponse
		result2 error
	}
	ValidateVolumeCapabilitiesStub        func(ctx context.Context, in *csi.ValidateVolumeCapabilitiesRequest, opts ...grpc.CallOption) (*csi.ValidateVolumeCapabilitiesResponse, error)
	validateVolumeCapabilitiesMutex       sync.RWMutex
	validateVolumeCapabilitiesArgsForCall []struct {
		ctx  context.Context
		in   *csi.ValidateVolumeCapabilitiesRequest
		opts []grpc.CallOption
	}
	validateVolumeCapabilitiesReturns struct {
		result1 *csi.ValidateVolumeCapabilitiesResponse
		result2 error
	}
	ListVolumesStub        func(ctx context.Context, in *csi.ListVolumesRequest, opts ...grpc.CallOption) (*csi.ListVolumesResponse, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct {
		ctx  context.Context
		in   *csi.ListVolumesRequest
		opts []grpc.CallOption
	}
	listVolumesReturns struct {
		result1 *csi.ListVolumesResponse
		result2 error
	}
	GetCapacityStub        func(ctx context.Context, in *csi.GetCapacityRequest, opts ...grpc.CallOption) (*csi.GetCapacityResponse, error)
	getCapacityMutex       sync.RWMutex
	getCapacityArgsForCall []struct {
		ctx  context.Context
		in   *csi.GetCapacityRequest
		opts []grpc.CallOption
	}
	getCapacityReturns struct {
		result1 *csi.GetCapacityResponse
		result2 error
	}
	ControllerGetCapabilitiesStub        func(ctx context.Context, in *csi.ControllerGetCapabilitiesRequest, opts ...grpc.CallOption) (*csi.ControllerGetCapabilitiesResponse, error)
	controllerGetCapabilitiesMutex       sync.RWMutex
	controllerGetCapabilitiesArgsForCall []struct {
		ctx  context.Context
		in   *csi.ControllerGetCapabilitiesRequest
		opts []grpc.CallOption
	}
	controllerGetCapabilitiesReturns struct {
		result1 *csi.ControllerGetCapabilitiesResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeControllerClient) CreateVolume(ctx context.Context, in *csi.CreateVolumeRequest, opts ...grpc.CallOption) (*csi.CreateVolumeResponse, error) {
	fake.createVolumeMutex.Lock()
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		ctx  context.Context
		in   *csi.CreateVolumeRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("CreateVolume", []interface{}{ctx, in, opts})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(ctx, in, opts...)
	}
	return fake.createVolumeReturns.result1, fake.createVolumeReturns.result2
}

func (fake *FakeControllerClient) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeControllerClient) CreateVolumeArgsForCall(i int) (context.Context, *csi.CreateVolumeRequest, []grpc.CallOption) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return fake.createVolumeArgsForCall[i].ctx, fake.createVolumeArgsForCall[i].in, fake.createVolumeArgsForCall[i].opts
}

func (fake *FakeControllerClient) CreateVolumeReturns(result1 *csi.CreateVolumeResponse, result2 error) {
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 *csi.CreateVolumeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeControllerClient) DeleteVolume(ctx context.Context, in *csi.DeleteVolumeRequest, opts ...grpc.CallOption) (*csi.DeleteVolumeResponse, error) {
	fake.deleteVolumeMutex.Lock()
	fake.deleteVolumeArgsForCall = append(fake.deleteVolumeArgsForCall, struct {
		ctx  context.Context
		in   *csi.DeleteVolumeRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("DeleteVolume", []interface{}{ctx, in, opts})
	fake.deleteVolumeMutex.Unlock()
	if fake.DeleteVolumeStub != nil {
		return fake.DeleteVolumeStub(ctx, in, opts...)
	}
	return fake.deleteVolumeReturns.result1, fake.deleteVolumeReturns.result2
}

func (fake *FakeControllerClient) DeleteVolumeCallCount() int {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return len(fake.deleteVolumeArgsForCall)
}

func (fake *FakeControllerClient) DeleteVolumeArgsForCall(i int) (context.Context, *csi.DeleteVolumeRequest, []grpc.CallOption) {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return fake.deleteVolumeArgsForCall[i].ctx, fake.deleteVolumeArgsForCall[i].in, fake.deleteVolumeArgsForCall[i].opts
}

func (fake *FakeControllerClient) DeleteVolumeReturns(result1 *csi.DeleteVolumeResponse, result2 error) {
	fake.DeleteVolumeStub = nil
	fake.deleteVolumeReturns = struct {
		result1 *csi.DeleteVolumeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeControllerClient) ControllerPublishVolume(ctx context.Context, in *csi.ControllerPublishVolumeRequest, opts ...grpc.CallOption) (*csi.ControllerPublishVolumeResponse, error) {
	fake.controllerPublishVolumeMutex.Lock()
	fake.controllerPublishVolumeArgsForCall = append(fake.controllerPublishVolumeArgsForCall, struct {
		ctx  context.Context
		in   *csi.ControllerPublishVolumeRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("ControllerPublishVolume", []interface{}{ctx, in, opts})
	fake.controllerPublishVolumeMutex.Unlock()
	if fake.ControllerPublishVolumeStub != nil {
		return fake.ControllerPublishVolumeStub(ctx, in, opts...)
	}
	return fake.controllerPublishVolumeReturns.result1, fake.controllerPublishVolumeReturns.result2
}

func (fake *FakeControllerClient) ControllerPublishVolumeCallCount() int {
	fake.controllerPublishVolumeMutex.RLock()
	defer fake.controllerPublishVolumeMutex.RUnlock()
	return len(fake.controllerPublishVolumeArgsForCall)
}

func (fake *FakeControllerClient) ControllerPublishVolumeArgsForCall(i int) (context.Context, *csi.ControllerPublishVolumeRequest, []grpc.CallOption) {
	fake.controllerPublishVolumeMutex.RLock()
	defer fake.controllerPublishVolumeMutex.RUnlock()
	return fake.controllerPublishVolumeArgsForCall[i].ctx, fake.controllerPublishVolumeArgsForCall[i].in, fake.controllerPublishVolumeArgsForCall[i].opts
}

func (fake *FakeControllerClient) ControllerPublishVolumeReturns(result1 *csi.ControllerPublishVolumeResponse, result2 error) {
	fake.ControllerPublishVolumeStub = nil
	fake.controllerPublishVolumeReturns = struct {
		result1 *csi.ControllerPublishVolumeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeControllerClient) ControllerUnpublishVolume(ctx context.Context, in *csi.ControllerUnpublishVolumeRequest, opts ...grpc.CallOption) (*csi.ControllerUnpublishVolumeResponse, error) {
	fake.controllerUnpublishVolumeMutex.Lock()
	fake.controllerUnpublishVolumeArgsForCall = append(fake.controllerUnpublishVolumeArgsForCall, struct {
		ctx  context.Context
		in   *csi.ControllerUnpublishVolumeRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("ControllerUnpublishVolume", []interface{}{ctx, in, opts})
	fake.controllerUnpublishVolumeMutex.Unlock()
	if fake.ControllerUnpublishVolumeStub != nil {
		return fake.ControllerUnpublishVolumeStub(ctx, in, opts...)
	}
	return fake.controllerUnpublishVolumeReturns.result1, fake.controllerUnpublishVolumeReturns.result2
}

func (fake *FakeControllerClient) ControllerUnpublishVolumeCallCount() int {
	fake.controllerUnpublishVolumeMutex.RLock()
	defer fake.controllerUnpublishVolumeMutex.RUnlock()
	return len(fake.controllerUnpublishVolumeArgsForCall)
}

func (fake *FakeControllerClient) ControllerUnpublishVolumeArgsForCall(i int) (context.Context, *csi.ControllerUnpublishVolumeRequest, []grpc.CallOption) {
	fake.controllerUnpublishVolumeMutex.RLock()
	defer fake.controllerUnpublishVolumeMutex.RUnlock()
	return fake.controllerUnpublishVolumeArgsForCall[i].ctx, fake.controllerUnpublishVolumeArgsForCall[i].in, fake.controllerUnpublishVolumeArgsForCall[i].opts
}

func (fake *FakeControllerClient) ControllerUnpublishVolumeReturns(result1 *csi.ControllerUnpublishVolumeResponse, result2 error) {
	fake.ControllerUnpublishVolumeStub = nil
	fake.controllerUnpublishVolumeReturns = struct {
		result1 *csi.ControllerUnpublishVolumeResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeControllerClient) ValidateVolumeCapabilities(ctx context.Context, in *csi.ValidateVolumeCapabilitiesRequest, opts ...grpc.CallOption) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	fake.validateVolumeCapabilitiesMutex.Lock()
	fake.validateVolumeCapabilitiesArgsForCall = append(fake.validateVolumeCapabilitiesArgsForCall, struct {
		ctx  context.Context
		in   *csi.ValidateVolumeCapabilitiesRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("ValidateVolumeCapabilities", []interface{}{ctx, in, opts})
	fake.validateVolumeCapabilitiesMutex.Unlock()
	if fake.ValidateVolumeCapabilitiesStub != nil {
		return fake.ValidateVolumeCapabilitiesStub(ctx, in, opts...)
	}
	return fake.validateVolumeCapabilitiesReturns.result1, fake.validateVolumeCapabilitiesReturns.result2
}

func (fake *FakeControllerClient) ValidateVolumeCapabilitiesCallCount() int {
	fake.validateVolumeCapabilitiesMutex.RLock()
	defer fake.validateVolumeCapabilitiesMutex.RUnlock()
	return len(fake.validateVolumeCapabilitiesArgsForCall)
}

func (fake *FakeControllerClient) ValidateVolumeCapabilitiesArgsForCall(i int) (context.Context, *csi.ValidateVolumeCapabilitiesRequest, []grpc.CallOption) {
	fake.validateVolumeCapabilitiesMutex.RLock()
	defer fake.validateVolumeCapabilitiesMutex.RUnlock()
	return fake.validateVolumeCapabilitiesArgsForCall[i].ctx, fake.validateVolumeCapabilitiesArgsForCall[i].in, fake.validateVolumeCapabilitiesArgsForCall[i].opts
}

func (fake *FakeControllerClient) ValidateVolumeCapabilitiesReturns(result1 *csi.ValidateVolumeCapabilitiesResponse, result2 error) {
	fake.ValidateVolumeCapabilitiesStub = nil
	fake.validateVolumeCapabilitiesReturns = struct {
		result1 *csi.ValidateVolumeCapabilitiesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeControllerClient) ListVolumes(ctx context.Context, in *csi.ListVolumesRequest, opts ...grpc.CallOption) (*csi.ListVolumesResponse, error) {
	fake.listVolumesMutex.Lock()
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct {
		ctx  context.Context
		in   *csi.ListVolumesRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("ListVolumes", []interface{}{ctx, in, opts})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub(ctx, in, opts...)
	}
	return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
}

func (fake *FakeControllerClient) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeControllerClient) ListVolumesArgsForCall(i int) (context.Context, *csi.ListVolumesRequest, []grpc.CallOption) {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return fake.listVolumesArgsForCall[i].ctx, fake.listVolumesArgsForCall[i].in, fake.listVolumesArgsForCall[i].opts
}

func (fake *FakeControllerClient) ListVolumesReturns(result1 *csi.ListVolumesResponse, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 *csi.ListVolumesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeControllerClient) GetCapacity(ctx context.Context, in *csi.GetCapacityRequest, opts ...grpc.CallOption) (*csi.GetCapacityResponse, error) {
	fake.getCapacityMutex.Lock()
	fake.getCapacityArgsForCall = append(fake.getCapacityArgsForCall, struct {
		ctx  context.Context
		in   *csi.GetCapacityRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("GetCapacity", []interface{}{ctx, in, opts})
	fake.getCapacityMutex.Unlock()
	if fake.GetCapacityStub != nil {
		return fake.GetCapacityStub(ctx, in, opts...)
	}
	return fake.getCapacityReturns.result1, fake.getCapacityReturns.result2
}

func (fake *FakeControllerClient) GetCapacityCallCount() int {
	fake.getCapacityMutex.RLock()
	defer fake.getCapacityMutex.RUnlock()
	return len(fake.getCapacityArgsForCall)
}

func (fake *FakeControllerClient) GetCapacityArgsForCall(i int) (context.Context, *csi.GetCapacityRequest, []grpc.CallOption) {
	fake.getCapacityMutex.RLock()
	defer fake.getCapacityMutex.RUnlock()
	return fake.getCapacityArgsForCall[i].ctx, fake.getCapacityArgsForCall[i].in, fake.getCapacityArgsForCall[i].opts
}

func (fake *FakeControllerClient) GetCapacityReturns(result1 *csi.GetCapacityResponse, result2 error) {
	fake.GetCapacityStub = nil
	fake.getCapacityReturns = struct {
		result1 *csi.GetCapacityResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeControllerClient) ControllerGetCapabilities(ctx context.Context, in *csi.ControllerGetCapabilitiesRequest, opts ...grpc.CallOption) (*csi.ControllerGetCapabilitiesResponse, error) {
	fake.controllerGetCapabilitiesMutex.Lock()
	fake.controllerGetCapabilitiesArgsForCall = append(fake.controllerGetCapabilitiesArgsForCall, struct {
		ctx  context.Context
		in   *csi.ControllerGetCapabilitiesRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("ControllerGetCapabilities", []interface{}{ctx, in, opts})
	fake.controllerGetCapabilitiesMutex.Unlock()
	if fake.ControllerGetCapabilitiesStub != nil {
		return fake.ControllerGetCapabilitiesStub(ctx, in, opts...)
	}
	return fake.controllerGetCapabilitiesReturns.result1, fake.controllerGetCapabilitiesReturns.result2
}

func (fake *FakeControllerClient) ControllerGetCapabilitiesCallCount() int {
	fake.controllerGetCapabilitiesMutex.RLock()
	defer fake.controllerGetCapabilitiesMutex.RUnlock()
	return len(fake.controllerGetCapabilitiesArgsForCall)
}

func (fake *FakeControllerClient) ControllerGetCapabilitiesArgsForCall(i int) (context.Context, *csi.ControllerGetCapabilitiesRequest, []grpc.CallOption) {
	fake.controllerGetCapabilitiesMutex.RLock()
	defer fake.controllerGetCapabilitiesMutex.RUnlock()
	return fake.controllerGetCapabilitiesArgsForCall[i].ctx, fake.controllerGetCapabilitiesArgsForCall[i].in, fake.controllerGetCapabilitiesArgsForCall[i].opts
}

func (fake *FakeControllerClient) ControllerGetCapabilitiesReturns(result1 *csi.ControllerGetCapabilitiesResponse, result2 error) {
	fake.ControllerGetCapabilitiesStub = nil
	fake.controllerGetCapabilitiesReturns = struct {
		result1 *csi.ControllerGetCapabilitiesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeControllerClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	fake.controllerPublishVolumeMutex.RLock()
	defer fake.controllerPublishVolumeMutex.RUnlock()
	fake.controllerUnpublishVolumeMutex.RLock()
	defer fake.controllerUnpublishVolumeMutex.RUnlock()
	fake.validateVolumeCapabilitiesMutex.RLock()
	defer fake.validateVolumeCapabilitiesMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	fake.getCapacityMutex.RLock()
	defer fake.getCapacityMutex.RUnlock()
	fake.controllerGetCapabilitiesMutex.RLock()
	defer fake.controllerGetCapabilitiesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeControllerClient) recordInvocation(key string, args []interface{}) {
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

var _ csi.ControllerClient = new(FakeControllerClient)
