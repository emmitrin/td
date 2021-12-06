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

// EditMessageSchedulingStateRequest represents TL type `editMessageSchedulingState#ae2a0bc0`.
type EditMessageSchedulingStateRequest struct {
	// The chat the message belongs to
	ChatID int64
	// Identifier of the message
	MessageID int64
	// The new message scheduling state; pass null to send the message immediately
	SchedulingState MessageSchedulingStateClass
}

// EditMessageSchedulingStateRequestTypeID is TL type id of EditMessageSchedulingStateRequest.
const EditMessageSchedulingStateRequestTypeID = 0xae2a0bc0

// Ensuring interfaces in compile-time for EditMessageSchedulingStateRequest.
var (
	_ bin.Encoder     = &EditMessageSchedulingStateRequest{}
	_ bin.Decoder     = &EditMessageSchedulingStateRequest{}
	_ bin.BareEncoder = &EditMessageSchedulingStateRequest{}
	_ bin.BareDecoder = &EditMessageSchedulingStateRequest{}
)

func (e *EditMessageSchedulingStateRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.MessageID == 0) {
		return false
	}
	if !(e.SchedulingState == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditMessageSchedulingStateRequest) String() string {
	if e == nil {
		return "EditMessageSchedulingStateRequest(nil)"
	}
	type Alias EditMessageSchedulingStateRequest
	return fmt.Sprintf("EditMessageSchedulingStateRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditMessageSchedulingStateRequest) TypeID() uint32 {
	return EditMessageSchedulingStateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditMessageSchedulingStateRequest) TypeName() string {
	return "editMessageSchedulingState"
}

// TypeInfo returns info about TL type.
func (e *EditMessageSchedulingStateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editMessageSchedulingState",
		ID:   EditMessageSchedulingStateRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "SchedulingState",
			SchemaName: "scheduling_state",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditMessageSchedulingStateRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageSchedulingState#ae2a0bc0 as nil")
	}
	b.PutID(EditMessageSchedulingStateRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditMessageSchedulingStateRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageSchedulingState#ae2a0bc0 as nil")
	}
	b.PutInt53(e.ChatID)
	b.PutInt53(e.MessageID)
	if e.SchedulingState == nil {
		return fmt.Errorf("unable to encode editMessageSchedulingState#ae2a0bc0: field scheduling_state is nil")
	}
	if err := e.SchedulingState.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editMessageSchedulingState#ae2a0bc0: field scheduling_state: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditMessageSchedulingStateRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageSchedulingState#ae2a0bc0 to nil")
	}
	if err := b.ConsumeID(EditMessageSchedulingStateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editMessageSchedulingState#ae2a0bc0: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditMessageSchedulingStateRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageSchedulingState#ae2a0bc0 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageSchedulingState#ae2a0bc0: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageSchedulingState#ae2a0bc0: field message_id: %w", err)
		}
		e.MessageID = value
	}
	{
		value, err := DecodeMessageSchedulingState(b)
		if err != nil {
			return fmt.Errorf("unable to decode editMessageSchedulingState#ae2a0bc0: field scheduling_state: %w", err)
		}
		e.SchedulingState = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditMessageSchedulingStateRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageSchedulingState#ae2a0bc0 as nil")
	}
	b.ObjStart()
	b.PutID("editMessageSchedulingState")
	b.FieldStart("chat_id")
	b.PutInt53(e.ChatID)
	b.FieldStart("message_id")
	b.PutInt53(e.MessageID)
	b.FieldStart("scheduling_state")
	if e.SchedulingState == nil {
		return fmt.Errorf("unable to encode editMessageSchedulingState#ae2a0bc0: field scheduling_state is nil")
	}
	if err := e.SchedulingState.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editMessageSchedulingState#ae2a0bc0: field scheduling_state: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditMessageSchedulingStateRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageSchedulingState#ae2a0bc0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editMessageSchedulingState"); err != nil {
				return fmt.Errorf("unable to decode editMessageSchedulingState#ae2a0bc0: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageSchedulingState#ae2a0bc0: field chat_id: %w", err)
			}
			e.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageSchedulingState#ae2a0bc0: field message_id: %w", err)
			}
			e.MessageID = value
		case "scheduling_state":
			value, err := DecodeTDLibJSONMessageSchedulingState(b)
			if err != nil {
				return fmt.Errorf("unable to decode editMessageSchedulingState#ae2a0bc0: field scheduling_state: %w", err)
			}
			e.SchedulingState = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (e *EditMessageSchedulingStateRequest) GetChatID() (value int64) {
	return e.ChatID
}

// GetMessageID returns value of MessageID field.
func (e *EditMessageSchedulingStateRequest) GetMessageID() (value int64) {
	return e.MessageID
}

// GetSchedulingState returns value of SchedulingState field.
func (e *EditMessageSchedulingStateRequest) GetSchedulingState() (value MessageSchedulingStateClass) {
	return e.SchedulingState
}

// EditMessageSchedulingState invokes method editMessageSchedulingState#ae2a0bc0 returning error if any.
func (c *Client) EditMessageSchedulingState(ctx context.Context, request *EditMessageSchedulingStateRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
