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

// HelpInviteText represents TL type `help.inviteText#18cb9f78`.
// Text of a text message with an invitation to install Telegram.
//
// See https://core.telegram.org/constructor/help.inviteText for reference.
type HelpInviteText struct {
	// Text of the message
	Message string `tl:"message"`
}

// HelpInviteTextTypeID is TL type id of HelpInviteText.
const HelpInviteTextTypeID = 0x18cb9f78

func (i *HelpInviteText) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Message == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *HelpInviteText) String() string {
	if i == nil {
		return "HelpInviteText(nil)"
	}
	type Alias HelpInviteText
	return fmt.Sprintf("HelpInviteText%+v", Alias(*i))
}

// FillFrom fills HelpInviteText from given interface.
func (i *HelpInviteText) FillFrom(from interface {
	GetMessage() (value string)
}) {
	i.Message = from.GetMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpInviteText) TypeID() uint32 {
	return HelpInviteTextTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpInviteText) TypeName() string {
	return "help.inviteText"
}

// TypeInfo returns info about TL type.
func (i *HelpInviteText) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.inviteText",
		ID:   HelpInviteTextTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Message",
			SchemaName: "message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *HelpInviteText) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode help.inviteText#18cb9f78 as nil")
	}
	b.PutID(HelpInviteTextTypeID)
	b.PutString(i.Message)
	return nil
}

// GetMessage returns value of Message field.
func (i *HelpInviteText) GetMessage() (value string) {
	return i.Message
}

// Decode implements bin.Decoder.
func (i *HelpInviteText) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode help.inviteText#18cb9f78 to nil")
	}
	if err := b.ConsumeID(HelpInviteTextTypeID); err != nil {
		return fmt.Errorf("unable to decode help.inviteText#18cb9f78: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.inviteText#18cb9f78: field message: %w", err)
		}
		i.Message = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpInviteText.
var (
	_ bin.Encoder = &HelpInviteText{}
	_ bin.Decoder = &HelpInviteText{}
)
