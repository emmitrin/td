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

// MessagesVotesList represents TL type `messages.votesList#823f649`.
// How users voted in a poll
//
// See https://core.telegram.org/constructor/messages.votesList for reference.
type MessagesVotesList struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Total number of votes for all options (or only for the chosen option, if provided to
	// messages.getPollVotes¹)
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.getPollVotes
	Count int
	// Vote info for each user
	Votes []MessageUserVoteClass
	// Info about users that voted in the poll
	Users []UserClass
	// Offset to use with the next messages.getPollVotes¹ request, empty string if no more
	// results are available.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.getPollVotes
	//
	// Use SetNextOffset and GetNextOffset helpers.
	NextOffset string
}

// MessagesVotesListTypeID is TL type id of MessagesVotesList.
const MessagesVotesListTypeID = 0x823f649

func (v *MessagesVotesList) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Flags.Zero()) {
		return false
	}
	if !(v.Count == 0) {
		return false
	}
	if !(v.Votes == nil) {
		return false
	}
	if !(v.Users == nil) {
		return false
	}
	if !(v.NextOffset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *MessagesVotesList) String() string {
	if v == nil {
		return "MessagesVotesList(nil)"
	}
	type Alias MessagesVotesList
	return fmt.Sprintf("MessagesVotesList%+v", Alias(*v))
}

// FillFrom fills MessagesVotesList from given interface.
func (v *MessagesVotesList) FillFrom(from interface {
	GetCount() (value int)
	GetVotes() (value []MessageUserVoteClass)
	GetUsers() (value []UserClass)
	GetNextOffset() (value string, ok bool)
}) {
	v.Count = from.GetCount()
	v.Votes = from.GetVotes()
	v.Users = from.GetUsers()
	if val, ok := from.GetNextOffset(); ok {
		v.NextOffset = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesVotesList) TypeID() uint32 {
	return MessagesVotesListTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesVotesList) TypeName() string {
	return "messages.votesList"
}

// TypeInfo returns info about TL type.
func (v *MessagesVotesList) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.votesList",
		ID:   MessagesVotesListTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Votes",
			SchemaName: "votes",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
			Null:       !v.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *MessagesVotesList) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode messages.votesList#823f649 as nil")
	}
	b.PutID(MessagesVotesListTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *MessagesVotesList) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode messages.votesList#823f649 as nil")
	}
	if !(v.NextOffset == "") {
		v.Flags.Set(0)
	}
	if err := v.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.votesList#823f649: field flags: %w", err)
	}
	b.PutInt(v.Count)
	b.PutVectorHeader(len(v.Votes))
	for idx, v := range v.Votes {
		if v == nil {
			return fmt.Errorf("unable to encode messages.votesList#823f649: field votes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.votesList#823f649: field votes element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(v.Users))
	for idx, v := range v.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.votesList#823f649: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.votesList#823f649: field users element with index %d: %w", idx, err)
		}
	}
	if v.Flags.Has(0) {
		b.PutString(v.NextOffset)
	}
	return nil
}

// GetCount returns value of Count field.
func (v *MessagesVotesList) GetCount() (value int) {
	return v.Count
}

// GetVotes returns value of Votes field.
func (v *MessagesVotesList) GetVotes() (value []MessageUserVoteClass) {
	return v.Votes
}

// MapVotes returns field Votes wrapped in MessageUserVoteClassArray helper.
func (v *MessagesVotesList) MapVotes() (value MessageUserVoteClassArray) {
	return MessageUserVoteClassArray(v.Votes)
}

// GetUsers returns value of Users field.
func (v *MessagesVotesList) GetUsers() (value []UserClass) {
	return v.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (v *MessagesVotesList) MapUsers() (value UserClassArray) {
	return UserClassArray(v.Users)
}

// SetNextOffset sets value of NextOffset conditional field.
func (v *MessagesVotesList) SetNextOffset(value string) {
	v.Flags.Set(0)
	v.NextOffset = value
}

// GetNextOffset returns value of NextOffset conditional field and
// boolean which is true if field was set.
func (v *MessagesVotesList) GetNextOffset() (value string, ok bool) {
	if !v.Flags.Has(0) {
		return value, false
	}
	return v.NextOffset, true
}

// Decode implements bin.Decoder.
func (v *MessagesVotesList) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode messages.votesList#823f649 to nil")
	}
	if err := b.ConsumeID(MessagesVotesListTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.votesList#823f649: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *MessagesVotesList) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode messages.votesList#823f649 to nil")
	}
	{
		if err := v.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.votesList#823f649: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.votesList#823f649: field count: %w", err)
		}
		v.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.votesList#823f649: field votes: %w", err)
		}

		if headerLen > 0 {
			v.Votes = make([]MessageUserVoteClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageUserVote(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.votesList#823f649: field votes: %w", err)
			}
			v.Votes = append(v.Votes, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.votesList#823f649: field users: %w", err)
		}

		if headerLen > 0 {
			v.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.votesList#823f649: field users: %w", err)
			}
			v.Users = append(v.Users, value)
		}
	}
	if v.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.votesList#823f649: field next_offset: %w", err)
		}
		v.NextOffset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesVotesList.
var (
	_ bin.Encoder     = &MessagesVotesList{}
	_ bin.Decoder     = &MessagesVotesList{}
	_ bin.BareEncoder = &MessagesVotesList{}
	_ bin.BareDecoder = &MessagesVotesList{}
)
