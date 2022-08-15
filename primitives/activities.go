package primitives

import (
	"math"
	"time"
)

//PresenceStatus documented at https://discord.com/developers/docs/topics/gateway#update-presence-status-types
type PresenceStatus string

//Const list of PresenceStatus possibilities
const (
	//PresenceStatusNil is a Nil Presence Status
	PresenceStatusNil PresenceStatus = ""
	//PresenceStatusIdle of PresenceUpdate.Status
	PresenceStatusIdle PresenceStatus = "idle"
	//PresenceStatusDnd of PresenceUpdate.Status
	PresenceStatusDnd PresenceStatus = "dnd"
	//PresenceStatusOnline of PresenceUpdate.Status
	PresenceStatusOnline PresenceStatus = "online"
	//PresenceStatusOffline of PresenceUpdate.Status
	PresenceStatusOffline PresenceStatus = "offline"
)

//IsValid PresenceStatus
func (status PresenceStatus) IsValid() bool {
	switch status {
	case PresenceStatusIdle,
		PresenceStatusDnd,
		PresenceStatusOnline,
		PresenceStatusOffline:
		return true
	default:
		return false
	}
}

//ActivityType documented at https://discord.com/developers/docs/topics/gateway#activity-object-activity-types
type ActivityType uint8

//todo: potentially make ActivityType default value not the same as the discord API; makes the default value not ActivityTypeGame
const (
	//ActivityTypeGame is when ActivityType is a Game
	ActivityTypeGame ActivityType = iota
	//ActivityTypeStreaming is when ActivityType is a Stream
	ActivityTypeStreaming
	//ActivityTypeListening is when ActivityType is Listening
	ActivityTypeListening
	//ActivityTypeWatching is when ActivityType is Watching
	ActivityTypeWatching
	//ActivityTypeCustom is when ActivityType is a Custom PresenceStatus
	ActivityTypeCustom
	//ActivityTypeCompeting is when ActivityType is Competing
	ActivityTypeCompeting
	//ActivityTypeInvalid is purposefully when ActivityType is Invalid
	ActivityTypeInvalid = math.MaxUint8
)

//IsValid ActivityType
func (activityType ActivityType) IsValid() bool {
	switch activityType {
	case ActivityTypeGame,
		ActivityTypeStreaming,
		ActivityTypeListening,
		ActivityTypeWatching,
		ActivityTypeCustom,
		ActivityTypeCompeting:
		return true
	default:
		return false
	}

}

//ActivityTimestamp documented at https://discord.com/developers/docs/topics/gateway#activity-object-activity-timestamps
type ActivityTimestamp struct {
	//Start of Activity
	Start time.Time `json:"start"`
	//End of Activity
	End time.Time `json:"end"`
}

//ActivityParty documented at https://discord.com/developers/docs/topics/gateway#activity-object-activity-party
type ActivityParty struct {
	//ID of ActivityParty
	ID string `json:"id"`
	//Size of ActivityParty; size[0] is current size, size[1] is max size
	Size [2]int `json:"size"`
}

//ActivityAssets documented at https://discord.com/developers/docs/topics/gateway#activity-object-activity-assets
type ActivityAssets struct {
	//LargeImage Asset todo: check if this is for CDN
	LargeImage Snowflake `json:"large_image"`
	//LargeText shown when hovering over large image of Activity
	LargeText string `json:"large_text"`
	//SmallImage Asset todo: check if this is for CDN
	SmallImage Snowflake `json:"small_image"`
	//SmallText shown when hovering over small image of Activity
	SmallText Snowflake `json:"small_text"`
}

//ActivitySecrets documented at https://discord.com/developers/docs/topics/gateway#activity-object-activity-secrets
type ActivitySecrets struct {
	//Join Secret for ActivityParty
	Join string `json:"join"`
	//Spectate Secret for ActivityParty
	Spectate string `json:"spectate"`
	//Match Secret for an Instanced Match
	Match string `json:"match"`
}

//ActivityFlag (bitwise, potential combination of flags) documented at https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
type ActivityFlag uint8

const (
	//ActivityFlagInvalid is a purposefully invalid Flag; for it is the default value of ActivityFlag
	ActivityFlagInvalid ActivityFlag = 0
	//ActivityFlagInstance of an Activity
	ActivityFlagInstance ActivityFlag = 1 << (iota - 1)
	//ActivityFlagJoin of an Activity
	ActivityFlagJoin
	//ActivityFlagSpectate of an Activity
	ActivityFlagSpectate
	//ActivityFlagJoinRequest of an Activity
	ActivityFlagJoinRequest
	//ActivityFlagSync of an Activity
	ActivityFlagSync
	//ActivityFlagPlay of an Activity
	ActivityFlagPlay
	//ActivityFlagAll ANDed bitmask of all ActivityFlag(s)
	ActivityFlagAll ActivityFlag = (1 << (iota - 1)) - 1
)

//IsValid ActivityFlag
func (activityFlag ActivityFlag) IsValid() bool {
	return ActivityFlagAll&activityFlag == activityFlag && activityFlag != ActivityFlagInvalid
}

//Contains a ActivityFlag
func (activityFlag ActivityFlag) Contains(flags ActivityFlag) bool {
	return flags&activityFlag == flags && flags != ActivityFlagInvalid
}

//Buttons documented at https://discord.com/developers/docs/topics/gateway#activity-object-activity-buttons
type Buttons struct {
	//Label of Button (1-32 characters)
	Label string `json:"label"`
	//Url opened when clicking button (1-512 characters)
	Url string `json:"url"`
}

//Activity struct documented at https://discord.com/developers/docs/topics/gateway#activity-object
type Activity struct {
	//Name of Activity
	Name string `json:"name"`
	//Type of Activity
	Type ActivityType `json:"type,string"`
	//Url of Stream when Type=Streaming
	Url string `json:"url"`
	//CreatedAt unix timestamp in milliseconds when activity was created in user session
	CreatedAt time.Time `json:"created_at"`
	//Timestamps for start and end of game
	Timestamps []ActivityTimestamp `json:"timestamps"`
	//ApplicationID of game
	ApplicationID Snowflake `json:"application_id"`
	//Details of what User is doing
	Details string `json:"details"`
	//State of current User party
	State string `json:"state"`
	//Emoji for in custom status
	Emoji Emoji `json:"emoji"`
	//Party of User
	Party ActivityParty `json:"party"`
	//Assets such as images and hover-texts for presence
	Assets ActivityAssets `json:"assets"`
	//Secrets for Rich Presence joining and spectating
	Secrets ActivitySecrets `json:"secrets"`
	//IsInstance of game session
	IsInstance bool `json:"is_instance,string"`
	//Flags of Activity ORed together
	Flags ActivityFlag `json:"flags,string"`
	//Buttons shown in RichPresence (max 2)
	Buttons []Buttons `json:"buttons"`
}
