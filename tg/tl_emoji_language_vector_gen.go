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

// EmojiLanguageVector is a box for Vector<EmojiLanguage>
type EmojiLanguageVector struct {
	// Elements of Vector<EmojiLanguage>
	Elems []EmojiLanguage `tl:"Elems"`
}

// EmojiLanguageVectorTypeID is TL type id of EmojiLanguageVector.
const EmojiLanguageVectorTypeID = bin.TypeVector

func (vec *EmojiLanguageVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *EmojiLanguageVector) String() string {
	if vec == nil {
		return "EmojiLanguageVector(nil)"
	}
	type Alias EmojiLanguageVector
	return fmt.Sprintf("EmojiLanguageVector%+v", Alias(*vec))
}

// FillFrom fills EmojiLanguageVector from given interface.
func (vec *EmojiLanguageVector) FillFrom(from interface {
	GetElems() (value []EmojiLanguage)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiLanguageVector) TypeID() uint32 {
	return EmojiLanguageVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiLanguageVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *EmojiLanguageVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   EmojiLanguageVectorTypeID,
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
func (vec *EmojiLanguageVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<EmojiLanguage> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<EmojiLanguage>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *EmojiLanguageVector) GetElems() (value []EmojiLanguage) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *EmojiLanguageVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<EmojiLanguage> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<EmojiLanguage>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value EmojiLanguage
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<EmojiLanguage>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for EmojiLanguageVector.
var (
	_ bin.Encoder = &EmojiLanguageVector{}
	_ bin.Decoder = &EmojiLanguageVector{}
)
