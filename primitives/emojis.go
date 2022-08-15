package primitives

//Emoji struct from https://discord.com/developers/docs/resources/emoji#emoji-object
type Emoji struct {
	//ID of Emoji
	ID Snowflake `json:"id,string"`
	//Name of Emoji; empty in reaction Emoji
	Name string `json:"name"`
	//Roles allowed to use this Emoji
	Roles []Role `json:"roles"`
	//Creator of this Emoji
	Creator User `json:"user"`
	//RequiresColons wrapped around Name to use
	RequiresColons bool `json:"require_colons"`
	//IsManaged Emoji?
	IsManaged bool `json:"managed"`
	//IsAnimated Emoji?
	IsAnimated bool `json:"animated"`
	//IsAvailable to use; may be false if Guild Boosts are removed
	IsAvailable bool `json:"available"`
}
