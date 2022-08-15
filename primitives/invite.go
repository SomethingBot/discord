package primitives

import "time"

//InviteTargetType documented at https://discord.com/developers/docs/resources/invite#invite-object-invite-target-types
type InviteTargetType int

const (
	//InviteTargetTypeNil is an intentionally nil InviteTargetType to show it is not set
	InviteTargetTypeNil InviteTargetType = 0
	//InviteTargetTypeStream is an Invite to a Stream
	InviteTargetTypeStream InviteTargetType = 1
	//InviteTargetTypeEmbeddedApplication is an Invite to a EmbeddedApplication
	InviteTargetTypeEmbeddedApplication InviteTargetType = 2
)

//IsValid InviteTargetType
func (i InviteTargetType) IsValid() bool {
	switch i {
	case InviteTargetTypeStream, InviteTargetTypeEmbeddedApplication:
		return true
	default:
		return false
	}
}

//InviteStageInstance documented at https://discord.com/developers/docs/resources/invite#invite-stage-instance-object-invite-stage-instance-structure
type InviteStageInstance struct {
	//Members speaking in the StageInstance
	Members []GuildMember `json:"members"`
	//ParticipantCount in the StageInstance
	ParticipantCount int `json:"participant_count"`
	//SpeakerCount in StageInstance
	SpeakerCount int `json:"speaker_count"`
	//Topic of StageInstance (1-120 characters)
	Topic string `json:"topic"`
}

//Invite documented at https://discord.com/developers/docs/resources/invite#invite-object
type Invite struct {
	//Code for Invite
	Code string `json:"code"`
	//Guild Invite is for
	Guild Guild `json:"guild"`
	//Channel Invite is for
	Channel Channel `json:"channel"`
	//Inviter is the User that created the Invite
	Inviter User `json:"inviter"`
	//TargetType is the target type for a ChannelTypeGuildVoice Invite
	TargetType InviteTargetType `json:"target_type"`
	//TargetUser is the User whose Stream to display for a ChannelTypeGuildVoice Invite
	TargetUser User `json:"target_user"`
	//TargetApplication is the EmbeddedApplication to open for a ChannelTypeGuildVoice embedded application Invite
	TargetApplication Application `json:"target_application"`
	//ApproximatePresenceCount of online members
	ApproximatePresenceCount int `json:"approximate_presence_count"`
	//ApproximateMemberCount of total members
	ApproximateMemberCount int `json:"approximate_member_count"`
	//ExpiresAt date
	ExpiresAt time.Time `json:"expires_at"`
	//StageInstance data if there is a public StageInstance in the ChannelTypeGuildStageVoice this Invite is for
	StageInstance InviteStageInstance `json:"stage_instance"`
	//GuildScheduledEvent data if this is to a GuildScheduledEvent
	GuildScheduledEvent GuildScheduledEvent `json:"guild_scheduled_event"`
}

//InviteMetadata documented at https://discord.com/developers/docs/resources/invite#invite-metadata-object
type InviteMetadata struct {
	//Uses of Invite
	Uses int `json:"uses"`
	//MaxUses of Invite
	MaxUses int `json:"max_uses"`
	//MaxAge of Invite
	MaxAge int `json:"max_age"`
	//IsTemporary Invite that only grants temporary membership
	IsTemporary bool `json:"is_temporary"`
	//CreatedAt is when the Invite was created
	CreatedAt time.Time `json:"created_at"`
}
