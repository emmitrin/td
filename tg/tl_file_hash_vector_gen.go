// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
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
)

// FileHashVector is a box for Vector<FileHash>
type FileHashVector struct {
	// Elements of Vector<FileHash>
	Elems []FileHash
}

// FileHashVectorTypeID is TL type id of FileHashVector.
const FileHashVectorTypeID = bin.TypeVector

func (vec *FileHashVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *FileHashVector) String() string {
	if vec == nil {
		return "FileHashVector(nil)"
	}
	type Alias FileHashVector
	return fmt.Sprintf("FileHashVector%+v", Alias(*vec))
}

// FillFrom fills FileHashVector from given interface.
func (vec *FileHashVector) FillFrom(from interface {
	GetElems() (value []FileHash)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FileHashVector) TypeID() uint32 {
	return FileHashVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*FileHashVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *FileHashVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   FileHashVectorTypeID,
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
func (vec *FileHashVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<FileHash> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *FileHashVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<FileHash> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<FileHash>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *FileHashVector) GetElems() (value []FileHash) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *FileHashVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<FileHash> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *FileHashVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<FileHash> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<FileHash>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]FileHash, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value FileHash
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<FileHash>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for FileHashVector.
var (
	_ bin.Encoder     = &FileHashVector{}
	_ bin.Decoder     = &FileHashVector{}
	_ bin.BareEncoder = &FileHashVector{}
	_ bin.BareDecoder = &FileHashVector{}
)
