package primitives

//ClientStatus struct documented at https://discord.com/developers/docs/topics/gateway#client-status-object
type ClientStatus struct {
	//IsDesktop application session
	Desktop string `json:"desktop"`
	//IsMobile application session
	Mobile string `json:"mobile"`
	//IsWeb or Bot application session
	Web string `json:"web"`
}

//PresenceUpdate struct documented at https://discord.com/developers/docs/topics/gateway#presence-update
type PresenceUpdate struct {
	//User presence is being updated for
	User User `json:"user"`
	//GuildID where PresenceUpdate is for
	GuildID Snowflake `json:"guild_id,string"`
	//Status that is being updated
	Status PresenceStatus `json:"status"`
	//Activities of User
	Activities []Activity `json:"activities"`
	//ClientStatus of User, platform-dependent
	ClientStatus ClientStatus `json:"client_status"`
}

//UserFlag documented at https://discord.com/developers/docs/resources/user#user-object-user-flags
type UserFlag uint32

const (
	//UserFlagNil is a User without any UserFlag(s)
	UserFlagNil UserFlag = 0
	//UserFlagDiscordEmployee is a Discord Employee
	UserFlagDiscordEmployee UserFlag = 1 << 0
	//UserFlagPartneredServerOwner is a Partnered Server Owner
	UserFlagPartneredServerOwner UserFlag = 1 << 1
	//UserFlagHypeSquadEvents is a HypeSquad Event coordinator
	UserFlagHypeSquadEvents UserFlag = 1 << 2
	//UserFlagBugHunterLevel1 is a Level 1 Bug Hunter
	UserFlagBugHunterLevel1 UserFlag = 1 << 3
	//UserFlagHouseBravery is a User that is part of House Bravery
	UserFlagHouseBravery UserFlag = 1 << 6
	//UserFlagHouseBrilliance is a User that is part of House Brilliance
	UserFlagHouseBrilliance UserFlag = 1 << 7
	//UserFlagHouseBalance is a User that is part of House Balance
	UserFlagHouseBalance UserFlag = 1 << 8
	//UserFlagEarlySupporter is an Early Nitro Supporter
	UserFlagEarlySupporter UserFlag = 1 << 9
	//UserFlagTeamUser is a bot team user? todo: don't actually know about this
	UserFlagTeamUser UserFlag = 1 << 10
	//UserFlagBugHunterLevel2 is Level 2 Bug Hunter
	UserFlagBugHunterLevel2 UserFlag = 1 << 14
	//UserFlagVerifiedBot is a Bot that has gone through the Verification process
	UserFlagVerifiedBot UserFlag = 1 << 16
	//UserFlagEarlyVerifiedBotDeveloper is a User who owns a Bot that has gone through the Verification Process when it just came out
	UserFlagEarlyVerifiedBotDeveloper UserFlag = 1 << 17
	//UserFlagDiscordCertifiedModerator is a User who has gone through the discord moderator academy and been active in the moderator Guild
	UserFlagDiscordCertifiedModerator UserFlag = 1 << 18
	//UserFlagAll is a UserFlag of all flags apart from UserFlagNil ANDed together
	UserFlagAll = UserFlagDiscordEmployee | UserFlagPartneredServerOwner | UserFlagHypeSquadEvents | UserFlagBugHunterLevel1 | UserFlagHouseBravery | UserFlagHouseBrilliance | UserFlagHouseBalance | UserFlagEarlySupporter | UserFlagTeamUser | UserFlagBugHunterLevel2 | UserFlagVerifiedBot | UserFlagEarlyVerifiedBotDeveloper | UserFlagDiscordCertifiedModerator
)

//IsValid UserFlag
func (userFlag UserFlag) IsValid() bool {
	return UserFlagAll.Contains(userFlag) && userFlag != UserFlagNil
}

//Contains another UserFlag
func (userFlag UserFlag) Contains(flags UserFlag) bool {
	return userFlag&flags == flags
}

//PremiumType documented at https://discord.com/developers/docs/resources/user#user-object-premium-types
type PremiumType uint8

const (
	//PremiumTypeNil is a User without a Nitro subscription
	PremiumTypeNil PremiumType = iota
	//PremiumTypeNitroClassic is a User with a Classic Nitro subscription
	PremiumTypeNitroClassic
	//PremiumTypeNitro is a User with a Nitro subscription
	PremiumTypeNitro
)

//IsValid PremiumType
func (premiumType PremiumType) IsValid() bool {
	switch premiumType {
	case PremiumTypeNitroClassic,
		PremiumTypeNitro:
		return true
	default:
		return false
	}
}

//User struct from https://discord.com/developers/docs/resources/user#user-object
type User struct {
	//ID of User
	ID Snowflake `json:"id,string"`
	//Username of User, not unique
	Username string `json:"username"`
	//Discriminator of User, 4 suffix digits
	Discriminator string `json:"discriminator"`
	//AvatarHash of User
	AvatarHash ImageHash `json:"avatar"`
	//IsBot User
	IsBot bool `json:"bot"`
	//IsSystemUser maintained by Discord for official communications
	IsSystemUser bool `json:"system"`
	//MFAEnabled is if a User has MultiFactorAuthenticated enabled
	MFAEnabled bool `json:"mfa_enabled"`
	//BannerHash of User
	BannerHash ImageHash `json:"banner"`
	//BannerAccentColor of User as a hexadecimal color code todo: maybe helper function or custom parse for a "color" package color
	BannerAccentColor int `json:"accent_color"`
	//Locale of User
	Locale string `json:"locale"`
	//IsVerified account (by email)
	IsVerified bool `json:"is_verified"`
	//Email of User
	Email string `json:"email"`
	//Flags for User (ex: Discord Employee, Early Supporter)
	Flags UserFlag `json:"flags"`
	//PremiumType of User (aka nitro type)
	PremiumType PremiumType `json:"premium_type"`
	//PublicFlags seen by all User(s)
	PublicFlags UserFlag `json:"public_flags"`
}
