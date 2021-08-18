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

// MessagesPeerDialogs represents TL type `messages.peerDialogs#3371c354`.
// Dialog info of multiple peers
//
// See https://core.telegram.org/constructor/messages.peerDialogs for reference.
type MessagesPeerDialogs struct {
	// Dialog info
	Dialogs []DialogClass
	// Messages mentioned in dialog info
	Messages []MessageClass
	// Chats
	Chats []ChatClass
	// Users
	Users []UserClass
	// Current update state of dialog¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	State UpdatesState
}

// MessagesPeerDialogsTypeID is TL type id of MessagesPeerDialogs.
const MessagesPeerDialogsTypeID = 0x3371c354

func (p *MessagesPeerDialogs) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Dialogs == nil) {
		return false
	}
	if !(p.Messages == nil) {
		return false
	}
	if !(p.Chats == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}
	if !(p.State.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *MessagesPeerDialogs) String() string {
	if p == nil {
		return "MessagesPeerDialogs(nil)"
	}
	type Alias MessagesPeerDialogs
	return fmt.Sprintf("MessagesPeerDialogs%+v", Alias(*p))
}

// FillFrom fills MessagesPeerDialogs from given interface.
func (p *MessagesPeerDialogs) FillFrom(from interface {
	GetDialogs() (value []DialogClass)
	GetMessages() (value []MessageClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
	GetState() (value UpdatesState)
}) {
	p.Dialogs = from.GetDialogs()
	p.Messages = from.GetMessages()
	p.Chats = from.GetChats()
	p.Users = from.GetUsers()
	p.State = from.GetState()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesPeerDialogs) TypeID() uint32 {
	return MessagesPeerDialogsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesPeerDialogs) TypeName() string {
	return "messages.peerDialogs"
}

// TypeInfo returns info about TL type.
func (p *MessagesPeerDialogs) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.peerDialogs",
		ID:   MessagesPeerDialogsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Dialogs",
			SchemaName: "dialogs",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
		{
			Name:       "State",
			SchemaName: "state",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *MessagesPeerDialogs) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode messages.peerDialogs#3371c354 as nil")
	}
	b.PutID(MessagesPeerDialogsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *MessagesPeerDialogs) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode messages.peerDialogs#3371c354 as nil")
	}
	b.PutVectorHeader(len(p.Dialogs))
	for idx, v := range p.Dialogs {
		if v == nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field dialogs element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field dialogs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Messages))
	for idx, v := range p.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Chats))
	for idx, v := range p.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field users element with index %d: %w", idx, err)
		}
	}
	if err := p.State.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field state: %w", err)
	}
	return nil
}

// GetDialogs returns value of Dialogs field.
func (p *MessagesPeerDialogs) GetDialogs() (value []DialogClass) {
	return p.Dialogs
}

// MapDialogs returns field Dialogs wrapped in DialogClassArray helper.
func (p *MessagesPeerDialogs) MapDialogs() (value DialogClassArray) {
	return DialogClassArray(p.Dialogs)
}

// GetMessages returns value of Messages field.
func (p *MessagesPeerDialogs) GetMessages() (value []MessageClass) {
	return p.Messages
}

// MapMessages returns field Messages wrapped in MessageClassArray helper.
func (p *MessagesPeerDialogs) MapMessages() (value MessageClassArray) {
	return MessageClassArray(p.Messages)
}

// GetChats returns value of Chats field.
func (p *MessagesPeerDialogs) GetChats() (value []ChatClass) {
	return p.Chats
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (p *MessagesPeerDialogs) MapChats() (value ChatClassArray) {
	return ChatClassArray(p.Chats)
}

// GetUsers returns value of Users field.
func (p *MessagesPeerDialogs) GetUsers() (value []UserClass) {
	return p.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (p *MessagesPeerDialogs) MapUsers() (value UserClassArray) {
	return UserClassArray(p.Users)
}

// GetState returns value of State field.
func (p *MessagesPeerDialogs) GetState() (value UpdatesState) {
	return p.State
}

// Decode implements bin.Decoder.
func (p *MessagesPeerDialogs) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode messages.peerDialogs#3371c354 to nil")
	}
	if err := b.ConsumeID(MessagesPeerDialogsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *MessagesPeerDialogs) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode messages.peerDialogs#3371c354 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field dialogs: %w", err)
		}

		if headerLen > 0 {
			p.Dialogs = make([]DialogClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDialog(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field dialogs: %w", err)
			}
			p.Dialogs = append(p.Dialogs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field messages: %w", err)
		}

		if headerLen > 0 {
			p.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field messages: %w", err)
			}
			p.Messages = append(p.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field chats: %w", err)
		}

		if headerLen > 0 {
			p.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field chats: %w", err)
			}
			p.Chats = append(p.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field users: %w", err)
		}

		if headerLen > 0 {
			p.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	{
		if err := p.State.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field state: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesPeerDialogs.
var (
	_ bin.Encoder     = &MessagesPeerDialogs{}
	_ bin.Decoder     = &MessagesPeerDialogs{}
	_ bin.BareEncoder = &MessagesPeerDialogs{}
	_ bin.BareDecoder = &MessagesPeerDialogs{}
)
