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

// CreateNewStickerSetRequest represents TL type `createNewStickerSet#aa878026`.
type CreateNewStickerSetRequest struct {
	// Sticker set owner; ignored for regular users
	UserID int64
	// Sticker set title; 1-64 characters
	Title string
	// Sticker set name. Can contain only English letters, digits and underscores. Must end
	// with *"_by_<bot username>"* (*<bot_username>* is case insensitive) for bots; 1-64
	// characters
	Name string
	// List of stickers to be added to the set; must be non-empty. All stickers must have the
	// same format. For TGS stickers, uploadStickerFile must be used before the sticker is
	// shown
	Stickers []InputSticker
	// Source of the sticker set; may be empty if unknown
	Source string
}

// CreateNewStickerSetRequestTypeID is TL type id of CreateNewStickerSetRequest.
const CreateNewStickerSetRequestTypeID = 0xaa878026

// Ensuring interfaces in compile-time for CreateNewStickerSetRequest.
var (
	_ bin.Encoder     = &CreateNewStickerSetRequest{}
	_ bin.Decoder     = &CreateNewStickerSetRequest{}
	_ bin.BareEncoder = &CreateNewStickerSetRequest{}
	_ bin.BareDecoder = &CreateNewStickerSetRequest{}
)

func (c *CreateNewStickerSetRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.Name == "") {
		return false
	}
	if !(c.Stickers == nil) {
		return false
	}
	if !(c.Source == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CreateNewStickerSetRequest) String() string {
	if c == nil {
		return "CreateNewStickerSetRequest(nil)"
	}
	type Alias CreateNewStickerSetRequest
	return fmt.Sprintf("CreateNewStickerSetRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CreateNewStickerSetRequest) TypeID() uint32 {
	return CreateNewStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CreateNewStickerSetRequest) TypeName() string {
	return "createNewStickerSet"
}

// TypeInfo returns info about TL type.
func (c *CreateNewStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "createNewStickerSet",
		ID:   CreateNewStickerSetRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Stickers",
			SchemaName: "stickers",
		},
		{
			Name:       "Source",
			SchemaName: "source",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CreateNewStickerSetRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createNewStickerSet#aa878026 as nil")
	}
	b.PutID(CreateNewStickerSetRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CreateNewStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createNewStickerSet#aa878026 as nil")
	}
	b.PutInt53(c.UserID)
	b.PutString(c.Title)
	b.PutString(c.Name)
	b.PutInt(len(c.Stickers))
	for idx, v := range c.Stickers {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare createNewStickerSet#aa878026: field stickers element with index %d: %w", idx, err)
		}
	}
	b.PutString(c.Source)
	return nil
}

// Decode implements bin.Decoder.
func (c *CreateNewStickerSetRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createNewStickerSet#aa878026 to nil")
	}
	if err := b.ConsumeID(CreateNewStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode createNewStickerSet#aa878026: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CreateNewStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createNewStickerSet#aa878026 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field title: %w", err)
		}
		c.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field name: %w", err)
		}
		c.Name = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field stickers: %w", err)
		}

		if headerLen > 0 {
			c.Stickers = make([]InputSticker, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputSticker
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare createNewStickerSet#aa878026: field stickers: %w", err)
			}
			c.Stickers = append(c.Stickers, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field source: %w", err)
		}
		c.Source = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CreateNewStickerSetRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode createNewStickerSet#aa878026 as nil")
	}
	b.ObjStart()
	b.PutID("createNewStickerSet")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(c.UserID)
	b.Comma()
	b.FieldStart("title")
	b.PutString(c.Title)
	b.Comma()
	b.FieldStart("name")
	b.PutString(c.Name)
	b.Comma()
	b.FieldStart("stickers")
	b.ArrStart()
	for idx, v := range c.Stickers {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode createNewStickerSet#aa878026: field stickers element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("source")
	b.PutString(c.Source)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CreateNewStickerSetRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode createNewStickerSet#aa878026 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("createNewStickerSet"); err != nil {
				return fmt.Errorf("unable to decode createNewStickerSet#aa878026: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field user_id: %w", err)
			}
			c.UserID = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field title: %w", err)
			}
			c.Title = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field name: %w", err)
			}
			c.Name = value
		case "stickers":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value InputSticker
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field stickers: %w", err)
				}
				c.Stickers = append(c.Stickers, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field stickers: %w", err)
			}
		case "source":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode createNewStickerSet#aa878026: field source: %w", err)
			}
			c.Source = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *CreateNewStickerSetRequest) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}

// GetTitle returns value of Title field.
func (c *CreateNewStickerSetRequest) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// GetName returns value of Name field.
func (c *CreateNewStickerSetRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

// GetStickers returns value of Stickers field.
func (c *CreateNewStickerSetRequest) GetStickers() (value []InputSticker) {
	if c == nil {
		return
	}
	return c.Stickers
}

// GetSource returns value of Source field.
func (c *CreateNewStickerSetRequest) GetSource() (value string) {
	if c == nil {
		return
	}
	return c.Source
}

// CreateNewStickerSet invokes method createNewStickerSet#aa878026 returning error if any.
func (c *Client) CreateNewStickerSet(ctx context.Context, request *CreateNewStickerSetRequest) (*StickerSet, error) {
	var result StickerSet

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
