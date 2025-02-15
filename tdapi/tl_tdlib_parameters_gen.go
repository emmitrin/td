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

// TdlibParameters represents TL type `tdlibParameters#d29c1d7b`.
type TdlibParameters struct {
	// If set to true, the Telegram test environment will be used instead of the production
	// environment
	UseTestDC bool
	// The path to the directory for the persistent database; if empty, the current working
	// directory will be used
	DatabaseDirectory string
	// The path to the directory for storing files; if empty, database_directory will be used
	FilesDirectory string
	// If set to true, information about downloaded and uploaded files will be saved between
	// application restarts
	UseFileDatabase bool
	// If set to true, the library will maintain a cache of users, basic groups, supergroups,
	// channels and secret chats. Implies use_file_database
	UseChatInfoDatabase bool
	// If set to true, the library will maintain a cache of chats and messages. Implies
	// use_chat_info_database
	UseMessageDatabase bool
	// If set to true, support for secret chats will be enabled
	UseSecretChats bool
	// Application identifier for Telegram API access, which can be obtained at https://my
	// telegram.org
	APIID int32
	// Application identifier hash for Telegram API access, which can be obtained at
	// https://my.telegram.org
	APIHash string
	// IETF language tag of the user's operating system language; must be non-empty
	SystemLanguageCode string
	// Model of the device the application is being run on; must be non-empty
	DeviceModel string
	// Version of the operating system the application is being run on. If empty, the version
	// is automatically detected by TDLib
	SystemVersion string
	// Application version; must be non-empty
	ApplicationVersion string
	// If set to true, old files will automatically be deleted
	EnableStorageOptimizer bool
	// If set to true, original file names will be ignored. Otherwise, downloaded files will
	// be saved under names as close as possible to the original name
	IgnoreFileNames bool
}

// TdlibParametersTypeID is TL type id of TdlibParameters.
const TdlibParametersTypeID = 0xd29c1d7b

// Ensuring interfaces in compile-time for TdlibParameters.
var (
	_ bin.Encoder     = &TdlibParameters{}
	_ bin.Decoder     = &TdlibParameters{}
	_ bin.BareEncoder = &TdlibParameters{}
	_ bin.BareDecoder = &TdlibParameters{}
)

func (t *TdlibParameters) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.UseTestDC == false) {
		return false
	}
	if !(t.DatabaseDirectory == "") {
		return false
	}
	if !(t.FilesDirectory == "") {
		return false
	}
	if !(t.UseFileDatabase == false) {
		return false
	}
	if !(t.UseChatInfoDatabase == false) {
		return false
	}
	if !(t.UseMessageDatabase == false) {
		return false
	}
	if !(t.UseSecretChats == false) {
		return false
	}
	if !(t.APIID == 0) {
		return false
	}
	if !(t.APIHash == "") {
		return false
	}
	if !(t.SystemLanguageCode == "") {
		return false
	}
	if !(t.DeviceModel == "") {
		return false
	}
	if !(t.SystemVersion == "") {
		return false
	}
	if !(t.ApplicationVersion == "") {
		return false
	}
	if !(t.EnableStorageOptimizer == false) {
		return false
	}
	if !(t.IgnoreFileNames == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TdlibParameters) String() string {
	if t == nil {
		return "TdlibParameters(nil)"
	}
	type Alias TdlibParameters
	return fmt.Sprintf("TdlibParameters%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TdlibParameters) TypeID() uint32 {
	return TdlibParametersTypeID
}

// TypeName returns name of type in TL schema.
func (*TdlibParameters) TypeName() string {
	return "tdlibParameters"
}

// TypeInfo returns info about TL type.
func (t *TdlibParameters) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "tdlibParameters",
		ID:   TdlibParametersTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UseTestDC",
			SchemaName: "use_test_dc",
		},
		{
			Name:       "DatabaseDirectory",
			SchemaName: "database_directory",
		},
		{
			Name:       "FilesDirectory",
			SchemaName: "files_directory",
		},
		{
			Name:       "UseFileDatabase",
			SchemaName: "use_file_database",
		},
		{
			Name:       "UseChatInfoDatabase",
			SchemaName: "use_chat_info_database",
		},
		{
			Name:       "UseMessageDatabase",
			SchemaName: "use_message_database",
		},
		{
			Name:       "UseSecretChats",
			SchemaName: "use_secret_chats",
		},
		{
			Name:       "APIID",
			SchemaName: "api_id",
		},
		{
			Name:       "APIHash",
			SchemaName: "api_hash",
		},
		{
			Name:       "SystemLanguageCode",
			SchemaName: "system_language_code",
		},
		{
			Name:       "DeviceModel",
			SchemaName: "device_model",
		},
		{
			Name:       "SystemVersion",
			SchemaName: "system_version",
		},
		{
			Name:       "ApplicationVersion",
			SchemaName: "application_version",
		},
		{
			Name:       "EnableStorageOptimizer",
			SchemaName: "enable_storage_optimizer",
		},
		{
			Name:       "IgnoreFileNames",
			SchemaName: "ignore_file_names",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TdlibParameters) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode tdlibParameters#d29c1d7b as nil")
	}
	b.PutID(TdlibParametersTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TdlibParameters) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode tdlibParameters#d29c1d7b as nil")
	}
	b.PutBool(t.UseTestDC)
	b.PutString(t.DatabaseDirectory)
	b.PutString(t.FilesDirectory)
	b.PutBool(t.UseFileDatabase)
	b.PutBool(t.UseChatInfoDatabase)
	b.PutBool(t.UseMessageDatabase)
	b.PutBool(t.UseSecretChats)
	b.PutInt32(t.APIID)
	b.PutString(t.APIHash)
	b.PutString(t.SystemLanguageCode)
	b.PutString(t.DeviceModel)
	b.PutString(t.SystemVersion)
	b.PutString(t.ApplicationVersion)
	b.PutBool(t.EnableStorageOptimizer)
	b.PutBool(t.IgnoreFileNames)
	return nil
}

// Decode implements bin.Decoder.
func (t *TdlibParameters) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode tdlibParameters#d29c1d7b to nil")
	}
	if err := b.ConsumeID(TdlibParametersTypeID); err != nil {
		return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TdlibParameters) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode tdlibParameters#d29c1d7b to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field use_test_dc: %w", err)
		}
		t.UseTestDC = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field database_directory: %w", err)
		}
		t.DatabaseDirectory = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field files_directory: %w", err)
		}
		t.FilesDirectory = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field use_file_database: %w", err)
		}
		t.UseFileDatabase = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field use_chat_info_database: %w", err)
		}
		t.UseChatInfoDatabase = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field use_message_database: %w", err)
		}
		t.UseMessageDatabase = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field use_secret_chats: %w", err)
		}
		t.UseSecretChats = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field api_id: %w", err)
		}
		t.APIID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field api_hash: %w", err)
		}
		t.APIHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field system_language_code: %w", err)
		}
		t.SystemLanguageCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field device_model: %w", err)
		}
		t.DeviceModel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field system_version: %w", err)
		}
		t.SystemVersion = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field application_version: %w", err)
		}
		t.ApplicationVersion = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field enable_storage_optimizer: %w", err)
		}
		t.EnableStorageOptimizer = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field ignore_file_names: %w", err)
		}
		t.IgnoreFileNames = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TdlibParameters) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode tdlibParameters#d29c1d7b as nil")
	}
	b.ObjStart()
	b.PutID("tdlibParameters")
	b.Comma()
	b.FieldStart("use_test_dc")
	b.PutBool(t.UseTestDC)
	b.Comma()
	b.FieldStart("database_directory")
	b.PutString(t.DatabaseDirectory)
	b.Comma()
	b.FieldStart("files_directory")
	b.PutString(t.FilesDirectory)
	b.Comma()
	b.FieldStart("use_file_database")
	b.PutBool(t.UseFileDatabase)
	b.Comma()
	b.FieldStart("use_chat_info_database")
	b.PutBool(t.UseChatInfoDatabase)
	b.Comma()
	b.FieldStart("use_message_database")
	b.PutBool(t.UseMessageDatabase)
	b.Comma()
	b.FieldStart("use_secret_chats")
	b.PutBool(t.UseSecretChats)
	b.Comma()
	b.FieldStart("api_id")
	b.PutInt32(t.APIID)
	b.Comma()
	b.FieldStart("api_hash")
	b.PutString(t.APIHash)
	b.Comma()
	b.FieldStart("system_language_code")
	b.PutString(t.SystemLanguageCode)
	b.Comma()
	b.FieldStart("device_model")
	b.PutString(t.DeviceModel)
	b.Comma()
	b.FieldStart("system_version")
	b.PutString(t.SystemVersion)
	b.Comma()
	b.FieldStart("application_version")
	b.PutString(t.ApplicationVersion)
	b.Comma()
	b.FieldStart("enable_storage_optimizer")
	b.PutBool(t.EnableStorageOptimizer)
	b.Comma()
	b.FieldStart("ignore_file_names")
	b.PutBool(t.IgnoreFileNames)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TdlibParameters) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode tdlibParameters#d29c1d7b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("tdlibParameters"); err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: %w", err)
			}
		case "use_test_dc":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field use_test_dc: %w", err)
			}
			t.UseTestDC = value
		case "database_directory":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field database_directory: %w", err)
			}
			t.DatabaseDirectory = value
		case "files_directory":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field files_directory: %w", err)
			}
			t.FilesDirectory = value
		case "use_file_database":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field use_file_database: %w", err)
			}
			t.UseFileDatabase = value
		case "use_chat_info_database":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field use_chat_info_database: %w", err)
			}
			t.UseChatInfoDatabase = value
		case "use_message_database":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field use_message_database: %w", err)
			}
			t.UseMessageDatabase = value
		case "use_secret_chats":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field use_secret_chats: %w", err)
			}
			t.UseSecretChats = value
		case "api_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field api_id: %w", err)
			}
			t.APIID = value
		case "api_hash":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field api_hash: %w", err)
			}
			t.APIHash = value
		case "system_language_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field system_language_code: %w", err)
			}
			t.SystemLanguageCode = value
		case "device_model":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field device_model: %w", err)
			}
			t.DeviceModel = value
		case "system_version":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field system_version: %w", err)
			}
			t.SystemVersion = value
		case "application_version":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field application_version: %w", err)
			}
			t.ApplicationVersion = value
		case "enable_storage_optimizer":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field enable_storage_optimizer: %w", err)
			}
			t.EnableStorageOptimizer = value
		case "ignore_file_names":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode tdlibParameters#d29c1d7b: field ignore_file_names: %w", err)
			}
			t.IgnoreFileNames = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUseTestDC returns value of UseTestDC field.
func (t *TdlibParameters) GetUseTestDC() (value bool) {
	if t == nil {
		return
	}
	return t.UseTestDC
}

// GetDatabaseDirectory returns value of DatabaseDirectory field.
func (t *TdlibParameters) GetDatabaseDirectory() (value string) {
	if t == nil {
		return
	}
	return t.DatabaseDirectory
}

// GetFilesDirectory returns value of FilesDirectory field.
func (t *TdlibParameters) GetFilesDirectory() (value string) {
	if t == nil {
		return
	}
	return t.FilesDirectory
}

// GetUseFileDatabase returns value of UseFileDatabase field.
func (t *TdlibParameters) GetUseFileDatabase() (value bool) {
	if t == nil {
		return
	}
	return t.UseFileDatabase
}

// GetUseChatInfoDatabase returns value of UseChatInfoDatabase field.
func (t *TdlibParameters) GetUseChatInfoDatabase() (value bool) {
	if t == nil {
		return
	}
	return t.UseChatInfoDatabase
}

// GetUseMessageDatabase returns value of UseMessageDatabase field.
func (t *TdlibParameters) GetUseMessageDatabase() (value bool) {
	if t == nil {
		return
	}
	return t.UseMessageDatabase
}

// GetUseSecretChats returns value of UseSecretChats field.
func (t *TdlibParameters) GetUseSecretChats() (value bool) {
	if t == nil {
		return
	}
	return t.UseSecretChats
}

// GetAPIID returns value of APIID field.
func (t *TdlibParameters) GetAPIID() (value int32) {
	if t == nil {
		return
	}
	return t.APIID
}

// GetAPIHash returns value of APIHash field.
func (t *TdlibParameters) GetAPIHash() (value string) {
	if t == nil {
		return
	}
	return t.APIHash
}

// GetSystemLanguageCode returns value of SystemLanguageCode field.
func (t *TdlibParameters) GetSystemLanguageCode() (value string) {
	if t == nil {
		return
	}
	return t.SystemLanguageCode
}

// GetDeviceModel returns value of DeviceModel field.
func (t *TdlibParameters) GetDeviceModel() (value string) {
	if t == nil {
		return
	}
	return t.DeviceModel
}

// GetSystemVersion returns value of SystemVersion field.
func (t *TdlibParameters) GetSystemVersion() (value string) {
	if t == nil {
		return
	}
	return t.SystemVersion
}

// GetApplicationVersion returns value of ApplicationVersion field.
func (t *TdlibParameters) GetApplicationVersion() (value string) {
	if t == nil {
		return
	}
	return t.ApplicationVersion
}

// GetEnableStorageOptimizer returns value of EnableStorageOptimizer field.
func (t *TdlibParameters) GetEnableStorageOptimizer() (value bool) {
	if t == nil {
		return
	}
	return t.EnableStorageOptimizer
}

// GetIgnoreFileNames returns value of IgnoreFileNames field.
func (t *TdlibParameters) GetIgnoreFileNames() (value bool) {
	if t == nil {
		return
	}
	return t.IgnoreFileNames
}
