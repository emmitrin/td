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

// MessagesGetOnlinesRequest represents TL type `messages.getOnlines#6e2be050`.
// Get count of online users in a chat
//
// See https://core.telegram.org/method/messages.getOnlines for reference.
type MessagesGetOnlinesRequest struct {
	// The chat
	Peer InputPeerClass `tl:"peer"`
}

// MessagesGetOnlinesRequestTypeID is TL type id of MessagesGetOnlinesRequest.
const MessagesGetOnlinesRequestTypeID = 0x6e2be050

func (g *MessagesGetOnlinesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetOnlinesRequest) String() string {
	if g == nil {
		return "MessagesGetOnlinesRequest(nil)"
	}
	type Alias MessagesGetOnlinesRequest
	return fmt.Sprintf("MessagesGetOnlinesRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetOnlinesRequest from given interface.
func (g *MessagesGetOnlinesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	g.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetOnlinesRequest) TypeID() uint32 {
	return MessagesGetOnlinesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetOnlinesRequest) TypeName() string {
	return "messages.getOnlines"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetOnlinesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getOnlines",
		ID:   MessagesGetOnlinesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetOnlinesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getOnlines#6e2be050 as nil")
	}
	b.PutID(MessagesGetOnlinesRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getOnlines#6e2be050: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getOnlines#6e2be050: field peer: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetOnlinesRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// Decode implements bin.Decoder.
func (g *MessagesGetOnlinesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getOnlines#6e2be050 to nil")
	}
	if err := b.ConsumeID(MessagesGetOnlinesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getOnlines#6e2be050: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getOnlines#6e2be050: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetOnlinesRequest.
var (
	_ bin.Encoder = &MessagesGetOnlinesRequest{}
	_ bin.Decoder = &MessagesGetOnlinesRequest{}
)

// MessagesGetOnlines invokes method messages.getOnlines#6e2be050 returning error if any.
// Get count of online users in a chat
//
// Possible errors:
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.getOnlines for reference.
func (c *Client) MessagesGetOnlines(ctx context.Context, peer InputPeerClass) (*ChatOnlines, error) {
	var result ChatOnlines

	request := &MessagesGetOnlinesRequest{
		Peer: peer,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
