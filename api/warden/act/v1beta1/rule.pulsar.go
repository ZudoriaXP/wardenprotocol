// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package actv1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	ast "github.com/warden-protocol/wardenprotocol/api/shield/ast"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Rule            protoreflect.MessageDescriptor
	fd_Rule_id         protoreflect.FieldDescriptor
	fd_Rule_creator    protoreflect.FieldDescriptor
	fd_Rule_name       protoreflect.FieldDescriptor
	fd_Rule_expression protoreflect.FieldDescriptor
)

func init() {
	file_warden_act_v1beta1_rule_proto_init()
	md_Rule = File_warden_act_v1beta1_rule_proto.Messages().ByName("Rule")
	fd_Rule_id = md_Rule.Fields().ByName("id")
	fd_Rule_creator = md_Rule.Fields().ByName("creator")
	fd_Rule_name = md_Rule.Fields().ByName("name")
	fd_Rule_expression = md_Rule.Fields().ByName("expression")
}

var _ protoreflect.Message = (*fastReflection_Rule)(nil)

type fastReflection_Rule Rule

func (x *Rule) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Rule)(x)
}

func (x *Rule) slowProtoReflect() protoreflect.Message {
	mi := &file_warden_act_v1beta1_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Rule_messageType fastReflection_Rule_messageType
var _ protoreflect.MessageType = fastReflection_Rule_messageType{}

type fastReflection_Rule_messageType struct{}

func (x fastReflection_Rule_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Rule)(nil)
}
func (x fastReflection_Rule_messageType) New() protoreflect.Message {
	return new(fastReflection_Rule)
}
func (x fastReflection_Rule_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Rule
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Rule) Descriptor() protoreflect.MessageDescriptor {
	return md_Rule
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Rule) Type() protoreflect.MessageType {
	return _fastReflection_Rule_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Rule) New() protoreflect.Message {
	return new(fastReflection_Rule)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Rule) Interface() protoreflect.ProtoMessage {
	return (*Rule)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Rule) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Rule_id, value) {
			return
		}
	}
	if x.Creator != "" {
		value := protoreflect.ValueOfString(x.Creator)
		if !f(fd_Rule_creator, value) {
			return
		}
	}
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_Rule_name, value) {
			return
		}
	}
	if x.Expression != nil {
		value := protoreflect.ValueOfMessage(x.Expression.ProtoReflect())
		if !f(fd_Rule_expression, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Rule) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "warden.act.v1beta1.Rule.id":
		return x.Id != uint64(0)
	case "warden.act.v1beta1.Rule.creator":
		return x.Creator != ""
	case "warden.act.v1beta1.Rule.name":
		return x.Name != ""
	case "warden.act.v1beta1.Rule.expression":
		return x.Expression != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.act.v1beta1.Rule"))
		}
		panic(fmt.Errorf("message warden.act.v1beta1.Rule does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rule) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "warden.act.v1beta1.Rule.id":
		x.Id = uint64(0)
	case "warden.act.v1beta1.Rule.creator":
		x.Creator = ""
	case "warden.act.v1beta1.Rule.name":
		x.Name = ""
	case "warden.act.v1beta1.Rule.expression":
		x.Expression = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.act.v1beta1.Rule"))
		}
		panic(fmt.Errorf("message warden.act.v1beta1.Rule does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Rule) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "warden.act.v1beta1.Rule.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "warden.act.v1beta1.Rule.creator":
		value := x.Creator
		return protoreflect.ValueOfString(value)
	case "warden.act.v1beta1.Rule.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "warden.act.v1beta1.Rule.expression":
		value := x.Expression
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.act.v1beta1.Rule"))
		}
		panic(fmt.Errorf("message warden.act.v1beta1.Rule does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rule) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "warden.act.v1beta1.Rule.id":
		x.Id = value.Uint()
	case "warden.act.v1beta1.Rule.creator":
		x.Creator = value.Interface().(string)
	case "warden.act.v1beta1.Rule.name":
		x.Name = value.Interface().(string)
	case "warden.act.v1beta1.Rule.expression":
		x.Expression = value.Message().Interface().(*ast.Expression)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.act.v1beta1.Rule"))
		}
		panic(fmt.Errorf("message warden.act.v1beta1.Rule does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rule) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "warden.act.v1beta1.Rule.expression":
		if x.Expression == nil {
			x.Expression = new(ast.Expression)
		}
		return protoreflect.ValueOfMessage(x.Expression.ProtoReflect())
	case "warden.act.v1beta1.Rule.id":
		panic(fmt.Errorf("field id of message warden.act.v1beta1.Rule is not mutable"))
	case "warden.act.v1beta1.Rule.creator":
		panic(fmt.Errorf("field creator of message warden.act.v1beta1.Rule is not mutable"))
	case "warden.act.v1beta1.Rule.name":
		panic(fmt.Errorf("field name of message warden.act.v1beta1.Rule is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.act.v1beta1.Rule"))
		}
		panic(fmt.Errorf("message warden.act.v1beta1.Rule does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Rule) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "warden.act.v1beta1.Rule.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "warden.act.v1beta1.Rule.creator":
		return protoreflect.ValueOfString("")
	case "warden.act.v1beta1.Rule.name":
		return protoreflect.ValueOfString("")
	case "warden.act.v1beta1.Rule.expression":
		m := new(ast.Expression)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.act.v1beta1.Rule"))
		}
		panic(fmt.Errorf("message warden.act.v1beta1.Rule does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Rule) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in warden.act.v1beta1.Rule", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Rule) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rule) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Rule) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Rule) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Rule)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Creator)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Expression != nil {
			l = options.Size(x.Expression)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Rule)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Expression != nil {
			encoded, err := options.Marshal(x.Expression)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Creator) > 0 {
			i -= len(x.Creator)
			copy(dAtA[i:], x.Creator)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Creator)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Rule)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Rule: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Rule: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Creator = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Expression", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Expression == nil {
					x.Expression = &ast.Expression{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Expression); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: warden/act/v1beta1/rule.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The expression to be evaluated for this rule.
	Expression *ast.Expression `protobuf:"bytes,4,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warden_act_v1beta1_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_warden_act_v1beta1_rule_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rule) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetExpression() *ast.Expression {
	if x != nil {
		return x.Expression
	}
	return nil
}

var File_warden_act_v1beta1_rule_proto protoreflect.FileDescriptor

var file_warden_act_v1beta1_rule_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2f, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x61, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0xda, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x2e, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x09, 0x52,
	0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e,
	0x2f, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x63, 0x74,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x41, 0x58, 0xaa, 0x02, 0x12,
	0x57, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x12, 0x57, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x5c, 0x41, 0x63, 0x74, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1e, 0x57, 0x61, 0x72, 0x64, 0x65, 0x6e,
	0x5c, 0x41, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x57, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x3a, 0x3a, 0x41, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_warden_act_v1beta1_rule_proto_rawDescOnce sync.Once
	file_warden_act_v1beta1_rule_proto_rawDescData = file_warden_act_v1beta1_rule_proto_rawDesc
)

func file_warden_act_v1beta1_rule_proto_rawDescGZIP() []byte {
	file_warden_act_v1beta1_rule_proto_rawDescOnce.Do(func() {
		file_warden_act_v1beta1_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_warden_act_v1beta1_rule_proto_rawDescData)
	})
	return file_warden_act_v1beta1_rule_proto_rawDescData
}

var file_warden_act_v1beta1_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_warden_act_v1beta1_rule_proto_goTypes = []interface{}{
	(*Rule)(nil),           // 0: warden.act.v1beta1.Rule
	(*ast.Expression)(nil), // 1: shield.ast.Expression
}
var file_warden_act_v1beta1_rule_proto_depIdxs = []int32{
	1, // 0: warden.act.v1beta1.Rule.expression:type_name -> shield.ast.Expression
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_warden_act_v1beta1_rule_proto_init() }
func file_warden_act_v1beta1_rule_proto_init() {
	if File_warden_act_v1beta1_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_warden_act_v1beta1_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
			RawDescriptor: file_warden_act_v1beta1_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_warden_act_v1beta1_rule_proto_goTypes,
		DependencyIndexes: file_warden_act_v1beta1_rule_proto_depIdxs,
		MessageInfos:      file_warden_act_v1beta1_rule_proto_msgTypes,
	}.Build()
	File_warden_act_v1beta1_rule_proto = out.File
	file_warden_act_v1beta1_rule_proto_rawDesc = nil
	file_warden_act_v1beta1_rule_proto_goTypes = nil
	file_warden_act_v1beta1_rule_proto_depIdxs = nil
}
