package primitives

//RoleTag struct from https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure
type RoleTag struct {
	//BotID that owns this Role
	BotID Snowflake `json:"bot_id"`
	//IntegrationID that this Role belongs to
	IntegrationID Snowflake `json:"integration_id"`
	//IsPremiumSubscriber Role aka booster Role
	IsPremiumSubscriber bool `json:"premium_subscriber"`
}

//Role struct from https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	//ID of Role
	ID Snowflake `json:"id,string"`
	//Name of Role
	Name string `json:"name"`
	//Color of Role
	Color int `json:"color"`
	//IsHoist aka pinned in GuildMember list on client
	IsHoist bool `json:"hoist"`
	//IconHash is the Role's IconHash for grabbing from CDN
	IconHash ImageHash `json:"icon"`
	//UnicodeEmoji equivalent
	UnicodeEmoji string `json:"unicode_emoji"`
	//Position of Role
	Position int `json:"position"`
	//Permissions for Role as a PermissionsBitSet
	Permissions PermissionFlag `json:"permissions,string"`
	//IsManaged by integration?
	IsManaged bool `json:"managed"`
	//IsMentionable Emoji?
	Mentionable bool `json:"mentionable"`
	//Tag for role
	Tag RoleTag `json:"tags"`
}
