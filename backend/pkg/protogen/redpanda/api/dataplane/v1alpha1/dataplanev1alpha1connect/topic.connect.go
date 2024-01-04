// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1alpha1/topic.proto

package dataplanev1alpha1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"

	v1alpha1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1alpha1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TopicServiceName is the fully-qualified name of the TopicService service.
	TopicServiceName = "redpanda.api.dataplane.v1alpha1.TopicService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TopicServiceCreateTopicProcedure is the fully-qualified name of the TopicService's CreateTopic
	// RPC.
	TopicServiceCreateTopicProcedure = "/redpanda.api.dataplane.v1alpha1.TopicService/CreateTopic"
	// TopicServiceListTopicsProcedure is the fully-qualified name of the TopicService's ListTopics RPC.
	TopicServiceListTopicsProcedure = "/redpanda.api.dataplane.v1alpha1.TopicService/ListTopics"
	// TopicServiceDeleteTopicProcedure is the fully-qualified name of the TopicService's DeleteTopic
	// RPC.
	TopicServiceDeleteTopicProcedure = "/redpanda.api.dataplane.v1alpha1.TopicService/DeleteTopic"
	// TopicServiceGetTopicConfigurationProcedure is the fully-qualified name of the TopicService's
	// GetTopicConfiguration RPC.
	TopicServiceGetTopicConfigurationProcedure = "/redpanda.api.dataplane.v1alpha1.TopicService/GetTopicConfiguration"
	// TopicServiceUpdateTopicConfigurationProcedure is the fully-qualified name of the TopicService's
	// UpdateTopicConfiguration RPC.
	TopicServiceUpdateTopicConfigurationProcedure = "/redpanda.api.dataplane.v1alpha1.TopicService/UpdateTopicConfiguration"
	// TopicServiceSetTopicConfigurationProcedure is the fully-qualified name of the TopicService's
	// SetTopicConfiguration RPC.
	TopicServiceSetTopicConfigurationProcedure = "/redpanda.api.dataplane.v1alpha1.TopicService/SetTopicConfiguration"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	topicServiceServiceDescriptor                        = v1alpha1.File_redpanda_api_dataplane_v1alpha1_topic_proto.Services().ByName("TopicService")
	topicServiceCreateTopicMethodDescriptor              = topicServiceServiceDescriptor.Methods().ByName("CreateTopic")
	topicServiceListTopicsMethodDescriptor               = topicServiceServiceDescriptor.Methods().ByName("ListTopics")
	topicServiceDeleteTopicMethodDescriptor              = topicServiceServiceDescriptor.Methods().ByName("DeleteTopic")
	topicServiceGetTopicConfigurationMethodDescriptor    = topicServiceServiceDescriptor.Methods().ByName("GetTopicConfiguration")
	topicServiceUpdateTopicConfigurationMethodDescriptor = topicServiceServiceDescriptor.Methods().ByName("UpdateTopicConfiguration")
	topicServiceSetTopicConfigurationMethodDescriptor    = topicServiceServiceDescriptor.Methods().ByName("SetTopicConfiguration")
)

// TopicServiceClient is a client for the redpanda.api.dataplane.v1alpha1.TopicService service.
type TopicServiceClient interface {
	CreateTopic(context.Context, *connect.Request[v1alpha1.CreateTopicRequest]) (*connect.Response[v1alpha1.CreateTopicResponse], error)
	ListTopics(context.Context, *connect.Request[v1alpha1.ListTopicsRequest]) (*connect.Response[v1alpha1.ListTopicsResponse], error)
	DeleteTopic(context.Context, *connect.Request[v1alpha1.DeleteTopicRequest]) (*connect.Response[v1alpha1.DeleteTopicResponse], error)
	GetTopicConfiguration(context.Context, *connect.Request[v1alpha1.GetTopicConfigurationRequest]) (*connect.Response[v1alpha1.GetTopicConfigurationResponse], error)
	UpdateTopicConfiguration(context.Context, *connect.Request[v1alpha1.UpdateTopicConfigurationRequest]) (*connect.Response[v1alpha1.UpdateTopicConfigurationResponse], error)
	SetTopicConfiguration(context.Context, *connect.Request[v1alpha1.SetTopicConfigurationRequest]) (*connect.Response[v1alpha1.SetTopicConfigurationResponse], error)
}

// NewTopicServiceClient constructs a client for the redpanda.api.dataplane.v1alpha1.TopicService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTopicServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TopicServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &topicServiceClient{
		createTopic: connect.NewClient[v1alpha1.CreateTopicRequest, v1alpha1.CreateTopicResponse](
			httpClient,
			baseURL+TopicServiceCreateTopicProcedure,
			connect.WithSchema(topicServiceCreateTopicMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listTopics: connect.NewClient[v1alpha1.ListTopicsRequest, v1alpha1.ListTopicsResponse](
			httpClient,
			baseURL+TopicServiceListTopicsProcedure,
			connect.WithSchema(topicServiceListTopicsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteTopic: connect.NewClient[v1alpha1.DeleteTopicRequest, v1alpha1.DeleteTopicResponse](
			httpClient,
			baseURL+TopicServiceDeleteTopicProcedure,
			connect.WithSchema(topicServiceDeleteTopicMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getTopicConfiguration: connect.NewClient[v1alpha1.GetTopicConfigurationRequest, v1alpha1.GetTopicConfigurationResponse](
			httpClient,
			baseURL+TopicServiceGetTopicConfigurationProcedure,
			connect.WithSchema(topicServiceGetTopicConfigurationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateTopicConfiguration: connect.NewClient[v1alpha1.UpdateTopicConfigurationRequest, v1alpha1.UpdateTopicConfigurationResponse](
			httpClient,
			baseURL+TopicServiceUpdateTopicConfigurationProcedure,
			connect.WithSchema(topicServiceUpdateTopicConfigurationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setTopicConfiguration: connect.NewClient[v1alpha1.SetTopicConfigurationRequest, v1alpha1.SetTopicConfigurationResponse](
			httpClient,
			baseURL+TopicServiceSetTopicConfigurationProcedure,
			connect.WithSchema(topicServiceSetTopicConfigurationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// topicServiceClient implements TopicServiceClient.
type topicServiceClient struct {
	createTopic              *connect.Client[v1alpha1.CreateTopicRequest, v1alpha1.CreateTopicResponse]
	listTopics               *connect.Client[v1alpha1.ListTopicsRequest, v1alpha1.ListTopicsResponse]
	deleteTopic              *connect.Client[v1alpha1.DeleteTopicRequest, v1alpha1.DeleteTopicResponse]
	getTopicConfiguration    *connect.Client[v1alpha1.GetTopicConfigurationRequest, v1alpha1.GetTopicConfigurationResponse]
	updateTopicConfiguration *connect.Client[v1alpha1.UpdateTopicConfigurationRequest, v1alpha1.UpdateTopicConfigurationResponse]
	setTopicConfiguration    *connect.Client[v1alpha1.SetTopicConfigurationRequest, v1alpha1.SetTopicConfigurationResponse]
}

// CreateTopic calls redpanda.api.dataplane.v1alpha1.TopicService.CreateTopic.
func (c *topicServiceClient) CreateTopic(ctx context.Context, req *connect.Request[v1alpha1.CreateTopicRequest]) (*connect.Response[v1alpha1.CreateTopicResponse], error) {
	return c.createTopic.CallUnary(ctx, req)
}

// ListTopics calls redpanda.api.dataplane.v1alpha1.TopicService.ListTopics.
func (c *topicServiceClient) ListTopics(ctx context.Context, req *connect.Request[v1alpha1.ListTopicsRequest]) (*connect.Response[v1alpha1.ListTopicsResponse], error) {
	return c.listTopics.CallUnary(ctx, req)
}

// DeleteTopic calls redpanda.api.dataplane.v1alpha1.TopicService.DeleteTopic.
func (c *topicServiceClient) DeleteTopic(ctx context.Context, req *connect.Request[v1alpha1.DeleteTopicRequest]) (*connect.Response[v1alpha1.DeleteTopicResponse], error) {
	return c.deleteTopic.CallUnary(ctx, req)
}

// GetTopicConfiguration calls redpanda.api.dataplane.v1alpha1.TopicService.GetTopicConfiguration.
func (c *topicServiceClient) GetTopicConfiguration(ctx context.Context, req *connect.Request[v1alpha1.GetTopicConfigurationRequest]) (*connect.Response[v1alpha1.GetTopicConfigurationResponse], error) {
	return c.getTopicConfiguration.CallUnary(ctx, req)
}

// UpdateTopicConfiguration calls
// redpanda.api.dataplane.v1alpha1.TopicService.UpdateTopicConfiguration.
func (c *topicServiceClient) UpdateTopicConfiguration(ctx context.Context, req *connect.Request[v1alpha1.UpdateTopicConfigurationRequest]) (*connect.Response[v1alpha1.UpdateTopicConfigurationResponse], error) {
	return c.updateTopicConfiguration.CallUnary(ctx, req)
}

// SetTopicConfiguration calls redpanda.api.dataplane.v1alpha1.TopicService.SetTopicConfiguration.
func (c *topicServiceClient) SetTopicConfiguration(ctx context.Context, req *connect.Request[v1alpha1.SetTopicConfigurationRequest]) (*connect.Response[v1alpha1.SetTopicConfigurationResponse], error) {
	return c.setTopicConfiguration.CallUnary(ctx, req)
}

// TopicServiceHandler is an implementation of the redpanda.api.dataplane.v1alpha1.TopicService
// service.
type TopicServiceHandler interface {
	CreateTopic(context.Context, *connect.Request[v1alpha1.CreateTopicRequest]) (*connect.Response[v1alpha1.CreateTopicResponse], error)
	ListTopics(context.Context, *connect.Request[v1alpha1.ListTopicsRequest]) (*connect.Response[v1alpha1.ListTopicsResponse], error)
	DeleteTopic(context.Context, *connect.Request[v1alpha1.DeleteTopicRequest]) (*connect.Response[v1alpha1.DeleteTopicResponse], error)
	GetTopicConfiguration(context.Context, *connect.Request[v1alpha1.GetTopicConfigurationRequest]) (*connect.Response[v1alpha1.GetTopicConfigurationResponse], error)
	UpdateTopicConfiguration(context.Context, *connect.Request[v1alpha1.UpdateTopicConfigurationRequest]) (*connect.Response[v1alpha1.UpdateTopicConfigurationResponse], error)
	SetTopicConfiguration(context.Context, *connect.Request[v1alpha1.SetTopicConfigurationRequest]) (*connect.Response[v1alpha1.SetTopicConfigurationResponse], error)
}

// NewTopicServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTopicServiceHandler(svc TopicServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	topicServiceCreateTopicHandler := connect.NewUnaryHandler(
		TopicServiceCreateTopicProcedure,
		svc.CreateTopic,
		connect.WithSchema(topicServiceCreateTopicMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	topicServiceListTopicsHandler := connect.NewUnaryHandler(
		TopicServiceListTopicsProcedure,
		svc.ListTopics,
		connect.WithSchema(topicServiceListTopicsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	topicServiceDeleteTopicHandler := connect.NewUnaryHandler(
		TopicServiceDeleteTopicProcedure,
		svc.DeleteTopic,
		connect.WithSchema(topicServiceDeleteTopicMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	topicServiceGetTopicConfigurationHandler := connect.NewUnaryHandler(
		TopicServiceGetTopicConfigurationProcedure,
		svc.GetTopicConfiguration,
		connect.WithSchema(topicServiceGetTopicConfigurationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	topicServiceUpdateTopicConfigurationHandler := connect.NewUnaryHandler(
		TopicServiceUpdateTopicConfigurationProcedure,
		svc.UpdateTopicConfiguration,
		connect.WithSchema(topicServiceUpdateTopicConfigurationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	topicServiceSetTopicConfigurationHandler := connect.NewUnaryHandler(
		TopicServiceSetTopicConfigurationProcedure,
		svc.SetTopicConfiguration,
		connect.WithSchema(topicServiceSetTopicConfigurationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/redpanda.api.dataplane.v1alpha1.TopicService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TopicServiceCreateTopicProcedure:
			topicServiceCreateTopicHandler.ServeHTTP(w, r)
		case TopicServiceListTopicsProcedure:
			topicServiceListTopicsHandler.ServeHTTP(w, r)
		case TopicServiceDeleteTopicProcedure:
			topicServiceDeleteTopicHandler.ServeHTTP(w, r)
		case TopicServiceGetTopicConfigurationProcedure:
			topicServiceGetTopicConfigurationHandler.ServeHTTP(w, r)
		case TopicServiceUpdateTopicConfigurationProcedure:
			topicServiceUpdateTopicConfigurationHandler.ServeHTTP(w, r)
		case TopicServiceSetTopicConfigurationProcedure:
			topicServiceSetTopicConfigurationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTopicServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTopicServiceHandler struct{}

func (UnimplementedTopicServiceHandler) CreateTopic(context.Context, *connect.Request[v1alpha1.CreateTopicRequest]) (*connect.Response[v1alpha1.CreateTopicResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha1.TopicService.CreateTopic is not implemented"))
}

func (UnimplementedTopicServiceHandler) ListTopics(context.Context, *connect.Request[v1alpha1.ListTopicsRequest]) (*connect.Response[v1alpha1.ListTopicsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha1.TopicService.ListTopics is not implemented"))
}

func (UnimplementedTopicServiceHandler) DeleteTopic(context.Context, *connect.Request[v1alpha1.DeleteTopicRequest]) (*connect.Response[v1alpha1.DeleteTopicResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha1.TopicService.DeleteTopic is not implemented"))
}

func (UnimplementedTopicServiceHandler) GetTopicConfiguration(context.Context, *connect.Request[v1alpha1.GetTopicConfigurationRequest]) (*connect.Response[v1alpha1.GetTopicConfigurationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha1.TopicService.GetTopicConfiguration is not implemented"))
}

func (UnimplementedTopicServiceHandler) UpdateTopicConfiguration(context.Context, *connect.Request[v1alpha1.UpdateTopicConfigurationRequest]) (*connect.Response[v1alpha1.UpdateTopicConfigurationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha1.TopicService.UpdateTopicConfiguration is not implemented"))
}

func (UnimplementedTopicServiceHandler) SetTopicConfiguration(context.Context, *connect.Request[v1alpha1.SetTopicConfigurationRequest]) (*connect.Response[v1alpha1.SetTopicConfigurationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha1.TopicService.SetTopicConfiguration is not implemented"))
}