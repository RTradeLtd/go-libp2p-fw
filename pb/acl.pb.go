// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: acl.proto

package pb

import (
	fmt "fmt"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// FILTER_TYPE lists different possible filter methods
type FILTER_TYPE int32

const (
	// PEER_ID is used to filter traffic based off the peerid it is coming from
	FILTER_TYPE_PEER_ID FILTER_TYPE = 0
	// IP_ADDRESS4 is used to filter traffic based off ipv4 addresses
	FILTER_TYPE_IP_ADDRESS4 FILTER_TYPE = 1
	// IP_ADDRESS6 is used to filter traffic based off ipv6 addresses
	FILTER_TYPE_IP_ADDRESS6 FILTER_TYPE = 2
	// PROTOCOL_ID is used to filter traffic based off the protocol it is using (ie, /ping/1.0.0)
	FILTER_TYPE_PROTOCOL_ID FILTER_TYPE = 3
)

var FILTER_TYPE_name = map[int32]string{
	0: "PEER_ID",
	1: "IP_ADDRESS4",
	2: "IP_ADDRESS6",
	3: "PROTOCOL_ID",
}

var FILTER_TYPE_value = map[string]int32{
	"PEER_ID":     0,
	"IP_ADDRESS4": 1,
	"IP_ADDRESS6": 2,
	"PROTOCOL_ID": 3,
}

func (x FILTER_TYPE) String() string {
	return proto.EnumName(FILTER_TYPE_name, int32(x))
}

func (FILTER_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{0}
}

// DIRECTION lists different traffic direction
type DIRECTION int32

const (
	DIRECTION_INBOUND  DIRECTION = 0
	DIRECTION_OUTBOUND DIRECTION = 1
	DIRECTION_IN_OUT   DIRECTION = 2
	DIRECTION_RELAY    DIRECTION = 3
)

var DIRECTION_name = map[int32]string{
	0: "INBOUND",
	1: "OUTBOUND",
	2: "IN_OUT",
	3: "RELAY",
}

var DIRECTION_value = map[string]int32{
	"INBOUND":  0,
	"OUTBOUND": 1,
	"IN_OUT":   2,
	"RELAY":    3,
}

func (x DIRECTION) String() string {
	return proto.EnumName(DIRECTION_name, int32(x))
}

func (DIRECTION) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{1}
}

// ACTION lists the different possibles actions to be taken on matched ACLs
type ACTION int32

const (
	// ACCEPT is used to accept the connection (no shit)
	ACTION_ACCEPT ACTION = 0
	// REJECT is used to reject the connection (no shit)
	ACTION_REJECT ACTION = 1
)

var ACTION_name = map[int32]string{
	0: "ACCEPT",
	1: "REJECT",
}

var ACTION_value = map[string]int32{
	"ACCEPT": 0,
	"REJECT": 1,
}

func (x ACTION) String() string {
	return proto.EnumName(ACTION_name, int32(x))
}

func (ACTION) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{2}
}

// List is a list of ACL rules
type List struct {
	// acls is an array of access control list entries
	// the entries are ordered by priority, with higher priority
	// rules being evaluated first, and lower priority rules being evaluated last
	Acls map[int64]*ACL `protobuf:"bytes,1,rep,name=acls,proto3" json:"acls,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *List) Reset()      { *m = List{} }
func (*List) ProtoMessage() {}
func (*List) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{0}
}
func (m *List) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_List.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List.Merge(m, src)
}
func (m *List) XXX_Size() int {
	return m.Size()
}
func (m *List) XXX_DiscardUnknown() {
	xxx_messageInfo_List.DiscardUnknown(m)
}

var xxx_messageInfo_List proto.InternalMessageInfo

func (m *List) GetAcls() map[int64]*ACL {
	if m != nil {
		return m.Acls
	}
	return nil
}

// ACL is an individual access control list entry
type ACL struct {
	// filterType denotes the type of filter this ACL is using
	FilterType FILTER_TYPE `protobuf:"varint,1,opt,name=filterType,proto3,enum=pb.FILTER_TYPE" json:"filterType,omitempty"`
	// direction is the traffic direction this ACL applies to
	Direction DIRECTION `protobuf:"varint,2,opt,name=direction,proto3,enum=pb.DIRECTION" json:"direction,omitempty"`
	// action defines the action to be taken against the traffic matched by this acl
	Action ACTION `protobuf:"varint,3,opt,name=action,proto3,enum=pb.ACTION" json:"action,omitempty"`
}

func (m *ACL) Reset()      { *m = ACL{} }
func (*ACL) ProtoMessage() {}
func (*ACL) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{1}
}
func (m *ACL) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ACL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ACL.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ACL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACL.Merge(m, src)
}
func (m *ACL) XXX_Size() int {
	return m.Size()
}
func (m *ACL) XXX_DiscardUnknown() {
	xxx_messageInfo_ACL.DiscardUnknown(m)
}

var xxx_messageInfo_ACL proto.InternalMessageInfo

func (m *ACL) GetFilterType() FILTER_TYPE {
	if m != nil {
		return m.FilterType
	}
	return FILTER_TYPE_PEER_ID
}

func (m *ACL) GetDirection() DIRECTION {
	if m != nil {
		return m.Direction
	}
	return DIRECTION_INBOUND
}

func (m *ACL) GetAction() ACTION {
	if m != nil {
		return m.Action
	}
	return ACTION_ACCEPT
}

func init() {
	proto.RegisterEnum("pb.FILTER_TYPE", FILTER_TYPE_name, FILTER_TYPE_value)
	proto.RegisterEnum("pb.DIRECTION", DIRECTION_name, DIRECTION_value)
	proto.RegisterEnum("pb.ACTION", ACTION_name, ACTION_value)
	proto.RegisterType((*List)(nil), "pb.List")
	proto.RegisterMapType((map[int64]*ACL)(nil), "pb.List.AclsEntry")
	proto.RegisterType((*ACL)(nil), "pb.ACL")
}

func init() { proto.RegisterFile("acl.proto", fileDescriptor_a452f070aeef01eb) }

var fileDescriptor_a452f070aeef01eb = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xb1, 0x6f, 0xd3, 0x40,
	0x14, 0xc6, 0xfd, 0xec, 0x36, 0xc5, 0xcf, 0xd0, 0x5a, 0x37, 0x45, 0x95, 0x38, 0x45, 0x19, 0x50,
	0x14, 0x84, 0x23, 0x05, 0x84, 0x2a, 0x24, 0x24, 0x5c, 0xe7, 0x90, 0x8c, 0x4c, 0x6c, 0x5d, 0x2f,
	0x43, 0xa7, 0x28, 0x36, 0x6e, 0xb0, 0x30, 0xb5, 0x95, 0x38, 0x48, 0xd9, 0xd8, 0x58, 0xf9, 0x33,
	0xf8, 0x13, 0x18, 0x19, 0x19, 0x3b, 0x76, 0xac, 0x9d, 0x7f, 0x80, 0x91, 0x11, 0xdd, 0x19, 0x95,
	0xb2, 0xbd, 0xf7, 0x7d, 0xbf, 0xf7, 0xdd, 0x27, 0x1d, 0x9a, 0x8b, 0x24, 0x77, 0xca, 0x55, 0x51,
	0x15, 0x44, 0x2f, 0xe3, 0xe3, 0x27, 0xcb, 0xac, 0x7a, 0xbf, 0x89, 0x9d, 0xa4, 0xf8, 0x38, 0x5a,
	0x16, 0xcb, 0x62, 0xa4, 0xac, 0x78, 0x73, 0xa1, 0x36, 0xb5, 0xa8, 0xa9, 0x3d, 0xe9, 0x97, 0xb8,
	0x17, 0x64, 0xeb, 0x8a, 0x3c, 0xc2, 0xbd, 0x45, 0x92, 0xaf, 0xbb, 0xd0, 0x33, 0x06, 0xd6, 0x98,
	0x38, 0x65, 0xec, 0x48, 0xdd, 0x71, 0x93, 0x7c, 0xcd, 0x2e, 0xab, 0xd5, 0x96, 0x2b, 0xff, 0xf8,
	0x15, 0x9a, 0xb7, 0x12, 0xb1, 0xd1, 0xf8, 0x90, 0x6e, 0xbb, 0xd0, 0x83, 0x81, 0xc1, 0xe5, 0x48,
	0x1e, 0xe2, 0xfe, 0xa7, 0x45, 0xbe, 0x49, 0xbb, 0x7a, 0x0f, 0x06, 0xd6, 0xf8, 0x40, 0xe6, 0xb8,
	0x5e, 0xc0, 0x5b, 0xf5, 0x85, 0x7e, 0x02, 0xfd, 0x2f, 0x80, 0x86, 0xeb, 0x05, 0x64, 0x84, 0x78,
	0x91, 0xe5, 0x55, 0xba, 0x12, 0xdb, 0x32, 0x55, 0x19, 0x87, 0xe3, 0x23, 0xc9, 0xbf, 0xf6, 0x03,
	0xc1, 0xf8, 0x5c, 0x9c, 0x47, 0x8c, 0xdf, 0x41, 0xc8, 0x63, 0x34, 0xdf, 0x65, 0xab, 0x34, 0xa9,
	0xb2, 0xe2, 0x52, 0xe5, 0x1f, 0x8e, 0x1f, 0x48, 0x7e, 0xe2, 0x73, 0xe6, 0x09, 0x3f, 0x9c, 0xf2,
	0x7f, 0x3e, 0xe9, 0x63, 0x67, 0xd1, 0x92, 0x86, 0x22, 0xb1, 0x6d, 0xa2, 0xb0, 0xbf, 0xce, 0xf0,
	0x2d, 0x5a, 0x77, 0xde, 0x22, 0x16, 0x1e, 0x44, 0x8c, 0xf1, 0xb9, 0x3f, 0xb1, 0x35, 0x72, 0x84,
	0x96, 0x1f, 0xcd, 0xdd, 0xc9, 0x84, 0xb3, 0xb3, 0xb3, 0x67, 0x36, 0xfc, 0x2f, 0x3c, 0xb7, 0x75,
	0x29, 0x44, 0x3c, 0x14, 0xa1, 0x17, 0x06, 0xf2, 0xc4, 0x18, 0xbe, 0x44, 0xf3, 0xb6, 0x8a, 0x0c,
	0xf3, 0xa7, 0xa7, 0xe1, 0x6c, 0x2a, 0xc3, 0xee, 0xe3, 0xbd, 0x70, 0x26, 0xda, 0x0d, 0x08, 0x62,
	0xc7, 0x9f, 0xce, 0xc3, 0x99, 0xb0, 0x75, 0x62, 0xe2, 0x3e, 0x67, 0x81, 0x7b, 0x6e, 0x1b, 0xc3,
	0x1e, 0x76, 0xda, 0x7e, 0x12, 0x70, 0x3d, 0x8f, 0x45, 0xc2, 0xd6, 0xe4, 0xcc, 0xd9, 0x1b, 0xe6,
	0x09, 0x1b, 0x4e, 0x4f, 0xae, 0x6b, 0xaa, 0xdd, 0xd4, 0x14, 0x7e, 0xd5, 0x14, 0x7e, 0xd7, 0x14,
	0x3e, 0x37, 0x14, 0xbe, 0x35, 0x14, 0xbe, 0x37, 0x14, 0x7e, 0x34, 0x14, 0x7e, 0x36, 0x14, 0xae,
	0x1a, 0x0a, 0x37, 0x0d, 0x85, 0xaf, 0x3b, 0xaa, 0x5d, 0xed, 0xa8, 0x76, 0xbd, 0xa3, 0x5a, 0xdc,
	0x51, 0x9f, 0xfd, 0xf4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x72, 0xdb, 0xf3, 0x2c, 0x02,
	0x00, 0x00,
}

func (this *List) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*List)
	if !ok {
		that2, ok := that.(List)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *List")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *List but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *List but is not nil && this == nil")
	}
	if len(this.Acls) != len(that1.Acls) {
		return fmt.Errorf("Acls this(%v) Not Equal that(%v)", len(this.Acls), len(that1.Acls))
	}
	for i := range this.Acls {
		if !this.Acls[i].Equal(that1.Acls[i]) {
			return fmt.Errorf("Acls this[%v](%v) Not Equal that[%v](%v)", i, this.Acls[i], i, that1.Acls[i])
		}
	}
	return nil
}
func (this *List) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*List)
	if !ok {
		that2, ok := that.(List)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Acls) != len(that1.Acls) {
		return false
	}
	for i := range this.Acls {
		if !this.Acls[i].Equal(that1.Acls[i]) {
			return false
		}
	}
	return true
}
func (this *ACL) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ACL)
	if !ok {
		that2, ok := that.(ACL)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ACL")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ACL but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ACL but is not nil && this == nil")
	}
	if this.FilterType != that1.FilterType {
		return fmt.Errorf("FilterType this(%v) Not Equal that(%v)", this.FilterType, that1.FilterType)
	}
	if this.Direction != that1.Direction {
		return fmt.Errorf("Direction this(%v) Not Equal that(%v)", this.Direction, that1.Direction)
	}
	if this.Action != that1.Action {
		return fmt.Errorf("Action this(%v) Not Equal that(%v)", this.Action, that1.Action)
	}
	return nil
}
func (this *ACL) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ACL)
	if !ok {
		that2, ok := that.(ACL)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FilterType != that1.FilterType {
		return false
	}
	if this.Direction != that1.Direction {
		return false
	}
	if this.Action != that1.Action {
		return false
	}
	return true
}
func (this *List) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.List{")
	keysForAcls := make([]int64, 0, len(this.Acls))
	for k := range this.Acls {
		keysForAcls = append(keysForAcls, k)
	}
	github_com_gogo_protobuf_sortkeys.Int64s(keysForAcls)
	mapStringForAcls := "map[int64]*ACL{"
	for _, k := range keysForAcls {
		mapStringForAcls += fmt.Sprintf("%#v: %#v,", k, this.Acls[k])
	}
	mapStringForAcls += "}"
	if this.Acls != nil {
		s = append(s, "Acls: "+mapStringForAcls+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ACL) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pb.ACL{")
	s = append(s, "FilterType: "+fmt.Sprintf("%#v", this.FilterType)+",\n")
	s = append(s, "Direction: "+fmt.Sprintf("%#v", this.Direction)+",\n")
	s = append(s, "Action: "+fmt.Sprintf("%#v", this.Action)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAcl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *List) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *List) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Acls) > 0 {
		for k := range m.Acls {
			dAtA[i] = 0xa
			i++
			v := m.Acls[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovAcl(uint64(msgSize))
			}
			mapSize := 1 + sovAcl(uint64(k)) + msgSize
			i = encodeVarintAcl(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintAcl(dAtA, i, uint64(k))
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintAcl(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	return i, nil
}

func (m *ACL) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ACL) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FilterType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAcl(dAtA, i, uint64(m.FilterType))
	}
	if m.Direction != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAcl(dAtA, i, uint64(m.Direction))
	}
	if m.Action != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAcl(dAtA, i, uint64(m.Action))
	}
	return i, nil
}

func encodeVarintAcl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedList(r randyAcl, easy bool) *List {
	this := &List{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(10)
		this.Acls = make(map[int64]*ACL)
		for i := 0; i < v1; i++ {
			this.Acls[int64(r.Int63())] = NewPopulatedACL(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedACL(r randyAcl, easy bool) *ACL {
	this := &ACL{}
	this.FilterType = FILTER_TYPE([]int32{0, 1, 2, 3}[r.Intn(4)])
	this.Direction = DIRECTION([]int32{0, 1, 2, 3}[r.Intn(4)])
	this.Action = ACTION([]int32{0, 1}[r.Intn(2)])
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyAcl interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAcl(r randyAcl) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAcl(r randyAcl) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneAcl(r)
	}
	return string(tmps)
}
func randUnrecognizedAcl(r randyAcl, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAcl(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAcl(dAtA []byte, r randyAcl, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAcl(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateAcl(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateAcl(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAcl(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAcl(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAcl(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAcl(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *List) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Acls) > 0 {
		for k, v := range m.Acls {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovAcl(uint64(l))
			}
			mapEntrySize := 1 + sovAcl(uint64(k)) + l
			n += mapEntrySize + 1 + sovAcl(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ACL) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FilterType != 0 {
		n += 1 + sovAcl(uint64(m.FilterType))
	}
	if m.Direction != 0 {
		n += 1 + sovAcl(uint64(m.Direction))
	}
	if m.Action != 0 {
		n += 1 + sovAcl(uint64(m.Action))
	}
	return n
}

func sovAcl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAcl(x uint64) (n int) {
	return sovAcl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *List) String() string {
	if this == nil {
		return "nil"
	}
	keysForAcls := make([]int64, 0, len(this.Acls))
	for k := range this.Acls {
		keysForAcls = append(keysForAcls, k)
	}
	github_com_gogo_protobuf_sortkeys.Int64s(keysForAcls)
	mapStringForAcls := "map[int64]*ACL{"
	for _, k := range keysForAcls {
		mapStringForAcls += fmt.Sprintf("%v: %v,", k, this.Acls[k])
	}
	mapStringForAcls += "}"
	s := strings.Join([]string{`&List{`,
		`Acls:` + mapStringForAcls + `,`,
		`}`,
	}, "")
	return s
}
func (this *ACL) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ACL{`,
		`FilterType:` + fmt.Sprintf("%v", this.FilterType) + `,`,
		`Direction:` + fmt.Sprintf("%v", this.Direction) + `,`,
		`Action:` + fmt.Sprintf("%v", this.Action) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAcl(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *List) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAcl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
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
			return fmt.Errorf("proto: List: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: List: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Acls", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAcl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAcl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Acls == nil {
				m.Acls = make(map[int64]*ACL)
			}
			var mapkey int64
			var mapvalue *ACL
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAcl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAcl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAcl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthAcl
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthAcl
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ACL{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipAcl(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthAcl
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Acls[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAcl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAcl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAcl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ACL) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAcl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
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
			return fmt.Errorf("proto: ACL: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ACL: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterType", wireType)
			}
			m.FilterType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FilterType |= FILTER_TYPE(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= DIRECTION(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= ACTION(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAcl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAcl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAcl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAcl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAcl
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAcl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAcl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAcl
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAcl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAcl
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAcl(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAcl
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAcl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAcl   = fmt.Errorf("proto: integer overflow")
)
