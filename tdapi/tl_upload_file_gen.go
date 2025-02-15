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

// UploadFileRequest represents TL type `uploadFile#d38f14a6`.
type UploadFileRequest struct {
	// File to upload
	File InputFileClass
	// File type; pass null if unknown
	FileType FileTypeClass
	// Priority of the upload (1-32). The higher the priority, the earlier the file will be
	// uploaded. If the priorities of two files are equal, then the first one for which
	// uploadFile was called will be uploaded first
	Priority int32
}

// UploadFileRequestTypeID is TL type id of UploadFileRequest.
const UploadFileRequestTypeID = 0xd38f14a6

// Ensuring interfaces in compile-time for UploadFileRequest.
var (
	_ bin.Encoder     = &UploadFileRequest{}
	_ bin.Decoder     = &UploadFileRequest{}
	_ bin.BareEncoder = &UploadFileRequest{}
	_ bin.BareDecoder = &UploadFileRequest{}
)

func (u *UploadFileRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.File == nil) {
		return false
	}
	if !(u.FileType == nil) {
		return false
	}
	if !(u.Priority == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UploadFileRequest) String() string {
	if u == nil {
		return "UploadFileRequest(nil)"
	}
	type Alias UploadFileRequest
	return fmt.Sprintf("UploadFileRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadFileRequest) TypeID() uint32 {
	return UploadFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadFileRequest) TypeName() string {
	return "uploadFile"
}

// TypeInfo returns info about TL type.
func (u *UploadFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "uploadFile",
		ID:   UploadFileRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "File",
			SchemaName: "file",
		},
		{
			Name:       "FileType",
			SchemaName: "file_type",
		},
		{
			Name:       "Priority",
			SchemaName: "priority",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UploadFileRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode uploadFile#d38f14a6 as nil")
	}
	b.PutID(UploadFileRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UploadFileRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode uploadFile#d38f14a6 as nil")
	}
	if u.File == nil {
		return fmt.Errorf("unable to encode uploadFile#d38f14a6: field file is nil")
	}
	if err := u.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode uploadFile#d38f14a6: field file: %w", err)
	}
	if u.FileType == nil {
		return fmt.Errorf("unable to encode uploadFile#d38f14a6: field file_type is nil")
	}
	if err := u.FileType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode uploadFile#d38f14a6: field file_type: %w", err)
	}
	b.PutInt32(u.Priority)
	return nil
}

// Decode implements bin.Decoder.
func (u *UploadFileRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode uploadFile#d38f14a6 to nil")
	}
	if err := b.ConsumeID(UploadFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode uploadFile#d38f14a6: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UploadFileRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode uploadFile#d38f14a6 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode uploadFile#d38f14a6: field file: %w", err)
		}
		u.File = value
	}
	{
		value, err := DecodeFileType(b)
		if err != nil {
			return fmt.Errorf("unable to decode uploadFile#d38f14a6: field file_type: %w", err)
		}
		u.FileType = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode uploadFile#d38f14a6: field priority: %w", err)
		}
		u.Priority = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UploadFileRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode uploadFile#d38f14a6 as nil")
	}
	b.ObjStart()
	b.PutID("uploadFile")
	b.Comma()
	b.FieldStart("file")
	if u.File == nil {
		return fmt.Errorf("unable to encode uploadFile#d38f14a6: field file is nil")
	}
	if err := u.File.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode uploadFile#d38f14a6: field file: %w", err)
	}
	b.Comma()
	b.FieldStart("file_type")
	if u.FileType == nil {
		return fmt.Errorf("unable to encode uploadFile#d38f14a6: field file_type is nil")
	}
	if err := u.FileType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode uploadFile#d38f14a6: field file_type: %w", err)
	}
	b.Comma()
	b.FieldStart("priority")
	b.PutInt32(u.Priority)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UploadFileRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode uploadFile#d38f14a6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("uploadFile"); err != nil {
				return fmt.Errorf("unable to decode uploadFile#d38f14a6: %w", err)
			}
		case "file":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode uploadFile#d38f14a6: field file: %w", err)
			}
			u.File = value
		case "file_type":
			value, err := DecodeTDLibJSONFileType(b)
			if err != nil {
				return fmt.Errorf("unable to decode uploadFile#d38f14a6: field file_type: %w", err)
			}
			u.FileType = value
		case "priority":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode uploadFile#d38f14a6: field priority: %w", err)
			}
			u.Priority = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFile returns value of File field.
func (u *UploadFileRequest) GetFile() (value InputFileClass) {
	if u == nil {
		return
	}
	return u.File
}

// GetFileType returns value of FileType field.
func (u *UploadFileRequest) GetFileType() (value FileTypeClass) {
	if u == nil {
		return
	}
	return u.FileType
}

// GetPriority returns value of Priority field.
func (u *UploadFileRequest) GetPriority() (value int32) {
	if u == nil {
		return
	}
	return u.Priority
}

// UploadFile invokes method uploadFile#d38f14a6 returning error if any.
func (c *Client) UploadFile(ctx context.Context, request *UploadFileRequest) (*File, error) {
	var result File

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
