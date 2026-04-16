
package order

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)


const _ = grpc.SupportPackageIsVersion9

const (
	OrderService_SubscribeToOrderUpdates_FullMethodName = "/order.OrderService/SubscribeToOrderUpdates"
)


type OrderServiceClient interface {
	SubscribeToOrderUpdates(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderStatusUpdate], error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) SubscribeToOrderUpdates(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderStatusUpdate], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], OrderService_SubscribeToOrderUpdates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[OrderRequest, OrderStatusUpdate]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}


type OrderService_SubscribeToOrderUpdatesClient = grpc.ServerStreamingClient[OrderStatusUpdate]


type OrderServiceServer interface {
	SubscribeToOrderUpdates(*OrderRequest, grpc.ServerStreamingServer[OrderStatusUpdate]) error
	mustEmbedUnimplementedOrderServiceServer()
}

type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) SubscribeToOrderUpdates(*OrderRequest, grpc.ServerStreamingServer[OrderStatusUpdate]) error {
	return status.Error(codes.Unimplemented, "method SubscribeToOrderUpdates not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}
func (UnimplementedOrderServiceServer) testEmbeddedByValue()                      {}


type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_SubscribeToOrderUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).SubscribeToOrderUpdates(m, &grpc.GenericServerStream[OrderRequest, OrderStatusUpdate]{ServerStream: stream})
}


type OrderService_SubscribeToOrderUpdatesServer = grpc.ServerStreamingServer[OrderStatusUpdate]


var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToOrderUpdates",
			Handler:       _OrderService_SubscribeToOrderUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "order.proto",
}
