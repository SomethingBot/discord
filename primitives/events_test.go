package primitives

import (
	"embed"
	"encoding/json"
	"testing"
)

func TestGetGatewayEventByName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		eventName string
		want      GatewayEventType
		wantErr   bool
	}{
		{
			eventName: "HELLO",
			want:      GatewayEventTypeHello,
			wantErr:   false,
		},
		{
			eventName: "INGESTOR_INTERNAL_HEARTBEAT_REQUEST",
			want:      GatewayEventTypeHeartbeatRequest,
			wantErr:   false,
		},
		{
			eventName: "INGESTOR_INTERNAL_CLIENT_SHUTDOWN",
			want:      GatewayEventTypeClientShutdown,
			wantErr:   false,
		},
		{
			eventName: "INGESTOR_INTERNAL_HEARTBEAT_ACK",
			want:      GatewayEventTypeHeartbeatACK,
			wantErr:   false,
		},
		{
			eventName: "READY",
			want:      GatewayEventTypeReady,
			wantErr:   false,
		},
		{
			eventName: "RESUMED",
			want:      GatewayEventTypeResumed,
			wantErr:   false,
		},
		{
			eventName: "RECONNECT",
			want:      GatewayEventTypeReconnect,
			wantErr:   false,
		},
		{
			eventName: "INVALID_SESSION",
			want:      GatewayEventTypeInvalidSession,
			wantErr:   false,
		},
		{
			eventName: "CHANNEL_CREATE",
			want:      GatewayEventTypeChannelCreate,
			wantErr:   false,
		},
		{
			eventName: "CHANNEL_UPDATE",
			want:      GatewayEventTypeChannelUpdate,
			wantErr:   false,
		},
		{
			eventName: "CHANNEL_DELETE",
			want:      GatewayEventTypeChannelDelete,
			wantErr:   false,
		},
		{
			eventName: "CHANNEL_PINS_UPDATE",
			want:      GatewayEventTypeChannelPinsUpdate,
			wantErr:   false,
		},
		{
			eventName: "THREAD_CREATE",
			want:      GatewayEventTypeThreadCreate,
			wantErr:   false,
		},
		{
			eventName: "THREAD_UPDATE",
			want:      GatewayEventTypeThreadUpdate,
			wantErr:   false,
		},
		{
			eventName: "THREAD_DELETE",
			want:      GatewayEventTypeThreadDelete,
			wantErr:   false,
		},
		{
			eventName: "THREAD_LIST_SYNC",
			want:      GatewayEventTypeThreadListSync,
			wantErr:   false,
		},
		{
			eventName: "THREAD_MEMBER_UPDATE",
			want:      GatewayEventTypeThreadMemberUpdate,
			wantErr:   false,
		},
		{
			eventName: "THREAD_MEMBERS_UPDATE",
			want:      GatewayEventTypeThreadMembersUpdate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_CREATE",
			want:      GatewayEventTypeGuildCreate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_UPDATE",
			want:      GatewayEventTypeGuildUpdate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_DELETE",
			want:      GatewayEventTypeGuildDelete,
			wantErr:   false,
		},
		{
			eventName: "GUILD_BAN_ADD",
			want:      GatewayEventTypeGuildBanAdd,
			wantErr:   false,
		},
		{
			eventName: "GUILD_BAN_REMOVE",
			want:      GatewayEventTypeGuildBanRemove,
			wantErr:   false,
		},
		{
			eventName: "GUILD_EMOJIS_UPDATE",
			want:      GatewayEventTypeGuildEmojisUpdate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_STICKERS_UPDATE",
			want:      GatewayEventTypeGuildStickersUpdate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_INTEGRATIONS_UPDATE",
			want:      GatewayEventTypeGuildIntegrationsUpdate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_MEMBER_ADD",
			want:      GatewayEventTypeGuildMemberAdd,
			wantErr:   false,
		},
		{
			eventName: "GUILD_MEMBER_REMOVE",
			want:      GatewayEventTypeGuildMemberRemove,
			wantErr:   false,
		},
		{
			eventName: "GUILD_MEMBER_UPDATE",
			want:      GatewayEventTypeGuildMemberUpdate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_MEMBERS_CHUNK",
			want:      GatewayEventTypeGuildMembersChunk,
			wantErr:   false,
		},
		{
			eventName: "GUILD_ROLE_CREATE",
			want:      GatewayEventTypeGuildRoleCreate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_ROLE_UPDATE",
			want:      GatewayEventTypeGuildRoleUpdate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_ROLE_DELETE",
			want:      GatewayEventTypeGuildRoleDelete,
			wantErr:   false,
		},
		{
			eventName: "GUILD_SCHEDULED_EVENT_CREATE",
			want:      GatewayEventTypeGuildScheduledEventCreate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_SCHEDULED_EVENT_UPDATE",
			want:      GatewayEventTypeGuildScheduledEventUpdate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_SCHEDULED_EVENT_DELETE",
			want:      GatewayEventTypeGuildScheduledEventDelete,
			wantErr:   false,
		},
		{
			eventName: "GUILD_SCHEDULED_EVENT_USER_ADD",
			want:      GatewayEventTypeGuildScheduledEventUserAdd,
			wantErr:   false,
		},
		{
			eventName: "GUILD_SCHEDULED_EVENT_USER_REMOVE",
			want:      GatewayEventTypeGuildScheduledEventUserRemove,
			wantErr:   false,
		},
		{
			eventName: "GUILD_INTEGRATION_CREATE",
			want:      GatewayEventTypeGuildIntegrationCreate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_INTEGRATION_UPDATE",
			want:      GatewayEventTypeGuildIntegrationUpdate,
			wantErr:   false,
		},
		{
			eventName: "GUILD_INTEGRATION_DELETE",
			want:      GatewayEventTypeGuildIntegrationDelete,
			wantErr:   false,
		},
		{
			eventName: "GUILD_INTEGRATION_CREATE",
			want:      GatewayEventTypeGuildInteractionCreate,
			wantErr:   false,
		},
		{
			eventName: "INVITE_CREATE",
			want:      GatewayEventTypeInviteCreate,
			wantErr:   false,
		},
		{
			eventName: "INVITE_DELETE",
			want:      GatewayEventTypeInviteDelete,
			wantErr:   false,
		},
		{
			eventName: "MESSAGE_CREATE",
			want:      GatewayEventTypeMessageCreate,
			wantErr:   false,
		},
		{
			eventName: "MESSAGE_UPDATE",
			want:      GatewayEventTypeMessageUpdate,
			wantErr:   false,
		},
		{
			eventName: "MESSAGE_DELETE",
			want:      GatewayEventTypeMessageDelete,
			wantErr:   false,
		},
		{
			eventName: "MESSAGE_DELETE_BULK",
			want:      GatewayEventTypeMessageDeleteBulk,
			wantErr:   false,
		},
		{
			eventName: "MESSAGE_REACTION_ADD",
			want:      GatewayEventTypeMessageReactionAdd,
			wantErr:   false,
		},
		{
			eventName: "MESSAGE_REACTION_REMOVE",
			want:      GatewayEventTypeMessageReactionRemove,
			wantErr:   false,
		},
		{
			eventName: "MESSAGE_REACTION_REMOVE_ALL",
			want:      GatewayEventTypeMessageReactionRemoveAll,
			wantErr:   false,
		},
		{
			eventName: "MESSAGE_REACTION_REMOVE_EMOJI",
			want:      GatewayEventTypeMessageReactionRemoveEmoji,
			wantErr:   false,
		},
		{
			eventName: "PRESENCE_UPDATE",
			want:      GatewayEventTypePresenceUpdate,
			wantErr:   false,
		},
		{
			eventName: "STAGE_INSTANCE_CREATE",
			want:      GatewayEventTypeStageInstanceCreate,
			wantErr:   false,
		},
		{
			eventName: "STAGE_INSTANCE_DELETE",
			want:      GatewayEventTypeStageInstanceDelete,
			wantErr:   false,
		},
		{
			eventName: "STAGE_INSTANCE_UPDATE",
			want:      GatewayEventTypeStageInstanceUpdate,
			wantErr:   false,
		},
		{
			eventName: "TYPING_START",
			want:      GatewayEventTypeTypingStart,
			wantErr:   false,
		},
		{
			eventName: "USER_UPDATE",
			want:      GatewayEventTypeUserUpdate,
			wantErr:   false,
		},
		{
			eventName: "VOICE_STATE_UPDATE",
			want:      GatewayEventTypeVoiceStateUpdate,
			wantErr:   false,
		},
		{
			eventName: "VOICE_SERVER_UPDATE",
			want:      GatewayEventTypeVoiceServerUpdate,
			wantErr:   false,
		},
		{
			eventName: "WEBHOOKS_UPDATE",
			want:      GatewayEventTypeWebhooksUpdate,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.eventName, func(t *testing.T) {
			tt := tt
			t.Parallel()
			got, err := GetGatewayEventByName(tt.eventName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGatewayEventByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Type() != tt.want {
				t.Errorf("GetGatewayEventByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:embed testdata
var testdata embed.FS

func TestGuildEventCreateValiditity(t *testing.T) {
	f, err := testdata.Open("testdata/guildcreateevent.json")
	if err != nil {
		t.Fatal(err)
	}

	decoder := json.NewDecoder(f)
	//decoder.DisallowUnknownFields()

	var e GatewayEventGuildCreate
	err = decoder.Decode(&e)
	if err != nil {
		t.Fatal(err)
	}
}
