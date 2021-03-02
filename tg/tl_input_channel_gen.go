// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// InputChannelEmpty represents TL type `inputChannelEmpty#ee8c1e86`.
// Represents the absence of a channel
//
// See https://core.telegram.org/constructor/inputChannelEmpty for reference.
type InputChannelEmpty struct {
}

// InputChannelEmptyTypeID is TL type id of InputChannelEmpty.
const InputChannelEmptyTypeID = 0xee8c1e86

func (i *InputChannelEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChannelEmpty) String() string {
	if i == nil {
		return "InputChannelEmpty(nil)"
	}
	type Alias InputChannelEmpty
	return fmt.Sprintf("InputChannelEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChannelEmpty) TypeID() uint32 {
	return InputChannelEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChannelEmpty) TypeName() string {
	return "inputChannelEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputChannelEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChannelEmpty",
		ID:   InputChannelEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChannelEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannelEmpty#ee8c1e86 as nil")
	}
	b.PutID(InputChannelEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChannelEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannelEmpty#ee8c1e86 to nil")
	}
	if err := b.ConsumeID(InputChannelEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannelEmpty#ee8c1e86: %w", err)
	}
	return nil
}

// construct implements constructor of InputChannelClass.
func (i InputChannelEmpty) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannelEmpty.
var (
	_ bin.Encoder = &InputChannelEmpty{}
	_ bin.Decoder = &InputChannelEmpty{}

	_ InputChannelClass = &InputChannelEmpty{}
)

// InputChannel represents TL type `inputChannel#afeb712e`.
// Represents a channel
//
// See https://core.telegram.org/constructor/inputChannel for reference.
type InputChannel struct {
	// Channel ID
	ChannelID int `tl:"channel_id"`
	// Access hash taken from the channel¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channel
	AccessHash int64 `tl:"access_hash"`
}

// InputChannelTypeID is TL type id of InputChannel.
const InputChannelTypeID = 0xafeb712e

func (i *InputChannel) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChannelID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChannel) String() string {
	if i == nil {
		return "InputChannel(nil)"
	}
	type Alias InputChannel
	return fmt.Sprintf("InputChannel%+v", Alias(*i))
}

// FillFrom fills InputChannel from given interface.
func (i *InputChannel) FillFrom(from interface {
	GetChannelID() (value int)
	GetAccessHash() (value int64)
}) {
	i.ChannelID = from.GetChannelID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChannel) TypeID() uint32 {
	return InputChannelTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChannel) TypeName() string {
	return "inputChannel"
}

// TypeInfo returns info about TL type.
func (i *InputChannel) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChannel",
		ID:   InputChannelTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChannelID",
			SchemaName: "channel_id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChannel) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannel#afeb712e as nil")
	}
	b.PutID(InputChannelTypeID)
	b.PutInt(i.ChannelID)
	b.PutLong(i.AccessHash)
	return nil
}

// GetChannelID returns value of ChannelID field.
func (i *InputChannel) GetChannelID() (value int) {
	return i.ChannelID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputChannel) GetAccessHash() (value int64) {
	return i.AccessHash
}

// Decode implements bin.Decoder.
func (i *InputChannel) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannel#afeb712e to nil")
	}
	if err := b.ConsumeID(InputChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannel#afeb712e: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannel#afeb712e: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannel#afeb712e: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputChannelClass.
func (i InputChannel) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannel.
var (
	_ bin.Encoder = &InputChannel{}
	_ bin.Decoder = &InputChannel{}

	_ InputChannelClass = &InputChannel{}
)

// InputChannelFromMessage represents TL type `inputChannelFromMessage#2a286531`.
// Defines a min¹ channel that was seen in a certain message of a certain chat.
//
// Links:
//  1) https://core.telegram.org/api/min
//
// See https://core.telegram.org/constructor/inputChannelFromMessage for reference.
type InputChannelFromMessage struct {
	// The chat where the channel was seen
	Peer InputPeerClass `tl:"peer"`
	// The message ID in the chat where the channel was seen
	MsgID int `tl:"msg_id"`
	// The channel ID
	ChannelID int `tl:"channel_id"`
}

// InputChannelFromMessageTypeID is TL type id of InputChannelFromMessage.
const InputChannelFromMessageTypeID = 0x2a286531

func (i *InputChannelFromMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}
	if !(i.MsgID == 0) {
		return false
	}
	if !(i.ChannelID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChannelFromMessage) String() string {
	if i == nil {
		return "InputChannelFromMessage(nil)"
	}
	type Alias InputChannelFromMessage
	return fmt.Sprintf("InputChannelFromMessage%+v", Alias(*i))
}

// FillFrom fills InputChannelFromMessage from given interface.
func (i *InputChannelFromMessage) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetChannelID() (value int)
}) {
	i.Peer = from.GetPeer()
	i.MsgID = from.GetMsgID()
	i.ChannelID = from.GetChannelID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChannelFromMessage) TypeID() uint32 {
	return InputChannelFromMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChannelFromMessage) TypeName() string {
	return "inputChannelFromMessage"
}

// TypeInfo returns info about TL type.
func (i *InputChannelFromMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChannelFromMessage",
		ID:   InputChannelFromMessageTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "ChannelID",
			SchemaName: "channel_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChannelFromMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannelFromMessage#2a286531 as nil")
	}
	b.PutID(InputChannelFromMessageTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputChannelFromMessage#2a286531: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChannelFromMessage#2a286531: field peer: %w", err)
	}
	b.PutInt(i.MsgID)
	b.PutInt(i.ChannelID)
	return nil
}

// GetPeer returns value of Peer field.
func (i *InputChannelFromMessage) GetPeer() (value InputPeerClass) {
	return i.Peer
}

// GetMsgID returns value of MsgID field.
func (i *InputChannelFromMessage) GetMsgID() (value int) {
	return i.MsgID
}

// GetChannelID returns value of ChannelID field.
func (i *InputChannelFromMessage) GetChannelID() (value int) {
	return i.ChannelID
}

// Decode implements bin.Decoder.
func (i *InputChannelFromMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannelFromMessage#2a286531 to nil")
	}
	if err := b.ConsumeID(InputChannelFromMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	return nil
}

// construct implements constructor of InputChannelClass.
func (i InputChannelFromMessage) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannelFromMessage.
var (
	_ bin.Encoder = &InputChannelFromMessage{}
	_ bin.Decoder = &InputChannelFromMessage{}

	_ InputChannelClass = &InputChannelFromMessage{}
)

// InputChannelClass represents InputChannel generic type.
//
// See https://core.telegram.org/type/InputChannel for reference.
//
// Example:
//  g, err := tg.DecodeInputChannel(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputChannelEmpty: // inputChannelEmpty#ee8c1e86
//  case *tg.InputChannel: // inputChannel#afeb712e
//  case *tg.InputChannelFromMessage: // inputChannelFromMessage#2a286531
//  default: panic(v)
//  }
type InputChannelClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputChannelClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsNotEmpty tries to map InputChannelClass to NotEmptyInputChannel.
	AsNotEmpty() (NotEmptyInputChannel, bool)
}

// NotEmptyInputChannel represents NotEmpty subset of InputChannelClass.
type NotEmptyInputChannel interface {
	bin.Encoder
	bin.Decoder
	construct() InputChannelClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Channel ID
	GetChannelID() (value int)
}

// AsNotEmpty tries to map InputChannelEmpty to NotEmptyInputChannel.
func (i *InputChannelEmpty) AsNotEmpty() (NotEmptyInputChannel, bool) {
	value, ok := (InputChannelClass(i)).(NotEmptyInputChannel)
	return value, ok
}

// AsNotEmpty tries to map InputChannel to NotEmptyInputChannel.
func (i *InputChannel) AsNotEmpty() (NotEmptyInputChannel, bool) {
	value, ok := (InputChannelClass(i)).(NotEmptyInputChannel)
	return value, ok
}

// AsNotEmpty tries to map InputChannelFromMessage to NotEmptyInputChannel.
func (i *InputChannelFromMessage) AsNotEmpty() (NotEmptyInputChannel, bool) {
	value, ok := (InputChannelClass(i)).(NotEmptyInputChannel)
	return value, ok
}

// DecodeInputChannel implements binary de-serialization for InputChannelClass.
func DecodeInputChannel(buf *bin.Buffer) (InputChannelClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputChannelEmptyTypeID:
		// Decoding inputChannelEmpty#ee8c1e86.
		v := InputChannelEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	case InputChannelTypeID:
		// Decoding inputChannel#afeb712e.
		v := InputChannel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	case InputChannelFromMessageTypeID:
		// Decoding inputChannelFromMessage#2a286531.
		v := InputChannelFromMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputChannelClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputChannel boxes the InputChannelClass providing a helper.
type InputChannelBox struct {
	InputChannel InputChannelClass
}

// Decode implements bin.Decoder for InputChannelBox.
func (b *InputChannelBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputChannelBox to nil")
	}
	v, err := DecodeInputChannel(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputChannel = v
	return nil
}

// Encode implements bin.Encode for InputChannelBox.
func (b *InputChannelBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputChannel == nil {
		return fmt.Errorf("unable to encode InputChannelClass as nil")
	}
	return b.InputChannel.Encode(buf)
}

// InputChannelClassArray is adapter for slice of InputChannelClass.
type InputChannelClassArray []InputChannelClass

// Sort sorts slice of InputChannelClass.
func (s InputChannelClassArray) Sort(less func(a, b InputChannelClass) bool) InputChannelClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputChannelClass.
func (s InputChannelClassArray) SortStable(less func(a, b InputChannelClass) bool) InputChannelClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputChannelClass.
func (s InputChannelClassArray) Retain(keep func(x InputChannelClass) bool) InputChannelClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputChannelClassArray) First() (v InputChannelClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputChannelClassArray) Last() (v InputChannelClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputChannelClassArray) PopFirst() (v InputChannelClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputChannelClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputChannelClassArray) Pop() (v InputChannelClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputChannel returns copy with only InputChannel constructors.
func (s InputChannelClassArray) AsInputChannel() (to InputChannelArray) {
	for _, elem := range s {
		value, ok := elem.(*InputChannel)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputChannelFromMessage returns copy with only InputChannelFromMessage constructors.
func (s InputChannelClassArray) AsInputChannelFromMessage() (to InputChannelFromMessageArray) {
	for _, elem := range s {
		value, ok := elem.(*InputChannelFromMessage)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s InputChannelClassArray) AppendOnlyNotEmpty(to []NotEmptyInputChannel) []NotEmptyInputChannel {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s InputChannelClassArray) AsNotEmpty() (to []NotEmptyInputChannel) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s InputChannelClassArray) FirstAsNotEmpty() (v NotEmptyInputChannel, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s InputChannelClassArray) LastAsNotEmpty() (v NotEmptyInputChannel, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *InputChannelClassArray) PopFirstAsNotEmpty() (v NotEmptyInputChannel, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *InputChannelClassArray) PopAsNotEmpty() (v NotEmptyInputChannel, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// InputChannelArray is adapter for slice of InputChannel.
type InputChannelArray []InputChannel

// Sort sorts slice of InputChannel.
func (s InputChannelArray) Sort(less func(a, b InputChannel) bool) InputChannelArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputChannel.
func (s InputChannelArray) SortStable(less func(a, b InputChannel) bool) InputChannelArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputChannel.
func (s InputChannelArray) Retain(keep func(x InputChannel) bool) InputChannelArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputChannelArray) First() (v InputChannel, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputChannelArray) Last() (v InputChannel, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputChannelArray) PopFirst() (v InputChannel, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputChannel
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputChannelArray) Pop() (v InputChannel, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputChannelFromMessageArray is adapter for slice of InputChannelFromMessage.
type InputChannelFromMessageArray []InputChannelFromMessage

// Sort sorts slice of InputChannelFromMessage.
func (s InputChannelFromMessageArray) Sort(less func(a, b InputChannelFromMessage) bool) InputChannelFromMessageArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputChannelFromMessage.
func (s InputChannelFromMessageArray) SortStable(less func(a, b InputChannelFromMessage) bool) InputChannelFromMessageArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputChannelFromMessage.
func (s InputChannelFromMessageArray) Retain(keep func(x InputChannelFromMessage) bool) InputChannelFromMessageArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputChannelFromMessageArray) First() (v InputChannelFromMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputChannelFromMessageArray) Last() (v InputChannelFromMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputChannelFromMessageArray) PopFirst() (v InputChannelFromMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputChannelFromMessage
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputChannelFromMessageArray) Pop() (v InputChannelFromMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
