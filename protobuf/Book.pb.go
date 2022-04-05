// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: Book.proto

package service

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

// Msg: to response with the code
type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` // 1
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_Book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_Book_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

// ParamInt: to post with an integer
type ParamInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param int32 `protobuf:"varint,1,opt,name=param,proto3" json:"param,omitempty"` // 1
}

func (x *ParamInt) Reset() {
	*x = ParamInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamInt) ProtoMessage() {}

func (x *ParamInt) ProtoReflect() protoreflect.Message {
	mi := &file_Book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamInt.ProtoReflect.Descriptor instead.
func (*ParamInt) Descriptor() ([]byte, []int) {
	return file_Book_proto_rawDescGZIP(), []int{1}
}

func (x *ParamInt) GetParam() int32 {
	if x != nil {
		return x.Param
	}
	return 0
}

// ParamString: to post with a String
type ParamString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param string `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"` // 1
}

func (x *ParamString) Reset() {
	*x = ParamString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamString) ProtoMessage() {}

func (x *ParamString) ProtoReflect() protoreflect.Message {
	mi := &file_Book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamString.ProtoReflect.Descriptor instead.
func (*ParamString) Descriptor() ([]byte, []int) {
	return file_Book_proto_rawDescGZIP(), []int{2}
}

func (x *ParamString) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

// Book: define the model of book
type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`              // the id
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`         // the title
	Author    string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`       // the author
	Publisher string `protobuf:"bytes,4,opt,name=publisher,proto3" json:"publisher,omitempty"` // the publisher
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_Book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_Book_proto_rawDescGZIP(), []int{3}
}

func (x *Book) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

// Book: define a list of books
type BookList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookList []*Book `protobuf:"bytes,1,rep,name=book_list,json=bookList,proto3" json:"book_list,omitempty"`
}

func (x *BookList) Reset() {
	*x = BookList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookList) ProtoMessage() {}

func (x *BookList) ProtoReflect() protoreflect.Message {
	mi := &file_Book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookList.ProtoReflect.Descriptor instead.
func (*BookList) Descriptor() ([]byte, []int) {
	return file_Book_proto_rawDescGZIP(), []int{4}
}

func (x *BookList) GetBookList() []*Book {
	if x != nil {
		return x.BookList
	}
	return nil
}

var File_Book_proto protoreflect.FileDescriptor

var file_Book_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x19, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x20, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x22, 0x23, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x62, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x08, 0x42,
	0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x32, 0xc7, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x74, 0x1a, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x36, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x74, 0x1a,
	0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Book_proto_rawDescOnce sync.Once
	file_Book_proto_rawDescData = file_Book_proto_rawDesc
)

func file_Book_proto_rawDescGZIP() []byte {
	file_Book_proto_rawDescOnce.Do(func() {
		file_Book_proto_rawDescData = protoimpl.X.CompressGZIP(file_Book_proto_rawDescData)
	})
	return file_Book_proto_rawDescData
}

var file_Book_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Book_proto_goTypes = []interface{}{
	(*Msg)(nil),         // 0: service.Msg
	(*ParamInt)(nil),    // 1: service.ParamInt
	(*ParamString)(nil), // 2: service.ParamString
	(*Book)(nil),        // 3: service.Book
	(*BookList)(nil),    // 4: service.BookList
}
var file_Book_proto_depIdxs = []int32{
	3, // 0: service.BookList.book_list:type_name -> service.Book
	3, // 1: service.BookService.Add:input_type -> service.Book
	1, // 2: service.BookService.QueryByID:input_type -> service.ParamInt
	2, // 3: service.BookService.QueryByName:input_type -> service.ParamString
	1, // 4: service.BookService.DeleteById:input_type -> service.ParamInt
	0, // 5: service.BookService.Add:output_type -> service.Msg
	3, // 6: service.BookService.QueryByID:output_type -> service.Book
	4, // 7: service.BookService.QueryByName:output_type -> service.BookList
	0, // 8: service.BookService.DeleteById:output_type -> service.Msg
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Book_proto_init() }
func file_Book_proto_init() {
	if File_Book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_Book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamInt); i {
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
		file_Book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamString); i {
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
		file_Book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_Book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookList); i {
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
			RawDescriptor: file_Book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Book_proto_goTypes,
		DependencyIndexes: file_Book_proto_depIdxs,
		MessageInfos:      file_Book_proto_msgTypes,
	}.Build()
	File_Book_proto = out.File
	file_Book_proto_rawDesc = nil
	file_Book_proto_goTypes = nil
	file_Book_proto_depIdxs = nil
}