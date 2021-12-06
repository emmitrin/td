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

// UploadStickerFileRequest represents TL type `uploadStickerFile#23a0f58e`.
type UploadStickerFileRequest struct {
	// Sticker file owner; ignored for regular users
	UserID int64
	// Sticker file to upload
	Sticker InputStickerClass
}

// UploadStickerFileRequestTypeID is TL type id of UploadStickerFileRequest.
const UploadStickerFileRequestTypeID = 0x23a0f58e

// Ensuring interfaces in compile-time for UploadStickerFileRequest.
var (
	_ bin.Encoder     = &UploadStickerFileRequest{}
	_ bin.Decoder     = &UploadStickerFileRequest{}
	_ bin.BareEncoder = &UploadStickerFileRequest{}
	_ bin.BareDecoder = &UploadStickerFileRequest{}
)

func (u *UploadStickerFileRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.UserID == 0) {
		return false
	}
	if !(u.Sticker == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UploadStickerFileRequest) String() string {
	if u == nil {
		return "UploadStickerFileRequest(nil)"
	}
	type Alias UploadStickerFileRequest
	return fmt.Sprintf("UploadStickerFileRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadStickerFileRequest) TypeID() uint32 {
	return UploadStickerFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadStickerFileRequest) TypeName() string {
	return "uploadStickerFile"
}

// TypeInfo returns info about TL type.
func (u *UploadStickerFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "uploadStickerFile",
		ID:   UploadStickerFileRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UploadStickerFileRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode uploadStickerFile#23a0f58e as nil")
	}
	b.PutID(UploadStickerFileRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UploadStickerFileRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode uploadStickerFile#23a0f58e as nil")
	}
	b.PutInt53(u.UserID)
	if u.Sticker == nil {
		return fmt.Errorf("unable to encode uploadStickerFile#23a0f58e: field sticker is nil")
	}
	if err := u.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode uploadStickerFile#23a0f58e: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UploadStickerFileRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode uploadStickerFile#23a0f58e to nil")
	}
	if err := b.ConsumeID(UploadStickerFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode uploadStickerFile#23a0f58e: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UploadStickerFileRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode uploadStickerFile#23a0f58e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode uploadStickerFile#23a0f58e: field user_id: %w", err)
		}
		u.UserID = value
	}
	{
		value, err := DecodeInputSticker(b)
		if err != nil {
			return fmt.Errorf("unable to decode uploadStickerFile#23a0f58e: field sticker: %w", err)
		}
		u.Sticker = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UploadStickerFileRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode uploadStickerFile#23a0f58e as nil")
	}
	b.ObjStart()
	b.PutID("uploadStickerFile")
	b.FieldStart("user_id")
	b.PutInt53(u.UserID)
	b.FieldStart("sticker")
	if u.Sticker == nil {
		return fmt.Errorf("unable to encode uploadStickerFile#23a0f58e: field sticker is nil")
	}
	if err := u.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode uploadStickerFile#23a0f58e: field sticker: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UploadStickerFileRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode uploadStickerFile#23a0f58e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("uploadStickerFile"); err != nil {
				return fmt.Errorf("unable to decode uploadStickerFile#23a0f58e: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode uploadStickerFile#23a0f58e: field user_id: %w", err)
			}
			u.UserID = value
		case "sticker":
			value, err := DecodeTDLibJSONInputSticker(b)
			if err != nil {
				return fmt.Errorf("unable to decode uploadStickerFile#23a0f58e: field sticker: %w", err)
			}
			u.Sticker = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (u *UploadStickerFileRequest) GetUserID() (value int64) {
	return u.UserID
}

// GetSticker returns value of Sticker field.
func (u *UploadStickerFileRequest) GetSticker() (value InputStickerClass) {
	return u.Sticker
}

// UploadStickerFile invokes method uploadStickerFile#23a0f58e returning error if any.
func (c *Client) UploadStickerFile(ctx context.Context, request *UploadStickerFileRequest) (*File, error) {
	var result File

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
