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

// ContactsTopPeersNotModified represents TL type `contacts.topPeersNotModified#de266ef5`.
// Top peer info hasn't changed
//
// See https://core.telegram.org/constructor/contacts.topPeersNotModified for reference.
type ContactsTopPeersNotModified struct {
}

// ContactsTopPeersNotModifiedTypeID is TL type id of ContactsTopPeersNotModified.
const ContactsTopPeersNotModifiedTypeID = 0xde266ef5

func (t *ContactsTopPeersNotModified) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ContactsTopPeersNotModified) String() string {
	if t == nil {
		return "ContactsTopPeersNotModified(nil)"
	}
	type Alias ContactsTopPeersNotModified
	return fmt.Sprintf("ContactsTopPeersNotModified%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsTopPeersNotModified) TypeID() uint32 {
	return ContactsTopPeersNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsTopPeersNotModified) TypeName() string {
	return "contacts.topPeersNotModified"
}

// TypeInfo returns info about TL type.
func (t *ContactsTopPeersNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.topPeersNotModified",
		ID:   ContactsTopPeersNotModifiedTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ContactsTopPeersNotModified) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeersNotModified#de266ef5 as nil")
	}
	b.PutID(ContactsTopPeersNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *ContactsTopPeersNotModified) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeersNotModified#de266ef5 to nil")
	}
	if err := b.ConsumeID(ContactsTopPeersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.topPeersNotModified#de266ef5: %w", err)
	}
	return nil
}

// construct implements constructor of ContactsTopPeersClass.
func (t ContactsTopPeersNotModified) construct() ContactsTopPeersClass { return &t }

// Ensuring interfaces in compile-time for ContactsTopPeersNotModified.
var (
	_ bin.Encoder = &ContactsTopPeersNotModified{}
	_ bin.Decoder = &ContactsTopPeersNotModified{}

	_ ContactsTopPeersClass = &ContactsTopPeersNotModified{}
)

// ContactsTopPeers represents TL type `contacts.topPeers#70b772a8`.
// Top peers
//
// See https://core.telegram.org/constructor/contacts.topPeers for reference.
type ContactsTopPeers struct {
	// Top peers by top peer category
	Categories []TopPeerCategoryPeers `tl:"categories"`
	// Chats
	Chats []ChatClass `tl:"chats"`
	// Users
	Users []UserClass `tl:"users"`
}

// ContactsTopPeersTypeID is TL type id of ContactsTopPeers.
const ContactsTopPeersTypeID = 0x70b772a8

func (t *ContactsTopPeers) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Categories == nil) {
		return false
	}
	if !(t.Chats == nil) {
		return false
	}
	if !(t.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ContactsTopPeers) String() string {
	if t == nil {
		return "ContactsTopPeers(nil)"
	}
	type Alias ContactsTopPeers
	return fmt.Sprintf("ContactsTopPeers%+v", Alias(*t))
}

// FillFrom fills ContactsTopPeers from given interface.
func (t *ContactsTopPeers) FillFrom(from interface {
	GetCategories() (value []TopPeerCategoryPeers)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	t.Categories = from.GetCategories()
	t.Chats = from.GetChats()
	t.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsTopPeers) TypeID() uint32 {
	return ContactsTopPeersTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsTopPeers) TypeName() string {
	return "contacts.topPeers"
}

// TypeInfo returns info about TL type.
func (t *ContactsTopPeers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.topPeers",
		ID:   ContactsTopPeersTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Categories",
			SchemaName: "categories",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ContactsTopPeers) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeers#70b772a8 as nil")
	}
	b.PutID(ContactsTopPeersTypeID)
	b.PutVectorHeader(len(t.Categories))
	for idx, v := range t.Categories {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field categories element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(t.Chats))
	for idx, v := range t.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(t.Users))
	for idx, v := range t.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetCategories returns value of Categories field.
func (t *ContactsTopPeers) GetCategories() (value []TopPeerCategoryPeers) {
	return t.Categories
}

// GetChats returns value of Chats field.
func (t *ContactsTopPeers) GetChats() (value []ChatClass) {
	return t.Chats
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (t *ContactsTopPeers) MapChats() (value ChatClassArray) {
	return ChatClassArray(t.Chats)
}

// GetUsers returns value of Users field.
func (t *ContactsTopPeers) GetUsers() (value []UserClass) {
	return t.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (t *ContactsTopPeers) MapUsers() (value UserClassArray) {
	return UserClassArray(t.Users)
}

// Decode implements bin.Decoder.
func (t *ContactsTopPeers) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeers#70b772a8 to nil")
	}
	if err := b.ConsumeID(ContactsTopPeersTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field categories: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TopPeerCategoryPeers
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field categories: %w", err)
			}
			t.Categories = append(t.Categories, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field chats: %w", err)
			}
			t.Chats = append(t.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field users: %w", err)
			}
			t.Users = append(t.Users, value)
		}
	}
	return nil
}

// construct implements constructor of ContactsTopPeersClass.
func (t ContactsTopPeers) construct() ContactsTopPeersClass { return &t }

// Ensuring interfaces in compile-time for ContactsTopPeers.
var (
	_ bin.Encoder = &ContactsTopPeers{}
	_ bin.Decoder = &ContactsTopPeers{}

	_ ContactsTopPeersClass = &ContactsTopPeers{}
)

// ContactsTopPeersDisabled represents TL type `contacts.topPeersDisabled#b52c939d`.
// Top peers disabled
//
// See https://core.telegram.org/constructor/contacts.topPeersDisabled for reference.
type ContactsTopPeersDisabled struct {
}

// ContactsTopPeersDisabledTypeID is TL type id of ContactsTopPeersDisabled.
const ContactsTopPeersDisabledTypeID = 0xb52c939d

func (t *ContactsTopPeersDisabled) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ContactsTopPeersDisabled) String() string {
	if t == nil {
		return "ContactsTopPeersDisabled(nil)"
	}
	type Alias ContactsTopPeersDisabled
	return fmt.Sprintf("ContactsTopPeersDisabled%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsTopPeersDisabled) TypeID() uint32 {
	return ContactsTopPeersDisabledTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsTopPeersDisabled) TypeName() string {
	return "contacts.topPeersDisabled"
}

// TypeInfo returns info about TL type.
func (t *ContactsTopPeersDisabled) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.topPeersDisabled",
		ID:   ContactsTopPeersDisabledTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ContactsTopPeersDisabled) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeersDisabled#b52c939d as nil")
	}
	b.PutID(ContactsTopPeersDisabledTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *ContactsTopPeersDisabled) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeersDisabled#b52c939d to nil")
	}
	if err := b.ConsumeID(ContactsTopPeersDisabledTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.topPeersDisabled#b52c939d: %w", err)
	}
	return nil
}

// construct implements constructor of ContactsTopPeersClass.
func (t ContactsTopPeersDisabled) construct() ContactsTopPeersClass { return &t }

// Ensuring interfaces in compile-time for ContactsTopPeersDisabled.
var (
	_ bin.Encoder = &ContactsTopPeersDisabled{}
	_ bin.Decoder = &ContactsTopPeersDisabled{}

	_ ContactsTopPeersClass = &ContactsTopPeersDisabled{}
)

// ContactsTopPeersClass represents contacts.TopPeers generic type.
//
// See https://core.telegram.org/type/contacts.TopPeers for reference.
//
// Example:
//  g, err := tg.DecodeContactsTopPeers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.ContactsTopPeersNotModified: // contacts.topPeersNotModified#de266ef5
//  case *tg.ContactsTopPeers: // contacts.topPeers#70b772a8
//  case *tg.ContactsTopPeersDisabled: // contacts.topPeersDisabled#b52c939d
//  default: panic(v)
//  }
type ContactsTopPeersClass interface {
	bin.Encoder
	bin.Decoder
	construct() ContactsTopPeersClass

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

// DecodeContactsTopPeers implements binary de-serialization for ContactsTopPeersClass.
func DecodeContactsTopPeers(buf *bin.Buffer) (ContactsTopPeersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ContactsTopPeersNotModifiedTypeID:
		// Decoding contacts.topPeersNotModified#de266ef5.
		v := ContactsTopPeersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", err)
		}
		return &v, nil
	case ContactsTopPeersTypeID:
		// Decoding contacts.topPeers#70b772a8.
		v := ContactsTopPeers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", err)
		}
		return &v, nil
	case ContactsTopPeersDisabledTypeID:
		// Decoding contacts.topPeersDisabled#b52c939d.
		v := ContactsTopPeersDisabled{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", bin.NewUnexpectedID(id))
	}
}

// ContactsTopPeers boxes the ContactsTopPeersClass providing a helper.
type ContactsTopPeersBox struct {
	TopPeers ContactsTopPeersClass
}

// Decode implements bin.Decoder for ContactsTopPeersBox.
func (b *ContactsTopPeersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ContactsTopPeersBox to nil")
	}
	v, err := DecodeContactsTopPeers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TopPeers = v
	return nil
}

// Encode implements bin.Encode for ContactsTopPeersBox.
func (b *ContactsTopPeersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TopPeers == nil {
		return fmt.Errorf("unable to encode ContactsTopPeersClass as nil")
	}
	return b.TopPeers.Encode(buf)
}

// ContactsTopPeersClassArray is adapter for slice of ContactsTopPeersClass.
type ContactsTopPeersClassArray []ContactsTopPeersClass

// Sort sorts slice of ContactsTopPeersClass.
func (s ContactsTopPeersClassArray) Sort(less func(a, b ContactsTopPeersClass) bool) ContactsTopPeersClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ContactsTopPeersClass.
func (s ContactsTopPeersClassArray) SortStable(less func(a, b ContactsTopPeersClass) bool) ContactsTopPeersClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ContactsTopPeersClass.
func (s ContactsTopPeersClassArray) Retain(keep func(x ContactsTopPeersClass) bool) ContactsTopPeersClassArray {
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
func (s ContactsTopPeersClassArray) First() (v ContactsTopPeersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ContactsTopPeersClassArray) Last() (v ContactsTopPeersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ContactsTopPeersClassArray) PopFirst() (v ContactsTopPeersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ContactsTopPeersClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ContactsTopPeersClassArray) Pop() (v ContactsTopPeersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsContactsTopPeers returns copy with only ContactsTopPeers constructors.
func (s ContactsTopPeersClassArray) AsContactsTopPeers() (to ContactsTopPeersArray) {
	for _, elem := range s {
		value, ok := elem.(*ContactsTopPeers)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ContactsTopPeersArray is adapter for slice of ContactsTopPeers.
type ContactsTopPeersArray []ContactsTopPeers

// Sort sorts slice of ContactsTopPeers.
func (s ContactsTopPeersArray) Sort(less func(a, b ContactsTopPeers) bool) ContactsTopPeersArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ContactsTopPeers.
func (s ContactsTopPeersArray) SortStable(less func(a, b ContactsTopPeers) bool) ContactsTopPeersArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ContactsTopPeers.
func (s ContactsTopPeersArray) Retain(keep func(x ContactsTopPeers) bool) ContactsTopPeersArray {
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
func (s ContactsTopPeersArray) First() (v ContactsTopPeers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ContactsTopPeersArray) Last() (v ContactsTopPeers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ContactsTopPeersArray) PopFirst() (v ContactsTopPeers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ContactsTopPeers
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ContactsTopPeersArray) Pop() (v ContactsTopPeers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
