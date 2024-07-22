// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package wardenv1beta3

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_Space_3_list)(nil)

type _Space_3_list struct {
	list *[]string
}

func (x *_Space_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Space_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_Space_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_Space_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_Space_3_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message Space at list field Owners as it is not of Message kind"))
}

func (x *_Space_3_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_Space_3_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_Space_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Space               protoreflect.MessageDescriptor
	fd_Space_id            protoreflect.FieldDescriptor
	fd_Space_creator       protoreflect.FieldDescriptor
	fd_Space_owners        protoreflect.FieldDescriptor
	fd_Space_admin_rule_id protoreflect.FieldDescriptor
	fd_Space_sign_rule_id  protoreflect.FieldDescriptor
)

func init() {
	file_warden_warden_v1beta3_space_proto_init()
	md_Space = File_warden_warden_v1beta3_space_proto.Messages().ByName("Space")
	fd_Space_id = md_Space.Fields().ByName("id")
	fd_Space_creator = md_Space.Fields().ByName("creator")
	fd_Space_owners = md_Space.Fields().ByName("owners")
	fd_Space_admin_rule_id = md_Space.Fields().ByName("admin_rule_id")
	fd_Space_sign_rule_id = md_Space.Fields().ByName("sign_rule_id")
}

var _ protoreflect.Message = (*fastReflection_Space)(nil)

type fastReflection_Space Space

func (x *Space) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Space)(x)
}

func (x *Space) slowProtoReflect() protoreflect.Message {
	mi := &file_warden_warden_v1beta3_space_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Space_messageType fastReflection_Space_messageType
var _ protoreflect.MessageType = fastReflection_Space_messageType{}

type fastReflection_Space_messageType struct{}

func (x fastReflection_Space_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Space)(nil)
}
func (x fastReflection_Space_messageType) New() protoreflect.Message {
	return new(fastReflection_Space)
}
func (x fastReflection_Space_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Space
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Space) Descriptor() protoreflect.MessageDescriptor {
	return md_Space
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Space) Type() protoreflect.MessageType {
	return _fastReflection_Space_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Space) New() protoreflect.Message {
	return new(fastReflection_Space)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Space) Interface() protoreflect.ProtoMessage {
	return (*Space)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Space) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Space_id, value) {
			return
		}
	}
	if x.Creator != "" {
		value := protoreflect.ValueOfString(x.Creator)
		if !f(fd_Space_creator, value) {
			return
		}
	}
	if len(x.Owners) != 0 {
		value := protoreflect.ValueOfList(&_Space_3_list{list: &x.Owners})
		if !f(fd_Space_owners, value) {
			return
		}
	}
	if x.AdminRuleId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AdminRuleId)
		if !f(fd_Space_admin_rule_id, value) {
			return
		}
	}
	if x.SignRuleId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SignRuleId)
		if !f(fd_Space_sign_rule_id, value) {
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
func (x *fastReflection_Space) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "warden.warden.v1beta3.Space.id":
		return x.Id != uint64(0)
	case "warden.warden.v1beta3.Space.creator":
		return x.Creator != ""
	case "warden.warden.v1beta3.Space.owners":
		return len(x.Owners) != 0
	case "warden.warden.v1beta3.Space.admin_rule_id":
		return x.AdminRuleId != uint64(0)
	case "warden.warden.v1beta3.Space.sign_rule_id":
		return x.SignRuleId != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.warden.v1beta3.Space"))
		}
		panic(fmt.Errorf("message warden.warden.v1beta3.Space does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Space) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "warden.warden.v1beta3.Space.id":
		x.Id = uint64(0)
	case "warden.warden.v1beta3.Space.creator":
		x.Creator = ""
	case "warden.warden.v1beta3.Space.owners":
		x.Owners = nil
	case "warden.warden.v1beta3.Space.admin_rule_id":
		x.AdminRuleId = uint64(0)
	case "warden.warden.v1beta3.Space.sign_rule_id":
		x.SignRuleId = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.warden.v1beta3.Space"))
		}
		panic(fmt.Errorf("message warden.warden.v1beta3.Space does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Space) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "warden.warden.v1beta3.Space.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "warden.warden.v1beta3.Space.creator":
		value := x.Creator
		return protoreflect.ValueOfString(value)
	case "warden.warden.v1beta3.Space.owners":
		if len(x.Owners) == 0 {
			return protoreflect.ValueOfList(&_Space_3_list{})
		}
		listValue := &_Space_3_list{list: &x.Owners}
		return protoreflect.ValueOfList(listValue)
	case "warden.warden.v1beta3.Space.admin_rule_id":
		value := x.AdminRuleId
		return protoreflect.ValueOfUint64(value)
	case "warden.warden.v1beta3.Space.sign_rule_id":
		value := x.SignRuleId
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.warden.v1beta3.Space"))
		}
		panic(fmt.Errorf("message warden.warden.v1beta3.Space does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Space) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "warden.warden.v1beta3.Space.id":
		x.Id = value.Uint()
	case "warden.warden.v1beta3.Space.creator":
		x.Creator = value.Interface().(string)
	case "warden.warden.v1beta3.Space.owners":
		lv := value.List()
		clv := lv.(*_Space_3_list)
		x.Owners = *clv.list
	case "warden.warden.v1beta3.Space.admin_rule_id":
		x.AdminRuleId = value.Uint()
	case "warden.warden.v1beta3.Space.sign_rule_id":
		x.SignRuleId = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.warden.v1beta3.Space"))
		}
		panic(fmt.Errorf("message warden.warden.v1beta3.Space does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Space) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "warden.warden.v1beta3.Space.owners":
		if x.Owners == nil {
			x.Owners = []string{}
		}
		value := &_Space_3_list{list: &x.Owners}
		return protoreflect.ValueOfList(value)
	case "warden.warden.v1beta3.Space.id":
		panic(fmt.Errorf("field id of message warden.warden.v1beta3.Space is not mutable"))
	case "warden.warden.v1beta3.Space.creator":
		panic(fmt.Errorf("field creator of message warden.warden.v1beta3.Space is not mutable"))
	case "warden.warden.v1beta3.Space.admin_rule_id":
		panic(fmt.Errorf("field admin_rule_id of message warden.warden.v1beta3.Space is not mutable"))
	case "warden.warden.v1beta3.Space.sign_rule_id":
		panic(fmt.Errorf("field sign_rule_id of message warden.warden.v1beta3.Space is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.warden.v1beta3.Space"))
		}
		panic(fmt.Errorf("message warden.warden.v1beta3.Space does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Space) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "warden.warden.v1beta3.Space.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "warden.warden.v1beta3.Space.creator":
		return protoreflect.ValueOfString("")
	case "warden.warden.v1beta3.Space.owners":
		list := []string{}
		return protoreflect.ValueOfList(&_Space_3_list{list: &list})
	case "warden.warden.v1beta3.Space.admin_rule_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "warden.warden.v1beta3.Space.sign_rule_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: warden.warden.v1beta3.Space"))
		}
		panic(fmt.Errorf("message warden.warden.v1beta3.Space does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Space) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in warden.warden.v1beta3.Space", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Space) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Space) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Space) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Space) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Space)
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
		if len(x.Owners) > 0 {
			for _, s := range x.Owners {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.AdminRuleId != 0 {
			n += 1 + runtime.Sov(uint64(x.AdminRuleId))
		}
		if x.SignRuleId != 0 {
			n += 1 + runtime.Sov(uint64(x.SignRuleId))
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
		x := input.Message.Interface().(*Space)
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
		if x.SignRuleId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SignRuleId))
			i--
			dAtA[i] = 0x30
		}
		if x.AdminRuleId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AdminRuleId))
			i--
			dAtA[i] = 0x28
		}
		if len(x.Owners) > 0 {
			for iNdEx := len(x.Owners) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Owners[iNdEx])
				copy(dAtA[i:], x.Owners[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owners[iNdEx])))
				i--
				dAtA[i] = 0x1a
			}
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
		x := input.Message.Interface().(*Space)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Space: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Space: illegal tag %d (wire type %d)", fieldNum, wire)
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owners", wireType)
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
				x.Owners = append(x.Owners, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AdminRuleId", wireType)
				}
				x.AdminRuleId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AdminRuleId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SignRuleId", wireType)
				}
				x.SignRuleId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SignRuleId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
// source: warden/warden/v1beta3/space.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Space is a collection of users (called owners) that manages a set of Keys.
type Space struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID of the space.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Address of the creator of the space.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// List of owners of the space.
	Owners []string `protobuf:"bytes,3,rep,name=owners,proto3" json:"owners,omitempty"`
	// Optional ID of the Rule to be applied to every *admin* operation.
	// If not specified, the default Rule is used.
	//
	// Admin operations are:
	// - warden.warden.Msg.AddSpaceOwner
	// - warden.warden.Msg.RemoveSpaceOwner
	// - warden.warden.Msg.UpdateSpace
	//
	// The default Rule is to allow any operation when at least one of its
	// owner approves it.
	AdminRuleId uint64 `protobuf:"varint,5,opt,name=admin_rule_id,json=adminRuleId,proto3" json:"admin_rule_id,omitempty"`
	// Optional ID of the Rule to be applied to every *sign* operation.
	// If not specified, the default Rule is used.
	//
	// Sign operations are:
	// - warden.warden.Msg.NewKeyRequest
	// - warden.warden.Msg.NewSignRequest
	// - warden.warden.Msg.UpdateKey
	//
	// The default Rule is to allow any operation when at least one of its
	// owner approves it.
	SignRuleId uint64 `protobuf:"varint,6,opt,name=sign_rule_id,json=signRuleId,proto3" json:"sign_rule_id,omitempty"`
}

func (x *Space) Reset() {
	*x = Space{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warden_warden_v1beta3_space_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Space) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Space) ProtoMessage() {}

// Deprecated: Use Space.ProtoReflect.Descriptor instead.
func (*Space) Descriptor() ([]byte, []int) {
	return file_warden_warden_v1beta3_space_proto_rawDescGZIP(), []int{0}
}

func (x *Space) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Space) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Space) GetOwners() []string {
	if x != nil {
		return x.Owners
	}
	return nil
}

func (x *Space) GetAdminRuleId() uint64 {
	if x != nil {
		return x.AdminRuleId
	}
	return 0
}

func (x *Space) GetSignRuleId() uint64 {
	if x != nil {
		return x.SignRuleId
	}
	return 0
}

var File_warden_warden_v1beta3_space_proto protoreflect.FileDescriptor

var file_warden_warden_v1beta3_space_proto_rawDesc = []byte{
	0x0a, 0x21, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2f, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x22, 0x8f, 0x01, 0x0a, 0x05, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x69,
	0x67, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x42, 0xf0, 0x01, 0x0a,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x42, 0x0a, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2f, 0x77,
	0x61, 0x72, 0x64, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x3b, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0xa2, 0x02, 0x03, 0x57, 0x57,
	0x58, 0xaa, 0x02, 0x15, 0x57, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x57, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0xca, 0x02, 0x15, 0x57, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x5c, 0x57, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x33, 0xe2, 0x02, 0x21, 0x57, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x5c, 0x57, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x57, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x3a, 0x3a,
	0x57, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_warden_warden_v1beta3_space_proto_rawDescOnce sync.Once
	file_warden_warden_v1beta3_space_proto_rawDescData = file_warden_warden_v1beta3_space_proto_rawDesc
)

func file_warden_warden_v1beta3_space_proto_rawDescGZIP() []byte {
	file_warden_warden_v1beta3_space_proto_rawDescOnce.Do(func() {
		file_warden_warden_v1beta3_space_proto_rawDescData = protoimpl.X.CompressGZIP(file_warden_warden_v1beta3_space_proto_rawDescData)
	})
	return file_warden_warden_v1beta3_space_proto_rawDescData
}

var file_warden_warden_v1beta3_space_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_warden_warden_v1beta3_space_proto_goTypes = []interface{}{
	(*Space)(nil), // 0: warden.warden.v1beta3.Space
}
var file_warden_warden_v1beta3_space_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_warden_warden_v1beta3_space_proto_init() }
func file_warden_warden_v1beta3_space_proto_init() {
	if File_warden_warden_v1beta3_space_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_warden_warden_v1beta3_space_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Space); i {
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
			RawDescriptor: file_warden_warden_v1beta3_space_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_warden_warden_v1beta3_space_proto_goTypes,
		DependencyIndexes: file_warden_warden_v1beta3_space_proto_depIdxs,
		MessageInfos:      file_warden_warden_v1beta3_space_proto_msgTypes,
	}.Build()
	File_warden_warden_v1beta3_space_proto = out.File
	file_warden_warden_v1beta3_space_proto_rawDesc = nil
	file_warden_warden_v1beta3_space_proto_goTypes = nil
	file_warden_warden_v1beta3_space_proto_depIdxs = nil
}