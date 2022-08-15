package primitives

import (
	"testing"
)

func TestGatewayIntent_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		gatewayIntent GatewayIntent
		want          bool
	}{
		{
			name:          "GatewayIntentNil",
			gatewayIntent: GatewayIntentNil,
			want:          false,
		},
		{
			name:          "GatewayIntentGuilds",
			gatewayIntent: GatewayIntentGuilds,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildMembers",
			gatewayIntent: GatewayIntentGuildMembers,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildBans",
			gatewayIntent: GatewayIntentGuildBans,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildEmojisAndStickers",
			gatewayIntent: GatewayIntentGuildEmojisAndStickers,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildIntegrations",
			gatewayIntent: GatewayIntentGuildIntegrations,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildWebhooks",
			gatewayIntent: GatewayIntentGuildWebhooks,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildInvites",
			gatewayIntent: GatewayIntentGuildInvites,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildVoiceStates",
			gatewayIntent: GatewayIntentGuildVoiceStates,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildPresences",
			gatewayIntent: GatewayIntentGuildPresences,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildMessages",
			gatewayIntent: GatewayIntentGuildMessages,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildMessageReactions",
			gatewayIntent: GatewayIntentGuildMessageReactions,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildMessageTyping",
			gatewayIntent: GatewayIntentGuildMessageTyping,
			want:          true,
		},
		{
			name:          "GatewayIntentDirectMessages",
			gatewayIntent: GatewayIntentDirectMessages,
			want:          true,
		},
		{
			name:          "GatewayIntentDirectMessageReactions",
			gatewayIntent: GatewayIntentDirectMessageReactions,
			want:          true,
		},
		{
			name:          "GatewayIntentDirectMessageTyping",
			gatewayIntent: GatewayIntentDirectMessageTyping,
			want:          true,
		},
		{
			name:          "GatewayIntentAll",
			gatewayIntent: GatewayIntentAll,
			want:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.gatewayIntent.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGatewayIntent_Contains(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		gatewayIntent GatewayIntent
		intent        GatewayIntent
		want          bool
	}{
		{
			name:          "GatewayIntentNil",
			gatewayIntent: GatewayIntentNil,
			intent:        GatewayIntentAll,
			want:          false,
		},
		{
			name:          "GatewayIntentGuilds",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuilds,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildMembers",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildMembers,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildBans",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildBans,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildEmojisAndStickers",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildEmojisAndStickers,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildIntegrations",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildIntegrations,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildWebhooks",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildWebhooks,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildInvites",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildInvites,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildVoiceStates",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildVoiceStates,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildPresences",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildPresences,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildMessages",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildMessages,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildMessageReactions",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildMessageReactions,
			want:          true,
		},
		{
			name:          "GatewayIntentGuildMessageTyping",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentGuildMessageTyping,
			want:          true,
		},
		{
			name:          "GatewayIntentDirectMessages",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentDirectMessages,
			want:          true,
		},
		{
			name:          "GatewayIntentDirectMessageReactions",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentDirectMessageReactions,
			want:          true,
		},
		{
			name:          "GatewayIntentDirectMessageTyping",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentDirectMessageTyping,
			want:          true,
		},
		{
			name:          "GatewayIntentAll",
			gatewayIntent: GatewayIntentAll,
			intent:        GatewayIntentAll,
			want:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.gatewayIntent.Contains(tt.intent); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
