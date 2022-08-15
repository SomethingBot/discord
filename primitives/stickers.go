package primitives

//StickerType documented at https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
type StickerType uint8

const (
	//StickerTypeNil is a zero-value StickerType
	StickerTypeNil StickerType = iota
	//StickerTypeStandard is a Discord-set Sticker in a pack
	StickerTypeStandard
	//StickerTypeGuild is a Sticker uploaded by a User to a Guild
	StickerTypeGuild
)

//IsValid StickerType
func (stickerType StickerType) IsValid() bool {
	switch stickerType {
	case StickerTypeStandard,
		StickerTypeGuild:
		return true
	default:
		return false
	}
}

//StickerFormatType documented at https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
type StickerFormatType uint8

const (
	//StickerFormatTypeNil is the default StickerFormatType
	StickerFormatTypeNil StickerFormatType = iota
	//StickerFormatTypePNG is a PNG format
	StickerFormatTypePNG
	//StickerFormatTypeAPNG is a APNG format
	StickerFormatTypeAPNG
	//StickerFormatTypeLOTTIE is a LOTTIE format
	StickerFormatTypeLOTTIE
)

func (stickerFormatType StickerFormatType) IsValid() bool {
	switch stickerFormatType {
	case StickerFormatTypePNG,
		StickerFormatTypeAPNG,
		StickerFormatTypeLOTTIE:
		return true
	default:
		return false
	}
}

//Sticker struct from json, documented at https://discord.com/developers/docs/resources/sticker#sticker-object
type Sticker struct {
	//ID of Sticker
	ID Snowflake `json:"id"`
	//PackID of Sticker, if a standard sticker
	PackID Snowflake `json:"pack_id"`
	//Name of Sticker
	Name string `json:"name"`
	//Description of Sticker
	Description string `json:"description"`
	//Tags for Autocomplete when searching for Sticker (max 200 characters)
	Tags string `json:"tags"`
	//Asset Deprecated, previously an Asset hash, now just empty
	Asset string `json:"asset"`
	//Type of Sticker
	Type StickerType `json:"type"`
	//FormatType of Sticker
	FormatType StickerFormatType `json:"format_type"`
	//IsAvailable false if fell below required PremiumTier
	IsAvailable bool `json:"is_available"`
	//GuildID of Guild this Sticker was uploaded to
	GuildID Snowflake `json:"guild_id"`
	//User that uploaded this Sticker
	User User `json:"user"`
	//SortValue this Sticker is in within its pack
	SortValue int `json:"sort_value"`
}
