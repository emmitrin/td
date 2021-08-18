// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
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
)

// MessagesCreateChatRequest represents TL type `messages.createChat#9cb126e`.
// Creates a new chat.
//
// See https://core.telegram.org/method/messages.createChat for reference.
type MessagesCreateChatRequest struct {
	// List of user IDs to be invited
	Users []InputUserClass
	// Chat name
	Title string
}

// MessagesCreateChatRequestTypeID is TL type id of MessagesCreateChatRequest.
const MessagesCreateChatRequestTypeID = 0x9cb126e

func (c *MessagesCreateChatRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Users == nil) {
		return false
	}
	if !(c.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesCreateChatRequest) String() string {
	if c == nil {
		return "MessagesCreateChatRequest(nil)"
	}
	type Alias MessagesCreateChatRequest
	return fmt.Sprintf("MessagesCreateChatRequest%+v", Alias(*c))
}

// FillFrom fills MessagesCreateChatRequest from given interface.
func (c *MessagesCreateChatRequest) FillFrom(from interface {
	GetUsers() (value []InputUserClass)
	GetTitle() (value string)
}) {
	c.Users = from.GetUsers()
	c.Title = from.GetTitle()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesCreateChatRequest) TypeID() uint32 {
	return MessagesCreateChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesCreateChatRequest) TypeName() string {
	return "messages.createChat"
}

// TypeInfo returns info about TL type.
func (c *MessagesCreateChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.createChat",
		ID:   MessagesCreateChatRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Users",
			SchemaName: "users",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesCreateChatRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.createChat#9cb126e as nil")
	}
	b.PutID(MessagesCreateChatRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesCreateChatRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.createChat#9cb126e as nil")
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.createChat#9cb126e: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.createChat#9cb126e: field users element with index %d: %w", idx, err)
		}
	}
	b.PutString(c.Title)
	return nil
}

// GetUsers returns value of Users field.
func (c *MessagesCreateChatRequest) GetUsers() (value []InputUserClass) {
	return c.Users
}

// MapUsers returns field Users wrapped in InputUserClassArray helper.
func (c *MessagesCreateChatRequest) MapUsers() (value InputUserClassArray) {
	return InputUserClassArray(c.Users)
}

// GetTitle returns value of Title field.
func (c *MessagesCreateChatRequest) GetTitle() (value string) {
	return c.Title
}

// Decode implements bin.Decoder.
func (c *MessagesCreateChatRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.createChat#9cb126e to nil")
	}
	if err := b.ConsumeID(MessagesCreateChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.createChat#9cb126e: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesCreateChatRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.createChat#9cb126e to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.createChat#9cb126e: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]InputUserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.createChat#9cb126e: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.createChat#9cb126e: field title: %w", err)
		}
		c.Title = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesCreateChatRequest.
var (
	_ bin.Encoder     = &MessagesCreateChatRequest{}
	_ bin.Decoder     = &MessagesCreateChatRequest{}
	_ bin.BareEncoder = &MessagesCreateChatRequest{}
	_ bin.BareDecoder = &MessagesCreateChatRequest{}
)

// MessagesCreateChat invokes method messages.createChat#9cb126e returning error if any.
// Creates a new chat.
//
// Possible errors:
//  400 CHAT_INVALID: Invalid chat
//  400 CHAT_TITLE_EMPTY: No chat title provided
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 USERS_TOO_FEW: Not enough users (to create a chat, for example)
//  403 USER_RESTRICTED: You're spamreported, you can't create channels or chats.
//
// See https://core.telegram.org/method/messages.createChat for reference.
func (c *Client) MessagesCreateChat(ctx context.Context, request *MessagesCreateChatRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
