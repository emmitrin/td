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

// DataJSON represents TL type `dataJSON#7d748d04`.
// Represents a json-encoded object
//
// See https://core.telegram.org/constructor/dataJSON for reference.
type DataJSON struct {
	// JSON-encoded object
	Data string `tl:"data"`
}

// DataJSONTypeID is TL type id of DataJSON.
const DataJSONTypeID = 0x7d748d04

func (d *DataJSON) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Data == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DataJSON) String() string {
	if d == nil {
		return "DataJSON(nil)"
	}
	type Alias DataJSON
	return fmt.Sprintf("DataJSON%+v", Alias(*d))
}

// FillFrom fills DataJSON from given interface.
func (d *DataJSON) FillFrom(from interface {
	GetData() (value string)
}) {
	d.Data = from.GetData()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DataJSON) TypeID() uint32 {
	return DataJSONTypeID
}

// TypeName returns name of type in TL schema.
func (*DataJSON) TypeName() string {
	return "dataJSON"
}

// TypeInfo returns info about TL type.
func (d *DataJSON) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dataJSON",
		ID:   DataJSONTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DataJSON) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dataJSON#7d748d04 as nil")
	}
	b.PutID(DataJSONTypeID)
	b.PutString(d.Data)
	return nil
}

// GetData returns value of Data field.
func (d *DataJSON) GetData() (value string) {
	return d.Data
}

// Decode implements bin.Decoder.
func (d *DataJSON) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dataJSON#7d748d04 to nil")
	}
	if err := b.ConsumeID(DataJSONTypeID); err != nil {
		return fmt.Errorf("unable to decode dataJSON#7d748d04: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dataJSON#7d748d04: field data: %w", err)
		}
		d.Data = value
	}
	return nil
}

// Ensuring interfaces in compile-time for DataJSON.
var (
	_ bin.Encoder = &DataJSON{}
	_ bin.Decoder = &DataJSON{}
)
