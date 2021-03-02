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

// MessagesSetGameScoreRequest represents TL type `messages.setGameScore#8ef8ecc0`.
// Use this method to set the score of the specified user in a game sent as a normal message (bots only).
//
// See https://core.telegram.org/method/messages.setGameScore for reference.
type MessagesSetGameScoreRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Set this flag if the game message should be automatically edited to include the current scoreboard
	EditMessage bool `tl:"edit_message"`
	// Set this flag if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	Force bool `tl:"force"`
	// Unique identifier of target chat
	Peer InputPeerClass `tl:"peer"`
	// Identifier of the sent message
	ID int `tl:"id"`
	// User identifier
	UserID InputUserClass `tl:"user_id"`
	// New score
	Score int `tl:"score"`
}

// MessagesSetGameScoreRequestTypeID is TL type id of MessagesSetGameScoreRequest.
const MessagesSetGameScoreRequestTypeID = 0x8ef8ecc0

func (s *MessagesSetGameScoreRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.EditMessage == false) {
		return false
	}
	if !(s.Force == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.UserID == nil) {
		return false
	}
	if !(s.Score == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetGameScoreRequest) String() string {
	if s == nil {
		return "MessagesSetGameScoreRequest(nil)"
	}
	type Alias MessagesSetGameScoreRequest
	return fmt.Sprintf("MessagesSetGameScoreRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetGameScoreRequest from given interface.
func (s *MessagesSetGameScoreRequest) FillFrom(from interface {
	GetEditMessage() (value bool)
	GetForce() (value bool)
	GetPeer() (value InputPeerClass)
	GetID() (value int)
	GetUserID() (value InputUserClass)
	GetScore() (value int)
}) {
	s.EditMessage = from.GetEditMessage()
	s.Force = from.GetForce()
	s.Peer = from.GetPeer()
	s.ID = from.GetID()
	s.UserID = from.GetUserID()
	s.Score = from.GetScore()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetGameScoreRequest) TypeID() uint32 {
	return MessagesSetGameScoreRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetGameScoreRequest) TypeName() string {
	return "messages.setGameScore"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetGameScoreRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setGameScore",
		ID:   MessagesSetGameScoreRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "EditMessage",
			SchemaName: "edit_message",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Force",
			SchemaName: "force",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Score",
			SchemaName: "score",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetGameScoreRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setGameScore#8ef8ecc0 as nil")
	}
	b.PutID(MessagesSetGameScoreRequestTypeID)
	if !(s.EditMessage == false) {
		s.Flags.Set(0)
	}
	if !(s.Force == false) {
		s.Flags.Set(1)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setGameScore#8ef8ecc0: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.setGameScore#8ef8ecc0: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setGameScore#8ef8ecc0: field peer: %w", err)
	}
	b.PutInt(s.ID)
	if s.UserID == nil {
		return fmt.Errorf("unable to encode messages.setGameScore#8ef8ecc0: field user_id is nil")
	}
	if err := s.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setGameScore#8ef8ecc0: field user_id: %w", err)
	}
	b.PutInt(s.Score)
	return nil
}

// SetEditMessage sets value of EditMessage conditional field.
func (s *MessagesSetGameScoreRequest) SetEditMessage(value bool) {
	if value {
		s.Flags.Set(0)
		s.EditMessage = true
	} else {
		s.Flags.Unset(0)
		s.EditMessage = false
	}
}

// GetEditMessage returns value of EditMessage conditional field.
func (s *MessagesSetGameScoreRequest) GetEditMessage() (value bool) {
	return s.Flags.Has(0)
}

// SetForce sets value of Force conditional field.
func (s *MessagesSetGameScoreRequest) SetForce(value bool) {
	if value {
		s.Flags.Set(1)
		s.Force = true
	} else {
		s.Flags.Unset(1)
		s.Force = false
	}
}

// GetForce returns value of Force conditional field.
func (s *MessagesSetGameScoreRequest) GetForce() (value bool) {
	return s.Flags.Has(1)
}

// GetPeer returns value of Peer field.
func (s *MessagesSetGameScoreRequest) GetPeer() (value InputPeerClass) {
	return s.Peer
}

// GetID returns value of ID field.
func (s *MessagesSetGameScoreRequest) GetID() (value int) {
	return s.ID
}

// GetUserID returns value of UserID field.
func (s *MessagesSetGameScoreRequest) GetUserID() (value InputUserClass) {
	return s.UserID
}

// GetScore returns value of Score field.
func (s *MessagesSetGameScoreRequest) GetScore() (value int) {
	return s.Score
}

// Decode implements bin.Decoder.
func (s *MessagesSetGameScoreRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setGameScore#8ef8ecc0 to nil")
	}
	if err := b.ConsumeID(MessagesSetGameScoreRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setGameScore#8ef8ecc0: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setGameScore#8ef8ecc0: field flags: %w", err)
		}
	}
	s.EditMessage = s.Flags.Has(0)
	s.Force = s.Flags.Has(1)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setGameScore#8ef8ecc0: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setGameScore#8ef8ecc0: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setGameScore#8ef8ecc0: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setGameScore#8ef8ecc0: field score: %w", err)
		}
		s.Score = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetGameScoreRequest.
var (
	_ bin.Encoder = &MessagesSetGameScoreRequest{}
	_ bin.Decoder = &MessagesSetGameScoreRequest{}
)

// MessagesSetGameScore invokes method messages.setGameScore#8ef8ecc0 returning error if any.
// Use this method to set the score of the specified user in a game sent as a normal message (bots only).
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 USER_BOT_REQUIRED: This method can only be called by a bot
//
// See https://core.telegram.org/method/messages.setGameScore for reference.
// Can be used by bots.
func (c *Client) MessagesSetGameScore(ctx context.Context, request *MessagesSetGameScoreRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
