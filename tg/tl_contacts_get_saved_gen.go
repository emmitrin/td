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

// ContactsGetSavedRequest represents TL type `contacts.getSaved#82f1e39f`.
// Get all contacts
//
// See https://core.telegram.org/method/contacts.getSaved for reference.
type ContactsGetSavedRequest struct {
}

// ContactsGetSavedRequestTypeID is TL type id of ContactsGetSavedRequest.
const ContactsGetSavedRequestTypeID = 0x82f1e39f

func (g *ContactsGetSavedRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *ContactsGetSavedRequest) String() string {
	if g == nil {
		return "ContactsGetSavedRequest(nil)"
	}
	type Alias ContactsGetSavedRequest
	return fmt.Sprintf("ContactsGetSavedRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsGetSavedRequest) TypeID() uint32 {
	return ContactsGetSavedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsGetSavedRequest) TypeName() string {
	return "contacts.getSaved"
}

// TypeInfo returns info about TL type.
func (g *ContactsGetSavedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.getSaved",
		ID:   ContactsGetSavedRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *ContactsGetSavedRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getSaved#82f1e39f as nil")
	}
	b.PutID(ContactsGetSavedRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetSavedRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getSaved#82f1e39f to nil")
	}
	if err := b.ConsumeID(ContactsGetSavedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getSaved#82f1e39f: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetSavedRequest.
var (
	_ bin.Encoder = &ContactsGetSavedRequest{}
	_ bin.Decoder = &ContactsGetSavedRequest{}
)

// ContactsGetSaved invokes method contacts.getSaved#82f1e39f returning error if any.
// Get all contacts
//
// Possible errors:
//  403 TAKEOUT_REQUIRED: A takeout session has to be initialized, first
//
// See https://core.telegram.org/method/contacts.getSaved for reference.
func (c *Client) ContactsGetSaved(ctx context.Context) ([]SavedPhoneContact, error) {
	var result SavedPhoneContactVector

	request := &ContactsGetSavedRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return []SavedPhoneContact(result.Elems), nil
}
