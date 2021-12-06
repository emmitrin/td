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

// GetChatAdministratorsRequest represents TL type `getChatAdministrators#5c0eb6bb`.
type GetChatAdministratorsRequest struct {
	// Chat identifier
	ChatID int64
}

// GetChatAdministratorsRequestTypeID is TL type id of GetChatAdministratorsRequest.
const GetChatAdministratorsRequestTypeID = 0x5c0eb6bb

// Ensuring interfaces in compile-time for GetChatAdministratorsRequest.
var (
	_ bin.Encoder     = &GetChatAdministratorsRequest{}
	_ bin.Decoder     = &GetChatAdministratorsRequest{}
	_ bin.BareEncoder = &GetChatAdministratorsRequest{}
	_ bin.BareDecoder = &GetChatAdministratorsRequest{}
)

func (g *GetChatAdministratorsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatAdministratorsRequest) String() string {
	if g == nil {
		return "GetChatAdministratorsRequest(nil)"
	}
	type Alias GetChatAdministratorsRequest
	return fmt.Sprintf("GetChatAdministratorsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatAdministratorsRequest) TypeID() uint32 {
	return GetChatAdministratorsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatAdministratorsRequest) TypeName() string {
	return "getChatAdministrators"
}

// TypeInfo returns info about TL type.
func (g *GetChatAdministratorsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatAdministrators",
		ID:   GetChatAdministratorsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatAdministratorsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatAdministrators#5c0eb6bb as nil")
	}
	b.PutID(GetChatAdministratorsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatAdministratorsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatAdministrators#5c0eb6bb as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatAdministratorsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatAdministrators#5c0eb6bb to nil")
	}
	if err := b.ConsumeID(GetChatAdministratorsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatAdministrators#5c0eb6bb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatAdministratorsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatAdministrators#5c0eb6bb to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatAdministrators#5c0eb6bb: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatAdministratorsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatAdministrators#5c0eb6bb as nil")
	}
	b.ObjStart()
	b.PutID("getChatAdministrators")
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatAdministratorsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatAdministrators#5c0eb6bb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatAdministrators"); err != nil {
				return fmt.Errorf("unable to decode getChatAdministrators#5c0eb6bb: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatAdministrators#5c0eb6bb: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatAdministratorsRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetChatAdministrators invokes method getChatAdministrators#5c0eb6bb returning error if any.
func (c *Client) GetChatAdministrators(ctx context.Context, chatid int64) (*ChatAdministrators, error) {
	var result ChatAdministrators

	request := &GetChatAdministratorsRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
