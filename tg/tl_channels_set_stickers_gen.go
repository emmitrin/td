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

// ChannelsSetStickersRequest represents TL type `channels.setStickers#ea8ca4f9`.
// Associate a stickerset to the supergroup
//
// See https://core.telegram.org/method/channels.setStickers for reference.
type ChannelsSetStickersRequest struct {
	// Supergroup
	Channel InputChannelClass `tl:"channel"`
	// The stickerset to associate
	Stickerset InputStickerSetClass `tl:"stickerset"`
}

// ChannelsSetStickersRequestTypeID is TL type id of ChannelsSetStickersRequest.
const ChannelsSetStickersRequestTypeID = 0xea8ca4f9

func (s *ChannelsSetStickersRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Channel == nil) {
		return false
	}
	if !(s.Stickerset == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ChannelsSetStickersRequest) String() string {
	if s == nil {
		return "ChannelsSetStickersRequest(nil)"
	}
	type Alias ChannelsSetStickersRequest
	return fmt.Sprintf("ChannelsSetStickersRequest%+v", Alias(*s))
}

// FillFrom fills ChannelsSetStickersRequest from given interface.
func (s *ChannelsSetStickersRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetStickerset() (value InputStickerSetClass)
}) {
	s.Channel = from.GetChannel()
	s.Stickerset = from.GetStickerset()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsSetStickersRequest) TypeID() uint32 {
	return ChannelsSetStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsSetStickersRequest) TypeName() string {
	return "channels.setStickers"
}

// TypeInfo returns info about TL type.
func (s *ChannelsSetStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.setStickers",
		ID:   ChannelsSetStickersRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ChannelsSetStickersRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.setStickers#ea8ca4f9 as nil")
	}
	b.PutID(ChannelsSetStickersRequestTypeID)
	if s.Channel == nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field channel is nil")
	}
	if err := s.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field channel: %w", err)
	}
	if s.Stickerset == nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field stickerset is nil")
	}
	if err := s.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field stickerset: %w", err)
	}
	return nil
}

// GetChannel returns value of Channel field.
func (s *ChannelsSetStickersRequest) GetChannel() (value InputChannelClass) {
	return s.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (s *ChannelsSetStickersRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return s.Channel.AsNotEmpty()
}

// GetStickerset returns value of Stickerset field.
func (s *ChannelsSetStickersRequest) GetStickerset() (value InputStickerSetClass) {
	return s.Stickerset
}

// Decode implements bin.Decoder.
func (s *ChannelsSetStickersRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.setStickers#ea8ca4f9 to nil")
	}
	if err := b.ConsumeID(ChannelsSetStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.setStickers#ea8ca4f9: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.setStickers#ea8ca4f9: field channel: %w", err)
		}
		s.Channel = value
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.setStickers#ea8ca4f9: field stickerset: %w", err)
		}
		s.Stickerset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsSetStickersRequest.
var (
	_ bin.Encoder = &ChannelsSetStickersRequest{}
	_ bin.Decoder = &ChannelsSetStickersRequest{}
)

// ChannelsSetStickers invokes method channels.setStickers#ea8ca4f9 returning error if any.
// Associate a stickerset to the supergroup
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 PARTICIPANTS_TOO_FEW: Not enough participants
//
// See https://core.telegram.org/method/channels.setStickers for reference.
// Can be used by bots.
func (c *Client) ChannelsSetStickers(ctx context.Context, request *ChannelsSetStickersRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
