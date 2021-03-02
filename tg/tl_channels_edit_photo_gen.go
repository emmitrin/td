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

// ChannelsEditPhotoRequest represents TL type `channels.editPhoto#f12e57c9`.
// Change the photo of a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.editPhoto for reference.
type ChannelsEditPhotoRequest struct {
	// Channel/supergroup whose photo should be edited
	Channel InputChannelClass `tl:"channel"`
	// New photo
	Photo InputChatPhotoClass `tl:"photo"`
}

// ChannelsEditPhotoRequestTypeID is TL type id of ChannelsEditPhotoRequest.
const ChannelsEditPhotoRequestTypeID = 0xf12e57c9

func (e *ChannelsEditPhotoRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Channel == nil) {
		return false
	}
	if !(e.Photo == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ChannelsEditPhotoRequest) String() string {
	if e == nil {
		return "ChannelsEditPhotoRequest(nil)"
	}
	type Alias ChannelsEditPhotoRequest
	return fmt.Sprintf("ChannelsEditPhotoRequest%+v", Alias(*e))
}

// FillFrom fills ChannelsEditPhotoRequest from given interface.
func (e *ChannelsEditPhotoRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetPhoto() (value InputChatPhotoClass)
}) {
	e.Channel = from.GetChannel()
	e.Photo = from.GetPhoto()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsEditPhotoRequest) TypeID() uint32 {
	return ChannelsEditPhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsEditPhotoRequest) TypeName() string {
	return "channels.editPhoto"
}

// TypeInfo returns info about TL type.
func (e *ChannelsEditPhotoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.editPhoto",
		ID:   ChannelsEditPhotoRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *ChannelsEditPhotoRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editPhoto#f12e57c9 as nil")
	}
	b.PutID(ChannelsEditPhotoRequestTypeID)
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editPhoto#f12e57c9: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editPhoto#f12e57c9: field channel: %w", err)
	}
	if e.Photo == nil {
		return fmt.Errorf("unable to encode channels.editPhoto#f12e57c9: field photo is nil")
	}
	if err := e.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editPhoto#f12e57c9: field photo: %w", err)
	}
	return nil
}

// GetChannel returns value of Channel field.
func (e *ChannelsEditPhotoRequest) GetChannel() (value InputChannelClass) {
	return e.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (e *ChannelsEditPhotoRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return e.Channel.AsNotEmpty()
}

// GetPhoto returns value of Photo field.
func (e *ChannelsEditPhotoRequest) GetPhoto() (value InputChatPhotoClass) {
	return e.Photo
}

// Decode implements bin.Decoder.
func (e *ChannelsEditPhotoRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editPhoto#f12e57c9 to nil")
	}
	if err := b.ConsumeID(ChannelsEditPhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editPhoto#f12e57c9: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editPhoto#f12e57c9: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editPhoto#f12e57c9: field photo: %w", err)
		}
		e.Photo = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsEditPhotoRequest.
var (
	_ bin.Encoder = &ChannelsEditPhotoRequest{}
	_ bin.Decoder = &ChannelsEditPhotoRequest{}
)

// ChannelsEditPhoto invokes method channels.editPhoto#f12e57c9 returning error if any.
// Change the photo of a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 PHOTO_CROP_SIZE_SMALL: Photo is too small
//  400 PHOTO_EXT_INVALID: The extension of the photo is invalid
//  400 PHOTO_INVALID: Photo invalid
//
// See https://core.telegram.org/method/channels.editPhoto for reference.
// Can be used by bots.
func (c *Client) ChannelsEditPhoto(ctx context.Context, request *ChannelsEditPhotoRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
