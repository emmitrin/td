// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// VoiceNote represents TL type `voiceNote#13ac718`.
type VoiceNote struct {
	// Duration of the voice note, in seconds; as defined by the sender
	Duration int32
	// A waveform representation of the voice note in 5-bit format
	Waveform []byte
	// MIME type of the file; as defined by the sender
	MimeType string
	// True, if speech recognition is completed; Premium users only
	IsRecognized bool
	// Recognized text of the voice note; Premium users only. Call recognizeSpeech to get
	// recognized text of the voice note
	RecognizedText string
	// File containing the voice note
	Voice File
}

// VoiceNoteTypeID is TL type id of VoiceNote.
const VoiceNoteTypeID = 0x13ac718

// Ensuring interfaces in compile-time for VoiceNote.
var (
	_ bin.Encoder     = &VoiceNote{}
	_ bin.Decoder     = &VoiceNote{}
	_ bin.BareEncoder = &VoiceNote{}
	_ bin.BareDecoder = &VoiceNote{}
)

func (v *VoiceNote) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Duration == 0) {
		return false
	}
	if !(v.Waveform == nil) {
		return false
	}
	if !(v.MimeType == "") {
		return false
	}
	if !(v.IsRecognized == false) {
		return false
	}
	if !(v.RecognizedText == "") {
		return false
	}
	if !(v.Voice.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *VoiceNote) String() string {
	if v == nil {
		return "VoiceNote(nil)"
	}
	type Alias VoiceNote
	return fmt.Sprintf("VoiceNote%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*VoiceNote) TypeID() uint32 {
	return VoiceNoteTypeID
}

// TypeName returns name of type in TL schema.
func (*VoiceNote) TypeName() string {
	return "voiceNote"
}

// TypeInfo returns info about TL type.
func (v *VoiceNote) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "voiceNote",
		ID:   VoiceNoteTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
		{
			Name:       "Waveform",
			SchemaName: "waveform",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "IsRecognized",
			SchemaName: "is_recognized",
		},
		{
			Name:       "RecognizedText",
			SchemaName: "recognized_text",
		},
		{
			Name:       "Voice",
			SchemaName: "voice",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *VoiceNote) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode voiceNote#13ac718 as nil")
	}
	b.PutID(VoiceNoteTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *VoiceNote) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode voiceNote#13ac718 as nil")
	}
	b.PutInt32(v.Duration)
	b.PutBytes(v.Waveform)
	b.PutString(v.MimeType)
	b.PutBool(v.IsRecognized)
	b.PutString(v.RecognizedText)
	if err := v.Voice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode voiceNote#13ac718: field voice: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *VoiceNote) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode voiceNote#13ac718 to nil")
	}
	if err := b.ConsumeID(VoiceNoteTypeID); err != nil {
		return fmt.Errorf("unable to decode voiceNote#13ac718: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *VoiceNote) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode voiceNote#13ac718 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode voiceNote#13ac718: field duration: %w", err)
		}
		v.Duration = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode voiceNote#13ac718: field waveform: %w", err)
		}
		v.Waveform = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode voiceNote#13ac718: field mime_type: %w", err)
		}
		v.MimeType = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode voiceNote#13ac718: field is_recognized: %w", err)
		}
		v.IsRecognized = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode voiceNote#13ac718: field recognized_text: %w", err)
		}
		v.RecognizedText = value
	}
	{
		if err := v.Voice.Decode(b); err != nil {
			return fmt.Errorf("unable to decode voiceNote#13ac718: field voice: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *VoiceNote) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode voiceNote#13ac718 as nil")
	}
	b.ObjStart()
	b.PutID("voiceNote")
	b.Comma()
	b.FieldStart("duration")
	b.PutInt32(v.Duration)
	b.Comma()
	b.FieldStart("waveform")
	b.PutBytes(v.Waveform)
	b.Comma()
	b.FieldStart("mime_type")
	b.PutString(v.MimeType)
	b.Comma()
	b.FieldStart("is_recognized")
	b.PutBool(v.IsRecognized)
	b.Comma()
	b.FieldStart("recognized_text")
	b.PutString(v.RecognizedText)
	b.Comma()
	b.FieldStart("voice")
	if err := v.Voice.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode voiceNote#13ac718: field voice: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *VoiceNote) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode voiceNote#13ac718 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("voiceNote"); err != nil {
				return fmt.Errorf("unable to decode voiceNote#13ac718: %w", err)
			}
		case "duration":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode voiceNote#13ac718: field duration: %w", err)
			}
			v.Duration = value
		case "waveform":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode voiceNote#13ac718: field waveform: %w", err)
			}
			v.Waveform = value
		case "mime_type":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode voiceNote#13ac718: field mime_type: %w", err)
			}
			v.MimeType = value
		case "is_recognized":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode voiceNote#13ac718: field is_recognized: %w", err)
			}
			v.IsRecognized = value
		case "recognized_text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode voiceNote#13ac718: field recognized_text: %w", err)
			}
			v.RecognizedText = value
		case "voice":
			if err := v.Voice.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode voiceNote#13ac718: field voice: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDuration returns value of Duration field.
func (v *VoiceNote) GetDuration() (value int32) {
	if v == nil {
		return
	}
	return v.Duration
}

// GetWaveform returns value of Waveform field.
func (v *VoiceNote) GetWaveform() (value []byte) {
	if v == nil {
		return
	}
	return v.Waveform
}

// GetMimeType returns value of MimeType field.
func (v *VoiceNote) GetMimeType() (value string) {
	if v == nil {
		return
	}
	return v.MimeType
}

// GetIsRecognized returns value of IsRecognized field.
func (v *VoiceNote) GetIsRecognized() (value bool) {
	if v == nil {
		return
	}
	return v.IsRecognized
}

// GetRecognizedText returns value of RecognizedText field.
func (v *VoiceNote) GetRecognizedText() (value string) {
	if v == nil {
		return
	}
	return v.RecognizedText
}

// GetVoice returns value of Voice field.
func (v *VoiceNote) GetVoice() (value File) {
	if v == nil {
		return
	}
	return v.Voice
}
