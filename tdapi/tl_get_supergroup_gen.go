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

// GetSupergroupRequest represents TL type `getSupergroup#3afd10e2`.
type GetSupergroupRequest struct {
	// Supergroup or channel identifier
	SupergroupID int64
}

// GetSupergroupRequestTypeID is TL type id of GetSupergroupRequest.
const GetSupergroupRequestTypeID = 0x3afd10e2

// Ensuring interfaces in compile-time for GetSupergroupRequest.
var (
	_ bin.Encoder     = &GetSupergroupRequest{}
	_ bin.Decoder     = &GetSupergroupRequest{}
	_ bin.BareEncoder = &GetSupergroupRequest{}
	_ bin.BareDecoder = &GetSupergroupRequest{}
)

func (g *GetSupergroupRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.SupergroupID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSupergroupRequest) String() string {
	if g == nil {
		return "GetSupergroupRequest(nil)"
	}
	type Alias GetSupergroupRequest
	return fmt.Sprintf("GetSupergroupRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSupergroupRequest) TypeID() uint32 {
	return GetSupergroupRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSupergroupRequest) TypeName() string {
	return "getSupergroup"
}

// TypeInfo returns info about TL type.
func (g *GetSupergroupRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSupergroup",
		ID:   GetSupergroupRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSupergroupRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupergroup#3afd10e2 as nil")
	}
	b.PutID(GetSupergroupRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSupergroupRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupergroup#3afd10e2 as nil")
	}
	b.PutInt53(g.SupergroupID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSupergroupRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupergroup#3afd10e2 to nil")
	}
	if err := b.ConsumeID(GetSupergroupRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSupergroup#3afd10e2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSupergroupRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupergroup#3afd10e2 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getSupergroup#3afd10e2: field supergroup_id: %w", err)
		}
		g.SupergroupID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSupergroupRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupergroup#3afd10e2 as nil")
	}
	b.ObjStart()
	b.PutID("getSupergroup")
	b.FieldStart("supergroup_id")
	b.PutInt53(g.SupergroupID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSupergroupRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupergroup#3afd10e2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSupergroup"); err != nil {
				return fmt.Errorf("unable to decode getSupergroup#3afd10e2: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getSupergroup#3afd10e2: field supergroup_id: %w", err)
			}
			g.SupergroupID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (g *GetSupergroupRequest) GetSupergroupID() (value int64) {
	return g.SupergroupID
}

// GetSupergroup invokes method getSupergroup#3afd10e2 returning error if any.
func (c *Client) GetSupergroup(ctx context.Context, supergroupid int64) (*Supergroup, error) {
	var result Supergroup

	request := &GetSupergroupRequest{
		SupergroupID: supergroupid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
