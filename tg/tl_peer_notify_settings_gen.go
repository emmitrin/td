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

// PeerNotifySettings represents TL type `peerNotifySettings#af509d20`.
// Notification settings.
//
// See https://core.telegram.org/constructor/peerNotifySettings for reference.
type PeerNotifySettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Display text in notifications
	//
	// Use SetShowPreviews and GetShowPreviews helpers.
	ShowPreviews bool `tl:"show_previews"`
	// Mute peer?
	//
	// Use SetSilent and GetSilent helpers.
	Silent bool `tl:"silent"`
	// Mute all notifications until this date
	//
	// Use SetMuteUntil and GetMuteUntil helpers.
	MuteUntil int `tl:"mute_until"`
	// Audio file name for notifications
	//
	// Use SetSound and GetSound helpers.
	Sound string `tl:"sound"`
}

// PeerNotifySettingsTypeID is TL type id of PeerNotifySettings.
const PeerNotifySettingsTypeID = 0xaf509d20

func (p *PeerNotifySettings) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.ShowPreviews == false) {
		return false
	}
	if !(p.Silent == false) {
		return false
	}
	if !(p.MuteUntil == 0) {
		return false
	}
	if !(p.Sound == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PeerNotifySettings) String() string {
	if p == nil {
		return "PeerNotifySettings(nil)"
	}
	type Alias PeerNotifySettings
	return fmt.Sprintf("PeerNotifySettings%+v", Alias(*p))
}

// FillFrom fills PeerNotifySettings from given interface.
func (p *PeerNotifySettings) FillFrom(from interface {
	GetShowPreviews() (value bool, ok bool)
	GetSilent() (value bool, ok bool)
	GetMuteUntil() (value int, ok bool)
	GetSound() (value string, ok bool)
}) {
	if val, ok := from.GetShowPreviews(); ok {
		p.ShowPreviews = val
	}

	if val, ok := from.GetSilent(); ok {
		p.Silent = val
	}

	if val, ok := from.GetMuteUntil(); ok {
		p.MuteUntil = val
	}

	if val, ok := from.GetSound(); ok {
		p.Sound = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PeerNotifySettings) TypeID() uint32 {
	return PeerNotifySettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*PeerNotifySettings) TypeName() string {
	return "peerNotifySettings"
}

// TypeInfo returns info about TL type.
func (p *PeerNotifySettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "peerNotifySettings",
		ID:   PeerNotifySettingsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "ShowPreviews",
			SchemaName: "show_previews",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "MuteUntil",
			SchemaName: "mute_until",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "Sound",
			SchemaName: "sound",
			Null:       !p.Flags.Has(3),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PeerNotifySettings) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerNotifySettings#af509d20 as nil")
	}
	b.PutID(PeerNotifySettingsTypeID)
	if !(p.ShowPreviews == false) {
		p.Flags.Set(0)
	}
	if !(p.Silent == false) {
		p.Flags.Set(1)
	}
	if !(p.MuteUntil == 0) {
		p.Flags.Set(2)
	}
	if !(p.Sound == "") {
		p.Flags.Set(3)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode peerNotifySettings#af509d20: field flags: %w", err)
	}
	if p.Flags.Has(0) {
		b.PutBool(p.ShowPreviews)
	}
	if p.Flags.Has(1) {
		b.PutBool(p.Silent)
	}
	if p.Flags.Has(2) {
		b.PutInt(p.MuteUntil)
	}
	if p.Flags.Has(3) {
		b.PutString(p.Sound)
	}
	return nil
}

// SetShowPreviews sets value of ShowPreviews conditional field.
func (p *PeerNotifySettings) SetShowPreviews(value bool) {
	p.Flags.Set(0)
	p.ShowPreviews = value
}

// GetShowPreviews returns value of ShowPreviews conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetShowPreviews() (value bool, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.ShowPreviews, true
}

// SetSilent sets value of Silent conditional field.
func (p *PeerNotifySettings) SetSilent(value bool) {
	p.Flags.Set(1)
	p.Silent = value
}

// GetSilent returns value of Silent conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetSilent() (value bool, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Silent, true
}

// SetMuteUntil sets value of MuteUntil conditional field.
func (p *PeerNotifySettings) SetMuteUntil(value int) {
	p.Flags.Set(2)
	p.MuteUntil = value
}

// GetMuteUntil returns value of MuteUntil conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetMuteUntil() (value int, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.MuteUntil, true
}

// SetSound sets value of Sound conditional field.
func (p *PeerNotifySettings) SetSound(value string) {
	p.Flags.Set(3)
	p.Sound = value
}

// GetSound returns value of Sound conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetSound() (value string, ok bool) {
	if !p.Flags.Has(3) {
		return value, false
	}
	return p.Sound, true
}

// Decode implements bin.Decoder.
func (p *PeerNotifySettings) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerNotifySettings#af509d20 to nil")
	}
	if err := b.ConsumeID(PeerNotifySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode peerNotifySettings#af509d20: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#af509d20: field flags: %w", err)
		}
	}
	if p.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#af509d20: field show_previews: %w", err)
		}
		p.ShowPreviews = value
	}
	if p.Flags.Has(1) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#af509d20: field silent: %w", err)
		}
		p.Silent = value
	}
	if p.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#af509d20: field mute_until: %w", err)
		}
		p.MuteUntil = value
	}
	if p.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#af509d20: field sound: %w", err)
		}
		p.Sound = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PeerNotifySettings.
var (
	_ bin.Encoder = &PeerNotifySettings{}
	_ bin.Decoder = &PeerNotifySettings{}
)
