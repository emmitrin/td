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

// AccountUpdateProfileRequest represents TL type `account.updateProfile#78515775`.
// Updates user profile.
//
// See https://core.telegram.org/method/account.updateProfile for reference.
type AccountUpdateProfileRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// New user first name
	//
	// Use SetFirstName and GetFirstName helpers.
	FirstName string `tl:"first_name"`
	// New user last name
	//
	// Use SetLastName and GetLastName helpers.
	LastName string `tl:"last_name"`
	// New bio
	//
	// Use SetAbout and GetAbout helpers.
	About string `tl:"about"`
}

// AccountUpdateProfileRequestTypeID is TL type id of AccountUpdateProfileRequest.
const AccountUpdateProfileRequestTypeID = 0x78515775

func (u *AccountUpdateProfileRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.FirstName == "") {
		return false
	}
	if !(u.LastName == "") {
		return false
	}
	if !(u.About == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateProfileRequest) String() string {
	if u == nil {
		return "AccountUpdateProfileRequest(nil)"
	}
	type Alias AccountUpdateProfileRequest
	return fmt.Sprintf("AccountUpdateProfileRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateProfileRequest from given interface.
func (u *AccountUpdateProfileRequest) FillFrom(from interface {
	GetFirstName() (value string, ok bool)
	GetLastName() (value string, ok bool)
	GetAbout() (value string, ok bool)
}) {
	if val, ok := from.GetFirstName(); ok {
		u.FirstName = val
	}

	if val, ok := from.GetLastName(); ok {
		u.LastName = val
	}

	if val, ok := from.GetAbout(); ok {
		u.About = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateProfileRequest) TypeID() uint32 {
	return AccountUpdateProfileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateProfileRequest) TypeName() string {
	return "account.updateProfile"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateProfileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateProfile",
		ID:   AccountUpdateProfileRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "FirstName",
			SchemaName: "first_name",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "LastName",
			SchemaName: "last_name",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "About",
			SchemaName: "about",
			Null:       !u.Flags.Has(2),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUpdateProfileRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateProfile#78515775 as nil")
	}
	b.PutID(AccountUpdateProfileRequestTypeID)
	if !(u.FirstName == "") {
		u.Flags.Set(0)
	}
	if !(u.LastName == "") {
		u.Flags.Set(1)
	}
	if !(u.About == "") {
		u.Flags.Set(2)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateProfile#78515775: field flags: %w", err)
	}
	if u.Flags.Has(0) {
		b.PutString(u.FirstName)
	}
	if u.Flags.Has(1) {
		b.PutString(u.LastName)
	}
	if u.Flags.Has(2) {
		b.PutString(u.About)
	}
	return nil
}

// SetFirstName sets value of FirstName conditional field.
func (u *AccountUpdateProfileRequest) SetFirstName(value string) {
	u.Flags.Set(0)
	u.FirstName = value
}

// GetFirstName returns value of FirstName conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateProfileRequest) GetFirstName() (value string, ok bool) {
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.FirstName, true
}

// SetLastName sets value of LastName conditional field.
func (u *AccountUpdateProfileRequest) SetLastName(value string) {
	u.Flags.Set(1)
	u.LastName = value
}

// GetLastName returns value of LastName conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateProfileRequest) GetLastName() (value string, ok bool) {
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.LastName, true
}

// SetAbout sets value of About conditional field.
func (u *AccountUpdateProfileRequest) SetAbout(value string) {
	u.Flags.Set(2)
	u.About = value
}

// GetAbout returns value of About conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateProfileRequest) GetAbout() (value string, ok bool) {
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.About, true
}

// Decode implements bin.Decoder.
func (u *AccountUpdateProfileRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateProfile#78515775 to nil")
	}
	if err := b.ConsumeID(AccountUpdateProfileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateProfile#78515775: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateProfile#78515775: field flags: %w", err)
		}
	}
	if u.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateProfile#78515775: field first_name: %w", err)
		}
		u.FirstName = value
	}
	if u.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateProfile#78515775: field last_name: %w", err)
		}
		u.LastName = value
	}
	if u.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateProfile#78515775: field about: %w", err)
		}
		u.About = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUpdateProfileRequest.
var (
	_ bin.Encoder = &AccountUpdateProfileRequest{}
	_ bin.Decoder = &AccountUpdateProfileRequest{}
)

// AccountUpdateProfile invokes method account.updateProfile#78515775 returning error if any.
// Updates user profile.
//
// Possible errors:
//  400 ABOUT_TOO_LONG: About string too long
//  400 FIRSTNAME_INVALID: The first name is invalid
//
// See https://core.telegram.org/method/account.updateProfile for reference.
func (c *Client) AccountUpdateProfile(ctx context.Context, request *AccountUpdateProfileRequest) (UserClass, error) {
	var result UserBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.User, nil
}
