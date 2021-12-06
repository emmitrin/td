// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// RequestQrCodeAuthenticationRequest represents TL type `requestQrCodeAuthentication#56fe3c4e`.
type RequestQrCodeAuthenticationRequest struct {
	// List of user identifiers of other users currently using the application
	OtherUserIDs []int64
}

// RequestQrCodeAuthenticationRequestTypeID is TL type id of RequestQrCodeAuthenticationRequest.
const RequestQrCodeAuthenticationRequestTypeID = 0x56fe3c4e

// Ensuring interfaces in compile-time for RequestQrCodeAuthenticationRequest.
var (
	_ bin.Encoder     = &RequestQrCodeAuthenticationRequest{}
	_ bin.Decoder     = &RequestQrCodeAuthenticationRequest{}
	_ bin.BareEncoder = &RequestQrCodeAuthenticationRequest{}
	_ bin.BareDecoder = &RequestQrCodeAuthenticationRequest{}
)

func (r *RequestQrCodeAuthenticationRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.OtherUserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RequestQrCodeAuthenticationRequest) String() string {
	if r == nil {
		return "RequestQrCodeAuthenticationRequest(nil)"
	}
	type Alias RequestQrCodeAuthenticationRequest
	return fmt.Sprintf("RequestQrCodeAuthenticationRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RequestQrCodeAuthenticationRequest) TypeID() uint32 {
	return RequestQrCodeAuthenticationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RequestQrCodeAuthenticationRequest) TypeName() string {
	return "requestQrCodeAuthentication"
}

// TypeInfo returns info about TL type.
func (r *RequestQrCodeAuthenticationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "requestQrCodeAuthentication",
		ID:   RequestQrCodeAuthenticationRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "OtherUserIDs",
			SchemaName: "other_user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RequestQrCodeAuthenticationRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode requestQrCodeAuthentication#56fe3c4e as nil")
	}
	b.PutID(RequestQrCodeAuthenticationRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RequestQrCodeAuthenticationRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode requestQrCodeAuthentication#56fe3c4e as nil")
	}
	b.PutInt(len(r.OtherUserIDs))
	for _, v := range r.OtherUserIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RequestQrCodeAuthenticationRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode requestQrCodeAuthentication#56fe3c4e to nil")
	}
	if err := b.ConsumeID(RequestQrCodeAuthenticationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode requestQrCodeAuthentication#56fe3c4e: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RequestQrCodeAuthenticationRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode requestQrCodeAuthentication#56fe3c4e to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode requestQrCodeAuthentication#56fe3c4e: field other_user_ids: %w", err)
		}

		if headerLen > 0 {
			r.OtherUserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode requestQrCodeAuthentication#56fe3c4e: field other_user_ids: %w", err)
			}
			r.OtherUserIDs = append(r.OtherUserIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RequestQrCodeAuthenticationRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode requestQrCodeAuthentication#56fe3c4e as nil")
	}
	b.ObjStart()
	b.PutID("requestQrCodeAuthentication")
	b.FieldStart("other_user_ids")
	b.ArrStart()
	for _, v := range r.OtherUserIDs {
		b.PutInt53(v)
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RequestQrCodeAuthenticationRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode requestQrCodeAuthentication#56fe3c4e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("requestQrCodeAuthentication"); err != nil {
				return fmt.Errorf("unable to decode requestQrCodeAuthentication#56fe3c4e: %w", err)
			}
		case "other_user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode requestQrCodeAuthentication#56fe3c4e: field other_user_ids: %w", err)
				}
				r.OtherUserIDs = append(r.OtherUserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode requestQrCodeAuthentication#56fe3c4e: field other_user_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetOtherUserIDs returns value of OtherUserIDs field.
func (r *RequestQrCodeAuthenticationRequest) GetOtherUserIDs() (value []int64) {
	return r.OtherUserIDs
}

// RequestQrCodeAuthentication invokes method requestQrCodeAuthentication#56fe3c4e returning error if any.
func (c *Client) RequestQrCodeAuthentication(ctx context.Context, otheruserids []int64) error {
	var ok Ok

	request := &RequestQrCodeAuthenticationRequest{
		OtherUserIDs: otheruserids,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
