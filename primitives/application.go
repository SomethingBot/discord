package primitives

//TeamMembershipState documented at https://discord.com/developers/docs/topics/teams#data-models-membership-state-enum
type TeamMembershipState uint8

const (
	//TeamMembershipStateInvalid is a zero-value from Discord (should never be this)
	TeamMembershipStateInvalid TeamMembershipState = iota
	//TeamMembershipStateInvited is a Member in a valid invite state
	TeamMembershipStateInvited
	//TeamMembershipStateAccepted is a Member who accepted the Membership
	TeamMembershipStateAccepted
)

//IsValid TeamMembershipState
func (teamMembershipState TeamMembershipState) IsValid() bool {
	switch teamMembershipState {
	case TeamMembershipStateInvited,
		TeamMembershipStateAccepted:
		return true
	default:
		return false
	}
}

//TeamMember documented at https://discord.com/developers/docs/topics/teams#data-models-team-member-object
type TeamMember struct {
	//MembershipState on Team
	MembershipState TeamMembershipState `json:"membership_state"`
	//Permissions of TeamMember; always "*"
	Permissions []string `json:"permissions"`
	//TeamID TeamMember is a part of
	TeamID Snowflake `json:"team_id,string"`
	//User this TeamMember is from (with discriminator, flags, id, and username fields filled)
	User User `json:"user"`
}

//Team documented at https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	//IconHash of Team Icon
	IconHash ImageHash `json:"icon"`
	//ID of Team
	ID Snowflake `json:"id,string"`
	//Members of Team
	Members []TeamMember `json:"members"`
	//Name of Team
	Name string `json:"name"`
	//OwnerID of Team
	OwnerID Snowflake `json:"owner_id,string"`
}

//ApplicationFlag documented at https://discord.com/developers/docs/resources/application#application-object-application-flags
type ApplicationFlag uint32

//todo: documented and write constants for ApplicationFlag, and helper functions

//Application documented at https://discord.com/developers/docs/resources/application#application-object
type Application struct {
	//ID of Application
	ID Snowflake `json:"id,string"`
	//Name of Application
	Name string `json:"name"`
	//IconHash of Application
	IconHash ImageHash `json:"icon"`
	//Description of Application
	Description string `json:"description"`
	//RPCOriginsURLs if RPC enabled
	RPCOriginURLS []string `json:"rpc_origin_urls"`
	//IsBotPublic or can only Owner User join this Bot to Guild(s)
	IsBotPublic bool `json:"bot_public"`
	//BotRequiresCodeGrant completion of oauth2 flow
	BotRequiresCodeGrant bool `json:"bot_require_code_grant"`
	//TermsOfServiceURL of Application
	TermsOfServiceURL string `json:"terms_of_service_url"`
	//PrivacyPolicyURL of Application
	PrivacyPolicyURL string `json:"privacy_policy_url"`
	//Owner of Application (with discriminator, flags, id, and username fields filled)
	Owner User `json:"owner"`
	//Summary of Application if sold on Discord Store
	Summary string `json:"summary"`
	//VerifyKey in HEX for Application Interactions and GameSDK GetTicker
	VerifyKey string `json:"verify_key"`
	//Team if Application belongs to a Team
	Team Team `json:"team"`
	//GuildID Application has been linked to
	GuildID Snowflake `json:"guild_id,string"`
	//PrimarySKUID of Game SKU
	PrimarySKUID Snowflake `json:"primary_sku_id,string"`
	//URLSlug if sold on Discord as a Game
	URLSlug string `json:"slug"`
	//CoverImageHash of Application's default rich presence Invite
	CoverImageHash ImageHash `json:"cover_image"`
	//Flags of Application that are public
	Flags ApplicationFlag `json:"flags"`
}
