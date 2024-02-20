// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: teleport/clusterconfig/v1/clusterconfig_service.proto

package clusterconfigv1

import (
	context "context"
	types "github.com/gravitational/teleport/api/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ClusterConfigService_GetClusterNetworkingConfig_FullMethodName    = "/teleport.clusterconfig.v1.ClusterConfigService/GetClusterNetworkingConfig"
	ClusterConfigService_UpdateClusterNetworkingConfig_FullMethodName = "/teleport.clusterconfig.v1.ClusterConfigService/UpdateClusterNetworkingConfig"
	ClusterConfigService_UpsertClusterNetworkingConfig_FullMethodName = "/teleport.clusterconfig.v1.ClusterConfigService/UpsertClusterNetworkingConfig"
	ClusterConfigService_ResetClusterNetworkingConfig_FullMethodName  = "/teleport.clusterconfig.v1.ClusterConfigService/ResetClusterNetworkingConfig"
	ClusterConfigService_GetSessionRecordingConfig_FullMethodName     = "/teleport.clusterconfig.v1.ClusterConfigService/GetSessionRecordingConfig"
	ClusterConfigService_UpdateSessionRecordingConfig_FullMethodName  = "/teleport.clusterconfig.v1.ClusterConfigService/UpdateSessionRecordingConfig"
	ClusterConfigService_UpsertSessionRecordingConfig_FullMethodName  = "/teleport.clusterconfig.v1.ClusterConfigService/UpsertSessionRecordingConfig"
	ClusterConfigService_ResetSessionRecordingConfig_FullMethodName   = "/teleport.clusterconfig.v1.ClusterConfigService/ResetSessionRecordingConfig"
	ClusterConfigService_GetAuthPreference_FullMethodName             = "/teleport.clusterconfig.v1.ClusterConfigService/GetAuthPreference"
	ClusterConfigService_UpdateAuthPreference_FullMethodName          = "/teleport.clusterconfig.v1.ClusterConfigService/UpdateAuthPreference"
	ClusterConfigService_UpsertAuthPreference_FullMethodName          = "/teleport.clusterconfig.v1.ClusterConfigService/UpsertAuthPreference"
	ClusterConfigService_ResetAuthPreference_FullMethodName           = "/teleport.clusterconfig.v1.ClusterConfigService/ResetAuthPreference"
	ClusterConfigService_GetClusterAuditConfig_FullMethodName         = "/teleport.clusterconfig.v1.ClusterConfigService/GetClusterAuditConfig"
	ClusterConfigService_UpdateClusterAuditConfig_FullMethodName      = "/teleport.clusterconfig.v1.ClusterConfigService/UpdateClusterAuditConfig"
	ClusterConfigService_UpsertClusterAuditConfig_FullMethodName      = "/teleport.clusterconfig.v1.ClusterConfigService/UpsertClusterAuditConfig"
	ClusterConfigService_ResetClusterAuditConfig_FullMethodName       = "/teleport.clusterconfig.v1.ClusterConfigService/ResetClusterAuditConfig"
)

// ClusterConfigServiceClient is the client API for ClusterConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterConfigServiceClient interface {
	// GetClusterNetworkingConfig retrieves the active cluster networking configuration.
	GetClusterNetworkingConfig(ctx context.Context, in *GetClusterNetworkingConfigRequest, opts ...grpc.CallOption) (*types.ClusterNetworkingConfigV2, error)
	// UpdateClusterNetworkingConfig updates the cluster networking configuration.
	UpdateClusterNetworkingConfig(ctx context.Context, in *UpdateClusterNetworkingConfigRequest, opts ...grpc.CallOption) (*types.ClusterNetworkingConfigV2, error)
	// UpsertClusterNetworkingConfig overwrites the active cluster networking configuration.
	UpsertClusterNetworkingConfig(ctx context.Context, in *UpsertClusterNetworkingConfigRequest, opts ...grpc.CallOption) (*types.ClusterNetworkingConfigV2, error)
	// ResetClusterNetworkingConfig restores the active cluster networking configuration to default settings.
	ResetClusterNetworkingConfig(ctx context.Context, in *ResetClusterNetworkingConfigRequest, opts ...grpc.CallOption) (*types.ClusterNetworkingConfigV2, error)
	// GetSessionRecordingConfig retrieves the active session recording configuration.
	GetSessionRecordingConfig(ctx context.Context, in *GetSessionRecordingConfigRequest, opts ...grpc.CallOption) (*types.SessionRecordingConfigV2, error)
	// UpdateSessionRecordingConfig updates the session recording configuration.
	UpdateSessionRecordingConfig(ctx context.Context, in *UpdateSessionRecordingConfigRequest, opts ...grpc.CallOption) (*types.SessionRecordingConfigV2, error)
	// UpsertSessionRecordingConfig overwrites the active session recording configuration.
	UpsertSessionRecordingConfig(ctx context.Context, in *UpsertSessionRecordingConfigRequest, opts ...grpc.CallOption) (*types.SessionRecordingConfigV2, error)
	// ResetSessionRecordingConfig restores the active session recording configuration to default settings.
	ResetSessionRecordingConfig(ctx context.Context, in *ResetSessionRecordingConfigRequest, opts ...grpc.CallOption) (*types.ClusterNetworkingConfigV2, error)
	// GetAuthPreference retrieves the active auth preference.
	GetAuthPreference(ctx context.Context, in *GetAuthPreferenceRequest, opts ...grpc.CallOption) (*types.AuthPreferenceV2, error)
	// UpdateAuthPreference updates the auth preference.
	UpdateAuthPreference(ctx context.Context, in *UpdateAuthPreferenceRequest, opts ...grpc.CallOption) (*types.AuthPreferenceV2, error)
	// UpsertAuthPreference overwrites the active auth preference.
	UpsertAuthPreference(ctx context.Context, in *UpsertAuthPreferenceRequest, opts ...grpc.CallOption) (*types.AuthPreferenceV2, error)
	// ResetAuthPreference restores the active auth preference to default settings.
	ResetAuthPreference(ctx context.Context, in *ResetAuthPreferenceRequest, opts ...grpc.CallOption) (*types.AuthPreferenceV2, error)
	// GetClusterAuditConfig retrieves the active cluster audit configuration.
	GetClusterAuditConfig(ctx context.Context, in *GetClusterAuditConfigRequest, opts ...grpc.CallOption) (*types.ClusterAuditConfigV2, error)
	// UpdateClusterAuditConfig updates the cluster audit configuration..
	UpdateClusterAuditConfig(ctx context.Context, in *UpdateClusterAuditConfigRequest, opts ...grpc.CallOption) (*types.ClusterAuditConfigV2, error)
	// UpsertClusterAuditConfig overwrites the active cluster audit configuration..
	UpsertClusterAuditConfig(ctx context.Context, in *UpsertClusterAuditConfigRequest, opts ...grpc.CallOption) (*types.ClusterAuditConfigV2, error)
	// ResetClusterAuditConfig restores the active cluster audit configuration. to default settings.
	ResetClusterAuditConfig(ctx context.Context, in *ResetClusterAuditConfigRequest, opts ...grpc.CallOption) (*types.ClusterAuditConfigV2, error)
}

type clusterConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterConfigServiceClient(cc grpc.ClientConnInterface) ClusterConfigServiceClient {
	return &clusterConfigServiceClient{cc}
}

func (c *clusterConfigServiceClient) GetClusterNetworkingConfig(ctx context.Context, in *GetClusterNetworkingConfigRequest, opts ...grpc.CallOption) (*types.ClusterNetworkingConfigV2, error) {
	out := new(types.ClusterNetworkingConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_GetClusterNetworkingConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) UpdateClusterNetworkingConfig(ctx context.Context, in *UpdateClusterNetworkingConfigRequest, opts ...grpc.CallOption) (*types.ClusterNetworkingConfigV2, error) {
	out := new(types.ClusterNetworkingConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_UpdateClusterNetworkingConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) UpsertClusterNetworkingConfig(ctx context.Context, in *UpsertClusterNetworkingConfigRequest, opts ...grpc.CallOption) (*types.ClusterNetworkingConfigV2, error) {
	out := new(types.ClusterNetworkingConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_UpsertClusterNetworkingConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) ResetClusterNetworkingConfig(ctx context.Context, in *ResetClusterNetworkingConfigRequest, opts ...grpc.CallOption) (*types.ClusterNetworkingConfigV2, error) {
	out := new(types.ClusterNetworkingConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_ResetClusterNetworkingConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) GetSessionRecordingConfig(ctx context.Context, in *GetSessionRecordingConfigRequest, opts ...grpc.CallOption) (*types.SessionRecordingConfigV2, error) {
	out := new(types.SessionRecordingConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_GetSessionRecordingConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) UpdateSessionRecordingConfig(ctx context.Context, in *UpdateSessionRecordingConfigRequest, opts ...grpc.CallOption) (*types.SessionRecordingConfigV2, error) {
	out := new(types.SessionRecordingConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_UpdateSessionRecordingConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) UpsertSessionRecordingConfig(ctx context.Context, in *UpsertSessionRecordingConfigRequest, opts ...grpc.CallOption) (*types.SessionRecordingConfigV2, error) {
	out := new(types.SessionRecordingConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_UpsertSessionRecordingConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) ResetSessionRecordingConfig(ctx context.Context, in *ResetSessionRecordingConfigRequest, opts ...grpc.CallOption) (*types.ClusterNetworkingConfigV2, error) {
	out := new(types.ClusterNetworkingConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_ResetSessionRecordingConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) GetAuthPreference(ctx context.Context, in *GetAuthPreferenceRequest, opts ...grpc.CallOption) (*types.AuthPreferenceV2, error) {
	out := new(types.AuthPreferenceV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_GetAuthPreference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) UpdateAuthPreference(ctx context.Context, in *UpdateAuthPreferenceRequest, opts ...grpc.CallOption) (*types.AuthPreferenceV2, error) {
	out := new(types.AuthPreferenceV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_UpdateAuthPreference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) UpsertAuthPreference(ctx context.Context, in *UpsertAuthPreferenceRequest, opts ...grpc.CallOption) (*types.AuthPreferenceV2, error) {
	out := new(types.AuthPreferenceV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_UpsertAuthPreference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) ResetAuthPreference(ctx context.Context, in *ResetAuthPreferenceRequest, opts ...grpc.CallOption) (*types.AuthPreferenceV2, error) {
	out := new(types.AuthPreferenceV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_ResetAuthPreference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) GetClusterAuditConfig(ctx context.Context, in *GetClusterAuditConfigRequest, opts ...grpc.CallOption) (*types.ClusterAuditConfigV2, error) {
	out := new(types.ClusterAuditConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_GetClusterAuditConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) UpdateClusterAuditConfig(ctx context.Context, in *UpdateClusterAuditConfigRequest, opts ...grpc.CallOption) (*types.ClusterAuditConfigV2, error) {
	out := new(types.ClusterAuditConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_UpdateClusterAuditConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) UpsertClusterAuditConfig(ctx context.Context, in *UpsertClusterAuditConfigRequest, opts ...grpc.CallOption) (*types.ClusterAuditConfigV2, error) {
	out := new(types.ClusterAuditConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_UpsertClusterAuditConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterConfigServiceClient) ResetClusterAuditConfig(ctx context.Context, in *ResetClusterAuditConfigRequest, opts ...grpc.CallOption) (*types.ClusterAuditConfigV2, error) {
	out := new(types.ClusterAuditConfigV2)
	err := c.cc.Invoke(ctx, ClusterConfigService_ResetClusterAuditConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterConfigServiceServer is the server API for ClusterConfigService service.
// All implementations must embed UnimplementedClusterConfigServiceServer
// for forward compatibility
type ClusterConfigServiceServer interface {
	// GetClusterNetworkingConfig retrieves the active cluster networking configuration.
	GetClusterNetworkingConfig(context.Context, *GetClusterNetworkingConfigRequest) (*types.ClusterNetworkingConfigV2, error)
	// UpdateClusterNetworkingConfig updates the cluster networking configuration.
	UpdateClusterNetworkingConfig(context.Context, *UpdateClusterNetworkingConfigRequest) (*types.ClusterNetworkingConfigV2, error)
	// UpsertClusterNetworkingConfig overwrites the active cluster networking configuration.
	UpsertClusterNetworkingConfig(context.Context, *UpsertClusterNetworkingConfigRequest) (*types.ClusterNetworkingConfigV2, error)
	// ResetClusterNetworkingConfig restores the active cluster networking configuration to default settings.
	ResetClusterNetworkingConfig(context.Context, *ResetClusterNetworkingConfigRequest) (*types.ClusterNetworkingConfigV2, error)
	// GetSessionRecordingConfig retrieves the active session recording configuration.
	GetSessionRecordingConfig(context.Context, *GetSessionRecordingConfigRequest) (*types.SessionRecordingConfigV2, error)
	// UpdateSessionRecordingConfig updates the session recording configuration.
	UpdateSessionRecordingConfig(context.Context, *UpdateSessionRecordingConfigRequest) (*types.SessionRecordingConfigV2, error)
	// UpsertSessionRecordingConfig overwrites the active session recording configuration.
	UpsertSessionRecordingConfig(context.Context, *UpsertSessionRecordingConfigRequest) (*types.SessionRecordingConfigV2, error)
	// ResetSessionRecordingConfig restores the active session recording configuration to default settings.
	ResetSessionRecordingConfig(context.Context, *ResetSessionRecordingConfigRequest) (*types.ClusterNetworkingConfigV2, error)
	// GetAuthPreference retrieves the active auth preference.
	GetAuthPreference(context.Context, *GetAuthPreferenceRequest) (*types.AuthPreferenceV2, error)
	// UpdateAuthPreference updates the auth preference.
	UpdateAuthPreference(context.Context, *UpdateAuthPreferenceRequest) (*types.AuthPreferenceV2, error)
	// UpsertAuthPreference overwrites the active auth preference.
	UpsertAuthPreference(context.Context, *UpsertAuthPreferenceRequest) (*types.AuthPreferenceV2, error)
	// ResetAuthPreference restores the active auth preference to default settings.
	ResetAuthPreference(context.Context, *ResetAuthPreferenceRequest) (*types.AuthPreferenceV2, error)
	// GetClusterAuditConfig retrieves the active cluster audit configuration.
	GetClusterAuditConfig(context.Context, *GetClusterAuditConfigRequest) (*types.ClusterAuditConfigV2, error)
	// UpdateClusterAuditConfig updates the cluster audit configuration..
	UpdateClusterAuditConfig(context.Context, *UpdateClusterAuditConfigRequest) (*types.ClusterAuditConfigV2, error)
	// UpsertClusterAuditConfig overwrites the active cluster audit configuration..
	UpsertClusterAuditConfig(context.Context, *UpsertClusterAuditConfigRequest) (*types.ClusterAuditConfigV2, error)
	// ResetClusterAuditConfig restores the active cluster audit configuration. to default settings.
	ResetClusterAuditConfig(context.Context, *ResetClusterAuditConfigRequest) (*types.ClusterAuditConfigV2, error)
	mustEmbedUnimplementedClusterConfigServiceServer()
}

// UnimplementedClusterConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClusterConfigServiceServer struct {
}

func (UnimplementedClusterConfigServiceServer) GetClusterNetworkingConfig(context.Context, *GetClusterNetworkingConfigRequest) (*types.ClusterNetworkingConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterNetworkingConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) UpdateClusterNetworkingConfig(context.Context, *UpdateClusterNetworkingConfigRequest) (*types.ClusterNetworkingConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClusterNetworkingConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) UpsertClusterNetworkingConfig(context.Context, *UpsertClusterNetworkingConfigRequest) (*types.ClusterNetworkingConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertClusterNetworkingConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) ResetClusterNetworkingConfig(context.Context, *ResetClusterNetworkingConfigRequest) (*types.ClusterNetworkingConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetClusterNetworkingConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) GetSessionRecordingConfig(context.Context, *GetSessionRecordingConfigRequest) (*types.SessionRecordingConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionRecordingConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) UpdateSessionRecordingConfig(context.Context, *UpdateSessionRecordingConfigRequest) (*types.SessionRecordingConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSessionRecordingConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) UpsertSessionRecordingConfig(context.Context, *UpsertSessionRecordingConfigRequest) (*types.SessionRecordingConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertSessionRecordingConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) ResetSessionRecordingConfig(context.Context, *ResetSessionRecordingConfigRequest) (*types.ClusterNetworkingConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetSessionRecordingConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) GetAuthPreference(context.Context, *GetAuthPreferenceRequest) (*types.AuthPreferenceV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthPreference not implemented")
}
func (UnimplementedClusterConfigServiceServer) UpdateAuthPreference(context.Context, *UpdateAuthPreferenceRequest) (*types.AuthPreferenceV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthPreference not implemented")
}
func (UnimplementedClusterConfigServiceServer) UpsertAuthPreference(context.Context, *UpsertAuthPreferenceRequest) (*types.AuthPreferenceV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertAuthPreference not implemented")
}
func (UnimplementedClusterConfigServiceServer) ResetAuthPreference(context.Context, *ResetAuthPreferenceRequest) (*types.AuthPreferenceV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetAuthPreference not implemented")
}
func (UnimplementedClusterConfigServiceServer) GetClusterAuditConfig(context.Context, *GetClusterAuditConfigRequest) (*types.ClusterAuditConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAuditConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) UpdateClusterAuditConfig(context.Context, *UpdateClusterAuditConfigRequest) (*types.ClusterAuditConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClusterAuditConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) UpsertClusterAuditConfig(context.Context, *UpsertClusterAuditConfigRequest) (*types.ClusterAuditConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertClusterAuditConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) ResetClusterAuditConfig(context.Context, *ResetClusterAuditConfigRequest) (*types.ClusterAuditConfigV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetClusterAuditConfig not implemented")
}
func (UnimplementedClusterConfigServiceServer) mustEmbedUnimplementedClusterConfigServiceServer() {}

// UnsafeClusterConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterConfigServiceServer will
// result in compilation errors.
type UnsafeClusterConfigServiceServer interface {
	mustEmbedUnimplementedClusterConfigServiceServer()
}

func RegisterClusterConfigServiceServer(s grpc.ServiceRegistrar, srv ClusterConfigServiceServer) {
	s.RegisterService(&ClusterConfigService_ServiceDesc, srv)
}

func _ClusterConfigService_GetClusterNetworkingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterNetworkingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).GetClusterNetworkingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_GetClusterNetworkingConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).GetClusterNetworkingConfig(ctx, req.(*GetClusterNetworkingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_UpdateClusterNetworkingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterNetworkingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).UpdateClusterNetworkingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_UpdateClusterNetworkingConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).UpdateClusterNetworkingConfig(ctx, req.(*UpdateClusterNetworkingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_UpsertClusterNetworkingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertClusterNetworkingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).UpsertClusterNetworkingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_UpsertClusterNetworkingConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).UpsertClusterNetworkingConfig(ctx, req.(*UpsertClusterNetworkingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_ResetClusterNetworkingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetClusterNetworkingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).ResetClusterNetworkingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_ResetClusterNetworkingConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).ResetClusterNetworkingConfig(ctx, req.(*ResetClusterNetworkingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_GetSessionRecordingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRecordingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).GetSessionRecordingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_GetSessionRecordingConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).GetSessionRecordingConfig(ctx, req.(*GetSessionRecordingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_UpdateSessionRecordingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionRecordingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).UpdateSessionRecordingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_UpdateSessionRecordingConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).UpdateSessionRecordingConfig(ctx, req.(*UpdateSessionRecordingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_UpsertSessionRecordingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertSessionRecordingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).UpsertSessionRecordingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_UpsertSessionRecordingConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).UpsertSessionRecordingConfig(ctx, req.(*UpsertSessionRecordingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_ResetSessionRecordingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetSessionRecordingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).ResetSessionRecordingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_ResetSessionRecordingConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).ResetSessionRecordingConfig(ctx, req.(*ResetSessionRecordingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_GetAuthPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).GetAuthPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_GetAuthPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).GetAuthPreference(ctx, req.(*GetAuthPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_UpdateAuthPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).UpdateAuthPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_UpdateAuthPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).UpdateAuthPreference(ctx, req.(*UpdateAuthPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_UpsertAuthPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertAuthPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).UpsertAuthPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_UpsertAuthPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).UpsertAuthPreference(ctx, req.(*UpsertAuthPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_ResetAuthPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetAuthPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).ResetAuthPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_ResetAuthPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).ResetAuthPreference(ctx, req.(*ResetAuthPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_GetClusterAuditConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAuditConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).GetClusterAuditConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_GetClusterAuditConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).GetClusterAuditConfig(ctx, req.(*GetClusterAuditConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_UpdateClusterAuditConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterAuditConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).UpdateClusterAuditConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_UpdateClusterAuditConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).UpdateClusterAuditConfig(ctx, req.(*UpdateClusterAuditConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_UpsertClusterAuditConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertClusterAuditConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).UpsertClusterAuditConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_UpsertClusterAuditConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).UpsertClusterAuditConfig(ctx, req.(*UpsertClusterAuditConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterConfigService_ResetClusterAuditConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetClusterAuditConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterConfigServiceServer).ResetClusterAuditConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterConfigService_ResetClusterAuditConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterConfigServiceServer).ResetClusterAuditConfig(ctx, req.(*ResetClusterAuditConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterConfigService_ServiceDesc is the grpc.ServiceDesc for ClusterConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.clusterconfig.v1.ClusterConfigService",
	HandlerType: (*ClusterConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClusterNetworkingConfig",
			Handler:    _ClusterConfigService_GetClusterNetworkingConfig_Handler,
		},
		{
			MethodName: "UpdateClusterNetworkingConfig",
			Handler:    _ClusterConfigService_UpdateClusterNetworkingConfig_Handler,
		},
		{
			MethodName: "UpsertClusterNetworkingConfig",
			Handler:    _ClusterConfigService_UpsertClusterNetworkingConfig_Handler,
		},
		{
			MethodName: "ResetClusterNetworkingConfig",
			Handler:    _ClusterConfigService_ResetClusterNetworkingConfig_Handler,
		},
		{
			MethodName: "GetSessionRecordingConfig",
			Handler:    _ClusterConfigService_GetSessionRecordingConfig_Handler,
		},
		{
			MethodName: "UpdateSessionRecordingConfig",
			Handler:    _ClusterConfigService_UpdateSessionRecordingConfig_Handler,
		},
		{
			MethodName: "UpsertSessionRecordingConfig",
			Handler:    _ClusterConfigService_UpsertSessionRecordingConfig_Handler,
		},
		{
			MethodName: "ResetSessionRecordingConfig",
			Handler:    _ClusterConfigService_ResetSessionRecordingConfig_Handler,
		},
		{
			MethodName: "GetAuthPreference",
			Handler:    _ClusterConfigService_GetAuthPreference_Handler,
		},
		{
			MethodName: "UpdateAuthPreference",
			Handler:    _ClusterConfigService_UpdateAuthPreference_Handler,
		},
		{
			MethodName: "UpsertAuthPreference",
			Handler:    _ClusterConfigService_UpsertAuthPreference_Handler,
		},
		{
			MethodName: "ResetAuthPreference",
			Handler:    _ClusterConfigService_ResetAuthPreference_Handler,
		},
		{
			MethodName: "GetClusterAuditConfig",
			Handler:    _ClusterConfigService_GetClusterAuditConfig_Handler,
		},
		{
			MethodName: "UpdateClusterAuditConfig",
			Handler:    _ClusterConfigService_UpdateClusterAuditConfig_Handler,
		},
		{
			MethodName: "UpsertClusterAuditConfig",
			Handler:    _ClusterConfigService_UpsertClusterAuditConfig_Handler,
		},
		{
			MethodName: "ResetClusterAuditConfig",
			Handler:    _ClusterConfigService_ResetClusterAuditConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/clusterconfig/v1/clusterconfig_service.proto",
}
