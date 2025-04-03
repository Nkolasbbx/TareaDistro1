// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: comunicacion.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Representación de un pirata
type Pirata struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Id                  int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                  // ID único del pirata
	Nombre              string                 `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`                           // Nombre del pirata
	Recompensa          int64                  `protobuf:"varint,3,opt,name=recompensa,proto3" json:"recompensa,omitempty"`                  // Recompensa ofrecida por el pirata
	NivelDePeligrosidad string                 `protobuf:"bytes,4,opt,name=nivelDePeligrosidad,proto3" json:"nivelDePeligrosidad,omitempty"` // Nivel de peligrosidad del pirata
	Estado              string                 `protobuf:"bytes,5,opt,name=estado,proto3" json:"estado,omitempty"`                           // Estado del pirata (Buscado, Capturado, Vendido, etc.)
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Pirata) Reset() {
	*x = Pirata{}
	mi := &file_comunicacion_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pirata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pirata) ProtoMessage() {}

func (x *Pirata) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pirata.ProtoReflect.Descriptor instead.
func (*Pirata) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{0}
}

func (x *Pirata) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pirata) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Pirata) GetRecompensa() int64 {
	if x != nil {
		return x.Recompensa
	}
	return 0
}

func (x *Pirata) GetNivelDePeligrosidad() string {
	if x != nil {
		return x.NivelDePeligrosidad
	}
	return ""
}

func (x *Pirata) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

// Solicitud para vender un pirata
type VentaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PirataId      int64                  `protobuf:"varint,1,opt,name=pirata_id,json=pirataId,proto3" json:"pirata_id,omitempty"`    // ID del pirata a vender
	CazadorId     int64                  `protobuf:"varint,2,opt,name=cazador_id,json=cazadorId,proto3" json:"cazador_id,omitempty"` // ID del cazarrecompensas que realiza la venta
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VentaRequest) Reset() {
	*x = VentaRequest{}
	mi := &file_comunicacion_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VentaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VentaRequest) ProtoMessage() {}

func (x *VentaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VentaRequest.ProtoReflect.Descriptor instead.
func (*VentaRequest) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{1}
}

func (x *VentaRequest) GetPirataId() int64 {
	if x != nil {
		return x.PirataId
	}
	return 0
}

func (x *VentaRequest) GetCazadorId() int64 {
	if x != nil {
		return x.CazadorId
	}
	return 0
}

// Respuesta de la venta
type VentaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exito         bool                   `protobuf:"varint,1,opt,name=exito,proto3" json:"exito,omitempty"`    // Indica si la venta fue exitosa
	Pago          int64                  `protobuf:"varint,2,opt,name=pago,proto3" json:"pago,omitempty"`      // Cantidad pagada al cazarrecompensas
	Mensaje       string                 `protobuf:"bytes,3,opt,name=mensaje,proto3" json:"mensaje,omitempty"` // Mensaje adicional (fraude, éxito, etc.)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VentaResponse) Reset() {
	*x = VentaResponse{}
	mi := &file_comunicacion_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VentaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VentaResponse) ProtoMessage() {}

func (x *VentaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VentaResponse.ProtoReflect.Descriptor instead.
func (*VentaResponse) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{2}
}

func (x *VentaResponse) GetExito() bool {
	if x != nil {
		return x.Exito
	}
	return false
}

func (x *VentaResponse) GetPago() int64 {
	if x != nil {
		return x.Pago
	}
	return 0
}

func (x *VentaResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

// Respuesta de la lista de piratas
type Respuesta struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Piratas       []*Pirata              `protobuf:"bytes,1,rep,name=piratas,proto3" json:"piratas,omitempty"` // Lista de piratas
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Respuesta) Reset() {
	*x = Respuesta{}
	mi := &file_comunicacion_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Respuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Respuesta) ProtoMessage() {}

func (x *Respuesta) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Respuesta.ProtoReflect.Descriptor instead.
func (*Respuesta) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{3}
}

func (x *Respuesta) GetPiratas() []*Pirata {
	if x != nil {
		return x.Piratas
	}
	return nil
}

// Mensaje vacío para solicitudes sin parámetros
type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_comunicacion_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{4}
}

var File_comunicacion_proto protoreflect.FileDescriptor

const file_comunicacion_proto_rawDesc = "" +
	"\n" +
	"\x12comunicacion.proto\x12\fcomunicacion\"\x9a\x01\n" +
	"\x06Pirata\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x16\n" +
	"\x06nombre\x18\x02 \x01(\tR\x06nombre\x12\x1e\n" +
	"\n" +
	"recompensa\x18\x03 \x01(\x03R\n" +
	"recompensa\x120\n" +
	"\x13nivelDePeligrosidad\x18\x04 \x01(\tR\x13nivelDePeligrosidad\x12\x16\n" +
	"\x06estado\x18\x05 \x01(\tR\x06estado\"J\n" +
	"\fVentaRequest\x12\x1b\n" +
	"\tpirata_id\x18\x01 \x01(\x03R\bpirataId\x12\x1d\n" +
	"\n" +
	"cazador_id\x18\x02 \x01(\x03R\tcazadorId\"S\n" +
	"\rVentaResponse\x12\x14\n" +
	"\x05exito\x18\x01 \x01(\bR\x05exito\x12\x12\n" +
	"\x04pago\x18\x02 \x01(\x03R\x04pago\x12\x18\n" +
	"\amensaje\x18\x03 \x01(\tR\amensaje\";\n" +
	"\tRespuesta\x12.\n" +
	"\apiratas\x18\x01 \x03(\v2\x14.comunicacion.PirataR\apiratas\"\a\n" +
	"\x05empty2\xf5\x01\n" +
	"\fComunicacion\x12E\n" +
	"\x15consultarListaPiratas\x12\x13.comunicacion.empty\x1a\x17.comunicacion.Respuesta\x12O\n" +
	"\x14venderPirataSubmundo\x12\x1a.comunicacion.VentaRequest\x1a\x1b.comunicacion.VentaResponse\x12M\n" +
	"\x12venderPirataMarina\x12\x1a.comunicacion.VentaRequest\x1a\x1b.comunicacion.VentaResponseB\x13Z\x11grpc-server/protob\x06proto3"

var (
	file_comunicacion_proto_rawDescOnce sync.Once
	file_comunicacion_proto_rawDescData []byte
)

func file_comunicacion_proto_rawDescGZIP() []byte {
	file_comunicacion_proto_rawDescOnce.Do(func() {
		file_comunicacion_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_comunicacion_proto_rawDesc), len(file_comunicacion_proto_rawDesc)))
	})
	return file_comunicacion_proto_rawDescData
}

var file_comunicacion_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_comunicacion_proto_goTypes = []any{
	(*Pirata)(nil),        // 0: comunicacion.Pirata
	(*VentaRequest)(nil),  // 1: comunicacion.VentaRequest
	(*VentaResponse)(nil), // 2: comunicacion.VentaResponse
	(*Respuesta)(nil),     // 3: comunicacion.Respuesta
	(*Empty)(nil),         // 4: comunicacion.empty
}
var file_comunicacion_proto_depIdxs = []int32{
	0, // 0: comunicacion.Respuesta.piratas:type_name -> comunicacion.Pirata
	4, // 1: comunicacion.Comunicacion.consultarListaPiratas:input_type -> comunicacion.empty
	1, // 2: comunicacion.Comunicacion.venderPirataSubmundo:input_type -> comunicacion.VentaRequest
	1, // 3: comunicacion.Comunicacion.venderPirataMarina:input_type -> comunicacion.VentaRequest
	3, // 4: comunicacion.Comunicacion.consultarListaPiratas:output_type -> comunicacion.Respuesta
	2, // 5: comunicacion.Comunicacion.venderPirataSubmundo:output_type -> comunicacion.VentaResponse
	2, // 6: comunicacion.Comunicacion.venderPirataMarina:output_type -> comunicacion.VentaResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_comunicacion_proto_init() }
func file_comunicacion_proto_init() {
	if File_comunicacion_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_comunicacion_proto_rawDesc), len(file_comunicacion_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comunicacion_proto_goTypes,
		DependencyIndexes: file_comunicacion_proto_depIdxs,
		MessageInfos:      file_comunicacion_proto_msgTypes,
	}.Build()
	File_comunicacion_proto = out.File
	file_comunicacion_proto_goTypes = nil
	file_comunicacion_proto_depIdxs = nil
}
