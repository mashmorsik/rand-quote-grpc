// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: rand_quote.proto

package randquotev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Name int32

const (
	Name_NAME_ALBUS_DUMBLEDORE     Name = 0
	Name_NAME_HERMIONE_GRANGER     Name = 1
	Name_NAME_RON_WEASLEY          Name = 2
	Name_NAME_SEVERUS_SNAPE        Name = 3
	Name_NAME_PROFESSOR_MCGONAGALL Name = 4
	Name_NAME_LUNA_LOVEGOOD        Name = 5
	Name_NAME_HAGRID               Name = 6
)

// Enum value maps for Name.
var (
	Name_name = map[int32]string{
		0: "NAME_ALBUS_DUMBLEDORE",
		1: "NAME_HERMIONE_GRANGER",
		2: "NAME_RON_WEASLEY",
		3: "NAME_SEVERUS_SNAPE",
		4: "NAME_PROFESSOR_MCGONAGALL",
		5: "NAME_LUNA_LOVEGOOD",
		6: "NAME_HAGRID",
	}
	Name_value = map[string]int32{
		"NAME_ALBUS_DUMBLEDORE":     0,
		"NAME_HERMIONE_GRANGER":     1,
		"NAME_RON_WEASLEY":          2,
		"NAME_SEVERUS_SNAPE":        3,
		"NAME_PROFESSOR_MCGONAGALL": 4,
		"NAME_LUNA_LOVEGOOD":        5,
		"NAME_HAGRID":               6,
	}
)

func (x Name) Enum() *Name {
	p := new(Name)
	*p = x
	return p
}

func (x Name) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Name) Descriptor() protoreflect.EnumDescriptor {
	return file_rand_quote_proto_enumTypes[0].Descriptor()
}

func (Name) Type() protoreflect.EnumType {
	return &file_rand_quote_proto_enumTypes[0]
}

func (x Name) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Name.Descriptor instead.
func (Name) EnumDescriptor() ([]byte, []int) {
	return file_rand_quote_proto_rawDescGZIP(), []int{0}
}

type ListQuotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name Name `protobuf:"varint,1,opt,name=name,proto3,enum=github.com.mashmorsik.rand.quote.grpc.v1.Name" json:"name,omitempty"`
}

func (x *ListQuotesRequest) Reset() {
	*x = ListQuotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rand_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuotesRequest) ProtoMessage() {}

func (x *ListQuotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rand_quote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuotesRequest.ProtoReflect.Descriptor instead.
func (*ListQuotesRequest) Descriptor() ([]byte, []int) {
	return file_rand_quote_proto_rawDescGZIP(), []int{0}
}

func (x *ListQuotesRequest) GetName() Name {
	if x != nil {
		return x.Name
	}
	return Name_NAME_ALBUS_DUMBLEDORE
}

type ListQuotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quotes string `protobuf:"bytes,1,opt,name=quotes,proto3" json:"quotes,omitempty"`
}

func (x *ListQuotesResponse) Reset() {
	*x = ListQuotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rand_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuotesResponse) ProtoMessage() {}

func (x *ListQuotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rand_quote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuotesResponse.ProtoReflect.Descriptor instead.
func (*ListQuotesResponse) Descriptor() ([]byte, []int) {
	return file_rand_quote_proto_rawDescGZIP(), []int{1}
}

func (x *ListQuotesResponse) GetQuotes() string {
	if x != nil {
		return x.Quotes
	}
	return ""
}

type RandQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name Name `protobuf:"varint,1,opt,name=name,proto3,enum=github.com.mashmorsik.rand.quote.grpc.v1.Name" json:"name,omitempty"`
}

func (x *RandQuoteRequest) Reset() {
	*x = RandQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rand_quote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandQuoteRequest) ProtoMessage() {}

func (x *RandQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rand_quote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandQuoteRequest.ProtoReflect.Descriptor instead.
func (*RandQuoteRequest) Descriptor() ([]byte, []int) {
	return file_rand_quote_proto_rawDescGZIP(), []int{2}
}

func (x *RandQuoteRequest) GetName() Name {
	if x != nil {
		return x.Name
	}
	return Name_NAME_ALBUS_DUMBLEDORE
}

type RandQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote string `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *RandQuoteResponse) Reset() {
	*x = RandQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rand_quote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandQuoteResponse) ProtoMessage() {}

func (x *RandQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rand_quote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandQuoteResponse.ProtoReflect.Descriptor instead.
func (*RandQuoteResponse) Descriptor() ([]byte, []int) {
	return file_rand_quote_proto_rawDescGZIP(), []int{3}
}

func (x *RandQuoteResponse) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

var File_rand_quote_proto protoreflect.FileDescriptor

var file_rand_quote_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x61, 0x73, 0x68, 0x6d, 0x6f, 0x72, 0x73, 0x69, 0x6b, 0x2e, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x22, 0x57, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x42, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x73,
	0x68, 0x6d, 0x6f, 0x72, 0x73, 0x69, 0x6b, 0x2e, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x73, 0x22, 0x56, 0x0a, 0x10, 0x52, 0x61, 0x6e, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x73, 0x68, 0x6d, 0x6f, 0x72, 0x73, 0x69, 0x6b, 0x2e, 0x72, 0x61,
	0x6e, 0x64, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x52,
	0x61, 0x6e, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2a, 0xb2, 0x01, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x15, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x4c, 0x42, 0x55, 0x53, 0x5f, 0x44, 0x55,
	0x4d, 0x42, 0x4c, 0x45, 0x44, 0x4f, 0x52, 0x45, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x48, 0x45, 0x52, 0x4d, 0x49, 0x4f, 0x4e, 0x45, 0x5f, 0x47, 0x52, 0x41, 0x4e,
	0x47, 0x45, 0x52, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x4f,
	0x4e, 0x5f, 0x57, 0x45, 0x41, 0x53, 0x4c, 0x45, 0x59, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x55, 0x53, 0x5f, 0x53, 0x4e, 0x41, 0x50,
	0x45, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x46,
	0x45, 0x53, 0x53, 0x4f, 0x52, 0x5f, 0x4d, 0x43, 0x47, 0x4f, 0x4e, 0x41, 0x47, 0x41, 0x4c, 0x4c,
	0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4c, 0x55, 0x4e, 0x41, 0x5f,
	0x4c, 0x4f, 0x56, 0x45, 0x47, 0x4f, 0x4f, 0x44, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x48, 0x41, 0x47, 0x52, 0x49, 0x44, 0x10, 0x06, 0x32, 0xa0, 0x02, 0x0a, 0x0a,
	0x52, 0x61, 0x6e, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x83, 0x01, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x3a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x73, 0x68, 0x6d, 0x6f, 0x72, 0x73, 0x69, 0x6b, 0x2e,
	0x72, 0x61, 0x6e, 0x64, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x6d, 0x61, 0x73, 0x68, 0x6d, 0x6f, 0x72, 0x73, 0x69, 0x6b, 0x2e, 0x72, 0x61, 0x6e, 0x64,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x61, 0x6e, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x8b, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x3b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x73,
	0x68, 0x6d, 0x6f, 0x72, 0x73, 0x69, 0x6b, 0x2e, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x73, 0x68, 0x6d, 0x6f,
	0x72, 0x73, 0x69, 0x6b, 0x2e, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x73,
	0x68, 0x6d, 0x6f, 0x72, 0x73, 0x69, 0x6b, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x2d, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x72, 0x61, 0x6e, 0x64, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rand_quote_proto_rawDescOnce sync.Once
	file_rand_quote_proto_rawDescData = file_rand_quote_proto_rawDesc
)

func file_rand_quote_proto_rawDescGZIP() []byte {
	file_rand_quote_proto_rawDescOnce.Do(func() {
		file_rand_quote_proto_rawDescData = protoimpl.X.CompressGZIP(file_rand_quote_proto_rawDescData)
	})
	return file_rand_quote_proto_rawDescData
}

var file_rand_quote_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rand_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rand_quote_proto_goTypes = []interface{}{
	(Name)(0),                  // 0: github.com.mashmorsik.rand.quote.grpc.v1.Name
	(*ListQuotesRequest)(nil),  // 1: github.com.mashmorsik.rand.quote.grpc.v1.ListQuotesRequest
	(*ListQuotesResponse)(nil), // 2: github.com.mashmorsik.rand.quote.grpc.v1.ListQuotesResponse
	(*RandQuoteRequest)(nil),   // 3: github.com.mashmorsik.rand.quote.grpc.v1.RandQuoteRequest
	(*RandQuoteResponse)(nil),  // 4: github.com.mashmorsik.rand.quote.grpc.v1.RandQuoteResponse
}
var file_rand_quote_proto_depIdxs = []int32{
	0, // 0: github.com.mashmorsik.rand.quote.grpc.v1.ListQuotesRequest.name:type_name -> github.com.mashmorsik.rand.quote.grpc.v1.Name
	0, // 1: github.com.mashmorsik.rand.quote.grpc.v1.RandQuoteRequest.name:type_name -> github.com.mashmorsik.rand.quote.grpc.v1.Name
	3, // 2: github.com.mashmorsik.rand.quote.grpc.v1.RandQuotes.GetQuote:input_type -> github.com.mashmorsik.rand.quote.grpc.v1.RandQuoteRequest
	1, // 3: github.com.mashmorsik.rand.quote.grpc.v1.RandQuotes.ListQuotes:input_type -> github.com.mashmorsik.rand.quote.grpc.v1.ListQuotesRequest
	4, // 4: github.com.mashmorsik.rand.quote.grpc.v1.RandQuotes.GetQuote:output_type -> github.com.mashmorsik.rand.quote.grpc.v1.RandQuoteResponse
	2, // 5: github.com.mashmorsik.rand.quote.grpc.v1.RandQuotes.ListQuotes:output_type -> github.com.mashmorsik.rand.quote.grpc.v1.ListQuotesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rand_quote_proto_init() }
func file_rand_quote_proto_init() {
	if File_rand_quote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rand_quote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuotesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rand_quote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuotesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rand_quote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandQuoteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rand_quote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandQuoteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rand_quote_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rand_quote_proto_goTypes,
		DependencyIndexes: file_rand_quote_proto_depIdxs,
		EnumInfos:         file_rand_quote_proto_enumTypes,
		MessageInfos:      file_rand_quote_proto_msgTypes,
	}.Build()
	File_rand_quote_proto = out.File
	file_rand_quote_proto_rawDesc = nil
	file_rand_quote_proto_goTypes = nil
	file_rand_quote_proto_depIdxs = nil
}
