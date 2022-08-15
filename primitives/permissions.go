package primitives

//PermissionFlag documented at https://discord.com/developers/docs/topics/permissions#permissions
//Check PermissionFlag.IsValid() then PermissionFlag.HasPermission()
type PermissionFlag uint64

const (
	//PermissionFlagNil is the default PermissionFlag value
	PermissionFlagNil PermissionFlag = 0
	//PermissionFlagCreateInstantInvite allows instant invite creation in ChannelTypeGuildText, ChannelTypeGuildVoice, & ChannelTypeGuildStageVoice
	PermissionFlagCreateInstantInvite PermissionFlag = 1 << (iota - 1)
	//PermissionFlagKickMembers allows kicking GuildMember(s)
	PermissionFlagKickMembers
	//PermissionFlagBanMembers allows banning GuildMember(s)
	PermissionFlagBanMembers
	//PermissionFlagAdministrator allows all PermissionFlag(s) and bypasses all Overwrite(s)
	PermissionFlagAdministrator
	//PermissionFlagManageChannels allows management and editing of Channel types ChannelTypeGuildText, ChannelTypeGuildVoice, & ChannelTypeGuildStageVoice
	PermissionFlagManageChannels
	//PermissionFlagManageGuild allows management and editing of a Guild
	PermissionFlagManageGuild
	//PermissionFlagAddReactions allows adding Reaction(s) (an emoji) to a message in a ChannelTypeGuildText
	PermissionFlagAddReactions
	//PermissionFlagViewAuditLog allows viewing AuditLog
	PermissionFlagViewAuditLog
	//PermissionFlagPrioritySpeaker allows being a PrioritySpeaker in a ChannelTypeGuildVoice
	PermissionFlagPrioritySpeaker
	//PermissionFlagStream allows User to start a Stream in a ChannelTypeGuildVoice
	PermissionFlagStream
	//PermissionFlagViewChannel allows a User to view a Channel
	PermissionFlagViewChannel
	//PermissionFlagSendMessages allows a User to send messages in a ChannelTypeGuildText
	PermissionFlagSendMessages
	//PermissionFlagSendTTSMessage allows a User to send Text-To-Speech messages in a ChannelTypeGuildText
	PermissionFlagSendTTSMessage
	//PermissionFlagManageMessages allows a User to manage a ChannelTypeGuildText messages
	PermissionFlagManageMessages
	//PermissionFlagEmbedLinks makes links sent by User with this PermissionFlag to be embedded automatically
	PermissionFlagEmbedLinks
	//PermissionFlagAttachFiles allows a User to upload images and files
	PermissionFlagAttachFiles
	//PermissionFlagReadMessageHistory allows a User to read message history in a ChannelTypeGuildText
	PermissionFlagReadMessageHistory
	//PermissionFlagMentionEveryone allows a User use the @everyone, @here, and all @roles
	PermissionFlagMentionEveryone
	//PermissionFlagUseExternalEmojis allows a User to use Emoji(s) from another Guild
	PermissionFlagUseExternalEmojis
	//PermissionFlagViewGuildInsights allows a User to view Guild Insights
	PermissionFlagViewGuildInsights
	//PermissionFlagConnect allows a User to connect to a ChannelTypeGuildVoice and ChannelTypeGuildStageVoice
	PermissionFlagConnect
	//PermissionFlagSpeak allows a User to speak in a ChannelTypeGuildVoice
	PermissionFlagSpeak
	//PermissionFlagMuteMembers allows a User to mute another User in a ChannelTypeGuildVoice and ChannelTypeGuildStageVoice
	PermissionFlagMuteMembers
	//PermissionFlagDeafenMembers allows a User to deafen another User in a ChannelTypeGuildVoice and ChannelTypeGuildStageVoice
	PermissionFlagDeafenMembers
	//PermissionFlagMoveMembers allows a User to move another User to another ChannelTypeGuildVoice or ChannelTypeGuildStageVoice
	PermissionFlagMoveMembers
	//PermissionFlagUseVoiceActivityDetection allows a User to use VoiceActivity to show when they are speaking
	PermissionFlagUseVoiceActivityDetection
	//PermissionFlagChangeNickname allows a User to change their User.Nickname
	PermissionFlagChangeNickname
	//PermissionFlagManageNicknames allows a User to change others User.Nickname
	PermissionFlagManageNicknames
	//PermissionFlagManageRoles allows a User to manage Role(s)
	PermissionFlagManageRoles
	//PermissionFlagManageWebhooks allows a User to manage Webhooks in a ChannelTypeGuildText
	PermissionFlagManageWebhooks
	//PermissionFlagManageEmojisAndStickers allows a User to manage Emoji(s) and Sticker(s)
	PermissionFlagManageEmojisAndStickers
	//PermissionFlagUseApplicationCommands allows a User to use application commands (slash commands and context menus)
	PermissionFlagUseApplicationCommands
	//PermissionFlagRequestToSpeak allows a User to request to speak in a ChannelTypeGuildStageVoice
	PermissionFlagRequestToSpeak
	//PermissionFlagSkippedNotReal is a fake flag because discord skips 1<<33
	PermissionFlagSkippedNotReal
	//PermissionFlagManageThreads allows a User to manage Thread(s)
	PermissionFlagManageThreads
	//PermissionFlagCreatePublicThreads allows a User create public and announcement threads
	PermissionFlagCreatePublicThreads
	//PermissionFlagCreatePrivateThreads allows a User to create a private Thread
	PermissionFlagCreatePrivateThreads
	//PermissionFlagUseExternalStickers allows a User to use another Guild's Sticker(s)
	PermissionFlagUseExternalStickers
	//PermissionFlagSendMessageInThreads allows a User to send messages in a Thread
	PermissionFlagSendMessageInThreads
	//PermissionFlagStartEmbeddedActivities allows a User to launch Activities (applications with the Embedded flag) in a ChannelTypeGuildVoice
	PermissionFlagStartEmbeddedActivities
	//PermissionModerateMembers allows a User to timeout users
	PermissionModerateMembers
	//PermissionFlagALL is a ANDed of all valid PermissionFlag
	PermissionFlagALL PermissionFlag = (1 << (iota - 1)) - 1
)

//IsValid PermissionFlag
func (permissionFlag PermissionFlag) IsValid() bool {
	return PermissionFlagALL.Contains(permissionFlag) && permissionFlag != PermissionFlagNil && !permissionFlag.Contains(PermissionFlagSkippedNotReal) || permissionFlag == PermissionFlagALL
}

//Contains another PermissionFlag
func (permissionFlag PermissionFlag) Contains(flags PermissionFlag) bool {
	return permissionFlag&flags == flags
}

//todo: implement permission base calculator
//https://discord.com/developers/docs/topics/permissions#permission-overwrites
