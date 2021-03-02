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

// PhotosGetUserPhotosRequest represents TL type `photos.getUserPhotos#91cd32a8`.
// Returns the list of user photos.
//
// See https://core.telegram.org/method/photos.getUserPhotos for reference.
type PhotosGetUserPhotosRequest struct {
	// User ID
	UserID InputUserClass `tl:"user_id"`
	// Number of list elements to be skipped
	Offset int `tl:"offset"`
	// If a positive value was transferred, the method will return only photos with IDs less than the set one
	MaxID int64 `tl:"max_id"`
	// Number of list elements to be returned
	Limit int `tl:"limit"`
}

// PhotosGetUserPhotosRequestTypeID is TL type id of PhotosGetUserPhotosRequest.
const PhotosGetUserPhotosRequestTypeID = 0x91cd32a8

func (g *PhotosGetUserPhotosRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.MaxID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhotosGetUserPhotosRequest) String() string {
	if g == nil {
		return "PhotosGetUserPhotosRequest(nil)"
	}
	type Alias PhotosGetUserPhotosRequest
	return fmt.Sprintf("PhotosGetUserPhotosRequest%+v", Alias(*g))
}

// FillFrom fills PhotosGetUserPhotosRequest from given interface.
func (g *PhotosGetUserPhotosRequest) FillFrom(from interface {
	GetUserID() (value InputUserClass)
	GetOffset() (value int)
	GetMaxID() (value int64)
	GetLimit() (value int)
}) {
	g.UserID = from.GetUserID()
	g.Offset = from.GetOffset()
	g.MaxID = from.GetMaxID()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotosGetUserPhotosRequest) TypeID() uint32 {
	return PhotosGetUserPhotosRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotosGetUserPhotosRequest) TypeName() string {
	return "photos.getUserPhotos"
}

// TypeInfo returns info about TL type.
func (g *PhotosGetUserPhotosRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photos.getUserPhotos",
		ID:   PhotosGetUserPhotosRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PhotosGetUserPhotosRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode photos.getUserPhotos#91cd32a8 as nil")
	}
	b.PutID(PhotosGetUserPhotosRequestTypeID)
	if g.UserID == nil {
		return fmt.Errorf("unable to encode photos.getUserPhotos#91cd32a8: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.getUserPhotos#91cd32a8: field user_id: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutLong(g.MaxID)
	b.PutInt(g.Limit)
	return nil
}

// GetUserID returns value of UserID field.
func (g *PhotosGetUserPhotosRequest) GetUserID() (value InputUserClass) {
	return g.UserID
}

// GetOffset returns value of Offset field.
func (g *PhotosGetUserPhotosRequest) GetOffset() (value int) {
	return g.Offset
}

// GetMaxID returns value of MaxID field.
func (g *PhotosGetUserPhotosRequest) GetMaxID() (value int64) {
	return g.MaxID
}

// GetLimit returns value of Limit field.
func (g *PhotosGetUserPhotosRequest) GetLimit() (value int) {
	return g.Limit
}

// Decode implements bin.Decoder.
func (g *PhotosGetUserPhotosRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode photos.getUserPhotos#91cd32a8 to nil")
	}
	if err := b.ConsumeID(PhotosGetUserPhotosRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field user_id: %w", err)
		}
		g.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhotosGetUserPhotosRequest.
var (
	_ bin.Encoder = &PhotosGetUserPhotosRequest{}
	_ bin.Decoder = &PhotosGetUserPhotosRequest{}
)

// PhotosGetUserPhotos invokes method photos.getUserPhotos#91cd32a8 returning error if any.
// Returns the list of user photos.
//
// Possible errors:
//  400 MAX_ID_INVALID: The provided max ID is invalid
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/photos.getUserPhotos for reference.
// Can be used by bots.
func (c *Client) PhotosGetUserPhotos(ctx context.Context, request *PhotosGetUserPhotosRequest) (PhotosPhotosClass, error) {
	var result PhotosPhotosBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Photos, nil
}
