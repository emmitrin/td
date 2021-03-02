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

// MessagesBotCallbackAnswer represents TL type `messages.botCallbackAnswer#36585ea4`.
// Callback answer sent by the bot in response to a button press
//
// See https://core.telegram.org/constructor/messages.botCallbackAnswer for reference.
type MessagesBotCallbackAnswer struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Whether an alert should be shown to the user instead of a toast notification
	Alert bool `tl:"alert"`
	// Whether an URL is present
	HasURL bool `tl:"has_url"`
	// Whether to show games in WebView or in native UI.
	NativeUI bool `tl:"native_ui"`
	// Alert to show
	//
	// Use SetMessage and GetMessage helpers.
	Message string `tl:"message"`
	// URL to open
	//
	// Use SetURL and GetURL helpers.
	URL string `tl:"url"`
	// For how long should this answer be cached
	CacheTime int `tl:"cache_time"`
}

// MessagesBotCallbackAnswerTypeID is TL type id of MessagesBotCallbackAnswer.
const MessagesBotCallbackAnswerTypeID = 0x36585ea4

func (b *MessagesBotCallbackAnswer) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.Alert == false) {
		return false
	}
	if !(b.HasURL == false) {
		return false
	}
	if !(b.NativeUI == false) {
		return false
	}
	if !(b.Message == "") {
		return false
	}
	if !(b.URL == "") {
		return false
	}
	if !(b.CacheTime == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *MessagesBotCallbackAnswer) String() string {
	if b == nil {
		return "MessagesBotCallbackAnswer(nil)"
	}
	type Alias MessagesBotCallbackAnswer
	return fmt.Sprintf("MessagesBotCallbackAnswer%+v", Alias(*b))
}

// FillFrom fills MessagesBotCallbackAnswer from given interface.
func (b *MessagesBotCallbackAnswer) FillFrom(from interface {
	GetAlert() (value bool)
	GetHasURL() (value bool)
	GetNativeUI() (value bool)
	GetMessage() (value string, ok bool)
	GetURL() (value string, ok bool)
	GetCacheTime() (value int)
}) {
	b.Alert = from.GetAlert()
	b.HasURL = from.GetHasURL()
	b.NativeUI = from.GetNativeUI()
	if val, ok := from.GetMessage(); ok {
		b.Message = val
	}

	if val, ok := from.GetURL(); ok {
		b.URL = val
	}

	b.CacheTime = from.GetCacheTime()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesBotCallbackAnswer) TypeID() uint32 {
	return MessagesBotCallbackAnswerTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesBotCallbackAnswer) TypeName() string {
	return "messages.botCallbackAnswer"
}

// TypeInfo returns info about TL type.
func (b *MessagesBotCallbackAnswer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.botCallbackAnswer",
		ID:   MessagesBotCallbackAnswerTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "Alert",
			SchemaName: "alert",
			Null:       !b.Flags.Has(1),
		},
		{
			Name:       "HasURL",
			SchemaName: "has_url",
			Null:       !b.Flags.Has(3),
		},
		{
			Name:       "NativeUI",
			SchemaName: "native_ui",
			Null:       !b.Flags.Has(4),
		},
		{
			Name:       "Message",
			SchemaName: "message",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "URL",
			SchemaName: "url",
			Null:       !b.Flags.Has(2),
		},
		{
			Name:       "CacheTime",
			SchemaName: "cache_time",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *MessagesBotCallbackAnswer) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode messages.botCallbackAnswer#36585ea4 as nil")
	}
	buf.PutID(MessagesBotCallbackAnswerTypeID)
	if !(b.Alert == false) {
		b.Flags.Set(1)
	}
	if !(b.HasURL == false) {
		b.Flags.Set(3)
	}
	if !(b.NativeUI == false) {
		b.Flags.Set(4)
	}
	if !(b.Message == "") {
		b.Flags.Set(0)
	}
	if !(b.URL == "") {
		b.Flags.Set(2)
	}
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode messages.botCallbackAnswer#36585ea4: field flags: %w", err)
	}
	if b.Flags.Has(0) {
		buf.PutString(b.Message)
	}
	if b.Flags.Has(2) {
		buf.PutString(b.URL)
	}
	buf.PutInt(b.CacheTime)
	return nil
}

// SetAlert sets value of Alert conditional field.
func (b *MessagesBotCallbackAnswer) SetAlert(value bool) {
	if value {
		b.Flags.Set(1)
		b.Alert = true
	} else {
		b.Flags.Unset(1)
		b.Alert = false
	}
}

// GetAlert returns value of Alert conditional field.
func (b *MessagesBotCallbackAnswer) GetAlert() (value bool) {
	return b.Flags.Has(1)
}

// SetHasURL sets value of HasURL conditional field.
func (b *MessagesBotCallbackAnswer) SetHasURL(value bool) {
	if value {
		b.Flags.Set(3)
		b.HasURL = true
	} else {
		b.Flags.Unset(3)
		b.HasURL = false
	}
}

// GetHasURL returns value of HasURL conditional field.
func (b *MessagesBotCallbackAnswer) GetHasURL() (value bool) {
	return b.Flags.Has(3)
}

// SetNativeUI sets value of NativeUI conditional field.
func (b *MessagesBotCallbackAnswer) SetNativeUI(value bool) {
	if value {
		b.Flags.Set(4)
		b.NativeUI = true
	} else {
		b.Flags.Unset(4)
		b.NativeUI = false
	}
}

// GetNativeUI returns value of NativeUI conditional field.
func (b *MessagesBotCallbackAnswer) GetNativeUI() (value bool) {
	return b.Flags.Has(4)
}

// SetMessage sets value of Message conditional field.
func (b *MessagesBotCallbackAnswer) SetMessage(value string) {
	b.Flags.Set(0)
	b.Message = value
}

// GetMessage returns value of Message conditional field and
// boolean which is true if field was set.
func (b *MessagesBotCallbackAnswer) GetMessage() (value string, ok bool) {
	if !b.Flags.Has(0) {
		return value, false
	}
	return b.Message, true
}

// SetURL sets value of URL conditional field.
func (b *MessagesBotCallbackAnswer) SetURL(value string) {
	b.Flags.Set(2)
	b.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (b *MessagesBotCallbackAnswer) GetURL() (value string, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.URL, true
}

// GetCacheTime returns value of CacheTime field.
func (b *MessagesBotCallbackAnswer) GetCacheTime() (value int) {
	return b.CacheTime
}

// Decode implements bin.Decoder.
func (b *MessagesBotCallbackAnswer) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode messages.botCallbackAnswer#36585ea4 to nil")
	}
	if err := buf.ConsumeID(MessagesBotCallbackAnswerTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.botCallbackAnswer#36585ea4: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode messages.botCallbackAnswer#36585ea4: field flags: %w", err)
		}
	}
	b.Alert = b.Flags.Has(1)
	b.HasURL = b.Flags.Has(3)
	b.NativeUI = b.Flags.Has(4)
	if b.Flags.Has(0) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.botCallbackAnswer#36585ea4: field message: %w", err)
		}
		b.Message = value
	}
	if b.Flags.Has(2) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.botCallbackAnswer#36585ea4: field url: %w", err)
		}
		b.URL = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.botCallbackAnswer#36585ea4: field cache_time: %w", err)
		}
		b.CacheTime = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesBotCallbackAnswer.
var (
	_ bin.Encoder = &MessagesBotCallbackAnswer{}
	_ bin.Decoder = &MessagesBotCallbackAnswer{}
)
