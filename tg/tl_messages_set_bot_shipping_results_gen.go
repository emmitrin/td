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

// MessagesSetBotShippingResultsRequest represents TL type `messages.setBotShippingResults#e5f672fa`.
// If you sent an invoice requesting a shipping address and the parameter is_flexible was
// specified, the bot will receive an updateBotShippingQuery¹ update. Use this method to
// reply to shipping queries.
//
// Links:
//  1) https://core.telegram.org/constructor/updateBotShippingQuery
//
// See https://core.telegram.org/method/messages.setBotShippingResults for reference.
type MessagesSetBotShippingResultsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Unique identifier for the query to be answered
	QueryID int64
	// Error message in human readable form that explains why it is impossible to complete
	// the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram
	// will display this message to the user.
	//
	// Use SetError and GetError helpers.
	Error string
	// A vector of available shipping options.
	//
	// Use SetShippingOptions and GetShippingOptions helpers.
	ShippingOptions []ShippingOption
}

// MessagesSetBotShippingResultsRequestTypeID is TL type id of MessagesSetBotShippingResultsRequest.
const MessagesSetBotShippingResultsRequestTypeID = 0xe5f672fa

func (s *MessagesSetBotShippingResultsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.QueryID == 0) {
		return false
	}
	if !(s.Error == "") {
		return false
	}
	if !(s.ShippingOptions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetBotShippingResultsRequest) String() string {
	if s == nil {
		return "MessagesSetBotShippingResultsRequest(nil)"
	}
	type Alias MessagesSetBotShippingResultsRequest
	return fmt.Sprintf("MessagesSetBotShippingResultsRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetBotShippingResultsRequest from given interface.
func (s *MessagesSetBotShippingResultsRequest) FillFrom(from interface {
	GetQueryID() (value int64)
	GetError() (value string, ok bool)
	GetShippingOptions() (value []ShippingOption, ok bool)
}) {
	s.QueryID = from.GetQueryID()
	if val, ok := from.GetError(); ok {
		s.Error = val
	}

	if val, ok := from.GetShippingOptions(); ok {
		s.ShippingOptions = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetBotShippingResultsRequest) TypeID() uint32 {
	return MessagesSetBotShippingResultsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetBotShippingResultsRequest) TypeName() string {
	return "messages.setBotShippingResults"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetBotShippingResultsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setBotShippingResults",
		ID:   MessagesSetBotShippingResultsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "Error",
			SchemaName: "error",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "ShippingOptions",
			SchemaName: "shipping_options",
			Null:       !s.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetBotShippingResultsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setBotShippingResults#e5f672fa as nil")
	}
	b.PutID(MessagesSetBotShippingResultsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetBotShippingResultsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setBotShippingResults#e5f672fa as nil")
	}
	if !(s.Error == "") {
		s.Flags.Set(0)
	}
	if !(s.ShippingOptions == nil) {
		s.Flags.Set(1)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setBotShippingResults#e5f672fa: field flags: %w", err)
	}
	b.PutLong(s.QueryID)
	if s.Flags.Has(0) {
		b.PutString(s.Error)
	}
	if s.Flags.Has(1) {
		b.PutVectorHeader(len(s.ShippingOptions))
		for idx, v := range s.ShippingOptions {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.setBotShippingResults#e5f672fa: field shipping_options element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// GetQueryID returns value of QueryID field.
func (s *MessagesSetBotShippingResultsRequest) GetQueryID() (value int64) {
	return s.QueryID
}

// SetError sets value of Error conditional field.
func (s *MessagesSetBotShippingResultsRequest) SetError(value string) {
	s.Flags.Set(0)
	s.Error = value
}

// GetError returns value of Error conditional field and
// boolean which is true if field was set.
func (s *MessagesSetBotShippingResultsRequest) GetError() (value string, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Error, true
}

// SetShippingOptions sets value of ShippingOptions conditional field.
func (s *MessagesSetBotShippingResultsRequest) SetShippingOptions(value []ShippingOption) {
	s.Flags.Set(1)
	s.ShippingOptions = value
}

// GetShippingOptions returns value of ShippingOptions conditional field and
// boolean which is true if field was set.
func (s *MessagesSetBotShippingResultsRequest) GetShippingOptions() (value []ShippingOption, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.ShippingOptions, true
}

// Decode implements bin.Decoder.
func (s *MessagesSetBotShippingResultsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setBotShippingResults#e5f672fa to nil")
	}
	if err := b.ConsumeID(MessagesSetBotShippingResultsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setBotShippingResults#e5f672fa: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetBotShippingResultsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setBotShippingResults#e5f672fa to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setBotShippingResults#e5f672fa: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotShippingResults#e5f672fa: field query_id: %w", err)
		}
		s.QueryID = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotShippingResults#e5f672fa: field error: %w", err)
		}
		s.Error = value
	}
	if s.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotShippingResults#e5f672fa: field shipping_options: %w", err)
		}

		if headerLen > 0 {
			s.ShippingOptions = make([]ShippingOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ShippingOption
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.setBotShippingResults#e5f672fa: field shipping_options: %w", err)
			}
			s.ShippingOptions = append(s.ShippingOptions, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetBotShippingResultsRequest.
var (
	_ bin.Encoder     = &MessagesSetBotShippingResultsRequest{}
	_ bin.Decoder     = &MessagesSetBotShippingResultsRequest{}
	_ bin.BareEncoder = &MessagesSetBotShippingResultsRequest{}
	_ bin.BareDecoder = &MessagesSetBotShippingResultsRequest{}
)

// MessagesSetBotShippingResults invokes method messages.setBotShippingResults#e5f672fa returning error if any.
// If you sent an invoice requesting a shipping address and the parameter is_flexible was
// specified, the bot will receive an updateBotShippingQuery¹ update. Use this method to
// reply to shipping queries.
//
// Links:
//  1) https://core.telegram.org/constructor/updateBotShippingQuery
//
// Possible errors:
//  400 QUERY_ID_INVALID: The query ID is invalid
//
// See https://core.telegram.org/method/messages.setBotShippingResults for reference.
// Can be used by bots.
func (c *Client) MessagesSetBotShippingResults(ctx context.Context, request *MessagesSetBotShippingResultsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
