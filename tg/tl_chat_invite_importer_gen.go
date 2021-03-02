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

// ChatInviteImporter represents TL type `chatInviteImporter#1e3e6680`.
//
// See https://core.telegram.org/constructor/chatInviteImporter for reference.
type ChatInviteImporter struct {
	// UserID field of ChatInviteImporter.
	UserID int `tl:"user_id"`
	// Date field of ChatInviteImporter.
	Date int `tl:"date"`
}

// ChatInviteImporterTypeID is TL type id of ChatInviteImporter.
const ChatInviteImporterTypeID = 0x1e3e6680

func (c *ChatInviteImporter) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteImporter) String() string {
	if c == nil {
		return "ChatInviteImporter(nil)"
	}
	type Alias ChatInviteImporter
	return fmt.Sprintf("ChatInviteImporter%+v", Alias(*c))
}

// FillFrom fills ChatInviteImporter from given interface.
func (c *ChatInviteImporter) FillFrom(from interface {
	GetUserID() (value int)
	GetDate() (value int)
}) {
	c.UserID = from.GetUserID()
	c.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteImporter) TypeID() uint32 {
	return ChatInviteImporterTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteImporter) TypeName() string {
	return "chatInviteImporter"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteImporter) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteImporter",
		ID:   ChatInviteImporterTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInviteImporter) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteImporter#1e3e6680 as nil")
	}
	b.PutID(ChatInviteImporterTypeID)
	b.PutInt(c.UserID)
	b.PutInt(c.Date)
	return nil
}

// GetUserID returns value of UserID field.
func (c *ChatInviteImporter) GetUserID() (value int) {
	return c.UserID
}

// GetDate returns value of Date field.
func (c *ChatInviteImporter) GetDate() (value int) {
	return c.Date
}

// Decode implements bin.Decoder.
func (c *ChatInviteImporter) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteImporter#1e3e6680 to nil")
	}
	if err := b.ConsumeID(ChatInviteImporterTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteImporter#1e3e6680: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteImporter#1e3e6680: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteImporter#1e3e6680: field date: %w", err)
		}
		c.Date = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChatInviteImporter.
var (
	_ bin.Encoder = &ChatInviteImporter{}
	_ bin.Decoder = &ChatInviteImporter{}
)
