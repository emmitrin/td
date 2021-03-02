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

// WallPaperClassVector is a box for Vector<WallPaper>
type WallPaperClassVector struct {
	// Elements of Vector<WallPaper>
	Elems []WallPaperClass `tl:"Elems"`
}

// WallPaperClassVectorTypeID is TL type id of WallPaperClassVector.
const WallPaperClassVectorTypeID = bin.TypeVector

func (vec *WallPaperClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *WallPaperClassVector) String() string {
	if vec == nil {
		return "WallPaperClassVector(nil)"
	}
	type Alias WallPaperClassVector
	return fmt.Sprintf("WallPaperClassVector%+v", Alias(*vec))
}

// FillFrom fills WallPaperClassVector from given interface.
func (vec *WallPaperClassVector) FillFrom(from interface {
	GetElems() (value []WallPaperClass)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WallPaperClassVector) TypeID() uint32 {
	return WallPaperClassVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*WallPaperClassVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *WallPaperClassVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   WallPaperClassVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *WallPaperClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<WallPaper> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<WallPaper>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<WallPaper>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *WallPaperClassVector) GetElems() (value []WallPaperClass) {
	return vec.Elems
}

// MapElems returns field Elems wrapped in WallPaperClassArray helper.
func (vec *WallPaperClassVector) MapElems() (value WallPaperClassArray) {
	return WallPaperClassArray(vec.Elems)
}

// Decode implements bin.Decoder.
func (vec *WallPaperClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<WallPaper> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<WallPaper>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeWallPaper(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<WallPaper>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for WallPaperClassVector.
var (
	_ bin.Encoder = &WallPaperClassVector{}
	_ bin.Decoder = &WallPaperClassVector{}
)
