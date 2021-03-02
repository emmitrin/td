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

// StatsGroupTopPoster represents TL type `statsGroupTopPoster#18f3d0f7`.
// Information about an active user in a supergroup
//
// See https://core.telegram.org/constructor/statsGroupTopPoster for reference.
type StatsGroupTopPoster struct {
	// User ID
	UserID int `tl:"user_id"`
	// Number of messages for statistics¹ period in consideration
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	Messages int `tl:"messages"`
	// Average number of characters per message
	AvgChars int `tl:"avg_chars"`
}

// StatsGroupTopPosterTypeID is TL type id of StatsGroupTopPoster.
const StatsGroupTopPosterTypeID = 0x18f3d0f7

func (s *StatsGroupTopPoster) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Messages == 0) {
		return false
	}
	if !(s.AvgChars == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsGroupTopPoster) String() string {
	if s == nil {
		return "StatsGroupTopPoster(nil)"
	}
	type Alias StatsGroupTopPoster
	return fmt.Sprintf("StatsGroupTopPoster%+v", Alias(*s))
}

// FillFrom fills StatsGroupTopPoster from given interface.
func (s *StatsGroupTopPoster) FillFrom(from interface {
	GetUserID() (value int)
	GetMessages() (value int)
	GetAvgChars() (value int)
}) {
	s.UserID = from.GetUserID()
	s.Messages = from.GetMessages()
	s.AvgChars = from.GetAvgChars()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGroupTopPoster) TypeID() uint32 {
	return StatsGroupTopPosterTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGroupTopPoster) TypeName() string {
	return "statsGroupTopPoster"
}

// TypeInfo returns info about TL type.
func (s *StatsGroupTopPoster) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsGroupTopPoster",
		ID:   StatsGroupTopPosterTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "AvgChars",
			SchemaName: "avg_chars",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsGroupTopPoster) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopPoster#18f3d0f7 as nil")
	}
	b.PutID(StatsGroupTopPosterTypeID)
	b.PutInt(s.UserID)
	b.PutInt(s.Messages)
	b.PutInt(s.AvgChars)
	return nil
}

// GetUserID returns value of UserID field.
func (s *StatsGroupTopPoster) GetUserID() (value int) {
	return s.UserID
}

// GetMessages returns value of Messages field.
func (s *StatsGroupTopPoster) GetMessages() (value int) {
	return s.Messages
}

// GetAvgChars returns value of AvgChars field.
func (s *StatsGroupTopPoster) GetAvgChars() (value int) {
	return s.AvgChars
}

// Decode implements bin.Decoder.
func (s *StatsGroupTopPoster) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopPoster#18f3d0f7 to nil")
	}
	if err := b.ConsumeID(StatsGroupTopPosterTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: field messages: %w", err)
		}
		s.Messages = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: field avg_chars: %w", err)
		}
		s.AvgChars = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGroupTopPoster.
var (
	_ bin.Encoder = &StatsGroupTopPoster{}
	_ bin.Decoder = &StatsGroupTopPoster{}
)
