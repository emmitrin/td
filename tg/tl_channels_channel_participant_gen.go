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

// ChannelsChannelParticipant represents TL type `channels.channelParticipant#d0d9b163`.
// Represents a channel participant
//
// See https://core.telegram.org/constructor/channels.channelParticipant for reference.
type ChannelsChannelParticipant struct {
	// The channel participant
	Participant ChannelParticipantClass `tl:"participant"`
	// Users
	Users []UserClass `tl:"users"`
}

// ChannelsChannelParticipantTypeID is TL type id of ChannelsChannelParticipant.
const ChannelsChannelParticipantTypeID = 0xd0d9b163

func (c *ChannelsChannelParticipant) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Participant == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelsChannelParticipant) String() string {
	if c == nil {
		return "ChannelsChannelParticipant(nil)"
	}
	type Alias ChannelsChannelParticipant
	return fmt.Sprintf("ChannelsChannelParticipant%+v", Alias(*c))
}

// FillFrom fills ChannelsChannelParticipant from given interface.
func (c *ChannelsChannelParticipant) FillFrom(from interface {
	GetParticipant() (value ChannelParticipantClass)
	GetUsers() (value []UserClass)
}) {
	c.Participant = from.GetParticipant()
	c.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsChannelParticipant) TypeID() uint32 {
	return ChannelsChannelParticipantTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsChannelParticipant) TypeName() string {
	return "channels.channelParticipant"
}

// TypeInfo returns info about TL type.
func (c *ChannelsChannelParticipant) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.channelParticipant",
		ID:   ChannelsChannelParticipantTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Participant",
			SchemaName: "participant",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelsChannelParticipant) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.channelParticipant#d0d9b163 as nil")
	}
	b.PutID(ChannelsChannelParticipantTypeID)
	if c.Participant == nil {
		return fmt.Errorf("unable to encode channels.channelParticipant#d0d9b163: field participant is nil")
	}
	if err := c.Participant.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.channelParticipant#d0d9b163: field participant: %w", err)
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode channels.channelParticipant#d0d9b163: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.channelParticipant#d0d9b163: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetParticipant returns value of Participant field.
func (c *ChannelsChannelParticipant) GetParticipant() (value ChannelParticipantClass) {
	return c.Participant
}

// GetUsers returns value of Users field.
func (c *ChannelsChannelParticipant) GetUsers() (value []UserClass) {
	return c.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (c *ChannelsChannelParticipant) MapUsers() (value UserClassArray) {
	return UserClassArray(c.Users)
}

// Decode implements bin.Decoder.
func (c *ChannelsChannelParticipant) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.channelParticipant#d0d9b163 to nil")
	}
	if err := b.ConsumeID(ChannelsChannelParticipantTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.channelParticipant#d0d9b163: %w", err)
	}
	{
		value, err := DecodeChannelParticipant(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.channelParticipant#d0d9b163: field participant: %w", err)
		}
		c.Participant = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.channelParticipant#d0d9b163: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.channelParticipant#d0d9b163: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsChannelParticipant.
var (
	_ bin.Encoder = &ChannelsChannelParticipant{}
	_ bin.Decoder = &ChannelsChannelParticipant{}
)
