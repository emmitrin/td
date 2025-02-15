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

// CheckChatUsernameResultOk represents TL type `checkChatUsernameResultOk#a6a7bb5c`.
type CheckChatUsernameResultOk struct {
}

// CheckChatUsernameResultOkTypeID is TL type id of CheckChatUsernameResultOk.
const CheckChatUsernameResultOkTypeID = 0xa6a7bb5c

// construct implements constructor of CheckChatUsernameResultClass.
func (c CheckChatUsernameResultOk) construct() CheckChatUsernameResultClass { return &c }

// Ensuring interfaces in compile-time for CheckChatUsernameResultOk.
var (
	_ bin.Encoder     = &CheckChatUsernameResultOk{}
	_ bin.Decoder     = &CheckChatUsernameResultOk{}
	_ bin.BareEncoder = &CheckChatUsernameResultOk{}
	_ bin.BareDecoder = &CheckChatUsernameResultOk{}

	_ CheckChatUsernameResultClass = &CheckChatUsernameResultOk{}
)

func (c *CheckChatUsernameResultOk) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckChatUsernameResultOk) String() string {
	if c == nil {
		return "CheckChatUsernameResultOk(nil)"
	}
	type Alias CheckChatUsernameResultOk
	return fmt.Sprintf("CheckChatUsernameResultOk%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckChatUsernameResultOk) TypeID() uint32 {
	return CheckChatUsernameResultOkTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckChatUsernameResultOk) TypeName() string {
	return "checkChatUsernameResultOk"
}

// TypeInfo returns info about TL type.
func (c *CheckChatUsernameResultOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkChatUsernameResultOk",
		ID:   CheckChatUsernameResultOkTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckChatUsernameResultOk) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultOk#a6a7bb5c as nil")
	}
	b.PutID(CheckChatUsernameResultOkTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckChatUsernameResultOk) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultOk#a6a7bb5c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckChatUsernameResultOk) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultOk#a6a7bb5c to nil")
	}
	if err := b.ConsumeID(CheckChatUsernameResultOkTypeID); err != nil {
		return fmt.Errorf("unable to decode checkChatUsernameResultOk#a6a7bb5c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckChatUsernameResultOk) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultOk#a6a7bb5c to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckChatUsernameResultOk) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultOk#a6a7bb5c as nil")
	}
	b.ObjStart()
	b.PutID("checkChatUsernameResultOk")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckChatUsernameResultOk) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultOk#a6a7bb5c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkChatUsernameResultOk"); err != nil {
				return fmt.Errorf("unable to decode checkChatUsernameResultOk#a6a7bb5c: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// CheckChatUsernameResultUsernameInvalid represents TL type `checkChatUsernameResultUsernameInvalid#da087756`.
type CheckChatUsernameResultUsernameInvalid struct {
}

// CheckChatUsernameResultUsernameInvalidTypeID is TL type id of CheckChatUsernameResultUsernameInvalid.
const CheckChatUsernameResultUsernameInvalidTypeID = 0xda087756

// construct implements constructor of CheckChatUsernameResultClass.
func (c CheckChatUsernameResultUsernameInvalid) construct() CheckChatUsernameResultClass { return &c }

// Ensuring interfaces in compile-time for CheckChatUsernameResultUsernameInvalid.
var (
	_ bin.Encoder     = &CheckChatUsernameResultUsernameInvalid{}
	_ bin.Decoder     = &CheckChatUsernameResultUsernameInvalid{}
	_ bin.BareEncoder = &CheckChatUsernameResultUsernameInvalid{}
	_ bin.BareDecoder = &CheckChatUsernameResultUsernameInvalid{}

	_ CheckChatUsernameResultClass = &CheckChatUsernameResultUsernameInvalid{}
)

func (c *CheckChatUsernameResultUsernameInvalid) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckChatUsernameResultUsernameInvalid) String() string {
	if c == nil {
		return "CheckChatUsernameResultUsernameInvalid(nil)"
	}
	type Alias CheckChatUsernameResultUsernameInvalid
	return fmt.Sprintf("CheckChatUsernameResultUsernameInvalid%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckChatUsernameResultUsernameInvalid) TypeID() uint32 {
	return CheckChatUsernameResultUsernameInvalidTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckChatUsernameResultUsernameInvalid) TypeName() string {
	return "checkChatUsernameResultUsernameInvalid"
}

// TypeInfo returns info about TL type.
func (c *CheckChatUsernameResultUsernameInvalid) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkChatUsernameResultUsernameInvalid",
		ID:   CheckChatUsernameResultUsernameInvalidTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckChatUsernameResultUsernameInvalid) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultUsernameInvalid#da087756 as nil")
	}
	b.PutID(CheckChatUsernameResultUsernameInvalidTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckChatUsernameResultUsernameInvalid) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultUsernameInvalid#da087756 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckChatUsernameResultUsernameInvalid) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultUsernameInvalid#da087756 to nil")
	}
	if err := b.ConsumeID(CheckChatUsernameResultUsernameInvalidTypeID); err != nil {
		return fmt.Errorf("unable to decode checkChatUsernameResultUsernameInvalid#da087756: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckChatUsernameResultUsernameInvalid) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultUsernameInvalid#da087756 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckChatUsernameResultUsernameInvalid) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultUsernameInvalid#da087756 as nil")
	}
	b.ObjStart()
	b.PutID("checkChatUsernameResultUsernameInvalid")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckChatUsernameResultUsernameInvalid) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultUsernameInvalid#da087756 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkChatUsernameResultUsernameInvalid"); err != nil {
				return fmt.Errorf("unable to decode checkChatUsernameResultUsernameInvalid#da087756: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// CheckChatUsernameResultUsernameOccupied represents TL type `checkChatUsernameResultUsernameOccupied#4ebb3729`.
type CheckChatUsernameResultUsernameOccupied struct {
}

// CheckChatUsernameResultUsernameOccupiedTypeID is TL type id of CheckChatUsernameResultUsernameOccupied.
const CheckChatUsernameResultUsernameOccupiedTypeID = 0x4ebb3729

// construct implements constructor of CheckChatUsernameResultClass.
func (c CheckChatUsernameResultUsernameOccupied) construct() CheckChatUsernameResultClass { return &c }

// Ensuring interfaces in compile-time for CheckChatUsernameResultUsernameOccupied.
var (
	_ bin.Encoder     = &CheckChatUsernameResultUsernameOccupied{}
	_ bin.Decoder     = &CheckChatUsernameResultUsernameOccupied{}
	_ bin.BareEncoder = &CheckChatUsernameResultUsernameOccupied{}
	_ bin.BareDecoder = &CheckChatUsernameResultUsernameOccupied{}

	_ CheckChatUsernameResultClass = &CheckChatUsernameResultUsernameOccupied{}
)

func (c *CheckChatUsernameResultUsernameOccupied) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckChatUsernameResultUsernameOccupied) String() string {
	if c == nil {
		return "CheckChatUsernameResultUsernameOccupied(nil)"
	}
	type Alias CheckChatUsernameResultUsernameOccupied
	return fmt.Sprintf("CheckChatUsernameResultUsernameOccupied%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckChatUsernameResultUsernameOccupied) TypeID() uint32 {
	return CheckChatUsernameResultUsernameOccupiedTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckChatUsernameResultUsernameOccupied) TypeName() string {
	return "checkChatUsernameResultUsernameOccupied"
}

// TypeInfo returns info about TL type.
func (c *CheckChatUsernameResultUsernameOccupied) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkChatUsernameResultUsernameOccupied",
		ID:   CheckChatUsernameResultUsernameOccupiedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckChatUsernameResultUsernameOccupied) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultUsernameOccupied#4ebb3729 as nil")
	}
	b.PutID(CheckChatUsernameResultUsernameOccupiedTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckChatUsernameResultUsernameOccupied) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultUsernameOccupied#4ebb3729 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckChatUsernameResultUsernameOccupied) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultUsernameOccupied#4ebb3729 to nil")
	}
	if err := b.ConsumeID(CheckChatUsernameResultUsernameOccupiedTypeID); err != nil {
		return fmt.Errorf("unable to decode checkChatUsernameResultUsernameOccupied#4ebb3729: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckChatUsernameResultUsernameOccupied) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultUsernameOccupied#4ebb3729 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckChatUsernameResultUsernameOccupied) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultUsernameOccupied#4ebb3729 as nil")
	}
	b.ObjStart()
	b.PutID("checkChatUsernameResultUsernameOccupied")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckChatUsernameResultUsernameOccupied) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultUsernameOccupied#4ebb3729 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkChatUsernameResultUsernameOccupied"); err != nil {
				return fmt.Errorf("unable to decode checkChatUsernameResultUsernameOccupied#4ebb3729: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// CheckChatUsernameResultPublicChatsTooMuch represents TL type `checkChatUsernameResultPublicChatsTooMuch#3327d23d`.
type CheckChatUsernameResultPublicChatsTooMuch struct {
}

// CheckChatUsernameResultPublicChatsTooMuchTypeID is TL type id of CheckChatUsernameResultPublicChatsTooMuch.
const CheckChatUsernameResultPublicChatsTooMuchTypeID = 0x3327d23d

// construct implements constructor of CheckChatUsernameResultClass.
func (c CheckChatUsernameResultPublicChatsTooMuch) construct() CheckChatUsernameResultClass {
	return &c
}

// Ensuring interfaces in compile-time for CheckChatUsernameResultPublicChatsTooMuch.
var (
	_ bin.Encoder     = &CheckChatUsernameResultPublicChatsTooMuch{}
	_ bin.Decoder     = &CheckChatUsernameResultPublicChatsTooMuch{}
	_ bin.BareEncoder = &CheckChatUsernameResultPublicChatsTooMuch{}
	_ bin.BareDecoder = &CheckChatUsernameResultPublicChatsTooMuch{}

	_ CheckChatUsernameResultClass = &CheckChatUsernameResultPublicChatsTooMuch{}
)

func (c *CheckChatUsernameResultPublicChatsTooMuch) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckChatUsernameResultPublicChatsTooMuch) String() string {
	if c == nil {
		return "CheckChatUsernameResultPublicChatsTooMuch(nil)"
	}
	type Alias CheckChatUsernameResultPublicChatsTooMuch
	return fmt.Sprintf("CheckChatUsernameResultPublicChatsTooMuch%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckChatUsernameResultPublicChatsTooMuch) TypeID() uint32 {
	return CheckChatUsernameResultPublicChatsTooMuchTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckChatUsernameResultPublicChatsTooMuch) TypeName() string {
	return "checkChatUsernameResultPublicChatsTooMuch"
}

// TypeInfo returns info about TL type.
func (c *CheckChatUsernameResultPublicChatsTooMuch) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkChatUsernameResultPublicChatsTooMuch",
		ID:   CheckChatUsernameResultPublicChatsTooMuchTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckChatUsernameResultPublicChatsTooMuch) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultPublicChatsTooMuch#3327d23d as nil")
	}
	b.PutID(CheckChatUsernameResultPublicChatsTooMuchTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckChatUsernameResultPublicChatsTooMuch) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultPublicChatsTooMuch#3327d23d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckChatUsernameResultPublicChatsTooMuch) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultPublicChatsTooMuch#3327d23d to nil")
	}
	if err := b.ConsumeID(CheckChatUsernameResultPublicChatsTooMuchTypeID); err != nil {
		return fmt.Errorf("unable to decode checkChatUsernameResultPublicChatsTooMuch#3327d23d: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckChatUsernameResultPublicChatsTooMuch) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultPublicChatsTooMuch#3327d23d to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckChatUsernameResultPublicChatsTooMuch) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultPublicChatsTooMuch#3327d23d as nil")
	}
	b.ObjStart()
	b.PutID("checkChatUsernameResultPublicChatsTooMuch")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckChatUsernameResultPublicChatsTooMuch) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultPublicChatsTooMuch#3327d23d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkChatUsernameResultPublicChatsTooMuch"); err != nil {
				return fmt.Errorf("unable to decode checkChatUsernameResultPublicChatsTooMuch#3327d23d: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// CheckChatUsernameResultPublicGroupsUnavailable represents TL type `checkChatUsernameResultPublicGroupsUnavailable#fce914d7`.
type CheckChatUsernameResultPublicGroupsUnavailable struct {
}

// CheckChatUsernameResultPublicGroupsUnavailableTypeID is TL type id of CheckChatUsernameResultPublicGroupsUnavailable.
const CheckChatUsernameResultPublicGroupsUnavailableTypeID = 0xfce914d7

// construct implements constructor of CheckChatUsernameResultClass.
func (c CheckChatUsernameResultPublicGroupsUnavailable) construct() CheckChatUsernameResultClass {
	return &c
}

// Ensuring interfaces in compile-time for CheckChatUsernameResultPublicGroupsUnavailable.
var (
	_ bin.Encoder     = &CheckChatUsernameResultPublicGroupsUnavailable{}
	_ bin.Decoder     = &CheckChatUsernameResultPublicGroupsUnavailable{}
	_ bin.BareEncoder = &CheckChatUsernameResultPublicGroupsUnavailable{}
	_ bin.BareDecoder = &CheckChatUsernameResultPublicGroupsUnavailable{}

	_ CheckChatUsernameResultClass = &CheckChatUsernameResultPublicGroupsUnavailable{}
)

func (c *CheckChatUsernameResultPublicGroupsUnavailable) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckChatUsernameResultPublicGroupsUnavailable) String() string {
	if c == nil {
		return "CheckChatUsernameResultPublicGroupsUnavailable(nil)"
	}
	type Alias CheckChatUsernameResultPublicGroupsUnavailable
	return fmt.Sprintf("CheckChatUsernameResultPublicGroupsUnavailable%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckChatUsernameResultPublicGroupsUnavailable) TypeID() uint32 {
	return CheckChatUsernameResultPublicGroupsUnavailableTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckChatUsernameResultPublicGroupsUnavailable) TypeName() string {
	return "checkChatUsernameResultPublicGroupsUnavailable"
}

// TypeInfo returns info about TL type.
func (c *CheckChatUsernameResultPublicGroupsUnavailable) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkChatUsernameResultPublicGroupsUnavailable",
		ID:   CheckChatUsernameResultPublicGroupsUnavailableTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckChatUsernameResultPublicGroupsUnavailable) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultPublicGroupsUnavailable#fce914d7 as nil")
	}
	b.PutID(CheckChatUsernameResultPublicGroupsUnavailableTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckChatUsernameResultPublicGroupsUnavailable) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultPublicGroupsUnavailable#fce914d7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckChatUsernameResultPublicGroupsUnavailable) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultPublicGroupsUnavailable#fce914d7 to nil")
	}
	if err := b.ConsumeID(CheckChatUsernameResultPublicGroupsUnavailableTypeID); err != nil {
		return fmt.Errorf("unable to decode checkChatUsernameResultPublicGroupsUnavailable#fce914d7: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckChatUsernameResultPublicGroupsUnavailable) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultPublicGroupsUnavailable#fce914d7 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckChatUsernameResultPublicGroupsUnavailable) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatUsernameResultPublicGroupsUnavailable#fce914d7 as nil")
	}
	b.ObjStart()
	b.PutID("checkChatUsernameResultPublicGroupsUnavailable")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckChatUsernameResultPublicGroupsUnavailable) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatUsernameResultPublicGroupsUnavailable#fce914d7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkChatUsernameResultPublicGroupsUnavailable"); err != nil {
				return fmt.Errorf("unable to decode checkChatUsernameResultPublicGroupsUnavailable#fce914d7: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// CheckChatUsernameResultClassName is schema name of CheckChatUsernameResultClass.
const CheckChatUsernameResultClassName = "CheckChatUsernameResult"

// CheckChatUsernameResultClass represents CheckChatUsernameResult generic type.
//
// Example:
//  g, err := tdapi.DecodeCheckChatUsernameResult(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.CheckChatUsernameResultOk: // checkChatUsernameResultOk#a6a7bb5c
//  case *tdapi.CheckChatUsernameResultUsernameInvalid: // checkChatUsernameResultUsernameInvalid#da087756
//  case *tdapi.CheckChatUsernameResultUsernameOccupied: // checkChatUsernameResultUsernameOccupied#4ebb3729
//  case *tdapi.CheckChatUsernameResultPublicChatsTooMuch: // checkChatUsernameResultPublicChatsTooMuch#3327d23d
//  case *tdapi.CheckChatUsernameResultPublicGroupsUnavailable: // checkChatUsernameResultPublicGroupsUnavailable#fce914d7
//  default: panic(v)
//  }
type CheckChatUsernameResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() CheckChatUsernameResultClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeCheckChatUsernameResult implements binary de-serialization for CheckChatUsernameResultClass.
func DecodeCheckChatUsernameResult(buf *bin.Buffer) (CheckChatUsernameResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case CheckChatUsernameResultOkTypeID:
		// Decoding checkChatUsernameResultOk#a6a7bb5c.
		v := CheckChatUsernameResultOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", err)
		}
		return &v, nil
	case CheckChatUsernameResultUsernameInvalidTypeID:
		// Decoding checkChatUsernameResultUsernameInvalid#da087756.
		v := CheckChatUsernameResultUsernameInvalid{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", err)
		}
		return &v, nil
	case CheckChatUsernameResultUsernameOccupiedTypeID:
		// Decoding checkChatUsernameResultUsernameOccupied#4ebb3729.
		v := CheckChatUsernameResultUsernameOccupied{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", err)
		}
		return &v, nil
	case CheckChatUsernameResultPublicChatsTooMuchTypeID:
		// Decoding checkChatUsernameResultPublicChatsTooMuch#3327d23d.
		v := CheckChatUsernameResultPublicChatsTooMuch{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", err)
		}
		return &v, nil
	case CheckChatUsernameResultPublicGroupsUnavailableTypeID:
		// Decoding checkChatUsernameResultPublicGroupsUnavailable#fce914d7.
		v := CheckChatUsernameResultPublicGroupsUnavailable{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONCheckChatUsernameResult implements binary de-serialization for CheckChatUsernameResultClass.
func DecodeTDLibJSONCheckChatUsernameResult(buf tdjson.Decoder) (CheckChatUsernameResultClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "checkChatUsernameResultOk":
		// Decoding checkChatUsernameResultOk#a6a7bb5c.
		v := CheckChatUsernameResultOk{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", err)
		}
		return &v, nil
	case "checkChatUsernameResultUsernameInvalid":
		// Decoding checkChatUsernameResultUsernameInvalid#da087756.
		v := CheckChatUsernameResultUsernameInvalid{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", err)
		}
		return &v, nil
	case "checkChatUsernameResultUsernameOccupied":
		// Decoding checkChatUsernameResultUsernameOccupied#4ebb3729.
		v := CheckChatUsernameResultUsernameOccupied{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", err)
		}
		return &v, nil
	case "checkChatUsernameResultPublicChatsTooMuch":
		// Decoding checkChatUsernameResultPublicChatsTooMuch#3327d23d.
		v := CheckChatUsernameResultPublicChatsTooMuch{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", err)
		}
		return &v, nil
	case "checkChatUsernameResultPublicGroupsUnavailable":
		// Decoding checkChatUsernameResultPublicGroupsUnavailable#fce914d7.
		v := CheckChatUsernameResultPublicGroupsUnavailable{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode CheckChatUsernameResultClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// CheckChatUsernameResult boxes the CheckChatUsernameResultClass providing a helper.
type CheckChatUsernameResultBox struct {
	CheckChatUsernameResult CheckChatUsernameResultClass
}

// Decode implements bin.Decoder for CheckChatUsernameResultBox.
func (b *CheckChatUsernameResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode CheckChatUsernameResultBox to nil")
	}
	v, err := DecodeCheckChatUsernameResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CheckChatUsernameResult = v
	return nil
}

// Encode implements bin.Encode for CheckChatUsernameResultBox.
func (b *CheckChatUsernameResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CheckChatUsernameResult == nil {
		return fmt.Errorf("unable to encode CheckChatUsernameResultClass as nil")
	}
	return b.CheckChatUsernameResult.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for CheckChatUsernameResultBox.
func (b *CheckChatUsernameResultBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode CheckChatUsernameResultBox to nil")
	}
	v, err := DecodeTDLibJSONCheckChatUsernameResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CheckChatUsernameResult = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for CheckChatUsernameResultBox.
func (b *CheckChatUsernameResultBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.CheckChatUsernameResult == nil {
		return fmt.Errorf("unable to encode CheckChatUsernameResultClass as nil")
	}
	return b.CheckChatUsernameResult.EncodeTDLibJSON(buf)
}
