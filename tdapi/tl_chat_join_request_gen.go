// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// ChatJoinRequest represents TL type `chatJoinRequest#3897a68`.
type ChatJoinRequest struct {
	// User identifier
	UserID int64
	// Point in time (Unix timestamp) when the user sent the join request
	Date int32
	// A short bio of the user
	Bio string
}

// ChatJoinRequestTypeID is TL type id of ChatJoinRequest.
const ChatJoinRequestTypeID = 0x3897a68

// Ensuring interfaces in compile-time for ChatJoinRequest.
var (
	_ bin.Encoder     = &ChatJoinRequest{}
	_ bin.Decoder     = &ChatJoinRequest{}
	_ bin.BareEncoder = &ChatJoinRequest{}
	_ bin.BareDecoder = &ChatJoinRequest{}
)

func (c *ChatJoinRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}
	if !(c.Bio == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatJoinRequest) String() string {
	if c == nil {
		return "ChatJoinRequest(nil)"
	}
	type Alias ChatJoinRequest
	return fmt.Sprintf("ChatJoinRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatJoinRequest) TypeID() uint32 {
	return ChatJoinRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatJoinRequest) TypeName() string {
	return "chatJoinRequest"
}

// TypeInfo returns info about TL type.
func (c *ChatJoinRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatJoinRequest",
		ID:   ChatJoinRequestTypeID,
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
		{
			Name:       "Bio",
			SchemaName: "bio",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatJoinRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatJoinRequest#3897a68 as nil")
	}
	b.PutID(ChatJoinRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatJoinRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatJoinRequest#3897a68 as nil")
	}
	b.PutInt53(c.UserID)
	b.PutInt32(c.Date)
	b.PutString(c.Bio)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatJoinRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatJoinRequest#3897a68 to nil")
	}
	if err := b.ConsumeID(ChatJoinRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode chatJoinRequest#3897a68: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatJoinRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatJoinRequest#3897a68 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chatJoinRequest#3897a68: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatJoinRequest#3897a68: field date: %w", err)
		}
		c.Date = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatJoinRequest#3897a68: field bio: %w", err)
		}
		c.Bio = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatJoinRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatJoinRequest#3897a68 as nil")
	}
	b.ObjStart()
	b.PutID("chatJoinRequest")
	b.FieldStart("user_id")
	b.PutInt53(c.UserID)
	b.FieldStart("date")
	b.PutInt32(c.Date)
	b.FieldStart("bio")
	b.PutString(c.Bio)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatJoinRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatJoinRequest#3897a68 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatJoinRequest"); err != nil {
				return fmt.Errorf("unable to decode chatJoinRequest#3897a68: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatJoinRequest#3897a68: field user_id: %w", err)
			}
			c.UserID = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatJoinRequest#3897a68: field date: %w", err)
			}
			c.Date = value
		case "bio":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatJoinRequest#3897a68: field bio: %w", err)
			}
			c.Bio = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *ChatJoinRequest) GetUserID() (value int64) {
	return c.UserID
}

// GetDate returns value of Date field.
func (c *ChatJoinRequest) GetDate() (value int32) {
	return c.Date
}

// GetBio returns value of Bio field.
func (c *ChatJoinRequest) GetBio() (value string) {
	return c.Bio
}
