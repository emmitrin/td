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

// PaymentsSendPaymentFormRequest represents TL type `payments.sendPaymentForm#2b8879b3`.
// Send compiled payment form
//
// See https://core.telegram.org/method/payments.sendPaymentForm for reference.
type PaymentsSendPaymentFormRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Message ID of form
	MsgID int `tl:"msg_id"`
	// ID of saved and validated order info¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/payments.validatedRequestedInfo
	//
	// Use SetRequestedInfoID and GetRequestedInfoID helpers.
	RequestedInfoID string `tl:"requested_info_id"`
	// Chosen shipping option ID
	//
	// Use SetShippingOptionID and GetShippingOptionID helpers.
	ShippingOptionID string `tl:"shipping_option_id"`
	// Payment credentials
	Credentials InputPaymentCredentialsClass `tl:"credentials"`
}

// PaymentsSendPaymentFormRequestTypeID is TL type id of PaymentsSendPaymentFormRequest.
const PaymentsSendPaymentFormRequestTypeID = 0x2b8879b3

func (s *PaymentsSendPaymentFormRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.MsgID == 0) {
		return false
	}
	if !(s.RequestedInfoID == "") {
		return false
	}
	if !(s.ShippingOptionID == "") {
		return false
	}
	if !(s.Credentials == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PaymentsSendPaymentFormRequest) String() string {
	if s == nil {
		return "PaymentsSendPaymentFormRequest(nil)"
	}
	type Alias PaymentsSendPaymentFormRequest
	return fmt.Sprintf("PaymentsSendPaymentFormRequest%+v", Alias(*s))
}

// FillFrom fills PaymentsSendPaymentFormRequest from given interface.
func (s *PaymentsSendPaymentFormRequest) FillFrom(from interface {
	GetMsgID() (value int)
	GetRequestedInfoID() (value string, ok bool)
	GetShippingOptionID() (value string, ok bool)
	GetCredentials() (value InputPaymentCredentialsClass)
}) {
	s.MsgID = from.GetMsgID()
	if val, ok := from.GetRequestedInfoID(); ok {
		s.RequestedInfoID = val
	}

	if val, ok := from.GetShippingOptionID(); ok {
		s.ShippingOptionID = val
	}

	s.Credentials = from.GetCredentials()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsSendPaymentFormRequest) TypeID() uint32 {
	return PaymentsSendPaymentFormRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsSendPaymentFormRequest) TypeName() string {
	return "payments.sendPaymentForm"
}

// TypeInfo returns info about TL type.
func (s *PaymentsSendPaymentFormRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.sendPaymentForm",
		ID:   PaymentsSendPaymentFormRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "RequestedInfoID",
			SchemaName: "requested_info_id",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "ShippingOptionID",
			SchemaName: "shipping_option_id",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Credentials",
			SchemaName: "credentials",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *PaymentsSendPaymentFormRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.sendPaymentForm#2b8879b3 as nil")
	}
	b.PutID(PaymentsSendPaymentFormRequestTypeID)
	if !(s.RequestedInfoID == "") {
		s.Flags.Set(0)
	}
	if !(s.ShippingOptionID == "") {
		s.Flags.Set(1)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2b8879b3: field flags: %w", err)
	}
	b.PutInt(s.MsgID)
	if s.Flags.Has(0) {
		b.PutString(s.RequestedInfoID)
	}
	if s.Flags.Has(1) {
		b.PutString(s.ShippingOptionID)
	}
	if s.Credentials == nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2b8879b3: field credentials is nil")
	}
	if err := s.Credentials.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2b8879b3: field credentials: %w", err)
	}
	return nil
}

// GetMsgID returns value of MsgID field.
func (s *PaymentsSendPaymentFormRequest) GetMsgID() (value int) {
	return s.MsgID
}

// SetRequestedInfoID sets value of RequestedInfoID conditional field.
func (s *PaymentsSendPaymentFormRequest) SetRequestedInfoID(value string) {
	s.Flags.Set(0)
	s.RequestedInfoID = value
}

// GetRequestedInfoID returns value of RequestedInfoID conditional field and
// boolean which is true if field was set.
func (s *PaymentsSendPaymentFormRequest) GetRequestedInfoID() (value string, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.RequestedInfoID, true
}

// SetShippingOptionID sets value of ShippingOptionID conditional field.
func (s *PaymentsSendPaymentFormRequest) SetShippingOptionID(value string) {
	s.Flags.Set(1)
	s.ShippingOptionID = value
}

// GetShippingOptionID returns value of ShippingOptionID conditional field and
// boolean which is true if field was set.
func (s *PaymentsSendPaymentFormRequest) GetShippingOptionID() (value string, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.ShippingOptionID, true
}

// GetCredentials returns value of Credentials field.
func (s *PaymentsSendPaymentFormRequest) GetCredentials() (value InputPaymentCredentialsClass) {
	return s.Credentials
}

// Decode implements bin.Decoder.
func (s *PaymentsSendPaymentFormRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.sendPaymentForm#2b8879b3 to nil")
	}
	if err := b.ConsumeID(PaymentsSendPaymentFormRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: field msg_id: %w", err)
		}
		s.MsgID = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: field requested_info_id: %w", err)
		}
		s.RequestedInfoID = value
	}
	if s.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: field shipping_option_id: %w", err)
		}
		s.ShippingOptionID = value
	}
	{
		value, err := DecodeInputPaymentCredentials(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2b8879b3: field credentials: %w", err)
		}
		s.Credentials = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsSendPaymentFormRequest.
var (
	_ bin.Encoder = &PaymentsSendPaymentFormRequest{}
	_ bin.Decoder = &PaymentsSendPaymentFormRequest{}
)

// PaymentsSendPaymentForm invokes method payments.sendPaymentForm#2b8879b3 returning error if any.
// Send compiled payment form
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//
// See https://core.telegram.org/method/payments.sendPaymentForm for reference.
func (c *Client) PaymentsSendPaymentForm(ctx context.Context, request *PaymentsSendPaymentFormRequest) (PaymentsPaymentResultClass, error) {
	var result PaymentsPaymentResultBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.PaymentResult, nil
}
