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

// MessagesStartBotRequest represents TL type `messages.startBot#e6df7378`.
// Start a conversation with a bot using a deep linking parameter¹
//
// Links:
//  1) https://core.telegram.org/bots#deep-linking
//
// See https://core.telegram.org/method/messages.startBot for reference.
type MessagesStartBotRequest struct {
	// The bot
	Bot InputUserClass `tl:"bot"`
	// The chat where to start the bot, can be the bot's private chat or a group
	Peer InputPeerClass `tl:"peer"`
	// Random ID to avoid resending the same message
	RandomID int64 `tl:"random_id"`
	// Deep linking parameter¹
	//
	// Links:
	//  1) https://core.telegram.org/bots#deep-linking
	StartParam string `tl:"start_param"`
}

// MessagesStartBotRequestTypeID is TL type id of MessagesStartBotRequest.
const MessagesStartBotRequestTypeID = 0xe6df7378

func (s *MessagesStartBotRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Bot == nil) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.StartParam == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesStartBotRequest) String() string {
	if s == nil {
		return "MessagesStartBotRequest(nil)"
	}
	type Alias MessagesStartBotRequest
	return fmt.Sprintf("MessagesStartBotRequest%+v", Alias(*s))
}

// FillFrom fills MessagesStartBotRequest from given interface.
func (s *MessagesStartBotRequest) FillFrom(from interface {
	GetBot() (value InputUserClass)
	GetPeer() (value InputPeerClass)
	GetRandomID() (value int64)
	GetStartParam() (value string)
}) {
	s.Bot = from.GetBot()
	s.Peer = from.GetPeer()
	s.RandomID = from.GetRandomID()
	s.StartParam = from.GetStartParam()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesStartBotRequest) TypeID() uint32 {
	return MessagesStartBotRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesStartBotRequest) TypeName() string {
	return "messages.startBot"
}

// TypeInfo returns info about TL type.
func (s *MessagesStartBotRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.startBot",
		ID:   MessagesStartBotRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "StartParam",
			SchemaName: "start_param",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesStartBotRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.startBot#e6df7378 as nil")
	}
	b.PutID(MessagesStartBotRequestTypeID)
	if s.Bot == nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field bot is nil")
	}
	if err := s.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field bot: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field peer: %w", err)
	}
	b.PutLong(s.RandomID)
	b.PutString(s.StartParam)
	return nil
}

// GetBot returns value of Bot field.
func (s *MessagesStartBotRequest) GetBot() (value InputUserClass) {
	return s.Bot
}

// GetPeer returns value of Peer field.
func (s *MessagesStartBotRequest) GetPeer() (value InputPeerClass) {
	return s.Peer
}

// GetRandomID returns value of RandomID field.
func (s *MessagesStartBotRequest) GetRandomID() (value int64) {
	return s.RandomID
}

// GetStartParam returns value of StartParam field.
func (s *MessagesStartBotRequest) GetStartParam() (value string) {
	return s.StartParam
}

// Decode implements bin.Decoder.
func (s *MessagesStartBotRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.startBot#e6df7378 to nil")
	}
	if err := b.ConsumeID(MessagesStartBotRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.startBot#e6df7378: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field bot: %w", err)
		}
		s.Bot = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field start_param: %w", err)
		}
		s.StartParam = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesStartBotRequest.
var (
	_ bin.Encoder = &MessagesStartBotRequest{}
	_ bin.Decoder = &MessagesStartBotRequest{}
)

// MessagesStartBot invokes method messages.startBot#e6df7378 returning error if any.
// Start a conversation with a bot using a deep linking parameter¹
//
// Links:
//  1) https://core.telegram.org/bots#deep-linking
//
// Possible errors:
//  400 BOT_INVALID: This is not a valid bot
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 START_PARAM_EMPTY: The start parameter is empty
//  400 START_PARAM_INVALID: Start parameter invalid
//  400 START_PARAM_TOO_LONG: Start parameter is too long
//
// See https://core.telegram.org/method/messages.startBot for reference.
func (c *Client) MessagesStartBot(ctx context.Context, request *MessagesStartBotRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
