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

// GetChatInviteLinkRequest represents TL type `getChatInviteLink#e36a41fd`.
type GetChatInviteLinkRequest struct {
	// Chat identifier
	ChatID int64
	// Invite link to get
	InviteLink string
}

// GetChatInviteLinkRequestTypeID is TL type id of GetChatInviteLinkRequest.
const GetChatInviteLinkRequestTypeID = 0xe36a41fd

// Ensuring interfaces in compile-time for GetChatInviteLinkRequest.
var (
	_ bin.Encoder     = &GetChatInviteLinkRequest{}
	_ bin.Decoder     = &GetChatInviteLinkRequest{}
	_ bin.BareEncoder = &GetChatInviteLinkRequest{}
	_ bin.BareDecoder = &GetChatInviteLinkRequest{}
)

func (g *GetChatInviteLinkRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.InviteLink == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatInviteLinkRequest) String() string {
	if g == nil {
		return "GetChatInviteLinkRequest(nil)"
	}
	type Alias GetChatInviteLinkRequest
	return fmt.Sprintf("GetChatInviteLinkRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatInviteLinkRequest) TypeID() uint32 {
	return GetChatInviteLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatInviteLinkRequest) TypeName() string {
	return "getChatInviteLink"
}

// TypeInfo returns info about TL type.
func (g *GetChatInviteLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatInviteLink",
		ID:   GetChatInviteLinkRequestTypeID,
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
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatInviteLinkRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLink#e36a41fd as nil")
	}
	b.PutID(GetChatInviteLinkRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatInviteLinkRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLink#e36a41fd as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutString(g.InviteLink)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatInviteLinkRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLink#e36a41fd to nil")
	}
	if err := b.ConsumeID(GetChatInviteLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatInviteLink#e36a41fd: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatInviteLinkRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLink#e36a41fd to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLink#e36a41fd: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLink#e36a41fd: field invite_link: %w", err)
		}
		g.InviteLink = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatInviteLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLink#e36a41fd as nil")
	}
	b.ObjStart()
	b.PutID("getChatInviteLink")
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.FieldStart("invite_link")
	b.PutString(g.InviteLink)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatInviteLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLink#e36a41fd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatInviteLink"); err != nil {
				return fmt.Errorf("unable to decode getChatInviteLink#e36a41fd: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatInviteLink#e36a41fd: field chat_id: %w", err)
			}
			g.ChatID = value
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getChatInviteLink#e36a41fd: field invite_link: %w", err)
			}
			g.InviteLink = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatInviteLinkRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetInviteLink returns value of InviteLink field.
func (g *GetChatInviteLinkRequest) GetInviteLink() (value string) {
	return g.InviteLink
}

// GetChatInviteLink invokes method getChatInviteLink#e36a41fd returning error if any.
func (c *Client) GetChatInviteLink(ctx context.Context, request *GetChatInviteLinkRequest) (*ChatInviteLink, error) {
	var result ChatInviteLink

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
