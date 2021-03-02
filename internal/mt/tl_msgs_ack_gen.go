// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// MsgsAck represents TL type `msgs_ack#62d6b459`.
type MsgsAck struct {
	// MsgIds field of MsgsAck.
	MsgIds []int64 `tl:"msg_ids"`
}

// MsgsAckTypeID is TL type id of MsgsAck.
const MsgsAckTypeID = 0x62d6b459

func (m *MsgsAck) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MsgIds == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgsAck) String() string {
	if m == nil {
		return "MsgsAck(nil)"
	}
	type Alias MsgsAck
	return fmt.Sprintf("MsgsAck%+v", Alias(*m))
}

// FillFrom fills MsgsAck from given interface.
func (m *MsgsAck) FillFrom(from interface {
	GetMsgIds() (value []int64)
}) {
	m.MsgIds = from.GetMsgIds()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MsgsAck) TypeID() uint32 {
	return MsgsAckTypeID
}

// TypeName returns name of type in TL schema.
func (*MsgsAck) TypeName() string {
	return "msgs_ack"
}

// TypeInfo returns info about TL type.
func (m *MsgsAck) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "msgs_ack",
		ID:   MsgsAckTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgIds",
			SchemaName: "msg_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MsgsAck) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_ack#62d6b459 as nil")
	}
	b.PutID(MsgsAckTypeID)
	b.PutVectorHeader(len(m.MsgIds))
	for _, v := range m.MsgIds {
		b.PutLong(v)
	}
	return nil
}

// GetMsgIds returns value of MsgIds field.
func (m *MsgsAck) GetMsgIds() (value []int64) {
	return m.MsgIds
}

// Decode implements bin.Decoder.
func (m *MsgsAck) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_ack#62d6b459 to nil")
	}
	if err := b.ConsumeID(MsgsAckTypeID); err != nil {
		return fmt.Errorf("unable to decode msgs_ack#62d6b459: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode msgs_ack#62d6b459: field msg_ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode msgs_ack#62d6b459: field msg_ids: %w", err)
			}
			m.MsgIds = append(m.MsgIds, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MsgsAck.
var (
	_ bin.Encoder = &MsgsAck{}
	_ bin.Decoder = &MsgsAck{}
)
