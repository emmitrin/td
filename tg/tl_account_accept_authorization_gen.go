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

// AccountAcceptAuthorizationRequest represents TL type `account.acceptAuthorization#e7027c94`.
// Sends a Telegram Passport authorization form, effectively sharing data with the service
//
// See https://core.telegram.org/method/account.acceptAuthorization for reference.
type AccountAcceptAuthorizationRequest struct {
	// Bot ID
	BotID int `tl:"bot_id"`
	// Telegram Passport element types requested by the service
	Scope string `tl:"scope"`
	// Service's public key
	PublicKey string `tl:"public_key"`
	// Types of values sent and their hashes
	ValueHashes []SecureValueHash `tl:"value_hashes"`
	// Encrypted values
	Credentials SecureCredentialsEncrypted `tl:"credentials"`
}

// AccountAcceptAuthorizationRequestTypeID is TL type id of AccountAcceptAuthorizationRequest.
const AccountAcceptAuthorizationRequestTypeID = 0xe7027c94

func (a *AccountAcceptAuthorizationRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.BotID == 0) {
		return false
	}
	if !(a.Scope == "") {
		return false
	}
	if !(a.PublicKey == "") {
		return false
	}
	if !(a.ValueHashes == nil) {
		return false
	}
	if !(a.Credentials.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AccountAcceptAuthorizationRequest) String() string {
	if a == nil {
		return "AccountAcceptAuthorizationRequest(nil)"
	}
	type Alias AccountAcceptAuthorizationRequest
	return fmt.Sprintf("AccountAcceptAuthorizationRequest%+v", Alias(*a))
}

// FillFrom fills AccountAcceptAuthorizationRequest from given interface.
func (a *AccountAcceptAuthorizationRequest) FillFrom(from interface {
	GetBotID() (value int)
	GetScope() (value string)
	GetPublicKey() (value string)
	GetValueHashes() (value []SecureValueHash)
	GetCredentials() (value SecureCredentialsEncrypted)
}) {
	a.BotID = from.GetBotID()
	a.Scope = from.GetScope()
	a.PublicKey = from.GetPublicKey()
	a.ValueHashes = from.GetValueHashes()
	a.Credentials = from.GetCredentials()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountAcceptAuthorizationRequest) TypeID() uint32 {
	return AccountAcceptAuthorizationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountAcceptAuthorizationRequest) TypeName() string {
	return "account.acceptAuthorization"
}

// TypeInfo returns info about TL type.
func (a *AccountAcceptAuthorizationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.acceptAuthorization",
		ID:   AccountAcceptAuthorizationRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotID",
			SchemaName: "bot_id",
		},
		{
			Name:       "Scope",
			SchemaName: "scope",
		},
		{
			Name:       "PublicKey",
			SchemaName: "public_key",
		},
		{
			Name:       "ValueHashes",
			SchemaName: "value_hashes",
		},
		{
			Name:       "Credentials",
			SchemaName: "credentials",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AccountAcceptAuthorizationRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode account.acceptAuthorization#e7027c94 as nil")
	}
	b.PutID(AccountAcceptAuthorizationRequestTypeID)
	b.PutInt(a.BotID)
	b.PutString(a.Scope)
	b.PutString(a.PublicKey)
	b.PutVectorHeader(len(a.ValueHashes))
	for idx, v := range a.ValueHashes {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.acceptAuthorization#e7027c94: field value_hashes element with index %d: %w", idx, err)
		}
	}
	if err := a.Credentials.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.acceptAuthorization#e7027c94: field credentials: %w", err)
	}
	return nil
}

// GetBotID returns value of BotID field.
func (a *AccountAcceptAuthorizationRequest) GetBotID() (value int) {
	return a.BotID
}

// GetScope returns value of Scope field.
func (a *AccountAcceptAuthorizationRequest) GetScope() (value string) {
	return a.Scope
}

// GetPublicKey returns value of PublicKey field.
func (a *AccountAcceptAuthorizationRequest) GetPublicKey() (value string) {
	return a.PublicKey
}

// GetValueHashes returns value of ValueHashes field.
func (a *AccountAcceptAuthorizationRequest) GetValueHashes() (value []SecureValueHash) {
	return a.ValueHashes
}

// GetCredentials returns value of Credentials field.
func (a *AccountAcceptAuthorizationRequest) GetCredentials() (value SecureCredentialsEncrypted) {
	return a.Credentials
}

// Decode implements bin.Decoder.
func (a *AccountAcceptAuthorizationRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode account.acceptAuthorization#e7027c94 to nil")
	}
	if err := b.ConsumeID(AccountAcceptAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field bot_id: %w", err)
		}
		a.BotID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field scope: %w", err)
		}
		a.Scope = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field public_key: %w", err)
		}
		a.PublicKey = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field value_hashes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SecureValueHash
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field value_hashes: %w", err)
			}
			a.ValueHashes = append(a.ValueHashes, value)
		}
	}
	{
		if err := a.Credentials.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field credentials: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountAcceptAuthorizationRequest.
var (
	_ bin.Encoder = &AccountAcceptAuthorizationRequest{}
	_ bin.Decoder = &AccountAcceptAuthorizationRequest{}
)

// AccountAcceptAuthorization invokes method account.acceptAuthorization#e7027c94 returning error if any.
// Sends a Telegram Passport authorization form, effectively sharing data with the service
//
// See https://core.telegram.org/method/account.acceptAuthorization for reference.
func (c *Client) AccountAcceptAuthorization(ctx context.Context, request *AccountAcceptAuthorizationRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
