// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: user.proto

/*
Package feedbacks is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package feedbacks

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_FeedbacksUsersService_List_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_FeedbacksUsersService_List_0(ctx context.Context, marshaler runtime.Marshaler, client FeedbacksUsersServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListUsersRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FeedbacksUsersService_List_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FeedbacksUsersService_List_0(ctx context.Context, marshaler runtime.Marshaler, server FeedbacksUsersServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListUsersRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FeedbacksUsersService_List_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.List(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_FeedbacksOperatorsService_List_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_FeedbacksOperatorsService_List_0(ctx context.Context, marshaler runtime.Marshaler, client FeedbacksOperatorsServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListUsersRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FeedbacksOperatorsService_List_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FeedbacksOperatorsService_List_0(ctx context.Context, marshaler runtime.Marshaler, server FeedbacksOperatorsServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListUsersRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FeedbacksOperatorsService_List_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.List(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterFeedbacksUsersServiceHandlerServer registers the http handlers for service FeedbacksUsersService to "mux".
// UnaryRPC     :call FeedbacksUsersServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterFeedbacksUsersServiceHandlerFromEndpoint instead.
func RegisterFeedbacksUsersServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server FeedbacksUsersServiceServer) error {

	mux.Handle("GET", pattern_FeedbacksUsersService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/feedbacks.FeedbacksUsersService/List", runtime.WithHTTPPathPattern("/api/v1/users"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FeedbacksUsersService_List_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FeedbacksUsersService_List_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterFeedbacksOperatorsServiceHandlerServer registers the http handlers for service FeedbacksOperatorsService to "mux".
// UnaryRPC     :call FeedbacksOperatorsServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterFeedbacksOperatorsServiceHandlerFromEndpoint instead.
func RegisterFeedbacksOperatorsServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server FeedbacksOperatorsServiceServer) error {

	mux.Handle("GET", pattern_FeedbacksOperatorsService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/feedbacks.FeedbacksOperatorsService/List", runtime.WithHTTPPathPattern("/api/v1/feedbacks/operators"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FeedbacksOperatorsService_List_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FeedbacksOperatorsService_List_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterFeedbacksUsersServiceHandlerFromEndpoint is same as RegisterFeedbacksUsersServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterFeedbacksUsersServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterFeedbacksUsersServiceHandler(ctx, mux, conn)
}

// RegisterFeedbacksUsersServiceHandler registers the http handlers for service FeedbacksUsersService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterFeedbacksUsersServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterFeedbacksUsersServiceHandlerClient(ctx, mux, NewFeedbacksUsersServiceClient(conn))
}

// RegisterFeedbacksUsersServiceHandlerClient registers the http handlers for service FeedbacksUsersService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "FeedbacksUsersServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "FeedbacksUsersServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "FeedbacksUsersServiceClient" to call the correct interceptors.
func RegisterFeedbacksUsersServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client FeedbacksUsersServiceClient) error {

	mux.Handle("GET", pattern_FeedbacksUsersService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/feedbacks.FeedbacksUsersService/List", runtime.WithHTTPPathPattern("/api/v1/users"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FeedbacksUsersService_List_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FeedbacksUsersService_List_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_FeedbacksUsersService_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "users"}, ""))
)

var (
	forward_FeedbacksUsersService_List_0 = runtime.ForwardResponseMessage
)

// RegisterFeedbacksOperatorsServiceHandlerFromEndpoint is same as RegisterFeedbacksOperatorsServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterFeedbacksOperatorsServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterFeedbacksOperatorsServiceHandler(ctx, mux, conn)
}

// RegisterFeedbacksOperatorsServiceHandler registers the http handlers for service FeedbacksOperatorsService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterFeedbacksOperatorsServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterFeedbacksOperatorsServiceHandlerClient(ctx, mux, NewFeedbacksOperatorsServiceClient(conn))
}

// RegisterFeedbacksOperatorsServiceHandlerClient registers the http handlers for service FeedbacksOperatorsService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "FeedbacksOperatorsServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "FeedbacksOperatorsServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "FeedbacksOperatorsServiceClient" to call the correct interceptors.
func RegisterFeedbacksOperatorsServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client FeedbacksOperatorsServiceClient) error {

	mux.Handle("GET", pattern_FeedbacksOperatorsService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/feedbacks.FeedbacksOperatorsService/List", runtime.WithHTTPPathPattern("/api/v1/feedbacks/operators"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FeedbacksOperatorsService_List_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FeedbacksOperatorsService_List_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_FeedbacksOperatorsService_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "feedbacks", "operators"}, ""))
)

var (
	forward_FeedbacksOperatorsService_List_0 = runtime.ForwardResponseMessage
)