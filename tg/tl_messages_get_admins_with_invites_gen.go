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

// MessagesGetAdminsWithInvitesRequest represents TL type `messages.getAdminsWithInvites#3920e6ef`.
//
// See https://core.telegram.org/method/messages.getAdminsWithInvites for reference.
type MessagesGetAdminsWithInvitesRequest struct {
	// Peer field of MessagesGetAdminsWithInvitesRequest.
	Peer InputPeerClass `tl:"peer"`
}

// MessagesGetAdminsWithInvitesRequestTypeID is TL type id of MessagesGetAdminsWithInvitesRequest.
const MessagesGetAdminsWithInvitesRequestTypeID = 0x3920e6ef

func (g *MessagesGetAdminsWithInvitesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAdminsWithInvitesRequest) String() string {
	if g == nil {
		return "MessagesGetAdminsWithInvitesRequest(nil)"
	}
	type Alias MessagesGetAdminsWithInvitesRequest
	return fmt.Sprintf("MessagesGetAdminsWithInvitesRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetAdminsWithInvitesRequest from given interface.
func (g *MessagesGetAdminsWithInvitesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	g.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetAdminsWithInvitesRequest) TypeID() uint32 {
	return MessagesGetAdminsWithInvitesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetAdminsWithInvitesRequest) TypeName() string {
	return "messages.getAdminsWithInvites"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetAdminsWithInvitesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getAdminsWithInvites",
		ID:   MessagesGetAdminsWithInvitesRequestTypeID,
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
func (g *MessagesGetAdminsWithInvitesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAdminsWithInvites#3920e6ef as nil")
	}
	b.PutID(MessagesGetAdminsWithInvitesRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getAdminsWithInvites#3920e6ef: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getAdminsWithInvites#3920e6ef: field peer: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetAdminsWithInvitesRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// Decode implements bin.Decoder.
func (g *MessagesGetAdminsWithInvitesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAdminsWithInvites#3920e6ef to nil")
	}
	if err := b.ConsumeID(MessagesGetAdminsWithInvitesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAdminsWithInvites#3920e6ef: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAdminsWithInvites#3920e6ef: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetAdminsWithInvitesRequest.
var (
	_ bin.Encoder = &MessagesGetAdminsWithInvitesRequest{}
	_ bin.Decoder = &MessagesGetAdminsWithInvitesRequest{}
)

// MessagesGetAdminsWithInvites invokes method messages.getAdminsWithInvites#3920e6ef returning error if any.
//
// See https://core.telegram.org/method/messages.getAdminsWithInvites for reference.
func (c *Client) MessagesGetAdminsWithInvites(ctx context.Context, peer InputPeerClass) (*MessagesChatAdminsWithInvites, error) {
	var result MessagesChatAdminsWithInvites

	request := &MessagesGetAdminsWithInvitesRequest{
		Peer: peer,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
