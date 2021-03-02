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

// InputTheme represents TL type `inputTheme#3c5693e9`.
// Theme
//
// See https://core.telegram.org/constructor/inputTheme for reference.
type InputTheme struct {
	// ID
	ID int64 `tl:"id"`
	// Access hash
	AccessHash int64 `tl:"access_hash"`
}

// InputThemeTypeID is TL type id of InputTheme.
const InputThemeTypeID = 0x3c5693e9

func (i *InputTheme) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputTheme) String() string {
	if i == nil {
		return "InputTheme(nil)"
	}
	type Alias InputTheme
	return fmt.Sprintf("InputTheme%+v", Alias(*i))
}

// FillFrom fills InputTheme from given interface.
func (i *InputTheme) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputTheme) TypeID() uint32 {
	return InputThemeTypeID
}

// TypeName returns name of type in TL schema.
func (*InputTheme) TypeName() string {
	return "inputTheme"
}

// TypeInfo returns info about TL type.
func (i *InputTheme) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputTheme",
		ID:   InputThemeTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputTheme) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputTheme#3c5693e9 as nil")
	}
	b.PutID(InputThemeTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// GetID returns value of ID field.
func (i *InputTheme) GetID() (value int64) {
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputTheme) GetAccessHash() (value int64) {
	return i.AccessHash
}

// Decode implements bin.Decoder.
func (i *InputTheme) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputTheme#3c5693e9 to nil")
	}
	if err := b.ConsumeID(InputThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode inputTheme#3c5693e9: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputTheme#3c5693e9: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputTheme#3c5693e9: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputThemeClass.
func (i InputTheme) construct() InputThemeClass { return &i }

// Ensuring interfaces in compile-time for InputTheme.
var (
	_ bin.Encoder = &InputTheme{}
	_ bin.Decoder = &InputTheme{}

	_ InputThemeClass = &InputTheme{}
)

// InputThemeSlug represents TL type `inputThemeSlug#f5890df1`.
// Theme by theme ID
//
// See https://core.telegram.org/constructor/inputThemeSlug for reference.
type InputThemeSlug struct {
	// Unique theme ID
	Slug string `tl:"slug"`
}

// InputThemeSlugTypeID is TL type id of InputThemeSlug.
const InputThemeSlugTypeID = 0xf5890df1

func (i *InputThemeSlug) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Slug == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputThemeSlug) String() string {
	if i == nil {
		return "InputThemeSlug(nil)"
	}
	type Alias InputThemeSlug
	return fmt.Sprintf("InputThemeSlug%+v", Alias(*i))
}

// FillFrom fills InputThemeSlug from given interface.
func (i *InputThemeSlug) FillFrom(from interface {
	GetSlug() (value string)
}) {
	i.Slug = from.GetSlug()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputThemeSlug) TypeID() uint32 {
	return InputThemeSlugTypeID
}

// TypeName returns name of type in TL schema.
func (*InputThemeSlug) TypeName() string {
	return "inputThemeSlug"
}

// TypeInfo returns info about TL type.
func (i *InputThemeSlug) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputThemeSlug",
		ID:   InputThemeSlugTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Slug",
			SchemaName: "slug",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputThemeSlug) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputThemeSlug#f5890df1 as nil")
	}
	b.PutID(InputThemeSlugTypeID)
	b.PutString(i.Slug)
	return nil
}

// GetSlug returns value of Slug field.
func (i *InputThemeSlug) GetSlug() (value string) {
	return i.Slug
}

// Decode implements bin.Decoder.
func (i *InputThemeSlug) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputThemeSlug#f5890df1 to nil")
	}
	if err := b.ConsumeID(InputThemeSlugTypeID); err != nil {
		return fmt.Errorf("unable to decode inputThemeSlug#f5890df1: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputThemeSlug#f5890df1: field slug: %w", err)
		}
		i.Slug = value
	}
	return nil
}

// construct implements constructor of InputThemeClass.
func (i InputThemeSlug) construct() InputThemeClass { return &i }

// Ensuring interfaces in compile-time for InputThemeSlug.
var (
	_ bin.Encoder = &InputThemeSlug{}
	_ bin.Decoder = &InputThemeSlug{}

	_ InputThemeClass = &InputThemeSlug{}
)

// InputThemeClass represents InputTheme generic type.
//
// See https://core.telegram.org/type/InputTheme for reference.
//
// Example:
//  g, err := tg.DecodeInputTheme(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputTheme: // inputTheme#3c5693e9
//  case *tg.InputThemeSlug: // inputThemeSlug#f5890df1
//  default: panic(v)
//  }
type InputThemeClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputThemeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputTheme implements binary de-serialization for InputThemeClass.
func DecodeInputTheme(buf *bin.Buffer) (InputThemeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputThemeTypeID:
		// Decoding inputTheme#3c5693e9.
		v := InputTheme{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputThemeClass: %w", err)
		}
		return &v, nil
	case InputThemeSlugTypeID:
		// Decoding inputThemeSlug#f5890df1.
		v := InputThemeSlug{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputThemeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputThemeClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputTheme boxes the InputThemeClass providing a helper.
type InputThemeBox struct {
	InputTheme InputThemeClass
}

// Decode implements bin.Decoder for InputThemeBox.
func (b *InputThemeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputThemeBox to nil")
	}
	v, err := DecodeInputTheme(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputTheme = v
	return nil
}

// Encode implements bin.Encode for InputThemeBox.
func (b *InputThemeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputTheme == nil {
		return fmt.Errorf("unable to encode InputThemeClass as nil")
	}
	return b.InputTheme.Encode(buf)
}

// InputThemeClassArray is adapter for slice of InputThemeClass.
type InputThemeClassArray []InputThemeClass

// Sort sorts slice of InputThemeClass.
func (s InputThemeClassArray) Sort(less func(a, b InputThemeClass) bool) InputThemeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputThemeClass.
func (s InputThemeClassArray) SortStable(less func(a, b InputThemeClass) bool) InputThemeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputThemeClass.
func (s InputThemeClassArray) Retain(keep func(x InputThemeClass) bool) InputThemeClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputThemeClassArray) First() (v InputThemeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputThemeClassArray) Last() (v InputThemeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputThemeClassArray) PopFirst() (v InputThemeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputThemeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputThemeClassArray) Pop() (v InputThemeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputTheme returns copy with only InputTheme constructors.
func (s InputThemeClassArray) AsInputTheme() (to InputThemeArray) {
	for _, elem := range s {
		value, ok := elem.(*InputTheme)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputThemeSlug returns copy with only InputThemeSlug constructors.
func (s InputThemeClassArray) AsInputThemeSlug() (to InputThemeSlugArray) {
	for _, elem := range s {
		value, ok := elem.(*InputThemeSlug)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputThemeArray is adapter for slice of InputTheme.
type InputThemeArray []InputTheme

// Sort sorts slice of InputTheme.
func (s InputThemeArray) Sort(less func(a, b InputTheme) bool) InputThemeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputTheme.
func (s InputThemeArray) SortStable(less func(a, b InputTheme) bool) InputThemeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputTheme.
func (s InputThemeArray) Retain(keep func(x InputTheme) bool) InputThemeArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputThemeArray) First() (v InputTheme, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputThemeArray) Last() (v InputTheme, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputThemeArray) PopFirst() (v InputTheme, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputTheme
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputThemeArray) Pop() (v InputTheme, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputThemeSlugArray is adapter for slice of InputThemeSlug.
type InputThemeSlugArray []InputThemeSlug

// Sort sorts slice of InputThemeSlug.
func (s InputThemeSlugArray) Sort(less func(a, b InputThemeSlug) bool) InputThemeSlugArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputThemeSlug.
func (s InputThemeSlugArray) SortStable(less func(a, b InputThemeSlug) bool) InputThemeSlugArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputThemeSlug.
func (s InputThemeSlugArray) Retain(keep func(x InputThemeSlug) bool) InputThemeSlugArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputThemeSlugArray) First() (v InputThemeSlug, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputThemeSlugArray) Last() (v InputThemeSlug, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputThemeSlugArray) PopFirst() (v InputThemeSlug, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputThemeSlug
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputThemeSlugArray) Pop() (v InputThemeSlug, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
