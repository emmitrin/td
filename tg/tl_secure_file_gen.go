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

// SecureFileEmpty represents TL type `secureFileEmpty#64199744`.
// Empty constructor
//
// See https://core.telegram.org/constructor/secureFileEmpty for reference.
type SecureFileEmpty struct {
}

// SecureFileEmptyTypeID is TL type id of SecureFileEmpty.
const SecureFileEmptyTypeID = 0x64199744

func (s *SecureFileEmpty) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureFileEmpty) String() string {
	if s == nil {
		return "SecureFileEmpty(nil)"
	}
	type Alias SecureFileEmpty
	return fmt.Sprintf("SecureFileEmpty%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecureFileEmpty) TypeID() uint32 {
	return SecureFileEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*SecureFileEmpty) TypeName() string {
	return "secureFileEmpty"
}

// TypeInfo returns info about TL type.
func (s *SecureFileEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secureFileEmpty",
		ID:   SecureFileEmptyTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecureFileEmpty) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureFileEmpty#64199744 as nil")
	}
	b.PutID(SecureFileEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureFileEmpty) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureFileEmpty#64199744 to nil")
	}
	if err := b.ConsumeID(SecureFileEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode secureFileEmpty#64199744: %w", err)
	}
	return nil
}

// construct implements constructor of SecureFileClass.
func (s SecureFileEmpty) construct() SecureFileClass { return &s }

// Ensuring interfaces in compile-time for SecureFileEmpty.
var (
	_ bin.Encoder = &SecureFileEmpty{}
	_ bin.Decoder = &SecureFileEmpty{}

	_ SecureFileClass = &SecureFileEmpty{}
)

// SecureFile represents TL type `secureFile#e0277a62`.
// Secure passport¹ file, for more info see the passport docs »²
//
// Links:
//  1) https://core.telegram.org/passport
//  2) https://core.telegram.org/passport/encryption#inputsecurefile
//
// See https://core.telegram.org/constructor/secureFile for reference.
type SecureFile struct {
	// ID
	ID int64 `tl:"id"`
	// Access hash
	AccessHash int64 `tl:"access_hash"`
	// File size
	Size int `tl:"size"`
	// DC ID
	DCID int `tl:"dc_id"`
	// Date of upload
	Date int `tl:"date"`
	// File hash
	FileHash []byte `tl:"file_hash"`
	// Secret
	Secret []byte `tl:"secret"`
}

// SecureFileTypeID is TL type id of SecureFile.
const SecureFileTypeID = 0xe0277a62

func (s *SecureFile) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.AccessHash == 0) {
		return false
	}
	if !(s.Size == 0) {
		return false
	}
	if !(s.DCID == 0) {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.FileHash == nil) {
		return false
	}
	if !(s.Secret == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureFile) String() string {
	if s == nil {
		return "SecureFile(nil)"
	}
	type Alias SecureFile
	return fmt.Sprintf("SecureFile%+v", Alias(*s))
}

// FillFrom fills SecureFile from given interface.
func (s *SecureFile) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetSize() (value int)
	GetDCID() (value int)
	GetDate() (value int)
	GetFileHash() (value []byte)
	GetSecret() (value []byte)
}) {
	s.ID = from.GetID()
	s.AccessHash = from.GetAccessHash()
	s.Size = from.GetSize()
	s.DCID = from.GetDCID()
	s.Date = from.GetDate()
	s.FileHash = from.GetFileHash()
	s.Secret = from.GetSecret()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecureFile) TypeID() uint32 {
	return SecureFileTypeID
}

// TypeName returns name of type in TL schema.
func (*SecureFile) TypeName() string {
	return "secureFile"
}

// TypeInfo returns info about TL type.
func (s *SecureFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secureFile",
		ID:   SecureFileTypeID,
	}
	if s == nil {
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
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "FileHash",
			SchemaName: "file_hash",
		},
		{
			Name:       "Secret",
			SchemaName: "secret",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecureFile) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureFile#e0277a62 as nil")
	}
	b.PutID(SecureFileTypeID)
	b.PutLong(s.ID)
	b.PutLong(s.AccessHash)
	b.PutInt(s.Size)
	b.PutInt(s.DCID)
	b.PutInt(s.Date)
	b.PutBytes(s.FileHash)
	b.PutBytes(s.Secret)
	return nil
}

// GetID returns value of ID field.
func (s *SecureFile) GetID() (value int64) {
	return s.ID
}

// GetAccessHash returns value of AccessHash field.
func (s *SecureFile) GetAccessHash() (value int64) {
	return s.AccessHash
}

// GetSize returns value of Size field.
func (s *SecureFile) GetSize() (value int) {
	return s.Size
}

// GetDCID returns value of DCID field.
func (s *SecureFile) GetDCID() (value int) {
	return s.DCID
}

// GetDate returns value of Date field.
func (s *SecureFile) GetDate() (value int) {
	return s.Date
}

// GetFileHash returns value of FileHash field.
func (s *SecureFile) GetFileHash() (value []byte) {
	return s.FileHash
}

// GetSecret returns value of Secret field.
func (s *SecureFile) GetSecret() (value []byte) {
	return s.Secret
}

// Decode implements bin.Decoder.
func (s *SecureFile) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureFile#e0277a62 to nil")
	}
	if err := b.ConsumeID(SecureFileTypeID); err != nil {
		return fmt.Errorf("unable to decode secureFile#e0277a62: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field access_hash: %w", err)
		}
		s.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field size: %w", err)
		}
		s.Size = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field dc_id: %w", err)
		}
		s.DCID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field file_hash: %w", err)
		}
		s.FileHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field secret: %w", err)
		}
		s.Secret = value
	}
	return nil
}

// construct implements constructor of SecureFileClass.
func (s SecureFile) construct() SecureFileClass { return &s }

// Ensuring interfaces in compile-time for SecureFile.
var (
	_ bin.Encoder = &SecureFile{}
	_ bin.Decoder = &SecureFile{}

	_ SecureFileClass = &SecureFile{}
)

// SecureFileClass represents SecureFile generic type.
//
// See https://core.telegram.org/type/SecureFile for reference.
//
// Example:
//  g, err := tg.DecodeSecureFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.SecureFileEmpty: // secureFileEmpty#64199744
//  case *tg.SecureFile: // secureFile#e0277a62
//  default: panic(v)
//  }
type SecureFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() SecureFileClass

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

	// AsNotEmpty tries to map SecureFileClass to SecureFile.
	AsNotEmpty() (*SecureFile, bool)
}

// AsInputSecureFileLocation tries to map SecureFile to InputSecureFileLocation.
func (s *SecureFile) AsInputSecureFileLocation() *InputSecureFileLocation {
	value := new(InputSecureFileLocation)
	value.ID = s.GetID()
	value.AccessHash = s.GetAccessHash()

	return value
}

// AsInput tries to map SecureFile to InputSecureFile.
func (s *SecureFile) AsInput() *InputSecureFile {
	value := new(InputSecureFile)
	value.ID = s.GetID()
	value.AccessHash = s.GetAccessHash()

	return value
}

// AsNotEmpty tries to map SecureFileEmpty to SecureFile.
func (s *SecureFileEmpty) AsNotEmpty() (*SecureFile, bool) {
	return nil, false
}

// AsNotEmpty tries to map SecureFile to SecureFile.
func (s *SecureFile) AsNotEmpty() (*SecureFile, bool) {
	return s, true
}

// DecodeSecureFile implements binary de-serialization for SecureFileClass.
func DecodeSecureFile(buf *bin.Buffer) (SecureFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SecureFileEmptyTypeID:
		// Decoding secureFileEmpty#64199744.
		v := SecureFileEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureFileClass: %w", err)
		}
		return &v, nil
	case SecureFileTypeID:
		// Decoding secureFile#e0277a62.
		v := SecureFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SecureFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// SecureFile boxes the SecureFileClass providing a helper.
type SecureFileBox struct {
	SecureFile SecureFileClass
}

// Decode implements bin.Decoder for SecureFileBox.
func (b *SecureFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SecureFileBox to nil")
	}
	v, err := DecodeSecureFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SecureFile = v
	return nil
}

// Encode implements bin.Encode for SecureFileBox.
func (b *SecureFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SecureFile == nil {
		return fmt.Errorf("unable to encode SecureFileClass as nil")
	}
	return b.SecureFile.Encode(buf)
}

// SecureFileClassArray is adapter for slice of SecureFileClass.
type SecureFileClassArray []SecureFileClass

// Sort sorts slice of SecureFileClass.
func (s SecureFileClassArray) Sort(less func(a, b SecureFileClass) bool) SecureFileClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureFileClass.
func (s SecureFileClassArray) SortStable(less func(a, b SecureFileClass) bool) SecureFileClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureFileClass.
func (s SecureFileClassArray) Retain(keep func(x SecureFileClass) bool) SecureFileClassArray {
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
func (s SecureFileClassArray) First() (v SecureFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureFileClassArray) Last() (v SecureFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureFileClassArray) PopFirst() (v SecureFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureFileClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureFileClassArray) Pop() (v SecureFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsSecureFile returns copy with only SecureFile constructors.
func (s SecureFileClassArray) AsSecureFile() (to SecureFileArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s SecureFileClassArray) AppendOnlyNotEmpty(to []*SecureFile) []*SecureFile {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s SecureFileClassArray) AsNotEmpty() (to []*SecureFile) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s SecureFileClassArray) FirstAsNotEmpty() (v *SecureFile, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s SecureFileClassArray) LastAsNotEmpty() (v *SecureFile, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *SecureFileClassArray) PopFirstAsNotEmpty() (v *SecureFile, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *SecureFileClassArray) PopAsNotEmpty() (v *SecureFile, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// SecureFileArray is adapter for slice of SecureFile.
type SecureFileArray []SecureFile

// Sort sorts slice of SecureFile.
func (s SecureFileArray) Sort(less func(a, b SecureFile) bool) SecureFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureFile.
func (s SecureFileArray) SortStable(less func(a, b SecureFile) bool) SecureFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureFile.
func (s SecureFileArray) Retain(keep func(x SecureFile) bool) SecureFileArray {
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
func (s SecureFileArray) First() (v SecureFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureFileArray) Last() (v SecureFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureFileArray) PopFirst() (v SecureFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureFileArray) Pop() (v SecureFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of SecureFile by Date.
func (s SecureFileArray) SortByDate() SecureFileArray {
	return s.Sort(func(a, b SecureFile) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of SecureFile by Date.
func (s SecureFileArray) SortStableByDate() SecureFileArray {
	return s.SortStable(func(a, b SecureFile) bool {
		return a.GetDate() < b.GetDate()
	})
}
