package primitives

import (
	"encoding/json"
)

//GatewayIntent from https://discord.com/developers/docs/topics/gateway#gateway-intents
type GatewayIntent uint16

const (
	//GatewayIntentNil is when no GatewayIntent is set
	GatewayIntentNil GatewayIntent = 0
	//GatewayIntentGuilds contains events:
	//- GuildCreate
	//- GuildUpdate
	//- GuildDelete
	//- GuildRoleCreate
	//- GuildRoleUpdate
	//- GuildRoleDelete
	//- ChannelCreate
	//- ChannelUpdate
	//- ChannelDelete
	//- ChannelPinsUpdate
	//- ThreadCreate
	//- ThreadUpdate
	//- ThreadDelete
	//- ThreadListSync
	//- ThreadMemberUpdate
	//- ThreadMembersUpdate; data is different depending on intents used
	//- StageInstanceCreate
	//- StageInstanceUpdate
	//- StageInstanceDelete
	GatewayIntentGuilds GatewayIntent = 1 << (iota - 1)
	//GatewayIntentGuildMembers contains events:
	//- GuildMemberAdd
	//- GuildMemberUpdate
	//- GuildMemberRemove
	//- ThreadMembersUpdate *
	GatewayIntentGuildMembers
	//GatewayIntentGuildBans contains events:
	//- GuildBanAdd
	//- GuildBanRemove
	GatewayIntentGuildBans
	//GatewayIntentGuildEmojisAndStickers contains events:
	//- GuildEmojisUpdate
	//- GuildStickersUpdate
	GatewayIntentGuildEmojisAndStickers
	//GatewayIntentGuildIntegrations contains events:
	//- GuildIntegrationsUpdate
	//- IntegrationCreate
	//- IntegrationUpdate
	//- IntegrationDelete
	GatewayIntentGuildIntegrations
	//GatewayIntentGuildWebhooks contains events:
	//- WebhooksUpdate
	GatewayIntentGuildWebhooks
	//GatewayIntentGuildInvites contains events:
	//- InviteCreate
	//- InviteDelete
	GatewayIntentGuildInvites
	//GatewayIntentGuildVoiceStates contains events:
	//- VoiceStateUpdate
	GatewayIntentGuildVoiceStates
	//GatewayIntentGuildPresences contains events:
	//- PresenceUpdate
	GatewayIntentGuildPresences
	//GatewayIntentGuildMessages contains events:
	//- MessageCreate
	//- MessageUpdate
	//- MessageDelete
	//- MessageDeleteBulk
	GatewayIntentGuildMessages
	//GatewayIntentGuildMessageReactions contains events:
	//- MessageReactionAdd
	//- MessageReactionRemove
	//- MessageReactionRemoveAll
	//- MessageReactionRemoveEmoji
	GatewayIntentGuildMessageReactions
	//GatewayIntentGuildMessageTyping contains events:
	//- TypingStart
	GatewayIntentGuildMessageTyping
	//GatewayIntentDirectMessages contains events:
	//- MessageCreate
	//- MessageUpdate
	//- MessageDelete
	//- ChannelPinsUpdate
	GatewayIntentDirectMessages
	//GatewayIntentDirectMessageReactions contains events:
	//- MessageReactionAdd
	//- MessageReactionRemove
	//- MessageReactionRemoveAll
	//- MessageReactionRemoveEmoji
	GatewayIntentDirectMessageReactions
	//GatewayIntentDirectMessageTyping contains events:
	//- TypingStart
	GatewayIntentDirectMessageTyping
	//GatewayIntentAll is a combination of all known GatewayIntents
	GatewayIntentAll GatewayIntent = (1 << (iota - 1)) - 1
)

//IsValid GatewayIntent
func (gatewayIntent GatewayIntent) IsValid() bool {
	return GatewayIntentAll&gatewayIntent == gatewayIntent && gatewayIntent != GatewayIntentNil
}

//Contains another GatewayIntent
func (gatewayIntent GatewayIntent) Contains(intent GatewayIntent) bool {
	return intent&gatewayIntent == intent && intent != GatewayIntentNil
}

//GatewayOpcode of payload sent by Gateway; documented at https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-opcodes
type GatewayOpcode int

const (
	//GatewayOpcodeNil is an intentionally invalid Opcode
	GatewayOpcodeNil GatewayOpcode = -1
	//GatewayOpcodeDispatch is received by a Client for a dispatched GatewayEvent
	GatewayOpcodeDispatch GatewayOpcode = iota - 1
	//GatewayOpcodeHeartbeat is sent or received by a Client to keep a connection alive
	GatewayOpcodeHeartbeat
	//GatewayOpcodeIdentify is sent by a Client to start a new Session during an initial handshake
	GatewayOpcodeIdentify
	//GatewayOpcodePresenceUpdate is sent by a Client to update their Presence
	GatewayOpcodePresenceUpdate
	//GatewayOpcodeVoiceStateUpdate is sent by a Client to move between ChannelTypeGuildVoice
	GatewayOpcodeVoiceStateUpdate
	//GatewayOpcodeResume is sent by a Client to resume a previous Session
	GatewayOpcodeResume GatewayOpcode = iota
	//GatewayOpcodeReconnect is received by a Client to inform them to disconnect and GatewayOpcodeResume
	GatewayOpcodeReconnect
	//GatewayOpcodeRequestGuildMembers is sent by a Client to request information about offline GuildMember(s) in a Guild.IsLarge
	GatewayOpcodeRequestGuildMembers
	//GatewayOpcodeRequestInvalidSession is received by a Client that a Session has been invalidated, Client should reconnect and GatewayOpcodeIdentify or GatewayOpcodeResume
	GatewayOpcodeRequestInvalidSession
	//GatewayOpcodeHello is received by a Client after connecting, containing the heartbeat_interval to use; also used by Client when sending a heartbeat
	GatewayOpcodeHello
	//GatewayOpcodeHeartbeatACK is received by a Client acknowledging a successful GatewayOpcodeHeartbeat
	GatewayOpcodeHeartbeatACK
)

//todo: isvalid functions for Opcodes (check if we send or receive a invalid opcode)

//GatewayErrorEventCode documented at https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-close-event-codes
type GatewayErrorEventCode int

const (
	GatewayErrorEventCodeUnknownError         GatewayErrorEventCode = 4000
	GatewayErrorEventCodeUnknownOpcode        GatewayErrorEventCode = 4001
	GatewayErrorEventCodeDecodeError          GatewayErrorEventCode = 4002
	GatewayErrorEventCodeNotAuthenticated     GatewayErrorEventCode = 4003
	GatewayErrorEventCodeAuthenticationFailed GatewayErrorEventCode = 4004
	GatewayErrorEventCodeAlreadyAuthenticated GatewayErrorEventCode = 4005
	GatewayErrorEventCodeInvalidSequence      GatewayErrorEventCode = 4007
	GatewayErrorEventCodeRateLimited          GatewayErrorEventCode = 4008
	GatewayErrorEventCodeSessionTimedOut      GatewayErrorEventCode = 4009
	GatewayErrorEventCodeInvalidShard         GatewayErrorEventCode = 4010
	GatewayErrorEventCodeSharingRequired      GatewayErrorEventCode = 4011
	GatewayErrorEventCodeInvalidAPIVersion    GatewayErrorEventCode = 4012
	GatewayErrorEventCodeInvalidIntents       GatewayErrorEventCode = 4013
	GatewayErrorEventCodeDisallowedIntents    GatewayErrorEventCode = 4014
)

//GEvent is an Opcode event from the Gateway, I really wish discord didn't send data like this, makes it essentially impossible to parse without multiple passes
//maybe look into seeing if ETF is any better
type GEvent struct {
	//Opcode for payload;
	Opcode         GatewayOpcode   `json:"op"`
	Data           json.RawMessage `json:"d,omitempty"`
	SequenceNumber int             `json:"s,omitempty"`
	Name           string          `json:"t,omitempty"`
}

//GatewayEventHeartbeatACK documented at https://discord.com/developers/docs/topics/gateway#heartbeating-example-gateway-heartbeat-ack
type GatewayEventHeartbeatACK struct{}

func (GatewayEventHeartbeatACK) Type() GatewayEventType {
	return GatewayEventTypeHeartbeatACK
}

func (GatewayEventHeartbeatACK) Opcode() GatewayOpcode {
	return GatewayOpcodeHeartbeatACK
}

type GatewayIdentifyProperties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

type GatewayIdentifyData struct {
	Token      string                    `json:"token"`
	Intents    GatewayIntent             `json:"intents"`
	Properties GatewayIdentifyProperties `json:"properties"`
}

type GatewayIdentify struct {
	Opcode GatewayOpcode       `json:"op"`
	Data   GatewayIdentifyData `json:"d"`
}
