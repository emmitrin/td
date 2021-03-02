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

// AccountDeleteAccountRequest represents TL type `account.deleteAccount#418d4e0b`.
// Delete the user's account from the telegram servers. Can be used, for example, to delete the account of a user that provided the login code, but forgot the 2FA password and no recovery method is configured¹.
//
// Links:
//  1) https://core.telegram.org/api/srp
//
// See https://core.telegram.org/method/account.deleteAccount for reference.
type AccountDeleteAccountRequest struct {
	// Why is the account being deleted, can be empty
	Reason string `tl:"reason"`
}

// AccountDeleteAccountRequestTypeID is TL type id of AccountDeleteAccountRequest.
const AccountDeleteAccountRequestTypeID = 0x418d4e0b

func (d *AccountDeleteAccountRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Reason == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *AccountDeleteAccountRequest) String() string {
	if d == nil {
		return "AccountDeleteAccountRequest(nil)"
	}
	type Alias AccountDeleteAccountRequest
	return fmt.Sprintf("AccountDeleteAccountRequest%+v", Alias(*d))
}

// FillFrom fills AccountDeleteAccountRequest from given interface.
func (d *AccountDeleteAccountRequest) FillFrom(from interface {
	GetReason() (value string)
}) {
	d.Reason = from.GetReason()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountDeleteAccountRequest) TypeID() uint32 {
	return AccountDeleteAccountRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountDeleteAccountRequest) TypeName() string {
	return "account.deleteAccount"
}

// TypeInfo returns info about TL type.
func (d *AccountDeleteAccountRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.deleteAccount",
		ID:   AccountDeleteAccountRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *AccountDeleteAccountRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode account.deleteAccount#418d4e0b as nil")
	}
	b.PutID(AccountDeleteAccountRequestTypeID)
	b.PutString(d.Reason)
	return nil
}

// GetReason returns value of Reason field.
func (d *AccountDeleteAccountRequest) GetReason() (value string) {
	return d.Reason
}

// Decode implements bin.Decoder.
func (d *AccountDeleteAccountRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode account.deleteAccount#418d4e0b to nil")
	}
	if err := b.ConsumeID(AccountDeleteAccountRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.deleteAccount#418d4e0b: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.deleteAccount#418d4e0b: field reason: %w", err)
		}
		d.Reason = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountDeleteAccountRequest.
var (
	_ bin.Encoder = &AccountDeleteAccountRequest{}
	_ bin.Decoder = &AccountDeleteAccountRequest{}
)

// AccountDeleteAccount invokes method account.deleteAccount#418d4e0b returning error if any.
// Delete the user's account from the telegram servers. Can be used, for example, to delete the account of a user that provided the login code, but forgot the 2FA password and no recovery method is configured¹.
//
// Links:
//  1) https://core.telegram.org/api/srp
//
// Possible errors:
//  420 2FA_CONFIRM_WAIT_X: Since this account is active and protected by a 2FA password, we will delete it in 1 week for security purposes. You can cancel this process at any time, you'll be able to reset your account in X seconds.
//
// See https://core.telegram.org/method/account.deleteAccount for reference.
func (c *Client) AccountDeleteAccount(ctx context.Context, reason string) (bool, error) {
	var result BoolBox

	request := &AccountDeleteAccountRequest{
		Reason: reason,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
