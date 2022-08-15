package primitives

import "time"

//VoiceRegion struct from https://discord.com/developers/docs/resources/voice#voice-region-object
type VoiceRegion struct {
	//ID of VoiceRegion
	ID string `json:"id"`
	//Name of VoiceRegion
	Name string `json:"name"`
	//IsOptimal VoiceRegion, in terms of distance for Bot
	IsOptimal bool `json:"optimal"`
	//IsDeprecated VoiceRegion
	IsDeprecated bool `json:"deprecated"`
	//IsCustom VoiceRegion
	IsCustom bool `json:"custom"`
}

//VoiceState struct from https://discord.com/developers/docs/resources/voice#voice-state-object
type VoiceState struct {
	//GuildID this VoiceState is from
	GuildID Snowflake `json:"guild_id,string"`
	//ChannelID this VoiceState is from
	ChannelID Snowflake `json:"channel_id,string"`
	//UserID this VoiceState is for
	UserID Snowflake `json:"user_id,string"`
	//GuildMember this VoiceState is for; only filled on GUILD_CREATE Event
	GuildMember GuildMember `json:"member"`
	//SessionID for VoiceState
	SessionID string `json:"session_id"`
	//IsDeafened by Guild
	IsDeafened bool `json:"deaf"`
	//IsMuted by Guild
	IsMuted bool `json:"mute"`
	//IsSelfDeafened by UserID
	IsSelfDeafened bool `json:"self_deaf"`
	//IsSelfMuted by UserID
	IsSelfMuted bool `json:"self_mute"`
	//IsSelfStreaming using "Go Live"
	IsSelfStreaming bool `json:"self_stream"`
	//IsSuppressed by Bot
	IsSuppressed bool `json:"suppress"`
	//RequestToSpeakTimestamp is when a UserID requested to speak
	RequestToSpeakTimestamp time.Time `json:"request_to_speak_timestamp"`
}
