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

// MessagesGetDhConfigRequest represents TL type `messages.getDhConfig#26cf8950`.
// Returns configuration parameters for Diffie-Hellman key generation. Can also return a random sequence of bytes of required length.
//
// See https://core.telegram.org/method/messages.getDhConfig for reference.
type MessagesGetDhConfigRequest struct {
	// Value of the version parameter from messages.dhConfig¹, avialable at the client
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messages.dhConfig
	Version int `tl:"version"`
	// Length of the required random sequence
	RandomLength int `tl:"random_length"`
}

// MessagesGetDhConfigRequestTypeID is TL type id of MessagesGetDhConfigRequest.
const MessagesGetDhConfigRequestTypeID = 0x26cf8950

func (g *MessagesGetDhConfigRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Version == 0) {
		return false
	}
	if !(g.RandomLength == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDhConfigRequest) String() string {
	if g == nil {
		return "MessagesGetDhConfigRequest(nil)"
	}
	type Alias MessagesGetDhConfigRequest
	return fmt.Sprintf("MessagesGetDhConfigRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetDhConfigRequest from given interface.
func (g *MessagesGetDhConfigRequest) FillFrom(from interface {
	GetVersion() (value int)
	GetRandomLength() (value int)
}) {
	g.Version = from.GetVersion()
	g.RandomLength = from.GetRandomLength()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetDhConfigRequest) TypeID() uint32 {
	return MessagesGetDhConfigRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetDhConfigRequest) TypeName() string {
	return "messages.getDhConfig"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetDhConfigRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getDhConfig",
		ID:   MessagesGetDhConfigRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Version",
			SchemaName: "version",
		},
		{
			Name:       "RandomLength",
			SchemaName: "random_length",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetDhConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDhConfig#26cf8950 as nil")
	}
	b.PutID(MessagesGetDhConfigRequestTypeID)
	b.PutInt(g.Version)
	b.PutInt(g.RandomLength)
	return nil
}

// GetVersion returns value of Version field.
func (g *MessagesGetDhConfigRequest) GetVersion() (value int) {
	return g.Version
}

// GetRandomLength returns value of RandomLength field.
func (g *MessagesGetDhConfigRequest) GetRandomLength() (value int) {
	return g.RandomLength
}

// Decode implements bin.Decoder.
func (g *MessagesGetDhConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDhConfig#26cf8950 to nil")
	}
	if err := b.ConsumeID(MessagesGetDhConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: field version: %w", err)
		}
		g.Version = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: field random_length: %w", err)
		}
		g.RandomLength = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDhConfigRequest.
var (
	_ bin.Encoder = &MessagesGetDhConfigRequest{}
	_ bin.Decoder = &MessagesGetDhConfigRequest{}
)

// MessagesGetDhConfig invokes method messages.getDhConfig#26cf8950 returning error if any.
// Returns configuration parameters for Diffie-Hellman key generation. Can also return a random sequence of bytes of required length.
//
// Possible errors:
//  400 RANDOM_LENGTH_INVALID: Random length invalid
//
// See https://core.telegram.org/method/messages.getDhConfig for reference.
func (c *Client) MessagesGetDhConfig(ctx context.Context, request *MessagesGetDhConfigRequest) (MessagesDhConfigClass, error) {
	var result MessagesDhConfigBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.DhConfig, nil
}
