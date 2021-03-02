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

// InputBotInlineMessageID represents TL type `inputBotInlineMessageID#890c3d89`.
// Represents a sent inline message from the perspective of a bot
//
// See https://core.telegram.org/constructor/inputBotInlineMessageID for reference.
type InputBotInlineMessageID struct {
	// DC ID to use when working with this inline message
	DCID int `tl:"dc_id"`
	// ID of message
	ID int64 `tl:"id"`
	// Access hash of message
	AccessHash int64 `tl:"access_hash"`
}

// InputBotInlineMessageIDTypeID is TL type id of InputBotInlineMessageID.
const InputBotInlineMessageIDTypeID = 0x890c3d89

func (i *InputBotInlineMessageID) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.DCID == 0) {
		return false
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputBotInlineMessageID) String() string {
	if i == nil {
		return "InputBotInlineMessageID(nil)"
	}
	type Alias InputBotInlineMessageID
	return fmt.Sprintf("InputBotInlineMessageID%+v", Alias(*i))
}

// FillFrom fills InputBotInlineMessageID from given interface.
func (i *InputBotInlineMessageID) FillFrom(from interface {
	GetDCID() (value int)
	GetID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.DCID = from.GetDCID()
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputBotInlineMessageID) TypeID() uint32 {
	return InputBotInlineMessageIDTypeID
}

// TypeName returns name of type in TL schema.
func (*InputBotInlineMessageID) TypeName() string {
	return "inputBotInlineMessageID"
}

// TypeInfo returns info about TL type.
func (i *InputBotInlineMessageID) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputBotInlineMessageID",
		ID:   InputBotInlineMessageIDTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputBotInlineMessageID) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineMessageID#890c3d89 as nil")
	}
	b.PutID(InputBotInlineMessageIDTypeID)
	b.PutInt(i.DCID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// GetDCID returns value of DCID field.
func (i *InputBotInlineMessageID) GetDCID() (value int) {
	return i.DCID
}

// GetID returns value of ID field.
func (i *InputBotInlineMessageID) GetID() (value int64) {
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputBotInlineMessageID) GetAccessHash() (value int64) {
	return i.AccessHash
}

// Decode implements bin.Decoder.
func (i *InputBotInlineMessageID) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineMessageID#890c3d89 to nil")
	}
	if err := b.ConsumeID(InputBotInlineMessageIDTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineMessageID#890c3d89: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageID#890c3d89: field dc_id: %w", err)
		}
		i.DCID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageID#890c3d89: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageID#890c3d89: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputBotInlineMessageID.
var (
	_ bin.Encoder = &InputBotInlineMessageID{}
	_ bin.Decoder = &InputBotInlineMessageID{}
)
