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

// ContactsBlockFromRepliesRequest represents TL type `contacts.blockFromReplies#29a8962c`.
// Stop getting notifications about thread replies¹ of a certain user in @replies
//
// Links:
//  1) https://core.telegram.org/api/threads
//
// See https://core.telegram.org/method/contacts.blockFromReplies for reference.
type ContactsBlockFromRepliesRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Whether to delete the specified message as well
	DeleteMessage bool `tl:"delete_message"`
	// Whether to delete all @replies messages from this user as well
	DeleteHistory bool `tl:"delete_history"`
	// Whether to also report this user for spam
	ReportSpam bool `tl:"report_spam"`
	// ID of the message in the @replies¹ chat
	//
	// Links:
	//  1) https://core.telegram.org/api/threads#replies
	MsgID int `tl:"msg_id"`
}

// ContactsBlockFromRepliesRequestTypeID is TL type id of ContactsBlockFromRepliesRequest.
const ContactsBlockFromRepliesRequestTypeID = 0x29a8962c

func (b *ContactsBlockFromRepliesRequest) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.DeleteMessage == false) {
		return false
	}
	if !(b.DeleteHistory == false) {
		return false
	}
	if !(b.ReportSpam == false) {
		return false
	}
	if !(b.MsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *ContactsBlockFromRepliesRequest) String() string {
	if b == nil {
		return "ContactsBlockFromRepliesRequest(nil)"
	}
	type Alias ContactsBlockFromRepliesRequest
	return fmt.Sprintf("ContactsBlockFromRepliesRequest%+v", Alias(*b))
}

// FillFrom fills ContactsBlockFromRepliesRequest from given interface.
func (b *ContactsBlockFromRepliesRequest) FillFrom(from interface {
	GetDeleteMessage() (value bool)
	GetDeleteHistory() (value bool)
	GetReportSpam() (value bool)
	GetMsgID() (value int)
}) {
	b.DeleteMessage = from.GetDeleteMessage()
	b.DeleteHistory = from.GetDeleteHistory()
	b.ReportSpam = from.GetReportSpam()
	b.MsgID = from.GetMsgID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsBlockFromRepliesRequest) TypeID() uint32 {
	return ContactsBlockFromRepliesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsBlockFromRepliesRequest) TypeName() string {
	return "contacts.blockFromReplies"
}

// TypeInfo returns info about TL type.
func (b *ContactsBlockFromRepliesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.blockFromReplies",
		ID:   ContactsBlockFromRepliesRequestTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "DeleteMessage",
			SchemaName: "delete_message",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "DeleteHistory",
			SchemaName: "delete_history",
			Null:       !b.Flags.Has(1),
		},
		{
			Name:       "ReportSpam",
			SchemaName: "report_spam",
			Null:       !b.Flags.Has(2),
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *ContactsBlockFromRepliesRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode contacts.blockFromReplies#29a8962c as nil")
	}
	buf.PutID(ContactsBlockFromRepliesRequestTypeID)
	if !(b.DeleteMessage == false) {
		b.Flags.Set(0)
	}
	if !(b.DeleteHistory == false) {
		b.Flags.Set(1)
	}
	if !(b.ReportSpam == false) {
		b.Flags.Set(2)
	}
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode contacts.blockFromReplies#29a8962c: field flags: %w", err)
	}
	buf.PutInt(b.MsgID)
	return nil
}

// SetDeleteMessage sets value of DeleteMessage conditional field.
func (b *ContactsBlockFromRepliesRequest) SetDeleteMessage(value bool) {
	if value {
		b.Flags.Set(0)
		b.DeleteMessage = true
	} else {
		b.Flags.Unset(0)
		b.DeleteMessage = false
	}
}

// GetDeleteMessage returns value of DeleteMessage conditional field.
func (b *ContactsBlockFromRepliesRequest) GetDeleteMessage() (value bool) {
	return b.Flags.Has(0)
}

// SetDeleteHistory sets value of DeleteHistory conditional field.
func (b *ContactsBlockFromRepliesRequest) SetDeleteHistory(value bool) {
	if value {
		b.Flags.Set(1)
		b.DeleteHistory = true
	} else {
		b.Flags.Unset(1)
		b.DeleteHistory = false
	}
}

// GetDeleteHistory returns value of DeleteHistory conditional field.
func (b *ContactsBlockFromRepliesRequest) GetDeleteHistory() (value bool) {
	return b.Flags.Has(1)
}

// SetReportSpam sets value of ReportSpam conditional field.
func (b *ContactsBlockFromRepliesRequest) SetReportSpam(value bool) {
	if value {
		b.Flags.Set(2)
		b.ReportSpam = true
	} else {
		b.Flags.Unset(2)
		b.ReportSpam = false
	}
}

// GetReportSpam returns value of ReportSpam conditional field.
func (b *ContactsBlockFromRepliesRequest) GetReportSpam() (value bool) {
	return b.Flags.Has(2)
}

// GetMsgID returns value of MsgID field.
func (b *ContactsBlockFromRepliesRequest) GetMsgID() (value int) {
	return b.MsgID
}

// Decode implements bin.Decoder.
func (b *ContactsBlockFromRepliesRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode contacts.blockFromReplies#29a8962c to nil")
	}
	if err := buf.ConsumeID(ContactsBlockFromRepliesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.blockFromReplies#29a8962c: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode contacts.blockFromReplies#29a8962c: field flags: %w", err)
		}
	}
	b.DeleteMessage = b.Flags.Has(0)
	b.DeleteHistory = b.Flags.Has(1)
	b.ReportSpam = b.Flags.Has(2)
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.blockFromReplies#29a8962c: field msg_id: %w", err)
		}
		b.MsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsBlockFromRepliesRequest.
var (
	_ bin.Encoder = &ContactsBlockFromRepliesRequest{}
	_ bin.Decoder = &ContactsBlockFromRepliesRequest{}
)

// ContactsBlockFromReplies invokes method contacts.blockFromReplies#29a8962c returning error if any.
// Stop getting notifications about thread replies¹ of a certain user in @replies
//
// Links:
//  1) https://core.telegram.org/api/threads
//
// See https://core.telegram.org/method/contacts.blockFromReplies for reference.
// Can be used by bots.
func (c *Client) ContactsBlockFromReplies(ctx context.Context, request *ContactsBlockFromRepliesRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
