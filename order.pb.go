
package order

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	mi := &file_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}


func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type OrderStatusUpdate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderStatusUpdate) Reset() {
	*x = OrderStatusUpdate{}
	mi := &file_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderStatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatusUpdate) ProtoMessage() {}

func (x *OrderStatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}


func (*OrderStatusUpdate) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderStatusUpdate) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderStatusUpdate) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderStatusUpdate) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

const file_order_proto_rawDesc = "" +
	"\n" +
	"\vorder.proto\x12\x05order\x1a\x1fgoogle/protobuf/timestamp.proto\")\n" +
	"\fOrderRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"\x81\x01\n" +
	"\x11OrderStatusUpdate\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\x129\n" +
	"\n" +
	"updated_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt2Z\n" +
	"\fOrderService\x12J\n" +
	"\x17SubscribeToOrderUpdates\x12\x13.order.OrderRequest\x1a\x18.order.OrderStatusUpdate0\x01B\x16Z\x14paymentservice/orderb\x06proto3"

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData []byte
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)))
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_order_proto_goTypes = []any{
	(*OrderRequest)(nil),          
	(*OrderStatusUpdate)(nil),     
	(*timestamppb.Timestamp)(nil), 
}
var file_order_proto_depIdxs = []int32{
	2, 
	0,
	1, 
	2, 
	1, 
	1, 
	1, 
	0, 
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
