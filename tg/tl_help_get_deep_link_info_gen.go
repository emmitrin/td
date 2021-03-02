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

// HelpGetDeepLinkInfoRequest represents TL type `help.getDeepLinkInfo#3fedc75f`.
// Get info about a t.me link
//
// See https://core.telegram.org/method/help.getDeepLinkInfo for reference.
type HelpGetDeepLinkInfoRequest struct {
	// Path in t.me/path
	Path string `tl:"path"`
}

// HelpGetDeepLinkInfoRequestTypeID is TL type id of HelpGetDeepLinkInfoRequest.
const HelpGetDeepLinkInfoRequestTypeID = 0x3fedc75f

func (g *HelpGetDeepLinkInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Path == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetDeepLinkInfoRequest) String() string {
	if g == nil {
		return "HelpGetDeepLinkInfoRequest(nil)"
	}
	type Alias HelpGetDeepLinkInfoRequest
	return fmt.Sprintf("HelpGetDeepLinkInfoRequest%+v", Alias(*g))
}

// FillFrom fills HelpGetDeepLinkInfoRequest from given interface.
func (g *HelpGetDeepLinkInfoRequest) FillFrom(from interface {
	GetPath() (value string)
}) {
	g.Path = from.GetPath()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetDeepLinkInfoRequest) TypeID() uint32 {
	return HelpGetDeepLinkInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetDeepLinkInfoRequest) TypeName() string {
	return "help.getDeepLinkInfo"
}

// TypeInfo returns info about TL type.
func (g *HelpGetDeepLinkInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getDeepLinkInfo",
		ID:   HelpGetDeepLinkInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Path",
			SchemaName: "path",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetDeepLinkInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getDeepLinkInfo#3fedc75f as nil")
	}
	b.PutID(HelpGetDeepLinkInfoRequestTypeID)
	b.PutString(g.Path)
	return nil
}

// GetPath returns value of Path field.
func (g *HelpGetDeepLinkInfoRequest) GetPath() (value string) {
	return g.Path
}

// Decode implements bin.Decoder.
func (g *HelpGetDeepLinkInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getDeepLinkInfo#3fedc75f to nil")
	}
	if err := b.ConsumeID(HelpGetDeepLinkInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getDeepLinkInfo#3fedc75f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.getDeepLinkInfo#3fedc75f: field path: %w", err)
		}
		g.Path = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetDeepLinkInfoRequest.
var (
	_ bin.Encoder = &HelpGetDeepLinkInfoRequest{}
	_ bin.Decoder = &HelpGetDeepLinkInfoRequest{}
)

// HelpGetDeepLinkInfo invokes method help.getDeepLinkInfo#3fedc75f returning error if any.
// Get info about a t.me link
//
// See https://core.telegram.org/method/help.getDeepLinkInfo for reference.
func (c *Client) HelpGetDeepLinkInfo(ctx context.Context, path string) (HelpDeepLinkInfoClass, error) {
	var result HelpDeepLinkInfoBox

	request := &HelpGetDeepLinkInfoRequest{
		Path: path,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.DeepLinkInfo, nil
}
