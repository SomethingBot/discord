package primitives

//PrivacyLevel documented at https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-privacy-level
type PrivacyLevel int8

const (
	PrivacyLevelNil PrivacyLevel = iota
	PrivacyLevelPublic
	PrivacyLevelGuildOnly
)

//IsValid PrivacyLevel
func (privacyLevel PrivacyLevel) IsValid() bool {
	switch privacyLevel {
	case PrivacyLevelPublic,
		PrivacyLevelGuildOnly:
		return true
	default:
		return false
	}
}

//StageInstance struct from json, documented at https://discord.com/developers/docs/resources/stage-instance#stage-instance-object
type StageInstance struct {
	//ID of StageInstance
	ID Snowflake `json:"id"`
	//GuildID of StageInstance
	GuildID Snowflake `json:"guild_id"`
	//ChannelID of StageInstance
	ChannelID Snowflake `json:"channel_id"`
	//Topic of StageInstance
	Topic string `json:"topic"`
	//PrivacyLevel of StageInstance
	PrivacyLevel PrivacyLevel `json:"privacy_level"`
	//IsDiscoverableDisabled aka stage discovery
	IsDiscoverableDisabled bool `json:"discoverable_disabled"`
}
