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

// TopPeer represents TL type `topPeer#edcdc05b`.
// Top peer
//
// See https://core.telegram.org/constructor/topPeer for reference.
type TopPeer struct {
	// Peer
	Peer PeerClass `tl:"peer"`
	// Rating as computed in top peer rating »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/top-rating
	Rating float64 `tl:"rating"`
}

// TopPeerTypeID is TL type id of TopPeer.
const TopPeerTypeID = 0xedcdc05b

func (t *TopPeer) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Peer == nil) {
		return false
	}
	if !(t.Rating == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeer) String() string {
	if t == nil {
		return "TopPeer(nil)"
	}
	type Alias TopPeer
	return fmt.Sprintf("TopPeer%+v", Alias(*t))
}

// FillFrom fills TopPeer from given interface.
func (t *TopPeer) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetRating() (value float64)
}) {
	t.Peer = from.GetPeer()
	t.Rating = from.GetRating()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopPeer) TypeID() uint32 {
	return TopPeerTypeID
}

// TypeName returns name of type in TL schema.
func (*TopPeer) TypeName() string {
	return "topPeer"
}

// TypeInfo returns info about TL type.
func (t *TopPeer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topPeer",
		ID:   TopPeerTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Rating",
			SchemaName: "rating",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopPeer) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeer#edcdc05b as nil")
	}
	b.PutID(TopPeerTypeID)
	if t.Peer == nil {
		return fmt.Errorf("unable to encode topPeer#edcdc05b: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode topPeer#edcdc05b: field peer: %w", err)
	}
	b.PutDouble(t.Rating)
	return nil
}

// GetPeer returns value of Peer field.
func (t *TopPeer) GetPeer() (value PeerClass) {
	return t.Peer
}

// GetRating returns value of Rating field.
func (t *TopPeer) GetRating() (value float64) {
	return t.Rating
}

// Decode implements bin.Decoder.
func (t *TopPeer) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeer#edcdc05b to nil")
	}
	if err := b.ConsumeID(TopPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeer#edcdc05b: %w", err)
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode topPeer#edcdc05b: field peer: %w", err)
		}
		t.Peer = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode topPeer#edcdc05b: field rating: %w", err)
		}
		t.Rating = value
	}
	return nil
}

// Ensuring interfaces in compile-time for TopPeer.
var (
	_ bin.Encoder = &TopPeer{}
	_ bin.Decoder = &TopPeer{}
)
