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

// UrlAuthResultRequest represents TL type `urlAuthResultRequest#92d33a0e`.
// Details about the authorization request, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/url-authorization
//
// See https://core.telegram.org/constructor/urlAuthResultRequest for reference.
type UrlAuthResultRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Whether the bot would like to send messages to the user
	RequestWriteAccess bool `tl:"request_write_access"`
	// Username of a bot, which will be used for user authorization. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot¹ for more details.
	//
	// Links:
	//  1) https://core.telegram.org/widgets/login#linking-your-domain-to-the-bot
	Bot UserClass `tl:"bot"`
	// The domain name of the website on which the user will log in.
	Domain string `tl:"domain"`
}

// UrlAuthResultRequestTypeID is TL type id of UrlAuthResultRequest.
const UrlAuthResultRequestTypeID = 0x92d33a0e

func (u *UrlAuthResultRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.RequestWriteAccess == false) {
		return false
	}
	if !(u.Bot == nil) {
		return false
	}
	if !(u.Domain == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UrlAuthResultRequest) String() string {
	if u == nil {
		return "UrlAuthResultRequest(nil)"
	}
	type Alias UrlAuthResultRequest
	return fmt.Sprintf("UrlAuthResultRequest%+v", Alias(*u))
}

// FillFrom fills UrlAuthResultRequest from given interface.
func (u *UrlAuthResultRequest) FillFrom(from interface {
	GetRequestWriteAccess() (value bool)
	GetBot() (value UserClass)
	GetDomain() (value string)
}) {
	u.RequestWriteAccess = from.GetRequestWriteAccess()
	u.Bot = from.GetBot()
	u.Domain = from.GetDomain()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UrlAuthResultRequest) TypeID() uint32 {
	return UrlAuthResultRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UrlAuthResultRequest) TypeName() string {
	return "urlAuthResultRequest"
}

// TypeInfo returns info about TL type.
func (u *UrlAuthResultRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "urlAuthResultRequest",
		ID:   UrlAuthResultRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "RequestWriteAccess",
			SchemaName: "request_write_access",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "Domain",
			SchemaName: "domain",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UrlAuthResultRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode urlAuthResultRequest#92d33a0e as nil")
	}
	b.PutID(UrlAuthResultRequestTypeID)
	if !(u.RequestWriteAccess == false) {
		u.Flags.Set(0)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode urlAuthResultRequest#92d33a0e: field flags: %w", err)
	}
	if u.Bot == nil {
		return fmt.Errorf("unable to encode urlAuthResultRequest#92d33a0e: field bot is nil")
	}
	if err := u.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode urlAuthResultRequest#92d33a0e: field bot: %w", err)
	}
	b.PutString(u.Domain)
	return nil
}

// SetRequestWriteAccess sets value of RequestWriteAccess conditional field.
func (u *UrlAuthResultRequest) SetRequestWriteAccess(value bool) {
	if value {
		u.Flags.Set(0)
		u.RequestWriteAccess = true
	} else {
		u.Flags.Unset(0)
		u.RequestWriteAccess = false
	}
}

// GetRequestWriteAccess returns value of RequestWriteAccess conditional field.
func (u *UrlAuthResultRequest) GetRequestWriteAccess() (value bool) {
	return u.Flags.Has(0)
}

// GetBot returns value of Bot field.
func (u *UrlAuthResultRequest) GetBot() (value UserClass) {
	return u.Bot
}

// GetDomain returns value of Domain field.
func (u *UrlAuthResultRequest) GetDomain() (value string) {
	return u.Domain
}

// Decode implements bin.Decoder.
func (u *UrlAuthResultRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode urlAuthResultRequest#92d33a0e to nil")
	}
	if err := b.ConsumeID(UrlAuthResultRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode urlAuthResultRequest#92d33a0e: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode urlAuthResultRequest#92d33a0e: field flags: %w", err)
		}
	}
	u.RequestWriteAccess = u.Flags.Has(0)
	{
		value, err := DecodeUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode urlAuthResultRequest#92d33a0e: field bot: %w", err)
		}
		u.Bot = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode urlAuthResultRequest#92d33a0e: field domain: %w", err)
		}
		u.Domain = value
	}
	return nil
}

// construct implements constructor of UrlAuthResultClass.
func (u UrlAuthResultRequest) construct() UrlAuthResultClass { return &u }

// Ensuring interfaces in compile-time for UrlAuthResultRequest.
var (
	_ bin.Encoder = &UrlAuthResultRequest{}
	_ bin.Decoder = &UrlAuthResultRequest{}

	_ UrlAuthResultClass = &UrlAuthResultRequest{}
)

// UrlAuthResultAccepted represents TL type `urlAuthResultAccepted#8f8c0e4e`.
// Details about an accepted authorization request, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/url-authorization
//
// See https://core.telegram.org/constructor/urlAuthResultAccepted for reference.
type UrlAuthResultAccepted struct {
	// The URL name of the website on which the user has logged in.
	URL string `tl:"url"`
}

// UrlAuthResultAcceptedTypeID is TL type id of UrlAuthResultAccepted.
const UrlAuthResultAcceptedTypeID = 0x8f8c0e4e

func (u *UrlAuthResultAccepted) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UrlAuthResultAccepted) String() string {
	if u == nil {
		return "UrlAuthResultAccepted(nil)"
	}
	type Alias UrlAuthResultAccepted
	return fmt.Sprintf("UrlAuthResultAccepted%+v", Alias(*u))
}

// FillFrom fills UrlAuthResultAccepted from given interface.
func (u *UrlAuthResultAccepted) FillFrom(from interface {
	GetURL() (value string)
}) {
	u.URL = from.GetURL()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UrlAuthResultAccepted) TypeID() uint32 {
	return UrlAuthResultAcceptedTypeID
}

// TypeName returns name of type in TL schema.
func (*UrlAuthResultAccepted) TypeName() string {
	return "urlAuthResultAccepted"
}

// TypeInfo returns info about TL type.
func (u *UrlAuthResultAccepted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "urlAuthResultAccepted",
		ID:   UrlAuthResultAcceptedTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UrlAuthResultAccepted) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode urlAuthResultAccepted#8f8c0e4e as nil")
	}
	b.PutID(UrlAuthResultAcceptedTypeID)
	b.PutString(u.URL)
	return nil
}

// GetURL returns value of URL field.
func (u *UrlAuthResultAccepted) GetURL() (value string) {
	return u.URL
}

// Decode implements bin.Decoder.
func (u *UrlAuthResultAccepted) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode urlAuthResultAccepted#8f8c0e4e to nil")
	}
	if err := b.ConsumeID(UrlAuthResultAcceptedTypeID); err != nil {
		return fmt.Errorf("unable to decode urlAuthResultAccepted#8f8c0e4e: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode urlAuthResultAccepted#8f8c0e4e: field url: %w", err)
		}
		u.URL = value
	}
	return nil
}

// construct implements constructor of UrlAuthResultClass.
func (u UrlAuthResultAccepted) construct() UrlAuthResultClass { return &u }

// Ensuring interfaces in compile-time for UrlAuthResultAccepted.
var (
	_ bin.Encoder = &UrlAuthResultAccepted{}
	_ bin.Decoder = &UrlAuthResultAccepted{}

	_ UrlAuthResultClass = &UrlAuthResultAccepted{}
)

// UrlAuthResultDefault represents TL type `urlAuthResultDefault#a9d6db1f`.
// Details about an accepted authorization request, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/url-authorization
//
// See https://core.telegram.org/constructor/urlAuthResultDefault for reference.
type UrlAuthResultDefault struct {
}

// UrlAuthResultDefaultTypeID is TL type id of UrlAuthResultDefault.
const UrlAuthResultDefaultTypeID = 0xa9d6db1f

func (u *UrlAuthResultDefault) Zero() bool {
	if u == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (u *UrlAuthResultDefault) String() string {
	if u == nil {
		return "UrlAuthResultDefault(nil)"
	}
	type Alias UrlAuthResultDefault
	return fmt.Sprintf("UrlAuthResultDefault%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UrlAuthResultDefault) TypeID() uint32 {
	return UrlAuthResultDefaultTypeID
}

// TypeName returns name of type in TL schema.
func (*UrlAuthResultDefault) TypeName() string {
	return "urlAuthResultDefault"
}

// TypeInfo returns info about TL type.
func (u *UrlAuthResultDefault) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "urlAuthResultDefault",
		ID:   UrlAuthResultDefaultTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (u *UrlAuthResultDefault) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode urlAuthResultDefault#a9d6db1f as nil")
	}
	b.PutID(UrlAuthResultDefaultTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UrlAuthResultDefault) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode urlAuthResultDefault#a9d6db1f to nil")
	}
	if err := b.ConsumeID(UrlAuthResultDefaultTypeID); err != nil {
		return fmt.Errorf("unable to decode urlAuthResultDefault#a9d6db1f: %w", err)
	}
	return nil
}

// construct implements constructor of UrlAuthResultClass.
func (u UrlAuthResultDefault) construct() UrlAuthResultClass { return &u }

// Ensuring interfaces in compile-time for UrlAuthResultDefault.
var (
	_ bin.Encoder = &UrlAuthResultDefault{}
	_ bin.Decoder = &UrlAuthResultDefault{}

	_ UrlAuthResultClass = &UrlAuthResultDefault{}
)

// UrlAuthResultClass represents UrlAuthResult generic type.
//
// See https://core.telegram.org/type/UrlAuthResult for reference.
//
// Example:
//  g, err := tg.DecodeUrlAuthResult(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.UrlAuthResultRequest: // urlAuthResultRequest#92d33a0e
//  case *tg.UrlAuthResultAccepted: // urlAuthResultAccepted#8f8c0e4e
//  case *tg.UrlAuthResultDefault: // urlAuthResultDefault#a9d6db1f
//  default: panic(v)
//  }
type UrlAuthResultClass interface {
	bin.Encoder
	bin.Decoder
	construct() UrlAuthResultClass

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

// DecodeUrlAuthResult implements binary de-serialization for UrlAuthResultClass.
func DecodeUrlAuthResult(buf *bin.Buffer) (UrlAuthResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UrlAuthResultRequestTypeID:
		// Decoding urlAuthResultRequest#92d33a0e.
		v := UrlAuthResultRequest{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UrlAuthResultClass: %w", err)
		}
		return &v, nil
	case UrlAuthResultAcceptedTypeID:
		// Decoding urlAuthResultAccepted#8f8c0e4e.
		v := UrlAuthResultAccepted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UrlAuthResultClass: %w", err)
		}
		return &v, nil
	case UrlAuthResultDefaultTypeID:
		// Decoding urlAuthResultDefault#a9d6db1f.
		v := UrlAuthResultDefault{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UrlAuthResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UrlAuthResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// UrlAuthResult boxes the UrlAuthResultClass providing a helper.
type UrlAuthResultBox struct {
	UrlAuthResult UrlAuthResultClass
}

// Decode implements bin.Decoder for UrlAuthResultBox.
func (b *UrlAuthResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UrlAuthResultBox to nil")
	}
	v, err := DecodeUrlAuthResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.UrlAuthResult = v
	return nil
}

// Encode implements bin.Encode for UrlAuthResultBox.
func (b *UrlAuthResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.UrlAuthResult == nil {
		return fmt.Errorf("unable to encode UrlAuthResultClass as nil")
	}
	return b.UrlAuthResult.Encode(buf)
}

// UrlAuthResultClassArray is adapter for slice of UrlAuthResultClass.
type UrlAuthResultClassArray []UrlAuthResultClass

// Sort sorts slice of UrlAuthResultClass.
func (s UrlAuthResultClassArray) Sort(less func(a, b UrlAuthResultClass) bool) UrlAuthResultClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UrlAuthResultClass.
func (s UrlAuthResultClassArray) SortStable(less func(a, b UrlAuthResultClass) bool) UrlAuthResultClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UrlAuthResultClass.
func (s UrlAuthResultClassArray) Retain(keep func(x UrlAuthResultClass) bool) UrlAuthResultClassArray {
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
func (s UrlAuthResultClassArray) First() (v UrlAuthResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UrlAuthResultClassArray) Last() (v UrlAuthResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UrlAuthResultClassArray) PopFirst() (v UrlAuthResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UrlAuthResultClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UrlAuthResultClassArray) Pop() (v UrlAuthResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUrlAuthResultRequest returns copy with only UrlAuthResultRequest constructors.
func (s UrlAuthResultClassArray) AsUrlAuthResultRequest() (to UrlAuthResultRequestArray) {
	for _, elem := range s {
		value, ok := elem.(*UrlAuthResultRequest)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUrlAuthResultAccepted returns copy with only UrlAuthResultAccepted constructors.
func (s UrlAuthResultClassArray) AsUrlAuthResultAccepted() (to UrlAuthResultAcceptedArray) {
	for _, elem := range s {
		value, ok := elem.(*UrlAuthResultAccepted)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// UrlAuthResultRequestArray is adapter for slice of UrlAuthResultRequest.
type UrlAuthResultRequestArray []UrlAuthResultRequest

// Sort sorts slice of UrlAuthResultRequest.
func (s UrlAuthResultRequestArray) Sort(less func(a, b UrlAuthResultRequest) bool) UrlAuthResultRequestArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UrlAuthResultRequest.
func (s UrlAuthResultRequestArray) SortStable(less func(a, b UrlAuthResultRequest) bool) UrlAuthResultRequestArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UrlAuthResultRequest.
func (s UrlAuthResultRequestArray) Retain(keep func(x UrlAuthResultRequest) bool) UrlAuthResultRequestArray {
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
func (s UrlAuthResultRequestArray) First() (v UrlAuthResultRequest, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UrlAuthResultRequestArray) Last() (v UrlAuthResultRequest, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UrlAuthResultRequestArray) PopFirst() (v UrlAuthResultRequest, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UrlAuthResultRequest
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UrlAuthResultRequestArray) Pop() (v UrlAuthResultRequest, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UrlAuthResultAcceptedArray is adapter for slice of UrlAuthResultAccepted.
type UrlAuthResultAcceptedArray []UrlAuthResultAccepted

// Sort sorts slice of UrlAuthResultAccepted.
func (s UrlAuthResultAcceptedArray) Sort(less func(a, b UrlAuthResultAccepted) bool) UrlAuthResultAcceptedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UrlAuthResultAccepted.
func (s UrlAuthResultAcceptedArray) SortStable(less func(a, b UrlAuthResultAccepted) bool) UrlAuthResultAcceptedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UrlAuthResultAccepted.
func (s UrlAuthResultAcceptedArray) Retain(keep func(x UrlAuthResultAccepted) bool) UrlAuthResultAcceptedArray {
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
func (s UrlAuthResultAcceptedArray) First() (v UrlAuthResultAccepted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UrlAuthResultAcceptedArray) Last() (v UrlAuthResultAccepted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UrlAuthResultAcceptedArray) PopFirst() (v UrlAuthResultAccepted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UrlAuthResultAccepted
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UrlAuthResultAcceptedArray) Pop() (v UrlAuthResultAccepted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
