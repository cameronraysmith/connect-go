// Code generated by protoc-gen-go-connect. DO NOT EDIT.
// versions:
// - protoc-gen-go-connect v0.0.1
// - protoc              v3.17.3
// source: connect/ping/v1test/ping.proto

package pingv1test

import (
	context "context"
	errors "errors"
	connect "github.com/bufbuild/connect"
	v1test "github.com/bufbuild/connect/internal/gen/proto/go/connect/ping/v1test"
	http "net/http"
	path "path"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the
// connect package are compatible. If you get a compiler error that this
// constant isn't defined, this code was generated with a version of connect
// newer than the one compiled into your binary. You can fix the problem by
// either regenerating this code with an older version of connect or updating
// the connect version compiled into your binary.
const _ = connect.IsAtLeastVersion0_0_1

// PingServiceClient is a client for the connect.ping.v1test.PingService
// service.
type PingServiceClient interface {
	// Ping sends a ping to the server to determine if it's reachable.
	Ping(context.Context, *connect.Envelope[v1test.PingRequest]) (*connect.Envelope[v1test.PingResponse], error)
	// Fail always fails.
	Fail(context.Context, *connect.Envelope[v1test.FailRequest]) (*connect.Envelope[v1test.FailResponse], error)
	// Sum calculates the sum of the numbers sent on the stream.
	Sum(context.Context) *connect.ClientStreamForClient[v1test.SumRequest, v1test.SumResponse]
	// CountUp returns a stream of the numbers up to the given request.
	CountUp(context.Context, *connect.Envelope[v1test.CountUpRequest]) (*connect.ServerStreamForClient[v1test.CountUpResponse], error)
	// CumSum determines the cumulative sum of all the numbers sent on the stream.
	CumSum(context.Context) *connect.BidiStreamForClient[v1test.CumSumRequest, v1test.CumSumResponse]
}

// NewPingServiceClient constructs a client for the
// connect.ping.v1test.PingService service. By default, it uses the binary
// protobuf codec.
//
// The URL supplied here should be the base URL for the gRPC server (e.g.,
// https://api.acme.com or https://acme.com/grpc).
func NewPingServiceClient(baseURL string, doer connect.Doer, opts ...connect.ClientOption) (PingServiceClient, error) {
	baseURL = strings.TrimRight(baseURL, "/")
	opts = append([]connect.ClientOption{
		connect.WithProtobufCodec(),
		connect.WithGzip(),
	}, opts...)
	var (
		client pingServiceClient
		err    error
	)
	client.ping, err = connect.NewUnaryClientImplementation[v1test.PingRequest, v1test.PingResponse](
		doer,
		baseURL,
		"connect.ping.v1test.PingService/Ping",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	client.fail, err = connect.NewUnaryClientImplementation[v1test.FailRequest, v1test.FailResponse](
		doer,
		baseURL,
		"connect.ping.v1test.PingService/Fail",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	client.sum, err = connect.NewStreamClientImplementation(
		doer,
		baseURL,
		"connect.ping.v1test.PingService/Sum",
		connect.StreamTypeClient,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	client.countUp, err = connect.NewStreamClientImplementation(
		doer,
		baseURL,
		"connect.ping.v1test.PingService/CountUp",
		connect.StreamTypeServer,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	client.cumSum, err = connect.NewStreamClientImplementation(
		doer,
		baseURL,
		"connect.ping.v1test.PingService/CumSum",
		connect.StreamTypeBidirectional,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

// pingServiceClient implements PingServiceClient.
type pingServiceClient struct {
	ping    func(context.Context, *connect.Envelope[v1test.PingRequest]) (*connect.Envelope[v1test.PingResponse], error)
	fail    func(context.Context, *connect.Envelope[v1test.FailRequest]) (*connect.Envelope[v1test.FailResponse], error)
	sum     func(context.Context) (connect.Sender, connect.Receiver)
	countUp func(context.Context) (connect.Sender, connect.Receiver)
	cumSum  func(context.Context) (connect.Sender, connect.Receiver)
}

var _ PingServiceClient = (*pingServiceClient)(nil) // verify interface implementation

// Ping calls connect.ping.v1test.PingService.Ping.
func (c *pingServiceClient) Ping(ctx context.Context, req *connect.Envelope[v1test.PingRequest]) (*connect.Envelope[v1test.PingResponse], error) {
	return c.ping(ctx, req)
}

// Fail calls connect.ping.v1test.PingService.Fail.
func (c *pingServiceClient) Fail(ctx context.Context, req *connect.Envelope[v1test.FailRequest]) (*connect.Envelope[v1test.FailResponse], error) {
	return c.fail(ctx, req)
}

// Sum calls connect.ping.v1test.PingService.Sum.
func (c *pingServiceClient) Sum(ctx context.Context) *connect.ClientStreamForClient[v1test.SumRequest, v1test.SumResponse] {
	sender, receiver := c.sum(ctx)
	return connect.NewClientStreamForClient[v1test.SumRequest, v1test.SumResponse](sender, receiver)
}

// CountUp calls connect.ping.v1test.PingService.CountUp.
func (c *pingServiceClient) CountUp(ctx context.Context, req *connect.Envelope[v1test.CountUpRequest]) (*connect.ServerStreamForClient[v1test.CountUpResponse], error) {
	sender, receiver := c.countUp(ctx)
	for key, values := range req.Header() {
		sender.Header()[key] = append(sender.Header()[key], values...)
	}
	for key, values := range req.Trailer() {
		sender.Trailer()[key] = append(sender.Trailer()[key], values...)
	}
	if err := sender.Send(req.Msg); err != nil {
		_ = sender.Close(err)
		_ = receiver.Close()
		return nil, err
	}
	if err := sender.Close(nil); err != nil {
		_ = receiver.Close()
		return nil, err
	}
	return connect.NewServerStreamForClient[v1test.CountUpResponse](receiver), nil
}

// CumSum calls connect.ping.v1test.PingService.CumSum.
func (c *pingServiceClient) CumSum(ctx context.Context) *connect.BidiStreamForClient[v1test.CumSumRequest, v1test.CumSumResponse] {
	sender, receiver := c.cumSum(ctx)
	return connect.NewBidiStreamForClient[v1test.CumSumRequest, v1test.CumSumResponse](sender, receiver)
}

// PingServiceHandler is an implementation of the
// connect.ping.v1test.PingService service.
type PingServiceHandler interface {
	// Ping sends a ping to the server to determine if it's reachable.
	Ping(context.Context, *connect.Envelope[v1test.PingRequest]) (*connect.Envelope[v1test.PingResponse], error)
	// Fail always fails.
	Fail(context.Context, *connect.Envelope[v1test.FailRequest]) (*connect.Envelope[v1test.FailResponse], error)
	// Sum calculates the sum of the numbers sent on the stream.
	Sum(context.Context, *connect.ClientStream[v1test.SumRequest, v1test.SumResponse]) error
	// CountUp returns a stream of the numbers up to the given request.
	CountUp(context.Context, *connect.Envelope[v1test.CountUpRequest], *connect.ServerStream[v1test.CountUpResponse]) error
	// CumSum determines the cumulative sum of all the numbers sent on the stream.
	CumSum(context.Context, *connect.BidiStream[v1test.CumSumRequest, v1test.CumSumResponse]) error
}

// NewPingServiceHandler builds an HTTP handler from the service implementation.
// It returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the gRPC and gRPC-Web protocols with the binary
// protobuf and JSON codecs.
func NewPingServiceHandler(svc PingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	var lastHandlerPath string
	mux := http.NewServeMux()
	opts = append([]connect.HandlerOption{
		connect.WithProtobufCodec(),
		connect.WithProtobufJSONCodec(),
		connect.WithGzip(),
	}, opts...)

	ping := connect.NewUnaryHandler(
		"connect.ping.v1test.PingService/Ping", // procedure name
		"connect.ping.v1test.PingService",      // reflection name
		svc.Ping,
		opts...,
	)
	mux.Handle(ping.Path(), ping)
	lastHandlerPath = ping.Path()

	fail := connect.NewUnaryHandler(
		"connect.ping.v1test.PingService/Fail", // procedure name
		"connect.ping.v1test.PingService",      // reflection name
		svc.Fail,
		opts...,
	)
	mux.Handle(fail.Path(), fail)
	lastHandlerPath = fail.Path()

	sum := connect.NewStreamHandler(
		"connect.ping.v1test.PingService/Sum", // procedure name
		"connect.ping.v1test.PingService",     // reflection name
		connect.StreamTypeClient,
		func(ctx context.Context, sender connect.Sender, receiver connect.Receiver) {
			typed := connect.NewClientStream[v1test.SumRequest, v1test.SumResponse](sender, receiver)
			err := svc.Sum(ctx, typed)
			_ = receiver.Close()
			_ = sender.Close(err)
		},
		opts...,
	)
	mux.Handle(sum.Path(), sum)
	lastHandlerPath = sum.Path()

	countUp := connect.NewStreamHandler(
		"connect.ping.v1test.PingService/CountUp", // procedure name
		"connect.ping.v1test.PingService",         // reflection name
		connect.StreamTypeServer,
		func(ctx context.Context, sender connect.Sender, receiver connect.Receiver) {
			typed := connect.NewServerStream[v1test.CountUpResponse](sender)
			req, err := connect.ReceiveUnaryEnvelope[v1test.CountUpRequest](receiver)
			if err != nil {
				_ = receiver.Close()
				_ = sender.Close(err)
				return
			}
			if err = receiver.Close(); err != nil {
				_ = sender.Close(err)
				return
			}
			err = svc.CountUp(ctx, req, typed)
			_ = sender.Close(err)
		},
		opts...,
	)
	mux.Handle(countUp.Path(), countUp)
	lastHandlerPath = countUp.Path()

	cumSum := connect.NewStreamHandler(
		"connect.ping.v1test.PingService/CumSum", // procedure name
		"connect.ping.v1test.PingService",        // reflection name
		connect.StreamTypeBidirectional,
		func(ctx context.Context, sender connect.Sender, receiver connect.Receiver) {
			typed := connect.NewBidiStream[v1test.CumSumRequest, v1test.CumSumResponse](sender, receiver)
			err := svc.CumSum(ctx, typed)
			_ = receiver.Close()
			_ = sender.Close(err)
		},
		opts...,
	)
	mux.Handle(cumSum.Path(), cumSum)
	lastHandlerPath = cumSum.Path()

	return path.Dir(lastHandlerPath) + "/", mux
}

// UnimplementedPingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPingServiceHandler struct{}

var _ PingServiceHandler = (*UnimplementedPingServiceHandler)(nil) // verify interface implementation

func (UnimplementedPingServiceHandler) Ping(context.Context, *connect.Envelope[v1test.PingRequest]) (*connect.Envelope[v1test.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("connect.ping.v1test.PingService.Ping isn't implemented"))
}

func (UnimplementedPingServiceHandler) Fail(context.Context, *connect.Envelope[v1test.FailRequest]) (*connect.Envelope[v1test.FailResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("connect.ping.v1test.PingService.Fail isn't implemented"))
}

func (UnimplementedPingServiceHandler) Sum(context.Context, *connect.ClientStream[v1test.SumRequest, v1test.SumResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("connect.ping.v1test.PingService.Sum isn't implemented"))
}

func (UnimplementedPingServiceHandler) CountUp(context.Context, *connect.Envelope[v1test.CountUpRequest], *connect.ServerStream[v1test.CountUpResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("connect.ping.v1test.PingService.CountUp isn't implemented"))
}

func (UnimplementedPingServiceHandler) CumSum(context.Context, *connect.BidiStream[v1test.CumSumRequest, v1test.CumSumResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("connect.ping.v1test.PingService.CumSum isn't implemented"))
}
