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

// PhoneCreateGroupCallRequest represents TL type `phone.createGroupCall#bd3dabe0`.
//
// See https://core.telegram.org/method/phone.createGroupCall for reference.
type PhoneCreateGroupCallRequest struct {
	// Peer field of PhoneCreateGroupCallRequest.
	Peer InputPeerClass `tl:"peer"`
	// RandomID field of PhoneCreateGroupCallRequest.
	RandomID int `tl:"random_id"`
}

// PhoneCreateGroupCallRequestTypeID is TL type id of PhoneCreateGroupCallRequest.
const PhoneCreateGroupCallRequestTypeID = 0xbd3dabe0

func (c *PhoneCreateGroupCallRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Peer == nil) {
		return false
	}
	if !(c.RandomID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *PhoneCreateGroupCallRequest) String() string {
	if c == nil {
		return "PhoneCreateGroupCallRequest(nil)"
	}
	type Alias PhoneCreateGroupCallRequest
	return fmt.Sprintf("PhoneCreateGroupCallRequest%+v", Alias(*c))
}

// FillFrom fills PhoneCreateGroupCallRequest from given interface.
func (c *PhoneCreateGroupCallRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetRandomID() (value int)
}) {
	c.Peer = from.GetPeer()
	c.RandomID = from.GetRandomID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneCreateGroupCallRequest) TypeID() uint32 {
	return PhoneCreateGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneCreateGroupCallRequest) TypeName() string {
	return "phone.createGroupCall"
}

// TypeInfo returns info about TL type.
func (c *PhoneCreateGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.createGroupCall",
		ID:   PhoneCreateGroupCallRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *PhoneCreateGroupCallRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode phone.createGroupCall#bd3dabe0 as nil")
	}
	b.PutID(PhoneCreateGroupCallRequestTypeID)
	if c.Peer == nil {
		return fmt.Errorf("unable to encode phone.createGroupCall#bd3dabe0: field peer is nil")
	}
	if err := c.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.createGroupCall#bd3dabe0: field peer: %w", err)
	}
	b.PutInt(c.RandomID)
	return nil
}

// GetPeer returns value of Peer field.
func (c *PhoneCreateGroupCallRequest) GetPeer() (value InputPeerClass) {
	return c.Peer
}

// GetRandomID returns value of RandomID field.
func (c *PhoneCreateGroupCallRequest) GetRandomID() (value int) {
	return c.RandomID
}

// Decode implements bin.Decoder.
func (c *PhoneCreateGroupCallRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode phone.createGroupCall#bd3dabe0 to nil")
	}
	if err := b.ConsumeID(PhoneCreateGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.createGroupCall#bd3dabe0: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.createGroupCall#bd3dabe0: field peer: %w", err)
		}
		c.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.createGroupCall#bd3dabe0: field random_id: %w", err)
		}
		c.RandomID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneCreateGroupCallRequest.
var (
	_ bin.Encoder = &PhoneCreateGroupCallRequest{}
	_ bin.Decoder = &PhoneCreateGroupCallRequest{}
)

// PhoneCreateGroupCall invokes method phone.createGroupCall#bd3dabe0 returning error if any.
//
// See https://core.telegram.org/method/phone.createGroupCall for reference.
func (c *Client) PhoneCreateGroupCall(ctx context.Context, request *PhoneCreateGroupCallRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
