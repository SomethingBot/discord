package primitives

import (
	"math"
	"time"
)

//SystemChannelFlag (bitwise, potential combination of flags) from https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
type SystemChannelFlag uint8

const (
	//SystemChannelFlagNil is a default Nil SystemChannelFlag
	SystemChannelFlagNil SystemChannelFlag = 0
	//SystemChannelFlagSuppressJoinNotifications SystemChannelFlag to suppress member join notifications
	SystemChannelFlagSuppressJoinNotifications SystemChannelFlag = 1 << (iota - 1)
	//SystemChannelFlagPremiumSubscriptions SystemChannelFlag to suppress guild boost notifications
	SystemChannelFlagPremiumSubscriptions
	//SystemChannelFlagSuppressGuildReminderNotifications SystemChannelFlag to suppress guild setup tips
	SystemChannelFlagSuppressGuildReminderNotifications
	//SystemChannelFlagALL ANDed bitmask of all SystemChannelFlag(s)
	SystemChannelFlagALL SystemChannelFlag = (1 << (iota - 1)) - 1
)

//IsValid SystemChannelFlag
func (systemChannelFlag SystemChannelFlag) IsValid() bool {
	return SystemChannelFlagALL&systemChannelFlag == systemChannelFlag && systemChannelFlag != SystemChannelFlagNil
}

//Contains a SystemChannelFlag
func (systemChannelFlag SystemChannelFlag) Contains(flags SystemChannelFlag) bool {
	return flags&systemChannelFlag == flags && flags != SystemChannelFlagNil
}

//ChannelType from https://discord.com/developers/docs/resources/channel#channel-object-channel-types
type ChannelType uint8

//todo: potentially change ChannelType to not follow discord, as default value is 0 and discord defines GuildText as 0
const (
	//ChannelTypeGuildText is a text Channel within a Guild
	ChannelTypeGuildText ChannelType = iota
	//ChannelTypeDM is a direct message between User(s)
	ChannelTypeDM
	//ChannelTypeGuildVoice is a voice Channel within a Guild
	ChannelTypeGuildVoice
	//ChannelTypeGroupDM is a direct message between multiple User(s)
	ChannelTypeGroupDM
	//ChannelTypeGuildCategory is an OrganizationalCategory (max 50 Channel(s)) documented at https://support.discord.com/hc/en-us/articles/115001580171-Channel-Categories-101
	ChannelTypeGuildCategory
	//ChannelTypeGuildNews is a followable Channel that duplications messages into another Guild
	ChannelTypeGuildNews
	//ChannelTypeGuildStore is a Channel for game sellers
	ChannelTypeGuildStore
	//ChannelTypeGuildNewsThread is a Thread within ChannelTypeGuildNews
	ChannelTypeGuildNewsThread ChannelType = iota + 3
	//ChannelTypeGuildPublicThread is a Thread within ChannelTypeGuildText
	ChannelTypeGuildPublicThread
	//ChannelTypeGuildPrivateThread is a private Thread shown to those with Permission(s)
	ChannelTypeGuildPrivateThread
	//ChannelTypeGuildStageVoice is a Voice Channel for a GuildMember to many GuildMember(s) with options for choosing hosts or raising hands
	ChannelTypeGuildStageVoice
	//ChannelTypeINVALID is a purposefully invalid and not used ChannelType
	ChannelTypeINVALID ChannelType = math.MaxUint8
)

//IsValid ChannelType
func (channelType ChannelType) IsValid() bool {
	switch channelType {
	case ChannelTypeGuildText,
		ChannelTypeDM,
		ChannelTypeGuildVoice,
		ChannelTypeGroupDM,
		ChannelTypeGuildCategory,
		ChannelTypeGuildNews,
		ChannelTypeGuildStore,
		ChannelTypeGuildNewsThread,
		ChannelTypeGuildPublicThread,
		ChannelTypeGuildPrivateThread,
		ChannelTypeGuildStageVoice:
	default:
		return false
	}
	return true
}

//VoiceQualityMode struct from https://discord.com/developers/docs/resources/channel#channel-object-video-quality-modes
type VoiceQualityMode uint8

const (
	//VoiceQualityModeInvalid is an invalid VoiceQualityMode
	VoiceQualityModeInvalid VoiceQualityMode = iota
	//VoiceQualityModeAuto is where VoiceQualityMode is automatically chosen by Discord
	VoiceQualityModeAuto
	//VoiceQualityModeFull is 720p
	VoiceQualityModeFull
)

//IsValid VoiceQualityMode
func (voiceQualityMode VoiceQualityMode) IsValid() bool {
	switch voiceQualityMode {
	case VoiceQualityModeAuto, VoiceQualityModeFull:
		return true
	default:
		return false
	}
}

//OverwriteIDType documented at https://discord.com/developers/docs/resources/channel#overwrite-object-overwrite-structure
type OverwriteIDType uint8

const (
	//OverwriteIDTypeRole is a Role ID
	OverwriteIDTypeRole OverwriteIDType = iota
	//OverwriteIDTypeUser is a User ID
	OverwriteIDTypeUser
)

//IsValid OverwriteIDType
func (overwriteIDType OverwriteIDType) IsValid() bool {
	switch overwriteIDType {
	case OverwriteIDTypeRole,
		OverwriteIDTypeUser:
		return true
	default:
		return false
	}
}

//Overwrite struct from https://discord.com/developers/docs/resources/channel#overwrite-object
type Overwrite struct {
	//ID of Role or User to Overwrite
	ID Snowflake `json:"id"`
	//Type of ID
	Type OverwriteIDType `json:"type"`
	//Allow is PermissionFlag for Overwrite to Allow a Permission
	Allow PermissionFlag `json:"allow,string"`
	//Deny is the PermissionFlag for Overwrite to Deny a Permission
	Deny PermissionFlag `json:"deny,string"`
}

//ThreadMetadata struct from https://discord.com/developers/docs/resources/channel#thread-metadata-object
type ThreadMetadata struct {
	//IsArchived Thread
	IsArchived bool `json:"archived"`
	//AutoArchiveDuration in minutes (possible values are: 60, 1440, 4320, 10080)
	AutoArchiveDuration int `json:"auto_archive_duration"`
	//ArchiveTimestamp is when Thread archive status was last set
	ArchiveTimestamp time.Time `json:"archive_timestamp"`
	//IsLocked Thread
	IsLocked bool `json:"locked"`
	//IsInvitable by non-moderators
	IsInvitable bool `json:"invitable"`
}

//ThreadMember struct from https://discord.com/developers/docs/resources/channel#thread-member-object
type ThreadMember struct {
	//ID of Thread; only sent in GUILD_CREATE event
	ID Snowflake `json:"id"`
	//UserID of User; only sent in GUILD_CREATE event
	UserID Snowflake `json:"user_id"`
	//JoinTimeStamp when User last joined Thread
	JoinTimeStamp Snowflake `json:"join_time_stamp"`
	//Flags for user-thread settings, currently only for notifications
	Flags MessageType `json:"flags"` //todo: find where this is documented, can't find the value possibilities; assuming MessageType
}

//Channel struct from https://discord.com/developers/docs/resources/channel#channel-object
type Channel struct {
	//ID of Channel
	ID Snowflake `json:"id,string"`
	//Type of Channel
	Type ChannelType `json:"type"`
	//GuildID of Guild; may be missing for some channel objects received over gateway guild dispatches)
	GuildID Snowflake `json:"guild_id,string"`
	//Position of sorting
	Position int `json:"position"`
	//PermissionOverwrites for GuildMembers and Roles; these are explicit
	PermissionOverwrites []Overwrite `json:"permission_overwrites"`
	//Name of Channel (1-100 characters)
	Name string `json:"name"`
	//Topic of Channel (0-1024 characters)
	Topic string `json:"topic"`
	//IsNSFW Channel
	IsNSFW bool `json:"nsfw"`
	//LastMessageID for Channel (may not point to actual message)
	LastMessageID Snowflake `json:"last_message_id,string"`
	//Bitrate of Channel if VC
	Bitrate int `json:"bitrate"`
	//MemberLimit of Channel if VC
	MemberLimit int `json:"user_limit"`
	//RateLimitPerMember that a GuildMember has to wait before sending another message (0-21600); members with MANAGE_MESSAGES or MANAGE_CHANNEL are unaffected; applies to both message and create creation
	RateLimitPerMember int `json:"rate_limit_per_user"`
	//Recipients of a DM
	Recipients []User `json:"recipients"`
	//IconHash for grabbing Icon from CDN
	IconHash ImageHash `json:"icon"`
	//OwnerID of group DM or Thread
	OwnerID Snowflake `json:"owner_id,string"`
	//ApplicationID if group DM is bot-created
	ApplicationID Snowflake `json:"application_id,string"`
	//ParentID for a Guild Channel: category, Thread: Channel where created
	ParentID Snowflake `json:"parent_id,string"`
	//LastPinTimestamp maybe nil when no Messages have been pinned
	LastPinTimestamp time.Time `json:"last_pin_timestamp"`
	//VoiceRegion, empty if automatic
	VoiceRegion VoiceRegion `json:"rtc_region"`
	//VideoQualityMode of Channel, 1 when not present
	VoiceQualityMode VoiceQualityMode `json:"voice_quality_mode"`
	//MessageCount, this is Approximate, stops after 50
	MessageCount int `json:"message_count"`
	//MemberCount, this is Approximate, stops after 50
	MemberCount int `json:"member_count"`
	//ThreadMetadata not needed by other channels
	ThreadMetadata ThreadMetadata `json:"thread_metadata"`
	//Member aka current user if they have joined thread; only available for certain API Endpoints
	Member ThreadMember `json:"member"`
	//DefaultAutoArchiveDuration that clients use for newly made threads in minutes after activity; can be set to 60, 1440, 4320, 10080
	DefaultAutoArchiveDuration int `json:"default_auto_archive_duration"`
	//Permissions of Bot user including overwrites; only included when part of RESOLVED data received from SlashCommandInteraction
	Permissions string `json:"permissions"`
}

//ChannelMention documented at https://discord.com/developers/docs/resources/channel#channel-mention-object
type ChannelMention struct {
	//ID of Channel
	ID Snowflake `json:"id,string"`
	//GuildID containing Channel
	GuildID Snowflake `json:"guild_id,string"`
	//Type of Channel
	Type ChannelType `json:"type"`
	//Name of Channel
	Name string `json:"name"`
}
