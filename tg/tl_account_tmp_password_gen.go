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

// AccountTmpPassword represents TL type `account.tmpPassword#db64fd34`.
// Temporary payment password
//
// See https://core.telegram.org/constructor/account.tmpPassword for reference.
type AccountTmpPassword struct {
	// Temporary password
	TmpPassword []byte `tl:"tmp_password"`
	// Validity period
	ValidUntil int `tl:"valid_until"`
}

// AccountTmpPasswordTypeID is TL type id of AccountTmpPassword.
const AccountTmpPasswordTypeID = 0xdb64fd34

func (t *AccountTmpPassword) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.TmpPassword == nil) {
		return false
	}
	if !(t.ValidUntil == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountTmpPassword) String() string {
	if t == nil {
		return "AccountTmpPassword(nil)"
	}
	type Alias AccountTmpPassword
	return fmt.Sprintf("AccountTmpPassword%+v", Alias(*t))
}

// FillFrom fills AccountTmpPassword from given interface.
func (t *AccountTmpPassword) FillFrom(from interface {
	GetTmpPassword() (value []byte)
	GetValidUntil() (value int)
}) {
	t.TmpPassword = from.GetTmpPassword()
	t.ValidUntil = from.GetValidUntil()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountTmpPassword) TypeID() uint32 {
	return AccountTmpPasswordTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountTmpPassword) TypeName() string {
	return "account.tmpPassword"
}

// TypeInfo returns info about TL type.
func (t *AccountTmpPassword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.tmpPassword",
		ID:   AccountTmpPasswordTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TmpPassword",
			SchemaName: "tmp_password",
		},
		{
			Name:       "ValidUntil",
			SchemaName: "valid_until",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *AccountTmpPassword) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.tmpPassword#db64fd34 as nil")
	}
	b.PutID(AccountTmpPasswordTypeID)
	b.PutBytes(t.TmpPassword)
	b.PutInt(t.ValidUntil)
	return nil
}

// GetTmpPassword returns value of TmpPassword field.
func (t *AccountTmpPassword) GetTmpPassword() (value []byte) {
	return t.TmpPassword
}

// GetValidUntil returns value of ValidUntil field.
func (t *AccountTmpPassword) GetValidUntil() (value int) {
	return t.ValidUntil
}

// Decode implements bin.Decoder.
func (t *AccountTmpPassword) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.tmpPassword#db64fd34 to nil")
	}
	if err := b.ConsumeID(AccountTmpPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode account.tmpPassword#db64fd34: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode account.tmpPassword#db64fd34: field tmp_password: %w", err)
		}
		t.TmpPassword = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.tmpPassword#db64fd34: field valid_until: %w", err)
		}
		t.ValidUntil = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountTmpPassword.
var (
	_ bin.Encoder = &AccountTmpPassword{}
	_ bin.Decoder = &AccountTmpPassword{}
)
