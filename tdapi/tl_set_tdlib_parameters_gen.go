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

// SetTdlibParametersRequest represents TL type `setTdlibParameters#8e00ae53`.
type SetTdlibParametersRequest struct {
	// Parameters for TDLib initialization
	Parameters TdlibParameters
}

// SetTdlibParametersRequestTypeID is TL type id of SetTdlibParametersRequest.
const SetTdlibParametersRequestTypeID = 0x8e00ae53

// Ensuring interfaces in compile-time for SetTdlibParametersRequest.
var (
	_ bin.Encoder     = &SetTdlibParametersRequest{}
	_ bin.Decoder     = &SetTdlibParametersRequest{}
	_ bin.BareEncoder = &SetTdlibParametersRequest{}
	_ bin.BareDecoder = &SetTdlibParametersRequest{}
)

func (s *SetTdlibParametersRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Parameters.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetTdlibParametersRequest) String() string {
	if s == nil {
		return "SetTdlibParametersRequest(nil)"
	}
	type Alias SetTdlibParametersRequest
	return fmt.Sprintf("SetTdlibParametersRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetTdlibParametersRequest) TypeID() uint32 {
	return SetTdlibParametersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetTdlibParametersRequest) TypeName() string {
	return "setTdlibParameters"
}

// TypeInfo returns info about TL type.
func (s *SetTdlibParametersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setTdlibParameters",
		ID:   SetTdlibParametersRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Parameters",
			SchemaName: "parameters",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetTdlibParametersRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setTdlibParameters#8e00ae53 as nil")
	}
	b.PutID(SetTdlibParametersRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetTdlibParametersRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setTdlibParameters#8e00ae53 as nil")
	}
	if err := s.Parameters.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setTdlibParameters#8e00ae53: field parameters: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetTdlibParametersRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setTdlibParameters#8e00ae53 to nil")
	}
	if err := b.ConsumeID(SetTdlibParametersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setTdlibParameters#8e00ae53: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetTdlibParametersRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setTdlibParameters#8e00ae53 to nil")
	}
	{
		if err := s.Parameters.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setTdlibParameters#8e00ae53: field parameters: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetTdlibParametersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setTdlibParameters#8e00ae53 as nil")
	}
	b.ObjStart()
	b.PutID("setTdlibParameters")
	b.Comma()
	b.FieldStart("parameters")
	if err := s.Parameters.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setTdlibParameters#8e00ae53: field parameters: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetTdlibParametersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setTdlibParameters#8e00ae53 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setTdlibParameters"); err != nil {
				return fmt.Errorf("unable to decode setTdlibParameters#8e00ae53: %w", err)
			}
		case "parameters":
			if err := s.Parameters.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setTdlibParameters#8e00ae53: field parameters: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetParameters returns value of Parameters field.
func (s *SetTdlibParametersRequest) GetParameters() (value TdlibParameters) {
	if s == nil {
		return
	}
	return s.Parameters
}

// SetTdlibParameters invokes method setTdlibParameters#8e00ae53 returning error if any.
func (c *Client) SetTdlibParameters(ctx context.Context, parameters TdlibParameters) error {
	var ok Ok

	request := &SetTdlibParametersRequest{
		Parameters: parameters,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
